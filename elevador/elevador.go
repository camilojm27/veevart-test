/*
Camilo Mezú Prueba Veevart
30 Abril, test 9
Se utilizará programación dinamica para resolver el ejercicio
Estado: Sin Terminar
*/
package main

import (
	"fmt"
	"sort"
)

type Elevator struct {
	A int // Hace referencia al piso actual
	F int // Hace referencial al piso furturo
}

func solver(l []Elevator, I int) {
	var q = l

	sort.Slice(q, func(i, j int) bool {
		return q[i].A > q[j].A
	})

	up := false
	stop := false

	for i := 0; i < 29; i++ {

		if l[0].A > I {
			up = true

			fmt.Println("Elevador en piso %v", i)
		}

		if l[0].A == I {
			stop = true
		}
		println(up, stop)

	}

}
func main() {
	input := []Elevator{{5, 2}, {29, 10}, {13, 1}, {10, 1}}
	initial_floor := 4
	solver(input, initial_floor)
}
