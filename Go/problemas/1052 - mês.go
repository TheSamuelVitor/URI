package problemas

func Mes(numero int64) string {

	meses := []string{
		"January",
		"February",
		"March",
		"April",
		"May",
		"June",
		"July",
		"August",
		"September",
		"October",
		"November",
		"December",
	}

	return meses[numero-1]

}
