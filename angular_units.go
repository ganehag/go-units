package units

import  (
	"math"
)

var (
	Angle = UnitOptionQuantity("Angle")

	// metric
	Radian = NewUnit("radians", "rad", Angle, SI)
	Degree = NewUnit("degree", "°", Angle, SI)
	Turn = NewUnit("turn", "rot", Angle, SI, UnitOptionAliases("revolution", "rotation"))
)

func init() {
	NewRatioConversion(Radian, Degree, 180.0 / math.Pi)
	NewRatioConversion(Turn, Degree, 360.0)
}
