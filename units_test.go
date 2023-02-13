package smetana

import (
	"fmt"
	"testing"
)

type UnitsTestCase struct {
	value    fmt.Stringer
	expected string
}

func TestCanFormatIntUnits(t *testing.T) {
	tests := []UnitsTestCase{
		{PX(5), "5px"},
		{EM(5), "5em"},
		{REM(5), "5rem"},
		{CM(5), "5cm"},
		{MM(5), "5mm"},
		{IN(5), "5in"},
		{PT(5), "5pt"},
		{PC(5), "5pc"},
		{EX(5), "5ex"},
		{CH(5), "5ch"},
		{VW(5), "5vw"},
		{VH(5), "5vh"},
		{VMin(5), "5vmin"},
		{VMax(5), "5vmax"},
		{Perc(5), "5%"},
	}

	for _, testCase := range tests {
		formatted := testCase.value.String()
		assertEqual(t, testCase.expected, formatted)
	}
}

func TestCanFormatFloatUnits(t *testing.T) {
	value := EM(4.3)
	formatted := value.String()
	assertEqual(t, "4.30em", formatted)
}

func TestFloatUnitsAreTruncatedTo2DP(t *testing.T) {
	value := EM(3.14159)
	formatted := value.String()
	assertEqual(t, "3.14em", formatted)
}
