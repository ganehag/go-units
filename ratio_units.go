package units

var (
	Ratio = UnitOptionQuantity("ratio")

	Percentage = NewUnit("percentage", "%", Ratio)
	Permille = NewUnit("permille", "‰", Ratio)
	Permyriad = NewUnit("permyriad", "‱", Ratio)
)

func init() {
	NewRatioConversion(Percentage, Permille, 10)
	NewRatioConversion(Permille, Permyriad, 10)
}
