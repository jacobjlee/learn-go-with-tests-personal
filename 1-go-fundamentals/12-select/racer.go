package racer

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

/*
net/http - HTTP 요청을 만드는데 사용
net/http/httptest - HTTP 요청을 테스팅하는데 사용
- 이 기본 라이브러리를 통해서 쉽게 mock HTTP 서버를 만들 수 있음.

time.Now()로 시각 기록
http.Get으로 url 접속 및 컨텐츠를 받아옴.
이 때 http.Response와 error를 리턴함.
time.Since는 시작 시간을 취하고 시작 시간과 차이나는 만큼의 시간을 반환함

By prefixing a function call with defer it will now call that function at the end of the containing function.
i.e. closing a file or a server

제로 값이 nil인 채널이 존재하고 <-로 해당 채널에 보내게 되면 그 채널은 영원히 블락된다 - nil 채널엔 아무것도 보낼 수 없기 때문
*/
