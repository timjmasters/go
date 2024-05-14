package OutputTracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutputTracker(t *testing.T) {
	t.Run("can track after Clear", func(t *testing.T) {
		var tracker = createTracker[string]()

		tracker.Add("foo")
		tracker.Clear()
		tracker.Add("bar")

		assert.Equal(t, []string{"bar"}, tracker.GetOutput())
	})

	t.Run("starts empty", func(t *testing.T) {
		var tracker = createTracker[string]()

		assert.Equal(t, []string{}, tracker.GetOutput())
	})

	t.Run("can get output after Add", func(t *testing.T) {
		var tracker = createTracker[string]()

		tracker.Add("foo")
		tracker.Add("bar")

		assert.Equal(t, []string{"foo", "bar"}, tracker.GetOutput())
	})

	t.Run("output is empty after Add then Clear", func(t *testing.T) {
		var tracker = createTracker[string]()

		tracker.Add("foo")
		tracker.Clear()

		assert.Equal(t, []string{}, tracker.GetOutput())
	})
}

func TestOutputTrackerIntegration(t *testing.T) {
	t.Run("Stop calls remove on listener", func(t *testing.T) {
		listener := &fakeListener[string]{
			removed: []Tracker[string]{},
		}

		tracker := createTracker[string](listener)

		tracker.Stop()

		assert.Equal(t, []Tracker[string]{tracker}, listener.removed)
	})
}

type fakeListener[T any] struct {
	removed []Tracker[T]
}

func (fake *fakeListener[T]) GetTracker() Tracker[T] {
	return createTracker[T]()
}

func (fake *fakeListener[T]) Track(output T) Listener[T] {
	return fake
}

func (fake *fakeListener[T]) Remove(tracker Tracker[T]) Listener[T] {
	fake.removed = append(fake.removed, tracker)
	return fake
}
