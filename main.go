package main

import (
	"fmt"

	"github.com/ParkSohyunee/go-study/greetings"
	"github.com/ParkSohyunee/go-study/study"
)


func main()  {
    // fmt 패키지의 PrintIn 함수를 호출 (대문자는 다른 패키지로부터 export되었음을 의미)
    fmt.Println("Hello world") 
    message := greetings.Hello("sohyun")
    fmt.Println(message)

    study.VariablesExample()
    study.FunctionExample()

    study.Loop()
}