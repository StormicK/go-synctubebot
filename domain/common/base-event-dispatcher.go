package common

type IBaseEventDispatcher interface {
	Dispatch(events []IBaseEvent)
}
