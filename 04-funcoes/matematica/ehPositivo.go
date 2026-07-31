package matematica

func EhPositivo(num int64) (bool, uint64) {
	var positivo bool
	var absoluto uint64
	if num <= 0 {
		positivo = false
		absoluto = uint64(-num)
	} else {
		positivo = true
		absoluto = uint64(num)
	}
	return positivo, absoluto
}
