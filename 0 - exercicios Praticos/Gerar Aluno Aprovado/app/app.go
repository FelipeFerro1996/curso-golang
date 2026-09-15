package app

func CalcularNota(n1, n2 float64) (media float64, aprovado bool) {

	media = (n1 + n2) / 2

	if media >= 6.0 {
		aprovado = true
	} else {
		aprovado = false
	}

	return

}
