package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Input the number of families : ")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		fmt.Println("Input jumlah keluarga tidak valid")
		return
	}

	fmt.Print("Input the number of members in the family (separated by a space) : ")
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Fields(line)

	if len(parts) != n {
		fmt.Println("Input must be equal with count of family")
		return
	}

	families := make([]int, n)
	for i, p := range parts {
		val, err := strconv.Atoi(p)
		if err != nil || val <= 0 {
			fmt.Println("Jumlah anggota keluarga tidak valid")
			return
		}
		families[i] = val
	}

	busCount := minBusRequired(families)
	fmt.Println("Minimum bus required is :", busCount)
}

func minBusRequired(families []int) int {
	count := 0
	used := make([]bool, len(families))

	for i := 0; i < len(families); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		count++ // minimal 1 bus untuk keluarga ini

		// Cari pasangan keluarga yang bisa digabung
		for j := i + 1; j < len(families); j++ {
			if !used[j] && families[i]+families[j] <= 4 {
				used[j] = true
				break // maksimal 2 keluarga per bus
			}
		}
	}

	return count
}
