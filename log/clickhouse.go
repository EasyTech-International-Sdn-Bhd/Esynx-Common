package log

import (
	"database/sql"
	"fmt"
	"github.com/easytech-international-sdn-bhd/esynx-common/entities"
	_ "github.com/mailru/go-clickhouse/v2"
	"time"
)

type ClickHouseLogStorage struct {
	conn       *sql.DB
	storageCnf *LogStorageConn
}

func NewClickHouseLogStorage(storageCnf *LogStorageConn) (*ClickHouseLogStorage, error) {
	var err error
	conn, err := initConnection(storageCnf)
	if err != nil {
		return nil, err
	}
	return &ClickHouseLogStorage{
		conn:       conn,
		storageCnf: storageCnf,
	}, nil
}

func initConnection(storageCnf *LogStorageConn) (*sql.DB, error) {
	conn, err := sql.Open("chhttp", fmt.Sprintf("http://%s:%s@%s:%d/%s",
		storageCnf.Username,
		storageCnf.Password,
		storageCnf.Host,
		storageCnf.Port,
		storageCnf.Database))
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(4)
	conn.SetMaxIdleConns(3)
	conn.SetConnMaxLifetime(time.Minute * 1)
	if err = conn.Ping(); err != nil {
		fmt.Printf("error ping clickhouse %s", err.Error())
		return nil, err
	}
	return conn, nil
}

func (s *ClickHouseLogStorage) isConnected() error {
	if s.conn == nil {
		var err error
		s.conn, err = initConnection(s.storageCnf)
		if err != nil {
			return err
		}
		return nil
	}
	if err := s.conn.Ping(); err != nil {
		s.conn, err = initConnection(s.storageCnf)
		if err != nil {
			s.conn = nil
			return err
		}
	}
	return nil
}

func (s *ClickHouseLogStorage) LogAudit(clientId string, data []*entities.AuditLog) {
	if s.isConnected() != nil {
		return
	}

	for _, payload := range data {
		insert := "INSERT INTO audit_log (ClientId,OperationType,RecordTable,RecordId,RecordBody,UserCode,AppName) VALUES (?,?,?,?,?,?,?)"
		tx, err := s.conn.Begin()
		if err != nil {
			fmt.Printf("error beginning transaction %s\n", err.Error())
			return
		}
		stmt, err := tx.Prepare(insert)
		if err != nil {
			fmt.Printf("error preparing statement %s\n", err.Error())
			return
		}
		defer stmt.Close()

		_, err = stmt.Exec(clientId, payload.OperationType, payload.RecordTable, payload.RecordId, payload.RecordBody, payload.UserCode, payload.AppName)
		if err != nil {
			fmt.Printf("clickhouse error inserting audit %s\n", err)
		}
		if err := tx.Commit(); err != nil {
			fmt.Printf("error committing transaction %s\n", err.Error())
		}
	}
}

func (s *ClickHouseLogStorage) LogHttpPayload(serverId, clientId, endpoint, payload string) {
	if s.isConnected() != nil {
		return
	}

	insert := "INSERT INTO http_log (ServerId, ClientId, Endpoint, Payload) VALUES (?,?,?,?)"
	tx, err := s.conn.Begin()
	if err != nil {
		fmt.Printf("error beginning transaction %s\n", err.Error())
		return
	}
	stmt, err := tx.Prepare(insert)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(serverId, clientId, endpoint, payload)
	if err != nil {
		fmt.Printf("clickhouse error http log %s\n", err)
	}

	if err := tx.Commit(); err != nil {
		fmt.Printf("error committing transaction %s\n", err.Error())
	}
}

func (s *ClickHouseLogStorage) LogTaskStarted(serverId, clientId, taskId string) {
	if s.isConnected() != nil {
		return
	}

	insert := "INSERT INTO task_log(ServerId, ClientId, TaskId, Producer, Status, StartedAt) VALUES (?,?,?,?,'Pending',now())"
	stmt, err := s.conn.Prepare(insert)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(serverId, clientId, taskId, "Esynx-Server")
	if err != nil {
		fmt.Printf("clickhouse error task log %s\n", err)
		return
	}
}

func (s *ClickHouseLogStorage) LogTaskConsumed(taskId, consumer string) {
	if s.isConnected() != nil {
		return
	}

	update := "UPDATE task_log SET Status = 'Started', ConsumedAt = now(), Consumer = ? WHERE TaskId = ? AND Status = 'Pending'"

	stmt, err := s.conn.Prepare(update)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(consumer, taskId)
	if err != nil {
		fmt.Printf("clickhouse error task log %s\n", err)
		return
	}
}

func (s *ClickHouseLogStorage) LogTaskFinished(taskId, result, consumer string) {
	if s.isConnected() != nil {
		return
	}

	update := "UPDATE task_log SET Status = 'Finished', FinishedAt = now(), Consumer = ?, Result = ? WHERE TaskId = ? AND Status = 'Started'"

	stmt, err := s.conn.Prepare(update)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(consumer, result, taskId)
	if err != nil {
		fmt.Printf("clickhouse error task log %s\n", err)
		return
	}
}

func (s *ClickHouseLogStorage) LogServerLog(serverId, payload, level string) {
	if s.isConnected() != nil {
		return
	}

	insert := "INSERT INTO server_app_log (ServerId, Payload, Severity) VALUES (?,?,?)"
	stmt, err := s.conn.Prepare(insert)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(serverId, payload, level)
	if err != nil {
		fmt.Printf("clickhouse error server log %s\n", err)
		return
	}
}

func (s *ClickHouseLogStorage) LogFailedPayload(clientId, operationType, serviceName, dataType, recordBody, userCode, appName string) {
	if s.isConnected() != nil {
		return
	}

	insert := "INSERT INTO failed_payload (ClientId, OperationType, ServiceName, DataType, RecordBody, UserCode, AppName) VALUES (?,?,?,?,?,?,?)"
	stmt, err := s.conn.Prepare(insert)
	if err != nil {
		fmt.Printf("error preparing statement %s\n", err.Error())
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(clientId, operationType, serviceName, dataType, recordBody, userCode, appName)
	if err != nil {
		fmt.Printf("clickhouse error failed payload %s\n", err)
	}
}

func (s *ClickHouseLogStorage) CloseLogStorage() {
	if s.conn.Ping() == nil {
		_ = s.conn.Close()
	}
}
