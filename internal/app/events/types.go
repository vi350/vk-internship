package events

type EventProcessor interface {
	Fetcher
	Processor
}

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process(event Event) error
}

type EventType int

const (
	Unknown EventType = iota
	Message
	CallbackQuery
)

type Event struct {
	Type EventType
	Text string
	Meta interface{}
}
