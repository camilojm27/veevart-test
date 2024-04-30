/*
Camilo Mezú Prueba Veevart
30 Abril 2024, test 19
Se utilizará programación dinamica para resolver el ejercicio
Estado: Terminado
*/

package main

import "fmt"

func printM(m [][]int) {
	for j := range m {
		//print("fila", j)
		for i := 0; i < len(m[j]); i++ {
			print(m[j][i])
		}
		println()
	}
}

func mochila(l []Item, C int) /*([]Item, int)*/ {
	//Creacion de la matriz
	matrix := make([][]int, C+1)
	cols := len(l)

	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]int, cols)
	}

	//Llenar la matriz con valores por defecto

	for i := 0; i <= C; i++ {
		if i >= l[0].W {
			matrix[i][0] = l[0].W
		}
	}

	//Tomar la decicion de llevar el elemento o no
	//De acuerdo a su utilidad maxima

	for j := 1; j < cols; j++ {
		for i := 0; i <= C; i++ {
			matrix[i][j] = max( //Elegir el de mejor utilidad
				matrix[i][j-1], // No llevar el item
				func() int {
					if i-l[j].W >= 0 {
						//fmt.Println("i = %v, l[j].W = %v  result = %v ", i, l[j].W, i-l[j].W)
						return matrix[i-l[j].W][j-1] + l[j].V // Llevarlo
					}
					return 0
				}(),
			)
		}
	}

	//Se crea una matriz para almacenar los resusltados

	result := make([]bool, cols)
	i := C
	j := cols - 1

	for j > 0 {
		if matrix[i][j] == matrix[i][j-1] {
			result[j] = false
		} else {
			result[j] = true
			i = i - l[j].W

		}
		//fmt.Println(i)

		j = j - 1
	}

	if matrix[i][0] > 0 {
		result[0] = true
	} else {
		result[0] = false
	}
	totalV := 0
	totalW := 0
	//printM(matrix)
	fmt.Println("Elementos que se llevan")
	for i := 0; i < cols; i++ {
		if result[i] {
			fmt.Print(l[i])
			totalV += l[i].V
			totalW += l[i].W
		}
	}

	fmt.Println("\nElementos que no llevan")
	for i := 0; i < cols; i++ {
		if !result[i] {
			fmt.Print(l[i])
		}
	}

	fmt.Println("\nEl valor total es de : ", totalV)
	fmt.Println("El peso total es de : ", totalW)
}

type Item struct {
	W int //Peso del objeto
	V int //Valor de objeto
}

func main() {

	input := []Item{{W: 2, V: 3}, {W: 3, V: 4}, {W: 4, V: 5}, {W: 5, V: 6}}
	capacity := 10
	mochila(input, capacity)
}
