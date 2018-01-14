package ymlconfig

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	assert := assert.New(t)

	c := new(conf)
	assert.NoError(New().Load(c))
	assert.Equal(3000, c.NATS.HTTPPort)
	assert.Equal(3001, c.NATS.Port)
	assert.Equal("USR", c.NATS.Username)
	assert.Equal("PSW", c.NATS.Password)
}

type conf struct {
	NATS struct {
		HTTPPort int
		Port     int
		Username string
		Password string
	}
}
