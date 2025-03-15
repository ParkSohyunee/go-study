package study

import "fmt"

func Loop()	{
	// for문
	// 초기값; 조건식; 증감식은 경우에 따라 생략 가능, 괄호 생략
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum) 

	// 조건식만 쓰는 for문 (JS while loop)
	n := 1
	for n < 50 {
		n++
	}
	fmt.Println(n) 

	// Infinite loop
	// for {
	// 	println("Infinite loop")
	// }

	// for-range (JS forEach)
	// for 인덱스,요소값 := range 컬렉션
	numbers := []int{1, 2, 3, 4, 5}
	for index, number := range numbers {
		fmt.Println(index, number)
	}

	// 반목문의 탈출
	// break, continue, goto
	i := 0
	for  {
		if i >= 10 {
			goto L1
		}
		i++
	}
	// Labels are declared by labeled statements 
	// and are used in the "break", "continue", and "goto" statements
L1: 
	fmt.Println(i) // 10

	j := 0
L2:
	for {
		if j == 0 {
			break L2
		}
	}
	println("OK") // break L2문이 for루프를 빠져나와서 L2레이블로 이동 후, for루프를 건너뛰고 println문으로 이동
}