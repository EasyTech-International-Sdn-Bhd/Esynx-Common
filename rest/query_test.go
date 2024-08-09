package rest

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlQueryExprBuilder(t *testing.T) {
	type Case struct {
		conditions    []string
		groupBy       string
		orderBy       string
		limit         int32
		inField       string
		inItems       []string
		operatorField string
		operatorOp    string
		operatorValue string
		likeField     string
		likeValue     string
	}
	cases := []struct {
		name string
		data Case
	}{
		{
			name: "Test1",
			data: Case{
				conditions:    []string{"cond1", "cond2", "cond3"},
				groupBy:       "groupByField",
				orderBy:       "orderByField",
				limit:         10,
				inField:       "inField",
				inItems:       []string{"item1", "item2", "item3"},
				operatorField: "operatorField",
				operatorOp:    "IN",
				operatorValue: "operatorValue",
				likeField:     "likeField",
				likeValue:     "%pattern%",
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			builder := NewSqlQueryExprBuilder()
			for _, cond := range tt.data.conditions {
				builder.AddCondition(cond)
			}
			builder.SetGroupBy(tt.data.groupBy)
			builder.SetOrderBy(tt.data.orderBy)
			builder.SetLimit(tt.data.limit)
			assert.Equal(t, tt.data.groupBy, builder.sqlQueryExpr.GroupBy)
			assert.Equal(t, tt.data.orderBy, builder.sqlQueryExpr.OrderBy)
			assert.Equal(t, tt.data.limit, builder.sqlQueryExpr.Limit)
			assert.Equal(t, tt.data.inField+" IN ("+tt.data.inItems[0]+", "+tt.data.inItems[1]+", "+tt.data.inItems[2]+")",
				builder.MakeInQuery(tt.data.inField, tt.data.inItems))
			assert.Equal(t, tt.data.operatorField+" "+tt.data.operatorOp+" "+tt.data.operatorValue,
				builder.OperatorQuery(tt.data.operatorField, tt.data.operatorOp, tt.data.operatorValue))
			assert.Equal(t, tt.data.likeField+" LIKE "+tt.data.likeValue, builder.LikeQuery(tt.data.likeField, tt.data.likeValue))
			_, err := builder.ToString()
			assert.NoError(t, err)
		})
	}
}

func TestParams(t *testing.T) {
	builder := NewSqlQueryExprBuilder()
	query, err := builder.AddCondition("age > 30").
		AddCondition(builder.OperatorQuery("age", "<=", "20"), "bday BETWEEN NOW() AND (NOW()-INTERVAL 1 YEAR)", builder.OperatorQuery("name", "=", "JUL")).
		AddCondition(builder.MakeInQuery("age", []string{"A", "N"}), builder.LikeQuery("name", "%JUL")).
		ToString()

	if err != nil {
		t.Error(err)
		t.Fail()
		return
	}
	t.Log(query)
}
