package main

// main.go file 은 프로젝트가 컴파일을 해야할 경우 필요하다. (진입점)
// 이 외의 파일명은 컴파일 되지 않는다. (ex. learning.go) 기능의 묶음 혹은 다른 사람들이 사용할 것들

import (
	"fmt" // go 가 가고있는 패키지들 중 하나

	"github.com/k-gn/learngo/go-nomad/something"
)

func main() {

	fmt.Println("Hello Go!!")

	// function 을 export 대문자로 시작
	something.SayHello()
	// something.sayBye()
}
