package smetana

import "fmt"

func formatFloatUnits[T ~float32](value T, suffix string) string {
	if value == T(int(value)) {
		return fmt.Sprintf("%d%s", int(value), suffix)
	}
	return fmt.Sprintf("%.2f%s", value, suffix)
}

// Utility type for marking CSS values as pixels (adds a "px" suffix)
type PX float32

func (value PX) String() string {
	return formatFloatUnits(value, "px")
}

// Utility type for marking CSS values as em (adds an "em" suffix)
type EM float32

func (value EM) String() string {
	return formatFloatUnits(value, "em")
}

// Utility type for marking CSS values as rem (adds a "rem" suffix)
type REM float32

func (value REM) String() string {
	return formatFloatUnits(value, "rem")
}

// Utility type for marking CSS values as centimeters (adds a "cm" suffix)
type CM float32

func (value CM) String() string {
	return formatFloatUnits(value, "cm")
}

// Utility type for marking CSS values as millimeters (adds an "mm" suffix)
type MM float32

func (value MM) String() string {
	return formatFloatUnits(value, "mm")
}

// Utility type for marking CSS values as inches (adds an "in" suffix)
type IN float32

func (value IN) String() string {
	return formatFloatUnits(value, "in")
}

// Utility type for marking CSS values as points (adds a "pt" suffix)
type PT float32

func (value PT) String() string {
	return formatFloatUnits(value, "pt")
}

// Utility type for marking CSS values as pc (adds a "pc" suffix)
type PC float32

func (value PC) String() string {
	return formatFloatUnits(value, "pc")
}

// Utility type for marking CSS values as ex (adds an "ex" suffix)
type EX float32

func (value EX) String() string {
	return formatFloatUnits(value, "ex")
}

// Utility type for marking CSS values as ch (adds a "ch" suffix)
type CH float32

func (value CH) String() string {
	return formatFloatUnits(value, "ch")
}

// Utility type for marking CSS values as vw (adds a "vw" suffix)
type VW float32

func (value VW) String() string {
	return formatFloatUnits(value, "vw")
}

// Utility type for marking CSS values as vh (adds a "vh" suffix)
type VH float32

func (value VH) String() string {
	return formatFloatUnits(value, "vh")
}

// Utility type for marking CSS values as vmin (adds a "vmin" suffix)
type VMin float32

func (value VMin) String() string {
	return formatFloatUnits(value, "vmin")
}

// Utility type for marking CSS values as vmax (adds a "vmax" suffix)
type VMax float32

func (value VMax) String() string {
	return formatFloatUnits(value, "vmax")
}

// Utility type for marking CSS values as a percentage (adds a "%" suffix)
type Perc float32

func (value Perc) String() string {
	return formatFloatUnits(value, "%")
}
