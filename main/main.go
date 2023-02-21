package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	vector := generarLista(1000)
	Menu(vector)
}

func generarLista(n int) []int {
	vector := make([]int, n)
	auxiliar := make(map[int]bool)
	for i := 0; i < n; i++ {
		var num int
		for {
			num = rand.Intn(n*10) + 1 // generar un número aleatorio entre 1 y 10000
			//fmt.Println(auxiliar[num])
			if !auxiliar[num] {
				auxiliar[num] = true
				break
			}
		}
		vector[i] = num
	}
	sort.Ints(vector)
	return vector
}

func Menu(vector []int) {
	var opcion int
	for {
		fmt.Println("Menú")
		fmt.Println("1. Mostrar lista")
		fmt.Println("2. Buscar elemento")
		fmt.Println("3. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scan(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Los elementos de la lista son: ")
			fmt.Println(vector)
			fmt.Println(len(vector))
		case 2:
			var buscar int
			fmt.Print("Ingrese el elemento que desea buscar: ")
			fmt.Scan(&buscar)
			index := busquedaBinaria(vector, buscar)
			if index == -1 {
				fmt.Println("El elemento no se encuentra en la lista :(")
			} else {
				fmt.Println("El elemento se encuentra :)")
			}
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
	}
}

func busquedaBinaria(vector []int, busqueda int) int {
	inicio := 0
	fin := len(vector) - 1
	for inicio <= fin {
		puntomedio := (inicio + fin) / 2
		if vector[puntomedio] == busqueda {
			return puntomedio
		} else if vector[puntomedio] < busqueda {
			inicio = puntomedio + 1
		} else {
			fin = puntomedio - 1
		}
	}
	return -1
}
