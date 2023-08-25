package application

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitEnv(t *testing.T) {
	env := initEnv()
	assert.Equal(t, env, "local")
}
