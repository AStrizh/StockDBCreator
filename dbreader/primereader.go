package dbreader

import (
	"bufio"
	"fmt"
	"os"
)

func Stockreader(filename string) {
	file, err := os.Open(filename)
	fmt.Sprintln("Hello")
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		fmt.Println("Failed reading file content.")
		fmt.Println(err)
		file.Close()
		return
	}

	//prices := make([]float64, len(lines))

	for _, line := range lines {
		fmt.Sprintln("hello2")
		//floatPrice, err := strconv.ParseFloat(line, 64)

		// if err != nil {
		// 	fmt.Println("Converting price to float failed.")
		// 	fmt.Println(err)
		// 	file.Close()
		// 	return
		// }
		// prices[linesIndex] = floatPrice

		fmt.Println(line)
	}
}
