package smetana

import "fmt"

func formatFloatUnits[T ~float32](value T, suffix string) string {
	if value == T(int(value)) {
		return fmt.Sprintf("%d%s", int(value), suffix)
	}
	return fmt.Sprintf("%.2f%s", value, suffix)
}

type PX float32

func (value PX) String() string {
	return formatFloatUnits(value, "px")
}

type EM float32

func (value EM) String() string {
	return formatFloatUnits(value, "em")
}

type REM float32

func (value REM) String() string {
	return formatFloatUnits(value, "rem")
}

type CM float32

func (value CM) String() string {
	return formatFloatUnits(value, "cm")
}

type MM float32

func (value MM) String() string {
	return formatFloatUnits(value, "mm")
}

type IN float32

func (value IN) String() string {
	return formatFloatUnits(value, "in")
}

type PT float32

func (value PT) String() string {
	return formatFloatUnits(value, "pt")
}

type PC float32

func (value PC) String() string {
	return formatFloatUnits(value, "pc")
}

type EX float32

func (value EX) String() string {
	return formatFloatUnits(value, "ex")
}

type CH float32

func (value CH) String() string {
	return formatFloatUnits(value, "ch")
}

type VW float32

func (value VW) String() string {
	return formatFloatUnits(value, "vw")
}

type VH float32

func (value VH) String() string {
	return formatFloatUnits(value, "vh")
}

type VMin float32

func (value VMin) String() string {
	return formatFloatUnits(value, "vmin")
}

type VMax float32

func (value VMax) String() string {
	return formatFloatUnits(value, "vmax")
}

type Perc float32

func (value Perc) String() string {
	return formatFloatUnits(value, "%")
}
