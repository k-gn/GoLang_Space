package something

/*
	main.go file 은 프로젝트가 컴파일을 해야할 경우 필요하다. (진입점)
	이 외의 파일명은 컴파일 되지 않는다. (ex. learning.go) 기능의 묶음 혹은 다른 사람들이 사용할 것들
	go 는 다양한 패키지들을 제공하고 있다. 거의 모든것이 패키지이다.
*/

func basic() {
	/*
		function 을 export 할 경우 대문자로 시작

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

		func superAdd(numbers ...int) int {
			total := 0
			for _, number := range numbers {
				total += number
			}
			return total
		}

		total := superAdd(1, 2, 3, 4, 5)
		fmt.Println(total)

		- for [var] := range [array]

			func superAdd(numbers ...int) {
				fmt.Println(numbers)

				for index, number := range numbers {
					fmt.Println(index, number)
				}

				for _, number := range numbers {
					fmt.Println(number)
				}
			}

		- basic for

			func superAdd(numbers ...int) {
				fmt.Println(numbers)

				for i := 0; i < len(numbers); i++ {
					fmt.Println(numbers[i])
				}
			}
	*/

	/*
		# If Else

		- basic if

			func catIDrink(age int) bool {
				if age < 18 {
					return false
				} else {
					return true
				}
			}

		- create variable (if 안에서만 사용)

			func catIDrink(age int) bool {
				if koreanAge := age + 2; koreanAge < 18 {
					return false
				} else {
					return true
				}
			}
	*/

	/*
		# switch

		- basic switch
			func catIDrink(age int) bool {
				switch age {
				case 10:
					return false
				case 18:
					return true
				}
				return false
			}

		- switch case 에 조건 넣기

			func catIDrink(age int) bool {
				switch {
				case age < 10:
					return false
				case age == 18:
					return true
				}
				return false
			}

		- add variable

			switch koreanAge := age + 2; koreanAge {
				case 10:
					return false
				case 18:
					return true
				}
				return false
			}
	*/

	/*
		# Pointers

		- 메모리 관련 기능 (Low-level Programming)
		- & 기호로 메모리 주소를 볼 수 있다.
		- * 기호로 메모리 주소의 값을 볼 수 있다.
		- 서로 연결되어 있는 개념 (메모리에 저장된 객체를 서로 똑같이 가지고 싶을 때 사용한다. - 계속 복사본을 만드는게 아님)

		a := 2
		b := &a // b 는 a 를 살펴보는 "포인터" 가 된다.
		fmt.Println(a, b)
		fmt.Println(*b)

		a = 5
		fmt.Println(a, b)
		fmt.Println(*b)

		*b = 20
		fmt.Println(a)

	*/

	/*
		# Array

		- basic array

			names := [5]string{"gyunam", "dong", "chan"}
			names[3] = "minsoo"
			names[4] = "gwang"
			fmt.Println(names)

		- slice (without length)
			names := []string{"gyunam", "dong", "chan"}
			names = append(names, "bin")
	*/

	/*
		# Map

		- map[key type]val type{ ... }

			gyul := map[string]string{"name": "gyunam", "age": "12"}
			for key, value := range gyul {
				fmt.Println(key, value)
			}
	*/

	/*
		# Structs

		- Go 는 생성자 메소드가 없다.

			type person struct {
				name    string
				age     int
				favFood []string
			}

			favFood := []string{"kimchi", "beer"}
			gyunam := person{name: "gyunam", age: 26, favFood: favFood} // 변수명을 쓸지 안쓸지 하나로 통일해야한다.

	*/
}
