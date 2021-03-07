package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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
