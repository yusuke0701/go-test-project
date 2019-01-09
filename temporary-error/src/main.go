package main

import (
	"fmt"
	"errors"
)

type temporary interface {
	Temporary() bool
}

//IsTemporary は一時エラーの場合は true を返す
func IsTemporary(err error) bool {
	te, ok := err.(temporary)
	return ok && te.Temporary()
}

type tmperr struct {
	original error
}

func (err *tmperr) Error() string {
	return err.original.Error()
}
func (err *tmperr) Temporary() bool {
	return true
}
func Temporary(err error) error {
	return &tmperr{original: err}
}

func f() error {
	return Temporary(errors.New("失敗しました"))
}

func main() {
	err := f()
	switch {
	case err == nil:
		fmt.Println("エラーなし")
	case IsTemporary(err):
		fmt.Println("一時エラー", err)
	default:
		fmt.Println("エラー", err)
	}
}