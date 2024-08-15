package main

import (
	"fmt"
	"math"
)

type Generator interface {
	Next(n int) int
}

type AscendingNumberGenerator struct{}

func (_ AscendingNumberGenerator) Next(n int) int {
	return n + 1
}

type DescendingNumberGenerator struct{}

func (_ DescendingNumberGenerator) Next(n int) int {
	return n - 1
}

type Predicate interface {
	Fullfill(n int) bool
}

type EvenNumberPredicate struct{}

func (_ EvenNumberPredicate) Fullfill(n int) bool {
	return n%2 == 0
}

type OddNumberPredicate struct {}

func (_ OddNumberPredicate) Fullfill(n int) bool {
	return !(n%2 == 0)
}

type Secuence struct {
	current         int
	numberGenerator Generator
	numberPredicate Predicate
}

func (s *Secuence) Next() int {
	var next int
	for n := s.numberGenerator.Next(s.current); n <= math.MaxInt32; n = s.numberGenerator.Next(n) {
		if s.numberPredicate.Fullfill(n) {
			next = n
			break
		}
	}
	s.current = next
	return next
}

func main() {
	var gen Generator = DescendingNumberGenerator{}
	var pre Predicate = OddNumberPredicate{}

	secuence := Secuence{
		current: 0,
		numberGenerator: gen,
		numberPredicate: pre,
	}

	for i := 0; i< 20; i ++ {
		fmt.Println(secuence.Next())
	}
}

