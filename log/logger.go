package log

import "github.com/easytech-international-sdn-bhd/esynx-common/entities"

type LogStorageConn struct {
	Username string
	Password string
	Host     string
	Port     int
	Database string
}

type IServerSideLogStorage interface {
	LogAudit(clientId string, data []*entities.AuditLog)
	LogHttpPayload(serverId, clientId, endpoint, payload string)
	LogTaskStarted(serverId, clientId, taskId string)
	LogTaskConsumed(taskId, consumer string)
	LogTaskFinished(taskId, result, consumer string)
	LogServerLog(serverId, payload, level string)
	LogFailedPayload(clientId, operationType, serviceName, dataType, recordBody, userCode, appName string)
	CloseLogStorage()
}
