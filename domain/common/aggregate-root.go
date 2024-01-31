package common

type IAggregateRoot interface {
	AddEvent(event IBaseEvent)
	GetEvents() []IBaseEvent
	ClearEvents()
}
