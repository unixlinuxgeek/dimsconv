package dimsconv

import "fmt"

type Inch float64
type Cm float64

func (i Inch) String() string {
	return fmt.Sprintf("%g In", i)
}

func (cm Cm) String() string {
	return fmt.Sprintf("%g Cm", cm)
}
