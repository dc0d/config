package reloader

import (
	"os"
	"testing"
	"time"

	"github.com/dc0d/config/ymlconfig"
	"github.com/fsnotify/fsnotify"
	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	assert := assert.New(t)

	fn := "sample.yml"
	f, err := os.OpenFile(fn, os.O_CREATE|os.O_RDWR, 0600)
	assert.NoError(err)
	defer f.Close()
	_, err = f.WriteString(`
nats:
  httpport: 10044
  port: 10041
  username: "USR"
  password: "PSW006"
`)
	assert.NoError(err)

	done := make(chan struct{})
	rl := New(func(e error) { panic(e) })
	rl.SetOnEvent(func(ev fsnotify.Event) {
		defer close(done)
		assert.Contains(ev.Name, f.Name())
	})
	err = rl.Add(fn)
	assert.NoError(err)

	go func() {
		<-time.After(time.Millisecond * 150)
		f.WriteString(`
nats:
  httpport: 10044
  port: 10041
  username: "USR"
  password: "PSW006"

`)
	}()

	select {
	case <-done:
	case <-time.After(time.Second * 5):
	}
}

type conf struct {
	NATS struct {
		HTTPPort int
		Port     int
		Username string
		Password string
	}
}

var f = ymlconfig.New
