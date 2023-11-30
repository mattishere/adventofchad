package adventofchad

import (
	"os"
	"strings"
)

// Returns the content of the challenge file.
// the part property is used to determine which file to read:
// - part 1: part1.aoc
// - part 2: part2.aoc
func Input(part Part) (string, error) {
	inp, err := os.ReadFile(part.filename)

	return string(inp), err
}

// Returns the content of the challenge file, split by a specific delimiter.
// the part property is used to determine which file to read:
// - part 1: part1.aoc
// - part 2: part2.aoc
func InputByDelimiter(part Part, delimiter string) ([]string, error) {
	inp, err := Input(part)

	return strings.Split(inp, delimiter), err
}

// Returns the content of the challenge file, split by new line.
// the part property is used to determine which file to read:
// - part 1: part1.aoc
// - part 2: part2.aoc
func InputByLine(part Part) ([]string, error) {
	inp, err := InputByDelimiter(part, "\n")

	return inp, err
}

// Returns the content of the challenge file, split by space.
// the part property is used to determine which file to read:
// - part 1: part1.aoc
// - part 2: part2.aoc
func InputBySpace(part Part) ([]string, error) {
	inp, err := InputByDelimiter(part, " ")

	return inp, err
}
