package units

var (
	Power = UnitOptionQuantity("power")

	Watt = NewUnit("watt", "W", Power, SI)
	ExaWatt   = Exa(Watt)
	PetaWatt  = Peta(Watt)
	TeraWatt  = Tera(Watt)
	GigaWatt  = Giga(Watt)
	MegaWatt  = Mega(Watt)
	KiloWatt  = Kilo(Watt)
	HectoWatt = Hecto(Watt)
	DecaWatt  = Deca(Watt)
	DeciWatt  = Deci(Watt)
	CentiWatt = Centi(Watt)
	MilliWatt = Milli(Watt)
	MicroWatt = Micro(Watt)
	NanoWatt  = Nano(Watt)
	PicoWatt  = Pico(Watt)
	FemtoWatt = Femto(Watt)
	AttoWatt  = Atto(Watt)
)

func init() {
}
