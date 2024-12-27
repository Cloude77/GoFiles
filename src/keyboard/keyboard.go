package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

/*************  ✨ Codeium Command ⭐  *************/
// GetFloat reads a line of input from standard input, trims any whitespace,
// and attempts to parse it as a float64. It returns the parsed float and an
// error if reading from input or parsing fails.

/******  3ba27099-dee5-4702-a585-786d60fc3876  *******/
func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}
