package main

import (
	"fmt"

	"github.com/ParkSohyunee/go-study/greetings"
)


func main()  {
    // fmt 패키지의 PrintIn 함수를 호출 (대문자는 다른 패키지로부터 export되었음을 의미)
    fmt.Println("Hello world") 
    message := greetings.Hello("sohyun")
    fmt.Println(message)

    // 1. Constant
    const name string = "sohyun"
    fmt.Println(name)

    // const (
    //   Visa = "Visa"
    //   Master = "MasterCard"
    //   Amex = "American Express"
    // )

    // 2. Variables
    var name1 string = "sohyun"
    var age int
    age = 100
    fmt.Println(name1, age)

    // Short Assignment Statement (함수 내부에서, 변수에서만)
    // 타입 자동 설정
    name2 := "sohyun"
    fmt.Println(name2)

    // 여러 변수 선언 및 초기화
    var one, two, three = true, "sohyun", 1
    fmt.Println(one, two, three)

    // Boolean
    // error would be occurred
    // if true == 1 {
    //   fmt.Println("value of true is 1")
    // } 
}