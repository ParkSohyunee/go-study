package study

import "fmt"

func VariablesExample(){
	    // 1. Constant
		const name string = "sohyun"
		fmt.Println(name)
	
		const (
		  Visa = "Visa"
		  Master = "MasterCard"
		  Amex = "American Express"
		)
        fmt.Println(Visa, Master, Amex)
	
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