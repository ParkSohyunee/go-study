package study

import (
	"fmt"
	"strings"
)

func multiply(a, b int) int {
    return a * b
}

// Return multiple value
func lenAndUpper(name string) (int, string)  {
    return len(name), strings.ToUpper(name)
}

// Naked return
func lenAndUpperNakedReturn(name string) (length int, uppercase string)  {
    length = len(name)
    uppercase = strings.ToUpper(name)
    return 
}

// Variadic Function
func repeatWord(words ...string)  {
    fmt.Println(words)
}

func FunctionExample() {
    multiplyResult := multiply(3, 4)
    fmt.Println(multiplyResult)

    // (참고) 하나의 값만 받고 싶다면 totalLength, _ (underscore)
    totalLength, upperName := lenAndUpper("sohyun")
    fmt.Println(totalLength, upperName)

    fmt.Println(lenAndUpperNakedReturn("sohyun"))

    repeatWord("a", "b", "c") // [a b c]
    repeatWord("a")

    // Anonymous Function
    // Go에서 함수 안에 다른 함수를 정의하려면 익명함수로 선언해야 함
    // 함수 전체를 변수에 할당하거나 다른 함수의 파라미터에 직접 정의되어서 사용
    result := func (a, b int) int {
        return a * b
    }(3, 4)

    fmt.Println(result)

    // Go에서 함수는 일급함수로 취급
    // 함수는 다른 함수에 인수로 전달될 수 있고, 다른 함수에서 반환될 수 있으며, 변수에 값으로 할당될 수 있음
}