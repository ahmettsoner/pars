package bus

import (
	internalBus "parsdevkit.net/internal/bus"
)

var b = internalBus.NewBus()

// func RegisterMessageHandler[T internalBus.Message](handler internalBus.MessageHandler[T]) {
// 	internalBus.RegisterMessageHandler(b, handler)
// }

// func SendMessage[T internalBus.Message](command T) error {
// 	return internalBus.SendMessage(b, command)
// }

func RegisterCommandHandler[T internalBus.Command](handler internalBus.CommandHandler[T]) {
	internalBus.RegisterCommandHandler(b, handler)
}

func SendCommand[T internalBus.Command](command T) error {
	return internalBus.SendCommand(b, command)
}

func RegisterEventHandler[T internalBus.Event](handler internalBus.EventHandler[T]) {
	internalBus.RegisterEventHandler(b, handler)
}
func PublishEvent[T internalBus.Event](event T) {
	internalBus.PublishEvent(b, event)
}
