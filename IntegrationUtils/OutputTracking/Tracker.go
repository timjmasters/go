package OutputTracking

type Tracker[T any] interface {
	GetOutput() []T
	Add(s T)
	Clear() Tracker[T]
	Stop() Tracker[T]
}

type tracker[T any] struct {
	output   []T
	listener Listener[T]
}

func createTracker[T any](listener ...Listener[T]) Tracker[T] {
	var l Listener[T]
	if len(listener) == 0 {
		l = CreateListener[T]()
	} else {
		l = listener[0]
	}
	return &tracker[T]{
		output:   []T{},
		listener: l,
	}
}

func (tracker *tracker[T]) GetOutput() []T {
	return tracker.output
}

func (tracker *tracker[T]) Add(o T) {
	tracker.output = append(tracker.output, o)
}

func (tracker *tracker[T]) Clear() Tracker[T] {
	tracker.output = []T{}
	return tracker
}

func (tracker *tracker[T]) Stop() Tracker[T] {
	tracker.listener.Remove(tracker)
	return tracker
}
