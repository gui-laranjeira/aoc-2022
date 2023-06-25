package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	var iterator int = 0
	elfos := make([][]int, 5)

	f, err := os.Open("calories.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "" {
			iterator++
			continue
		}

		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// ... handle error
			fmt.Println(err)
			panic(err)
		}

		elfos[iterator] = append(elfos[iterator], i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var elfo1 int
	var elfo2 int
	var elfo3 int
	var elfo4 int
	var elfo5 int

	for j := 0; j < len(elfos); j++ {
		var tempSoma int = 0

		for t := 0; t < len(elfos[j]); t++ {
			tempSoma += elfos[j][t]
		}
		switch j {
		case 0:
			elfo1 = tempSoma
		case 1:
			elfo2 = tempSoma
		case 2:
			elfo3 = tempSoma
		case 3:
			elfo4 = tempSoma
		case 4:
			elfo5 = tempSoma
		}
	}
	var caloriasTotais []int
	caloriasTotais = append(caloriasTotais, elfo1, elfo2, elfo3, elfo4, elfo5)
	sort.Slice(caloriasTotais, func(i, j int) bool {
		return caloriasTotais[i] < caloriasTotais[j]
	})
	caloriasMapeado := map[string]int{
		"Elf 1": elfo1,
		"Elf 2": elfo2,
		"Elf 3": elfo3,
		"Elf 4": elfo4,
		"Elf 5": elfo5,
	}
	max := caloriasTotais[len(caloriasTotais)-1]
	for nome, caloria := range caloriasMapeado {
		if caloria == max {
			println(nome, ": ", caloria, " calories")
		}
	}
}
