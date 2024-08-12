package checker

import (
	"bufio"
	"fmt"
	"os"

	"github.com/viking311/algo/solver"
)

type Checker struct{}

func (c *Checker) Test(s solver.SolverInterface) {
	for i := 0; i < 10; i++ {
		filePath := fmt.Sprintf("./tests/test.%d.in", i)
		input := c.readFile(filePath)
		filePath = fmt.Sprintf("./tests/test.%d.out", i)
		expected := c.readFile(filePath)

		result := s.Run(input)

		if c.testEq(result, expected) {
			fmt.Printf("Test %d: Ok\n", i)
		} else {
			fmt.Printf("Test %d: Faild\nExpected: %v\nActual: %v\n", i, expected, result)
		}

	}
}

func (c *Checker) readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return text
}

func (c *Checker) testEq(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func NewChecker() *Checker {
	return &Checker{}
}
