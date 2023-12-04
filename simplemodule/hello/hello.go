package main

import (
	"fmt"
)

var globalVar = 12

type Phone interface {
	call() string
}

type NokiaPhone struct {
	name   string
	telNum int32
}

func (nokiaPhone NokiaPhone) call() string {
	fmt.Println("I am Nokia, I can call you!")
	return "Nokia"
}

type IPhone struct {
	name   string
	telNum int32
}

func (iphone IPhone) call() string {
	fmt.Println("I am IPhone, I can call you!")
	return "IPhone"
}

/**
错误处理示例
*/

type DivideError struct {
	dividee int
	divider int
}

func (divideError DivideError) Error() string {
	errorMsg := "Cannot proceed, the divider is zero.\n\t\tdividee: %d\ndivider: 0"
	return fmt.Sprintf(errorMsg, divideError.dividee)
}

func Divide(varDividee int, varDivider int) (result int, errorMsg string) {
	if varDivider == 0 {
		dData := DivideError{
			dividee: varDividee,
			divider: varDivider,
		}
		errorMsg = dData.Error()
		return
	} else {
		return varDividee / varDivider, ""
	}
}

func printStr() {
	fmt.Println("我GOTO到这里了")
}

func multiArr(aArr []int, aLen int, bArr []int, bLen int, str string) float32 {

	var i, sum1, sum2 int
	var avg1, avg2 float32

	aArr = append(aArr, 666)
	fmt.Println(aArr)
	for i = 0; i < aLen; i++ {
		sum1 += aArr[i]
	}

	avg1 = float32(sum1 / aLen)

	for i = 0; i < bLen; i++ {
		sum2 += bArr[i]
	}

	avg2 = float32(sum2 / bLen)
	fmt.Println(str)
	avg := avg1 + avg2
	fmt.Println(avg)
	result := avg / 2
	fmt.Println(result)

	return (avg1 + avg2) / 2
}

func main() {
	var balance1 = []int{1000, 2, 3, 17, 50}
	var balance2 = []int{1000, 2, 3, 17, 50}
	fmt.Println(multiArr(balance1, 6, balance2, 5, "Hi"))

	//	var a = 11
	//	for a < 17 {
	//		if a == 15 {
	//			a++
	//			continue
	//		}
	//		fmt.Printf("a的值为 : %d\n", a)
	//		a++
	//	}
	//
	//	var b = 11
	//LOOP:
	//	for b < 17 {
	//		if b == 15 {
	//			b++
	//			goto LOOP
	//		}
	//		fmt.Printf("b的值为 : %d\n", b)
	//		b++
	//	}
	//
	//	var c = 11
	//	for c < 17 {
	//
	//	ELEVEN:
	//		printStr()
	//		if c == 11 {
	//			c++
	//			goto ELEVEN
	//		}
	//		if b == 15 {
	//			b++
	//			goto ELEVEN
	//		}
	//		fmt.Printf("c的值为 : %d\n", c)
	//		c++
	//	}
	//
	//	result, errorMsg := Divide(100, 10)
	//	// 正常情况
	//	if errorMsg == "" {
	//		fmt.Println("100/10 = ", result)
	//	}
	//	result1, errorMsg1 := Divide(100, 0)
	//	// 当被除数为零的时候会返回错误信息
	//	if errorMsg1 != "" {
	//		fmt.Println("errorMsg is: \n", errorMsg1)
	//	} else {
	//		fmt.Println("100/0 = ", result1)
	//	}

	//var numbers []int
	//
	//numbers = append(numbers, 0)
	//
	//fmt.Printf("原始切片的长度=%d 容量=%d 切片内容=%v\n", len(numbers), cap(numbers), numbers)
	//
	//numbers = append(numbers, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	//
	//fmt.Printf("原始切片的长度=%d 容量=%d 切片内容=%v\n", len(numbers), cap(numbers), numbers)
	//
	//numbers = append(numbers, 10)
	//
	//fmt.Printf("原始切片的长度=%d 容量=%d 切片内容=%v\n", len(numbers), cap(numbers), numbers)
	//
	//numbers = append(numbers, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21)
	//fmt.Printf("原始切片的长度=%d 容量=%d 切片内容=%v\n", len(numbers), cap(numbers), numbers)
	//
	//numbers1 := make([]int, len(numbers), cap(numbers)*2)
	//copy(numbers1, numbers)
	//fmt.Printf("numbers1 原始切片的长度=%d 容量=%d 切片内容=%v\n", len(numbers1), cap(numbers1), numbers1)
	//var phone Phone
	//
	//phone = new(NokiaPhone)
	//phone.call()
	//
	//phone = new(IPhone)
	//phone.call()
	//
	//s := 20
	//var spr *int
	//spr = &s
	//fmt.Printf("指针spr指向的变量s的地址：%x\n", &s)
	//fmt.Printf("指针spr的存储地址：%x\n", spr)
	//fmt.Printf("*spr指针变量的值：%d\n", *spr)
	//str := "Hello world!"
	//fmt.Println(len(str))
	//strArr := []string{"hello", "world"}
	//fmt.Println(strings.Join(strArr, " "))
	//
	//var vname1, vname2, vname3 = 1, 2, 3
	//fmt.Println(vname1, vname2, vname3)

	//vname4, vname5, vname6 := vname1, vname2, vname3
	//fmt.Println(vname4, vname5, vname6)
	//
	//fmt.Println("string length is ", len("日本/x80語"))
	//for pos, char := range "日本/x80語" { // \x80 is an illegal UTF-8 encoding
	//	fmt.Printf("character %#U starts at byte position %d\n", char, pos)
	//}
	//fmt.Println(globalVar)
	//add := vname1 + vname2
	//sub := vname2 - vname1
	//multi := vname1 * vname2
	//divide := float64(vname1) / float64(vname2)
	//vname3++
	//vname3--
	//num := maxNum(vname1, vname2)
	//
	//fmt.Println(add, sub, multi, divide, vname3, num)
	//
	//log.SetPrefix("greetings: ")
	//log.SetFlags(0)
	//// Simple.
	//message, err := greetings.Hello("Gladys")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(message)
	////Multiple.
	//names := []string{"张三", "李四", "王五"}
	//messages, err := greetings.Hellos(names)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(messages)
}

func maxNum(var1, var2 int) int {
	fmt.Println(globalVar)
	if var1 > var2 {
		return var1
	} else {
		return var2
	}

}
