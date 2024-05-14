package OutputTracking

import "github.com/timjmasters/go/Utils/Slice"

type Listener[T any] interface {
	GetTracker() Tracker[T]
	Track(output T) Listener[T]
	Remove(tracker Tracker[T]) Listener[T]
}

type listener[T any] struct {
	trackers []Tracker[T]
}

func CreateListener[T any]() Listener[T] {
	return &listener[T]{}
}

func (listener *listener[T]) GetTracker() Tracker[T] {
	var tracker = createTracker[T]()
	listener.trackers = append(listener.trackers, tracker)
	return tracker
}

func (listener *listener[T]) Track(output T) Listener[T] {
	for _, t := range listener.trackers {
		t.Add(output)
	}
	return listener
}

func (listener *listener[T]) Remove(tracker Tracker[T]) Listener[T] {
	listener.trackers = Slice.RemoveFromSlice(tracker, listener.trackers)
	return listener
}
