package handler

import (
	"fmt"
	"github.com/4ER0/go-ted/estatestics/model"
	"testing"
)

func BenchmarkMain(b *testing.B) {
	runTest(b, "Calc", 10, doCalcBench)
	runTest(b, "Calc", 1000, doCalcBench)
	runTest(b, "Calc", 1000000, doCalcBench)
	runTest(b, "CalcNoParallel", 10, doCalcBenchNoParallel)
	runTest(b, "CalcNoParallel", 1000, doCalcBenchNoParallel)
	runTest(b, "CalcNoParallel", 1000000, doCalcBenchNoParallel)
}

func runTest(b *testing.B, s string, c int, testFunc func(b *testing.B, es []model.RealEstate)) {
	es := makeEstates(c)
	b.Run(fmt.Sprintf("%s%d", s, c), func(b *testing.B) {
		testFunc(b, es)
	})
}

func makeEstates(c int) []model.RealEstate {
	estates := make([]model.RealEstate, c)
	for i := 0; i < c; i++ {
		estates[i] = model.RealEstate{
			Key:         1000,
			City:        "City",
			Total:       4000,
			OwnedLiving: 1000,
			Rented:      1000,
			SecondHome:  1000,
			Empty:       1000,
		}
	}
	return estates
}

func doCalcBench(b *testing.B, es []model.RealEstate) {
	for i := 0; i < b.N; i++ {
		doCalc(es, "")
	}
}

func doCalcBenchNoParallel(b *testing.B, es []model.RealEstate) {
	for i := 0; i < b.N; i++ {
		doCalcNoParallel(es, "")
	}
}
