package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	isPM := strings.HasSuffix(s, "PM")
	s = s[:len(s)-2]
	parts := strings.Split(s, ":")

	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	second, _ := strconv.Atoi(parts[2])

	if isPM && hour != 12 {
		hour += 12
	} else if !isPM && hour == 12 {
		hour = 0
	}

	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukkan waktu dalam format 12 jam (hh:mm:ssAM/PM): ")
	s, _ := reader.ReadString('\n')
	s = strings.TrimSpace(s)

	result := timeConversion(s)
	fmt.Println("Waktu dalam format 24 jam:", result)
}
