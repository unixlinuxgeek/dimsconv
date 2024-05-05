package dimsconv

// InchToCm Умножаем значение длины на 2.54
func InchToCm(i Inch) Cm {
	return Cm(i * 2.54)
}

// CmToInch делим значение длины на 2.54
func CmToInch(cm Cm) Inch {
	return Inch(cm / 2.54)
}
