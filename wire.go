//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/ken1kasap/learning-wire/msg"
)

func InitializeEvent() msg.Event {
	wire.Build(msg.NewEvent, msg.NewGreeter, msg.NewMessage)
	return msg.Event{}
}
