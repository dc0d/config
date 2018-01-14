package reloader

import (
	"fmt"
)

type sentinelErr string

func (v sentinelErr) Error() string { return string(v) }
func errorf(format string, a ...interface{}) error {
	return sentinelErr(fmt.Sprintf(format, a...))
}
