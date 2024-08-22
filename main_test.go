package main

import "testing"

func TestSayHello(t *testing.T) {
 name := "John Doe"
    expected := "Hello John Doe"

    actual := sayhello(name)

    if actual != expected {
        t.Errorf("Expected '%s', got '%s'", expected, actual)
    }
}
