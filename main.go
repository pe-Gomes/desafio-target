package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
 * PERGUNTA 3
 *
 * Trata-se de estratégias para ordenar elementos de um vetor de números em ordem crescente
 * e decrescente, ignorando os iguais a zero, bem como tratando a média anual.
 *
 * Para tanto, pode-se utilizar a técnica de merge sort em conjunto com a soma
 * dos valores.
 */

func main() {
	// Criando um vetor/slice aleatório representando 365 dias de receita
	arr := make([]float64, 0, 365)

	for i := range 365 {
		if i%7 == 0 {
			arr = append(arr, 0)
			continue
		}

		arr = append(arr, rand.ExpFloat64()*10_000)
	}

	start := time.Now()

	sortedArr, sum := mergeSortAndSumWithoutZeros(arr)

	// fmt.Println("Vetor filtrado:", sortedArr)
	fmt.Println("Maior receita:", sortedArr[len(sortedArr)-1])
	fmt.Println("Menor receita:", sortedArr[0])

	annualAvg := sum / float64(len(sortedArr))

	fmt.Println("\nMédia anual:", annualAvg)

	var daysOverAvg int

	for _, v := range sortedArr {
		if v > annualAvg {
			daysOverAvg++
		}
	}

	fmt.Println("Dias em que o valor foi superior à média anual:", daysOverAvg)
	fmt.Println("\nTempo de execução:", time.Since(start))
}

type MergeSortNumbers interface {
	~float64 | ~float32 | ~int64 | ~int32 | ~int
}

func merge[T MergeSortNumbers](left, right []T) ([]T, T) {
	result := make([]T, 0, len(left)+len(right))

	var sum T

	for len(left) > 0 && len(right) > 0 {
		if left[0] == 0 {
			left = left[1:] // Pulando valores zero no slice da esquerda
			continue
		}
		if right[0] == 0 {
			right = right[1:] // Pular valores zero no slice da direita
			continue
		}

		if left[0] <= right[0] {
			result = append(result, left[0])
			sum += left[0]
			left = left[1:]
			continue

		}

		result = append(result, right[0])
		sum += right[0]
		right = right[1:]
	}

	for len(left) > 0 {
		if left[0] != 0 {
			result = append(result, left[0])
			sum += left[0]
		}
		left = left[1:]
	}

	for len(right) > 0 {
		if right[0] != 0 {
			result = append(result, right[0])
			sum += right[0]
		}
		right = right[1:]
	}

	return result, sum
}

func mergeSortAndSumWithoutZeros[T MergeSortNumbers](arr []T) ([]T, T) {
	if len(arr) <= 1 {
		if len(arr) == 1 && arr[0] != 0 {
			return arr, arr[0]
		}

		return arr, 0
	}

	mid := len(arr) / 2
	left, _ := mergeSortAndSumWithoutZeros(arr[:mid])
	right, _ := mergeSortAndSumWithoutZeros(arr[mid:])

	merged, sum := merge(left, right)

	return merged, sum
}
