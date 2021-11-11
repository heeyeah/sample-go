package main

import (
	"fmt"
	"github.com/heeyeah/heepack"
)

func main() {
	fmt.Println("sample-go, heeye")
	val := heepack.Reverse("November")
	fmt.Println("this is November reverse word : " + val)
}
