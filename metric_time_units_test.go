package units

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMetricTimeConversion(t *testing.T) {
	assert.Equal(t, 1.0, NewValue(1, MeterPerHour).MustConvert(MeterPerHour).Float())
	assert.Equal(t, .001, NewValue(1, MeterPerSecond).MustConvert(KiloMeterPerSecond).Float())
	assert.Equal(t, 3600000.0, NewValue(1, KiloMeterPerSecond).MustConvert(MeterPerHour).Float())
	assert.Equal(t, 3600.0, NewValue(1, MeterPerHour).MustConvert(MeterPerSecond).Float())
	assert.Equal(t, 1000.0, NewValue(1, KiloMeterPerSecond).MustConvert(MeterPerSecond).Float())
}
