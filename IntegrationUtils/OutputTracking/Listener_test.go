package OutputTracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListener(t *testing.T) {
	t.Run("multiple Trackers get output", func(t *testing.T) {
		var listener = CreateListener[bool]()

		var a = listener.GetTracker()
		var b = listener.GetTracker()

		listener.Track(true)
		listener.Track(false)

		var expected = []bool{true, false}
		assert.Equal(t, expected, a.GetOutput())
		assert.Equal(t, expected, b.GetOutput())
	})

	t.Run("does not add output on removed tracker", func(t *testing.T) {
		var listener = CreateListener[string]()

		var tracker = listener.GetTracker()

		listener.Remove(tracker)

		listener.Track("foo")

		assert.Equal(t, []string{}, tracker.GetOutput())
	})

	t.Run("Track adds output to OutputTracker", func(t *testing.T) {
		var listener = CreateListener[string]()

		var tracker = listener.GetTracker()

		listener.Track("foo")

		assert.Equal(t, []string{"foo"}, tracker.GetOutput())
	})

	t.Run("Track without trackers doesn't cause error", func(t *testing.T) {
		var listener = CreateListener[int]()

		listener.Track(0)
	})
}
