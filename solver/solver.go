package solver

import (
	"fmt"
	"strconv"
)

type Solver struct{}

func (s *Solver) Run(st []string) []string {
	n, _ := strconv.ParseUint(st[0], 0, 64)
	var arr []uint64 = []uint64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	for i := 1; uint64(i) < n; i++ {
		arr = s.newArr(arr)
	}

	var summ uint64

	for _, v := range arr {
		summ += v * v
	}

	var result []string
	result = append(result, fmt.Sprintf("%d", summ))
	return result

}

func (s *Solver) newArr(prevArr []uint64) []uint64 {
	var tmp []uint64
	newLen := len(prevArr) + 9
	for i := 0; i < newLen; i++ {
		q := uint64(0)
		for j := 0; j < 10; j++ {
			if (i-j) >= 0 && (i-j) < len(prevArr) {
				q += prevArr[i-j]
			}
		}
		tmp = append(tmp, q)
	}

	return tmp
}

func NewSolver() *Solver {
	return &Solver{}
}
