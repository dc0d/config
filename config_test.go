package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	assert := assert.New(t)

	p, err := RelativeResource("functions.go")
	assert.NoError(err)
	assert.Contains(p, "/src/github.com/dc0d/config/functions.go")
}

func Test02(t *testing.T) {
	assert := assert.New(t)

	p, err := RelativeResource("app.yml", "./ymlconfig")
	assert.NoError(err)
	assert.Contains(p, "/src/github.com/dc0d/config/ymlconfig/app.yml")
}

func Test03(t *testing.T) {
	assert := assert.New(t)

	assert.NoError(os.Chdir("./hclconfig"))

	p, err := RelativeResource("app.yml", "./../ymlconfig")
	assert.NoError(err)
	assert.Contains(p, "/src/github.com/dc0d/config/ymlconfig/app.yml")
}
