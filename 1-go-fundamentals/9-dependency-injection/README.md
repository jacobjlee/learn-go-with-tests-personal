## Dependency Injection (의존 관계 주입)

누군가를 환영하는 함수 작성
```go
func greet(name string) {
	fmt.Printf("Hello, %s", name)
}
```

이 함수를 테스팅하기는 매우 어렵다 - `fmt.Printf`를 `stdout`에 호출하기 때문에 포착하기가 어려움.

이러한 니즈가 있을 때 프린팅의 의존성을 주입해줄 수 있음.

작성한 함수는 어떻게 혹은 어디에서 프린팅이 작동하는지 알 필요가 없기때문에 콘크리트 타입이 아닌 인터페이스를 받아들여야한다.

소스 코드를 파고 들어가보기
```go
// fmt.Printf

// It returns the number of bytes written and any write error encountered.
func Printf(format string, a ...interface{}) (n int, err error) {
    return Fprintf(os.Stdout, format, a...)
}
```

`Printf`가 `os.Stdout`에 전달하기 위해 `Fprintf`를 사용하는 것 확인
-> `os.Stdout`이 뭔지 들여다보기

```go
// Fprintf
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	p := newPrinter()
	p.doPrintf(format, a)
	n, err = w.Write(p.buf)
	p.free()
	return
}

// io.Writer
type Writer interface {
Write(p []byte) (n int, err error)
}
```

여기에서 이제 `os.Stdout`이 `io.Writer`를 구현하고, `Printf`가 `io.Writer`를 기대하는 `Fprintf`에 `os.Stdout`을 전달한다고 추정할 수 있다.
