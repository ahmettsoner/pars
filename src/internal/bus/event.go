package bus

type Event interface{}

type EventHandler[T Event] interface {
	Handle(event T)
}
