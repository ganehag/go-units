package units

var (
	MetricTime = UnitOptionQuantity("metric/time")

	MeterPerSecond = NewUnit("meter/second", "m/s", MetricTime, SI, UnitOptionAliases("metre/second"))
	MeterPerHour = NewUnit("meter/hour", "m/h", MetricTime, SI, UnitOptionAliases("metre/hour"))
        KiloMeterPerSecond = NewUnit("kilometer/second", "km/s", MetricTime, SI, UnitOptionAliases("kilometre/second"))
        KiloMeterPerHour = NewUnit("kilometer/hour", "km/h", MetricTime, SI, UnitOptionAliases("kilometre/hour"))
)

func init() {
	NewRatioConversion(MeterPerSecond, MeterPerHour, 1 / 3600.0)
	NewRatioConversion(MeterPerSecond, KiloMeterPerSecond, 1 / 1000.0)
	NewRatioConversion(MeterPerSecond, KiloMeterPerHour, 3.6)
	NewRatioConversion(KiloMeterPerSecond, MeterPerHour, 3600000)
}
