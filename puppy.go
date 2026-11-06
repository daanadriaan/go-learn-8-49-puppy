package puppy

import (
    "fmt"
    "strconv"

    "github.com/advdv/stdgo/stdlo"
)

func Bark() string {
    return "Bark"
}

func TryAdamsPackage() string {
    // Externe package use
    n := stdlo.Must1(strconv.Atoi("42")) // panics if there's an error, otherwise returns the value
    p := stdlo.ToPtr("hello")            // gives you a pointer to the value
    return fmt.Sprint(n, " ", *p)
}
