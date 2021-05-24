package units

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVolumeTimeConversion(t *testing.T) {
	assert.Equal(t, 1000.0, NewValue(1, M3PerSecond).MustConvert(LiterPerSecond).Float())
	assert.InDelta(t, 0.000277778, NewValue(1, LiterPerHour).MustConvert(LiterPerSecond).Float(), 0.00001)
	assert.InDelta(t, 0.001, NewValue(1, LiterPerHour).MustConvert(M3PerHour).Float(), 0.0001)
}
