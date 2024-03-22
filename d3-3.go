package main

import "fmt"
import "time"


func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}











func main4() {
	fmt.Println("main4 working",)
	go say("Hello")
	say("Hi")
	// print output -> main4 working
					// Hello
					// Hi
					// Hello
					// Hi
					// Hi
					// Hello
					// Hi
					// Hello
					// Hello
					// Hi
}