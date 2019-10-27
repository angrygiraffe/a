package a

import (
	log "github.com/angrygiraffe/go-log"
)

type A string

func (a A) Out() string {
	log.Debug("test")
	return string(a)
}
