package main


import (
    "fmt"
    "github.com/fatih/color"
    greetingv1 "hexlet-go/greeting"

)


func main() {
    fmt.Println(greetingv1.Hello())
    color.Magenta("Hello World!")
}
