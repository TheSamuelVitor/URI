package problemas

import (
	"fmt"
)

func IdadeemDias() {
	var dias, meses, anos int

	fmt.Scanln(&dias)

	anos = dias / 365
	dias = dias - (anos * 365)
	meses = dias / 30
	dias = dias - (meses * 30)

	fmt.Printf("%d ano(s)\n%d mes(es)\n%d dia(s)\n", anos, meses, dias)
}