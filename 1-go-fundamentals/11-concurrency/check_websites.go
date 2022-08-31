package concurrency

import "time"

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// Send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// Receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	time.Sleep(2 * time.Second)

	return results
}

/*
doSomething()과 같은 함수를 호출한다고 했을 때, 해당 함수에 리턴 값이 없더라도 보통 해당 함수에 응답을 기다리게 된다.
이런 동작을 블로킹이라고 부른다 - 해당 함수가 동작이 끝나기 까지 우리를 기다리게하기 때문.

Go에서 이런 방식과 다르게 블로킹하지 않고 각 프로세스를 실행하게끔 해주는 것을 고루틴이라고 부름.
Go에게 새로운 고루틴을 시작하게끔하려면 함수 앞에 go 키워드를 붙혀 함수를 호출해주면 된다.
예: go doSomething()

고루틴을 시작할 수 있는 유일한 방법은 함수 콜 앞에 go를 붙이는 방법뿐이기 때문에, 고루틴을 시작할 때 익명함수를 주로 사용한다.
익명함수는 보통 함수와 흡사하게 보이지만 이름에서 알 수 있듯이, 함수 이름을 갖지 않는다.

마지막에 ()이 위치함으로써 선언과 동시에 병렬적으로 실행된다.

병렬적으로 실행하여 동시에 두번 이상으로 map을 작성하려고하면 fatal error가 발생함

**채널**
채널은 Go 자료구조중에 하나로써, 값을 받고 전달할 수 있다.
즉, 다른 프로세스 사이에서 커뮤니케이션이 가능해진다.

** Wrapping up **
- go루틴을 통해 동시에 하나 이상의 웹사이트를 체크할 수 있음
- 익명함수를 통해 웹사이트를 체크하기 위한 동시 프로세스를 시작하는데 사용할 수 있음
- 채널은 다른 프로세스 사이에서 커뮤니케이션을 정리 및 제어할 수 있게 되고 레이스 컨디션 버그를 피할 수 있게 해줌
- 테스트 레이스 디텍터를 통해 컨큐렌트 코드 문제를 디버깅할 수 있음
*/
