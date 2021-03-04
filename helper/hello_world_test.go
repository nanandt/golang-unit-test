package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Harun")

	if result != "Hello Harun" {
		// unit test failed
		panic("result is not related")
	}
}
func TestHelloWorldHasan(t *testing.T) {
	result := HelloWorld("Hasan")

	if result != "Hello Hasan" {
		// unit test failed
		panic("result is not related")
	}
}
