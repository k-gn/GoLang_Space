package main

/*
	main.go file 은 프로젝트가 컴파일을 해야할 경우 필요하다. (진입점)
	이 외의 파일명은 컴파일 되지 않는다. (ex. learning.go) 기능의 묶음 혹은 다른 사람들이 사용할 것들
	go 는 다양한 패키지들을 제공하고 있다.
*/

func main() {
	/*
		function 을 export 대문자로 시작

			something.SayHello() (O)
			something.sayBye()	(X)
	*/

	/*
		Go 는 타입 언어이다.

		- 상수 : const
			const name string = "gyunam"

		- 변수 : var
			var name string = "gyunam"
			name = "donggyu"

		위 방식을 축약할 수 있다. (이 때 Go 가 자동으로 타입을 추론해준다.)
			name := "gyunam"

		처음 선언된 타입에 의존한다.
			name = false (X)

		함수 밖에선 축약할 수 없다.
			var flag bool = false
	*/

}
