package frontend

import (
	"github.com/zeebo/assert"
	"testing"
)

func TestFrontend_SqlToLogicalPlan(t *testing.T) {
	sql := "select abs(1)"

	fe := NewFrontend()
	builder, err := fe.SqlToLogicalPlan(nil, sql)
	assert.Nil(t, err)

	assert.NotNil(t, builder)
}
