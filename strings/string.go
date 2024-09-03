package main

func inverterString(s string) string {
	// Cria um slice de bytes para a string invertida
	runas := []rune(s)
	n := len(runas)

	// Inverte os caracteres no slice
	for i := 0; i < n/2; i++ {
		runas[i], runas[n-i-1] = runas[n-i-1], runas[i]
	}

	// Converte o slice de runes de volta para uma string
	return string(runas)
}
