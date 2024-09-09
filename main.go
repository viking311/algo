package main

import (
	"github.com/viking311/algo/checker"
	"github.com/viking311/algo/solver"
)

func main() {
	t := solver.NewSolver()
	checker := checker.NewChecker()
	checker.Test(t)
}
