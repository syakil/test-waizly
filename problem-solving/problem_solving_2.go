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

	fmt.Print("Masukkan ukuran array (jumlah elemen): ")
	nStr, _ := reader.ReadString('\n')
	nStr = strings.TrimSpace(nStr)
	n, err := strconv.Atoi(nStr)
	if err != nil {
		fmt.Println("Error: Ukuran array harus berupa angka.")
		return
	}

	fmt.Print("Masukkan elemen-elemen array (dipisahkan spasi): ")
	arrStr, _ := reader.ReadString('\n')
	arrStr = strings.TrimSpace(arrStr)
	arrTokens := strings.Split(arrStr, " ")

	if len(arrTokens) != n {
		fmt.Println("Error: Jumlah elemen tidak sesuai dengan ukuran array.")
		return
	}

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i], err = strconv.Atoi(arrTokens[i])
		if err != nil {
			fmt.Println("Error: Elemen array harus berupa angka.")
			return
		}
	}

	positif, negatif, nol := 0, 0, 0
	for _, num := range arr {
		if num > 0 {
			positif++
		} else if num < 0 {
			negatif++
		} else {
			nol++
		}
	}

	fmt.Printf("Proporsi bilangan positif: %.6f\n", float64(positif)/float64(n))
	fmt.Printf("Proporsi bilangan negatif: %.6f\n", float64(negatif)/float64(n))
	fmt.Printf("Proporsi bilangan nol: %.6f\n", float64(nol)/float64(n))
}
