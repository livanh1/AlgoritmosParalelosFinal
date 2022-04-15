package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func ArregloOrdenado(a *[100]int) {

	//c := 207

	for i := 0; i < 100; i++ {
		a[i] = i + 1
		//c += 7

	}
}

func ArregloDesordenado2(n int) []int {

	a := make([]int, n, n)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(999) - rand.Intn(999)
	}
	return a
}

func MedirTiempo(start time.Time, algoritmo string) {

	elapsedTime := time.Since(start)
	fmt.Println("Algoritmo: " + algoritmo + " | Tiempo de ejecucion: " + elapsedTime.String() + "\n")
}

func QuickSort(a []int, izq int, der int) []int {

	pivote := a[izq]
	i := izq
	j := der
	var aux int

	for i < j {
		for a[i] <= pivote && i < j {
			i++
		}
		for a[j] > pivote {
			j--
		}
		if i < j {
			aux = a[i]
			a[i] = a[j]
			a[j] = aux
		}
	}

	a[izq] = a[j]
	a[j] = pivote

	if izq < j-1 {
		QuickSort(a, izq, j-1)
	}
	if j+1 < der {
		QuickSort(a, j+1, der)
	}

	return a
}

func main() {

	var a = []int{13, 25, 3, 1, 110, 102, 58, 91, 150, 33, 8, 69, 137, 54, 76}
	fmt.Println("\nARREGLO PARA ALGORITMOS DE ORDENAMIENTO: \n", a)

	var b [100]int
	ArregloOrdenado(&b)
	fmt.Println("\nARREGLO PARA ALGORITMOS DE BUSQUEDA: \n", b)

	runtime.GOMAXPROCS(4)

	go func() {

		var a [100]int
		var dato int
		defer MedirTiempo(time.Now(), "Busqueda Secuencial")

		ArregloOrdenado(&a)
		dato = 60 //Dato a buscar
		i := 0

		for (a[i] != dato) && (i < 100) {
			i++
			if a[i] == dato {
				fmt.Printf("\n\nEl dato %d, fue encontrado en la posicion: %d del arreglo!!\n", dato, i)

			}
		}
	}()

	go func() {

		var a [100]int
		ArregloOrdenado(&a)
		defer MedirTiempo(time.Now(), "Búsqueda Binaria")

		inicio := 0
		fin := len(a) - 1
		var mitad, dato int

		dato = 60 //Dato a buscar
		mitad = (inicio + fin) / 2
		for (inicio <= fin) && (a[mitad] != dato) {
			if dato < a[mitad] {
				fin = mitad - 1
			} else {
				inicio = mitad + 1

			}
			mitad = (inicio + fin) / 2
		}

		if dato == a[mitad] {
			fmt.Printf("\n\nEl dato %d, fue encontrado en la posicion: %d del arreglo!!\n", dato, mitad)

		}
	}()

	go func() {

		var a = []int{13, 25, 3, 1, 110, 102, 58, 91, 150, 33, 8, 69, 137, 54, 76}
		var n int = len(a)
		defer MedirTiempo(time.Now(), "QuickSort")

		b := QuickSort(a, 0, n-1)
		fmt.Println(b)
	}()

	go func() {

		var a = []int{13, 25, 3, 1, 110, 102, 58, 91, 150, 33, 8, 69, 137, 54, 76}
		var n int = len(a)
		defer MedirTiempo(time.Now(), "Ordenamiento de Burbuja")
		var e, i, auxiliar int

		for e = 0; e < n; e++ {
			for i = 0; i < n-1-e; i++ {
				if a[i] > a[i+1] {
					auxiliar = a[i+1]
					a[i+1] = a[i]
					a[i] = auxiliar
				}
			}
		}

		fmt.Println(a)
	}()

	go func() {

		var a = []int{13, 25, 3, 1, 110, 102, 58, 91, 150, 33, 8, 69, 137, 54, 76}
		var n int = len(a)
		defer MedirTiempo(time.Now(), "Insercion")
		var auxiliar int

		for i := 1; i < n; i++ {
			auxiliar = a[i]
			for j := i - 1; j >= 0 && a[j] > auxiliar; j-- {
				a[j+1] = a[j]
				a[j] = auxiliar
			}
		}
		fmt.Println(a)
	}()

	fmt.Scanln()
}
