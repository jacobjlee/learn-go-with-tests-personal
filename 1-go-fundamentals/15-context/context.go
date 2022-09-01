package context

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: 에러 로깅하기
		}

		fmt.Fprint(w, data)
	}
}

/*
컨텍스트 패키지는 새로운 컨텍스트 값을 이미 존재하던 곳에서 가져오는 역활을 한다.
트리와 같은 구조를 형성하는데(?) - 하나의 컨텍스트가 취소되면 하나의 컨텍스트로부터 파생된 모든 컨텍스트가 취소된다.
*/
