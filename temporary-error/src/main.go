package main

import (
	"errors"
	"fmt"
)

func main() {
	err := doSomething()
	if err != nil {
		if IsTemporary(err) {
			fmt.Println("一時エラー")
		}
		fmt.Println(err.Error())
	}
}

func doSomething() error {
	return &tmperr{original: errors.New("エラー発生")}
}

//IsTemporary は一時エラーの場合は true を返す
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.IsTemporary()
}

type temporary interface {
	IsTemporary() bool
}

type tmperr struct {
	original error
}

func (err *tmperr) Error() string {
	return err.original.Error()
}
func (err *tmperr) IsTemporary() bool {
	return true
}
