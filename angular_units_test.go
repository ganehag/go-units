package units

import (
	"math"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAngularConversion(t *testing.T) {
	assert.Equal(t, NewValue(1, Radian).MustConvert(Degree).Float(), 180.0 / math.Pi)
	assert.Equal(t, NewValue(1, Turn).MustConvert(Degree).Float(), 360.0)
	assert.Equal(t, NewValue(1, Turn).MustConvert(Radian).Float(), math.Pi * 2)
}