package main

import (
	"fmt"
	"strings"
	"unicode"

	//	"strings"
	//	"unicode"

	//	"unicode"

	//"unicode/utf8"

)

func main() {
	bootStrap()
}

func bootStrap() {


	//words:=[4]string{"the", "bla"} //unknown num of elements
	//fmt.Println(words[1])
	//
	//slice_example:=[]string{"bla", "bla" ,"blu"} // slice-- no length specified
	//fmt.Println(slice_example)

	//make:
	//arr:= make([]string, 0, 3) // capacity == 3, the maximum it can have
	//arr = append(arr, "bla")
	//fmt.Println(len(arr))

	//map == dict
	jsonDict := make(map[string]int)
	jsonDict["1"] = 1
	jsonDict["2"] = 2

	//jsonElement, ok := jsonDict["3"]
	//if !ok {
	//	fmt.Println("not in list")
	//} else {
	//	fmt.Println(jsonElement)
	//}
	s := "1 2 3-a 4     5-6-7-8-9-0"
	splitted := strings.Fields(s)
	var result = make([]string,0, len(splitted))
	for _, str := range splitted{
		//fmt.Printf("%T\n",[]rune(digit))
		for _, char := range(str){
			if unicode.IsDigit(rune(char)) {
				result = append(result, string(char))
			}
		}
	}
	fmt.Println(result)

	var res strings.Builder
	res.WriteString("{123")
	array2 := append(result[0:3])
	array3 := append(result[3:6])
	array4 := append(result[6:10])
	result[3] = "5"
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(res.String())


}
func buildStringFromArray(arr []string){

}
func calculator(a int, b string) int {
	fmt.Println(b)
	return a * 2
}

