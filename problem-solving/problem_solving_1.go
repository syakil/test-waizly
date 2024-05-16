package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func miniMaxSum(arr []int32) {
	if len(arr) < 5 {
		fmt.Println("Error: Array should have at least five elements.")
		return
	}

	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	var minSum, maxSum int64
	var n = len(arr) - 1
	for i := 0; i < n; i++ {
		minSum += int64(arr[i])
		maxSum += int64(arr[len(arr)-n+i])
	}

	fmt.Println(minSum, maxSum)

}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan elemen-elemen array (dipisahkan dengan spasi): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	strValues := strings.Split(input, " ")
	var arr []int32
	for _, strVal := range strValues {
		num, err := strconv.ParseInt(strVal, 10, 32)
		if err != nil {
			fmt.Println("Error: Elemen array harus berupa angka.")
			return
		}
		arr = append(arr, int32(num))
	}

	miniMaxSum(arr)
}
