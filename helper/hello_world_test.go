package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Benchmark
func BenchmarkHelloWorldBayu(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Bayu")
	}
}

func BenchmarkHelloWorldSedana(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Sedana")
	}
}

// Sub Benchmark
func BenchmarkSub(b *testing.B) {
	b.Run("Ninja", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ninja")
		}
	})
	b.Run("Naruto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Naruto")
		}
	})
}

// Table Benchmark
func BenchmarkTable(b *testing.B) {
	benchmark := []struct {
		name    string
		request string
	}{
		{
			name:    "Eta",
			request: "Eta",
		},
		{
			name:    "Edo",
			request: "Edo",
		},
		{
			name:    "Eduardo",
			request: "Ronaldo",
		},
		{
			name:    "Messi",
			request: "Lionel",
		},
	}
	for _, benchmarks := range benchmark {
		b.Run(benchmarks.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmarks.request)
			}
		})
	}
}

// Testing M test
func TestMain(m *testing.M) {
	//before
	fmt.Println("Before unit test")

	m.Run()
	//after
	fmt.Println("After unit test")
}

// Manual unit test

func TestHelloBima(t *testing.T) {
	result := HelloWorld("Bima")

	if result != "HelloBima" {
		// if unit test failed
		t.Fail()
	}
	fmt.Println("Test for Bima is done")
}

func TestHelloBanu(t *testing.T) {
	result := HelloWorld("Banu")

	if result != "HelloBanu" {
		// if unit test failed
		t.FailNow()
	}
	fmt.Println("Test for Banu is done")
}
func TestHelloBayu(t *testing.T) {
	result := HelloWorld("Bayu")

	if result != "HelloBayu" {
		// if unit test failed
		t.Error("Supposed to be HelloBayu")
	}
	fmt.Println("Test for Bayu is done")
}

func TestHelloTania(t *testing.T) {
	result := HelloWorld("Tania")

	if result != "HelloTania" {
		// if unit test failed
		t.Fatal("Supposed to be HelloTania")
	}
	fmt.Println("Test for Tania is done")
}

// Use librabry testify don't manual

func TestHelloAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "HelloAssert", result, "Result must be HelloAssert")
	fmt.Println("Test using Assert Done") // Fail()
}

func TestHelloRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "HelloRequire", result, "Result must be HelloRequire")
	fmt.Println("Test using Require Done") // FailNow()
}

// Skip Test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot be run on Mac OS")
	}
	result := HelloWorld("Gale")
	require.Equal(t, "HelloGale", result, "Result must be HelloGale")
}

// test SubTest
func TestSubTest(t *testing.T) {
	t.Run("Bayu", func(t *testing.T) {
		result := HelloWorld("Bayu")
		require.Equal(t, "HelloBayu", result, "Result must be HelloBayu")
	})
	t.Run("Gale", func(t *testing.T) {
		result := HelloWorld("Gale")
		require.Equal(t, "HelloGale", result, "Result must be HelloGale")
	})
}

// Table Test
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Kyema",
			request:  "Kyema",
			expected: "HelloKyema",
		},
		{
			name:     "Ezyar",
			request:  "Eyzar",
			expected: "HelloEyzar",
		},
		{
			name:     "Das",
			request:  "Das",
			expected: "HelloDas",
		},
		{
			name:     "Roti",
			request:  "Roti",
			expected: "HelloRoti",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}
