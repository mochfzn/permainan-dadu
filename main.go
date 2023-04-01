package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	var (
		n          int     // jumlah pemain
		m          int     // jumlah dadu
		score      []int   // poin bagi setiap pemain
		dice       [][]int // dadu yang dimiliki pemain
		diceNumber string
	)

	n = 3
	m = 4
	score = make([]int, n)
	dice = make([][]int, n)

	for i := 0; i < n; i++ {
		score[i] = 0
	}

	for j := range dice {
		dice[j] = make([]int, m, m*n)
	}

	for k := 0; k < n; k++ {
		for l := 0; l < m; l++ {
			dice[k][l] = 0
		}
	}

	l := 1

	for {
		fmt.Printf("Pemain = %d, Dadu = %d\n", n, m)
		fmt.Printf("=============\n")
		fmt.Printf("Giliran %d lempar dadu:\n", l)

		acakDadu(dice)

		for i := 0; i < n; i++ {
			fmt.Printf("\tPemain #%d (%d): ", i+1, score[i])

			for j, number := range dice[i] {
				if j != 0 {
					diceNumber += ","
				}
				diceNumber += strconv.Itoa(number)
			}

			fmt.Printf("%s\n", diceNumber)

			diceNumber = ""
		}

		fmt.Println("Setelah evaluasi:")

		evaluasi(dice, score)

		for j := 0; j < n; j++ {
			fmt.Printf("\tPemain #%d (%d): ", j+1, score[j])

			for k, number := range dice[j] {
				if k != 0 {
					diceNumber += ","
				}
				diceNumber += strconv.Itoa(number)
			}

			fmt.Printf("%s\n", diceNumber)

			diceNumber = ""
		}

		if berhentiPermainan(dice) {
			break
		}

		l++
	}

	akhirGame(dice, score)
}

func akhirGame(dice [][]int, score []int) {
	var sisaPemain int
	var pemenang int
	var max int

	for i := range dice {
		if len(dice) != 0 {
			sisaPemain = i
		}
	}

	max = score[0]
	for j, value := range score {
		if value > max {
			max = value
			pemenang = j
		}
	}

	fmt.Printf("Game berakhir karena hanya pemain #%d yang memiliki dadu\n", sisaPemain)
	fmt.Printf("Game dimenangkan oleh pemain #%d karena memiliki poin lebih banyak dari pemain lainnya.", pemenang+1)
}

func berhentiPermainan(dice [][]int) bool {
	total := len(dice)
	jumlah := 0

	for i := range dice {
		if len(dice[i]) == 0 {
			jumlah++
		}

		if jumlah == total-1 {
			return true
		}
	}

	return false
}

func acakDadu(dice [][]int) {
	for i := range dice {
		for j := range dice[i] {
			dice[i][j] = rand.Intn(7-1) + 1
		}
	}
}

func evaluasi(dice [][]int, score []int) {
	var jumlah int
	var posisi int
	var poin int

	addOne := make(map[int]int)

	for i := range dice {
		if len(dice[i]) == 0 {
			continue
		}

		dice[i], poin = kurangiAngkaEnam(dice[i])
		score[i] = score[i] + poin

		dice[i], jumlah = kurangiAngkaSatu(dice[i])

		posisi = i + 1
		if posisi == len(dice) {
			posisi = 0
		}

		addOne[posisi] = jumlah
	}

	for key, value := range addOne {
		dice[key] = tambah(dice[key], value)
	}
}

func kurangiAngkaSatu(dice []int) ([]int, int) {
	total := len(dice)
	index := 0
	jumlah := 0

	if total == 0 {
		return dice, jumlah
	}

	for {
		if dice[index] != 1 {
			index++
		} else if dice[index] == 1 {
			dice = append(dice[:index], dice[index+1:]...)
			total--
			jumlah++
		}

		if total == index {
			break
		}
	}

	return dice, jumlah
}

func kurangiAngkaEnam(dice []int) ([]int, int) {
	index := 0
	total := len(dice)
	jumlah := 0

	if total == 0 {
		return dice, jumlah
	}

	for {
		if dice[index] == 6 {
			dice = append(dice[:index], dice[index+1:]...)
			total--
			jumlah++
		} else {
			index++
		}

		if total == index {
			break
		}
	}

	return dice, jumlah
}

func tambah(dice []int, jumlah int) []int {
	for i := 0; i < jumlah; i++ {
		dice = append(dice, 1)
	}

	return dice
}
