package fitcha_test

import (
	"testing"

	"github.com/nndi-oss/fitcha"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	ctx := fitcha.NewContext("test", "test-org", map[string]any{"hello": "world"})

	assert.NotNil(t, ctx.Value(fitcha.UserCtx))
	assert.NotNil(t, ctx.Value(fitcha.OrgCtx))
	assert.NotNil(t, ctx.Value(fitcha.ExtraCtx))
}
