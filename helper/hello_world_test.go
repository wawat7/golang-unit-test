package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

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
