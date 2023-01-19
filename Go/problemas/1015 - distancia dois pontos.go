package problemas

import (
	"fmt"
	"math"
)

func DistanciaEntreDoisPontos() {

	var x1, y1, x2, y2 float64

	fmt.Scanf("%f %f\n", &x1, &y1)
	fmt.Scanf("%f %f", &x2, &y2)

	distancia := math.Sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1))

	fmt.Printf("%.4f\n", distancia)

}
