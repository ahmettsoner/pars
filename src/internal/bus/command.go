package bus

type Command interface{}

type CommandHandler[T Command] interface {
	Handle(command T) error
}
