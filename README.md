# 🎅😎 Advent of Chad
The neat solution to importing your challenge input for Advent of Code with Go!

## 🎄🎁 Usage
So, hey! Glad you're checking out my project.

AoChad is fairly easy to use! It can read two files, `part1.aoc` and `part2.aoc`, in which you put your challenge input. These two files should sit in the root of your project (or the directory from which you'll be running the programs).

The general structure of an AoChad *compatible* program looks something like this (but not limited to):
- `main.go` -> you run both parts of a challenge here!
- `part1.go` & `part2.go` -> the code for the separate parts sits in these files (alternatively, you can just make two different functions in `main.go`, but to each their own).
- `part1.aoc` & `part2.aoc` -> the challenge input for each separate part should be in these two files - they're read by AoChad.
- `go.mod` -> well, of course go.mod is there.

Anyways, you obviously have to install AoChad before using it. You can do that by running `go get github.com/mattishere/adventofchad` inside your project (make sure to initialize it with `go mod init` first).

Then, it's as simple as importing it in a file. For example:

```go
package main

import (
	aoc "github.com/mattishere/adventofchad"
)

func part1() {
	// Input() returns the content of the challenge file (as type string).
	input, err := aoc.Input(aoc.Part1)

	// InputByDelimiter() returns the content of the challenge file, split by delimiter (as type []string).
	inputByDelimiter, err := aoc.InputByDelimiter(aoc.Part1, ",")

	// InputByLine() returns the content of the challenge file, split by new line (as type []string).
	inputByLine, err := aoc.InputByLine(aoc.Part1)

	// InputBySpace() returns the content of the challenge file, split by space (as type []string).
	inputBySpace, err := aoc.InputBySpace(aoc.Part1)


	if err != nil {
		panic(err)
	}
    ...
}
```

Rather basic, but it does what it's supposed to!

If you haven't noticed yet, this is not a groundbreaking library that will make your advent of code experience 1000% easier. Instead, it's just meant to make your solutions look better and fit nicer in [a carbon screenshot](https://carbon.now.sh) 😎.

## 👋👋 Conclusion
Merry Christmas and a fun advent of code! I wish you the best of luck, I'll see you on the leaderboards (I'm not getting on those myself of course)!