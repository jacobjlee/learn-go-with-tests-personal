package main

type Dictionary map[string]string

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists)")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

/*
- map은 &myMap같이 주소를 전달하지 않아도 수정이 가능하다(마치 레퍼런트 타입 같이)
- 그러므로 맵을 함수나 메서드에 전달할 때 포인터면에서 복사한다고 할 수 있다.
- map은 nil값을 가질 수 있으며, nil 맵은 읽을 때 빈 맵처럼 작동한다.
- 하지만 nill map을 작성하려고하면 런타임 패닉을 야기할 수 있다.
- map은 이미 존재하는 값을 추가하려고 해도 에러를 반환하지 않고, 해당 코드를 진행하며 새로운 값으로 덮어쓴다.

Map 선언
- var m map[string]string
- var dictionary = map[string]string{}
- var dictionary = make(map[string]string)

- go map은 delete라는 내장함수를 가짐
- 형식: delete(map, key)
*/
