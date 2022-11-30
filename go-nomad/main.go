package main

/*
	main.go file 은 프로젝트가 컴파일을 해야할 경우 필요하다. (진입점)
	이 외의 파일명은 컴파일 되지 않는다. (ex. learning.go) 기능의 묶음 혹은 다른 사람들이 사용할 것들
	go 는 다양한 패키지들을 제공하고 있다. 거의 모든것이 패키지이다.
*/

func main() {
	/*
		function 을 export 대문자로 시작

			something.SayHello() (O)
			something.sayBye()	(X)
	*/

	/*
		Go 는 타입 언어이다.
		## Go 언어는 무언가를 만들고 사용하지 않으면 에러를 발생시킨다!

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

	/*
		# Function

		- return int
			func multiply(a int, b int) int {
				return a * b
			}

		- 같은 타입의 파라미터인 경우
			func multiply(a, b int) int {
				return a * b
			}

		- void
			func multiply(a, b int) {
				fmt.Println(a * b)
			}

		- multiple return
			func lenAndUpper(name string) (int, string) {
				return len(name), strings.ToUpper(name)
			}
			totalLength, upperName := lenAndUpper("gyunam")
			fmt.Println(totalLength, upperName)

			- 2개를 return 하는 함수에서 하나의 값만 받는 방법 ( _ 를 사용하면 무시한다는 의미 )
			totalLength, _ := lenAndUpper("gyunam")

		- multiple argument
			func repeatMe(words ...string) {
				fmt.Println(words)
			}
			repeatMe("gyunam", "donggyu", "techan")

		- naked return (return 할 값을 꼭 명시하지 않아도 된다.)
			func lenAndUpper(name string) (length int, uppercase string) {
				// 이 때 := 로 생성하는 것이 아니라 업데이트 하는 것!
				length = len(name)
				uppercase = strings.ToUpper(name)
				return
			}

		- defer (함수가 끝날 때 추가적으로 무엇인가 동작시킬 수 있다! Cool)
			func lenAndUpper(name string) (length int, uppercase string) {
				defer fmt.Println("is done")

				length = len(name)
				uppercase = strings.ToUpper(name)
				return
			}


	*/

	/*
		# Loop
	*/
}
