package units

var (
	Energy = UnitOptionQuantity("energy")

	// metric
	Joule = NewUnit("joule", "J", Energy, SI)
	ExaJoule   = Exa(Joule)
	PetaJoule  = Peta(Joule)
	TeraJoule  = Tera(Joule)
	GigaJoule  = Giga(Joule)
	MegaJoule  = Mega(Joule)
	KiloJoule  = Kilo(Joule)
	HectoJoule = Hecto(Joule)
	DecaJoule  = Deca(Joule)
	DeciJoule  = Deci(Joule)
	CentiJoule = Centi(Joule)
	MilliJoule = Milli(Joule)
	MicroJoule = Micro(Joule)
	NanoJoule  = Nano(Joule)
	PicoJoule  = Pico(Joule)
	FemtoJoule = Femto(Joule)
	AttoJoule  = Atto(Joule)

	WattHour      = NewUnit("watthour", "Wh", Energy, SI)
	ExaWattHour   = Exa(WattHour)
	PetaWattHour  = Peta(WattHour)
	TeraWattHour  = Tera(WattHour)
	GigaWattHour  = Giga(WattHour)
	MegaWattHour  = Mega(WattHour)
	KiloWattHour  = Kilo(WattHour)
	HectoWattHour = Hecto(WattHour)
	DecaWattHour  = Deca(WattHour)
	DeciWattHour  = Deci(WattHour)
	CentiWattHour = Centi(WattHour)
	MilliWattHour = Milli(WattHour)
	MicroWattHour = Micro(WattHour)
	NanoWattHour  = Nano(WattHour)
	PicoWattHour  = Pico(WattHour)
	FemtoWattHour = Femto(WattHour)
	AttoWattHour  = Atto(WattHour)
)

func init() {
	NewRatioConversion(WattHour, Joule, 3600.0)
}
