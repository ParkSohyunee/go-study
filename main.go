package main

import (
	"fmt"

	"github.com/ParkSohyunee/go-study/greetings"
)

func main()  {
		fmt.Println("Hello world") // fmt 패키지의 PrintIn 함수를 호출 (대문자는 다른 패키지로부터 export되었음을 의미)
    message := greetings.Hello("sohyun")
    fmt.Println(message)
}