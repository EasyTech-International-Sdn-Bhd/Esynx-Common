package rest

import (
	"bytes"
	"fmt"
	"github.com/goccy/go-json"
	"strings"
)

type SqlQueryExpr struct {
	Cond    [][]string `json:"cond"`
	GroupBy string     `json:"groupBy"`
	OrderBy string     `json:"orderBy"`
	Limit   int32      `json:"limit"`
}

type SqlQueryExprBuilder struct {
	sqlQueryExpr SqlQueryExpr
}

// NewSqlQueryExprBuilder will create a new SqlQueryExprBuilder
func NewSqlQueryExprBuilder() *SqlQueryExprBuilder {
	return &SqlQueryExprBuilder{}
}

func (b *SqlQueryExprBuilder) AddCondition(condition ...string) *SqlQueryExprBuilder {
	b.sqlQueryExpr.Cond = append(b.sqlQueryExpr.Cond, condition)
	return b
}

func (b *SqlQueryExprBuilder) SetGroupBy(groupBy string) *SqlQueryExprBuilder {
	b.sqlQueryExpr.GroupBy = groupBy
	return b
}

func (b *SqlQueryExprBuilder) SetOrderBy(orderBy string) *SqlQueryExprBuilder {
	b.sqlQueryExpr.OrderBy = orderBy
	return b
}

func (b *SqlQueryExprBuilder) SetLimit(limit int32) *SqlQueryExprBuilder {
	b.sqlQueryExpr.Limit = limit
	return b
}

func (b *SqlQueryExprBuilder) MakeInQuery(field string, items []string) string {
	return fmt.Sprintf("%s IN (%s)", field, strings.Join(items, ", "))
}

func (b *SqlQueryExprBuilder) OperatorQuery(field, operator, value string) string {
	return fmt.Sprintf("%s %s %s", field, operator, value)
}

func (b *SqlQueryExprBuilder) LikeQuery(field, value string) string {
	return fmt.Sprintf("%s LIKE %s", field, value)
}

func (b *SqlQueryExprBuilder) Build() SqlQueryExpr {
	return b.sqlQueryExpr
}

// ToString creates the SqlQueryExpr with supplied conditions
func (b *SqlQueryExprBuilder) ToString() (string, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	err := encoder.Encode(b.sqlQueryExpr)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
