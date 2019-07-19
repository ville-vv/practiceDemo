// Package other
package other

import "testing"

func Benchmark_TypeSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		TypeSwitch(a)
	}
}

func Benchmark_NormalSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		NormalSwitch(a)
	}
}

func Benchmark_InterfaceSwitch(b *testing.B) {
	var a = new(A)

	for i := 0; i < b.N; i++ {
		InterfaceSwitch(a)
	}
}

func Benchmark_TypeString(b *testing.B) {
	var a = "bajign"

	for i := 0; i < b.N; i++ {
		TypeString(a)
	}
}

func Benchmark_TypeNoString(b *testing.B) {
	var a = "bajign"

	for i := 0; i < b.N; i++ {
		switch a {
		case a:
			continue
		}
	}
}
