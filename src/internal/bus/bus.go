package bus

import (
	"fmt"
	"reflect"
	"sync"
)

type MessageBus struct {
	commandHandlers map[reflect.Type]any
	eventHandlers   map[reflect.Type][]any
	lock            sync.RWMutex
}

func NewBus() *MessageBus {
	return &MessageBus{
		commandHandlers: make(map[reflect.Type]any),
		eventHandlers:   make(map[reflect.Type][]any),
	}
}

// === Command ===

func RegisterCommandHandler[T Command](b *MessageBus, handler CommandHandler[T]) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	b.lock.Lock()
	defer b.lock.Unlock()

	b.commandHandlers[t] = handler
}

func SendCommand[T Command](b *MessageBus, command T) error {
	t := reflect.TypeOf(command)
	b.lock.RLock()
	handlerRaw, ok := b.commandHandlers[t]
	b.lock.RUnlock()

	if !ok {
		return fmt.Errorf("no command handler for %v", t)
	}

	handler := handlerRaw.(CommandHandler[T])
	return handler.Handle(command)
}

// === Event ===

func RegisterEventHandler[T Event](b *MessageBus, handler EventHandler[T]) {
	t := reflect.TypeOf((*T)(nil)).Elem()
	b.lock.Lock()
	defer b.lock.Unlock()

	b.eventHandlers[t] = append(b.eventHandlers[t], handler)
}

func PublishEvent[T Event](b *MessageBus, event T) {
	t := reflect.TypeOf(event)
	b.lock.RLock()
	handlers := b.eventHandlers[t]
	b.lock.RUnlock()

	for _, h := range handlers {
		go h.(EventHandler[T]).Handle(event)
	}
}
