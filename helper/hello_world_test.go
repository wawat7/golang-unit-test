package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("before unit test")

	m.Run()

	//after
	fmt.Println("after unit test")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Wawat")
	if result != "Hello Wawat" {
		t.Fatal("Harusnya Hello Wawat")
	}

	fmt.Println("world beres")
}

func TestHelloWorldGanteng(t *testing.T) {
	result := HelloWorld("Ganteng")
	if result != "Hello Ganteng" {
		t.Error("Harusnya Hello Ganteng")
	}

	fmt.Println("ganteng beres")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Wawat")
	assert.Equal(t, "Hello Wawat", result, "Result must be 'Hello Wawat'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ganteng")
	require.Equal(t, "Hello Ganteng", result, "Result must be 'Hello Ganteng'")
}

func TestSubTest(t *testing.T) {
	t.Run("Wawat", func(t *testing.T) {
		result := HelloWorld("Wawat")
		assert.Equal(t, "Hello Wawat", result, "Result must be 'Hello Wawat'")
	})

	t.Run("Ganteng", func(t *testing.T) {
		result := HelloWorld("Ganteng")
		assert.Equal(t, "Hello Ganteng", result, "Result must be 'Hello Ganteng'")
	})
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Wawat",
			request:  "Wawat",
			expected: "Hello Wawat",
		},
		{
			name:     "Ganteng",
			request:  "Ganteng",
			expected: "Hello Ganteng",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, test.expected, result)
		})
	}
}
