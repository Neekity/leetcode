package concurrency

import (
	"testing"
)

func testAdd(x, y int) int {
	return x + y
}
func TestAdd(t *testing.T) {
	if testAdd(1, 2) != 3 {
		t.FailNow()
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testAdd(1, 2)
	}
}

func repeatedDNA(DNA string) []string {
	var results []string
	dictionary := make(map[string]int, 0)
	lenDNA := len(DNA)
	if lenDNA <= 10 {
		return results
	}
	for i := 0; i < lenDNA-10; i++ {
		for j := i + 10; j <= lenDNA; j++ {
			dictionary[DNA[i:j]] += 1
		}

	}

	for key, value := range dictionary {
		if value > 1 {
			results = append(results, key)
		}
	}
	return results
}
