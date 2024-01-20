package frontend

import (
	"github.com/polarsignals/frostdb/query"
	"github.com/zeebo/assert"
	"testing"
)

func TestFrontend_SqlToLogicalPlan(t *testing.T) {
	sql := "select a,b from t1"

	fe := NewFrontend()
	builder, err := fe.SqlToLogicalPlan(query.LocalQueryBuilder{}, sql)
	assert.Nil(t, err)

	assert.NotNil(t, builder)
}
