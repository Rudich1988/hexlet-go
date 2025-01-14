package main


import (
    "fmt"
    greetingv1 "./greeting"
    //greetingv2 "./greeting/v2"
)


func main() {
    fmt.Println(greetingv1.Hello())
}
