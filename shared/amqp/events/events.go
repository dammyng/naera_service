package events

type Event interface {
	EventName() string
}

type EventListener interface {
	Listen(exchange string, eventNames ...string) (<-chan Event, <-chan error, error)
}