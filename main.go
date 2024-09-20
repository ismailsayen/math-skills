package main

import (
	"bufio"
	"fmt"
	mathskills "mathskills/functions"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("The running command isn't correct")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("The command must be: go run . <file>.txt")
		return
	}

	var data []int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		nb, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("Error under adding %v maybe isn't a int number\n", scanner.Text())
			continue
		}
		data = append(data, nb)
	}
	fmt.Println("Average: ", mathskills.Average(data))
	fmt.Println("Median: ", mathskills.Median(data))
	fmt.Println("Variance: ", mathskills.Variance(data))
	fmt.Println("Standard Deviation: ", mathskills.StandardDeviation(data))
}
