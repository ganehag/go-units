package units

var (
	VolumeTime = UnitOptionQuantity("volume/time")

	LiterPerSecond = NewUnit("liter/second", "l/s", VolumeTime, SI, UnitOptionAliases("litre/second"))
	M3PerSecond = NewUnit("cubicmeter/second", "m3/s", VolumeTime, SI, UnitOptionAliases("cubicmetre/second"))
	LiterPerHour = NewUnit("liter/hour", "l/h", VolumeTime, SI, UnitOptionAliases("litre/hour"))
	M3PerHour = NewUnit("cubicmeter/hour", "m3/h", VolumeTime, SI, UnitOptionAliases("cubicmetre/hour"))
)

func init() {
	// There are #Z Y:s in an X
	NewRatioConversion(M3PerSecond, LiterPerSecond, 1000)
	NewRatioConversion(LiterPerHour, LiterPerSecond, 1 / 3600.0)
	NewRatioConversion(M3PerHour, LiterPerSecond, 0.277778)
}
