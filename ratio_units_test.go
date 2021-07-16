package units

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRatioConversion(t *testing.T) {
	assert.Equal(t, NewValue(1, Percentage).MustConvert(Permille).Float(), 10.0)
	assert.Equal(t, NewValue(1, Permille).MustConvert(Permyriad).Float(), 10.0)
	assert.Equal(t, NewValue(1, Percentage).MustConvert(Permyriad).Float(), 100.0)
}
