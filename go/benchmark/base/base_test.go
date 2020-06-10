package base

//执行所有测试 go test -bench=.
import (
	"fmt"
	"testing"
)

// 执行和Calculate相关的时间 go test -run=Calculate -bench=.
func TestCalculate(t *testing.T) {
	fmt.Println("Test Calculate")
	expected := 4
	result := Calculate(2)
	if expected != result {
		t.Error("Failed")
	}
}

//单独执行某个性能测试时 go test -run=BenchmarkC -bench=.
func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
	}
}

func benchmarkCalculate(input int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Calculate(input)
	}
}

func BenchmarkCalculate100(b *testing.B)         { benchmarkCalculate(100, b) }
func BenchmarkCalculateNegative100(b *testing.B) { benchmarkCalculate(-100, b) }
func BenchmarkCalculateNegative1(b *testing.B)   { benchmarkCalculate(-1, b) }

func TestOther(t *testing.T) {
	fmt.Println("Testing something else")
	fmt.Println("This shouldn't run with -run=calc")
}
