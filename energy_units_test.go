package units

import (
        "github.com/stretchr/testify/assert"
        "testing"
)

func TestEnergyConversion(t *testing.T) {
	assert.Equal(t, NewValue(1, WattHour).MustConvert(Joule).Float(), 3600.0)
}
