package event

type Event interface {
	EventDirChange | EventFileCreate
}

type Subject[T Event] interface {
	register(observer Observer[T])
	// deregister(observer Observer[T])
	notifyAll()
}

type Observer[T Event] interface {
	update(event T)
}
