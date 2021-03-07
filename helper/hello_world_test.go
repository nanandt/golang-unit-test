package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "Rizky",
			request:  "Rizky",
			expected: "Hello Rizky",
		},
		{
			name:     "Bani",
			request:  "Bani",
			expected: "Hello Bani",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			require.Equal(t, test.expected, result)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Jaka", func(t *testing.T) {
		result := HelloWorld("Jaka")
		assert.Equal(t, "Hello Jaka", result, "Result must be 'Hello Jaka'")
	})
	t.Run("Tingkir", func(t *testing.T) {
		result := HelloWorld("Tingkir")
		assert.Equal(t, "Hello Tingkir", result, "Result must be 'Hello Tingkir'")
	})
}

func TestMain(m *testing.M) {
	// before
	fmt.Println("BEFORE UNIT TEST")
	m.Run() // memanggil semua unit test di sebuat package
	fmt.Println("AFTER UNIT TEST")

}
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Rizky")
	assert.Equal(t, "Hello Rizky", result, "Result must be 'Hello Rizky'")
	fmt.Println("TestHelloWorldAssert Done!")
}
func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Ucup")
	require.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
	fmt.Println("TestHelloWorldRequire Done!")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on Windows")
	}
	result := HelloWorld("Ucup")
	require.Equal(t, "Hello Ucup", result, "Result must be 'Hello Ucup'")
	fmt.Println("TestHelloWorldRequire Done!")
}

func TestHelloWorldHarun(t *testing.T) {
	result := HelloWorld("Harun")

	if result != "Hello Harun" {
		// unit test failed
		t.Error("Result must be 'Hello Harun'")
	}

	fmt.Println("TestHelloWorldHarun Done!")
}
func TestHelloWorldHasan(t *testing.T) {
	result := HelloWorld("Hasan")

	if result != "Hello Hasan" {
		// unit test failed
		t.Fatal("Result must be 'Hello Hasan'")
	}

	fmt.Println("TestHelloWorldHasan Done!")
}
func TestHelloWorldSirun(t *testing.T) {
	result := HelloWorld("Hasan")

	if result != "Hello Hasan" {
		// unit test failed
		t.FailNow()
	}

	fmt.Println("TestHelloWorldHasan Done!")
}
