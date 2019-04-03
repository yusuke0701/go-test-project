package main

import (
	"fmt"
)

func main() {
	// カリー化
	val := curryingSum(3)(5)
	fmt.Println(val)

	// 部分適用
	tasu3 := curryingSum(3)
	fmt.Println(tasu3(10))
}

func curryingSum(val1 int) func(val2 int) int {
	return func(val2 int) int {
		return val1 + val2
	}
}

/*
以下の記事を参考にした。
https://qiita.com/jzmstrjp/items/99cc1c8ebcfc703b1410?utm_source=Qiita%E3%83%8B%E3%83%A5%E3%83%BC%E3%82%B9&utm_campaign=c79ac2dc65-Qiita_newsletter_356_04_03_2019&utm_medium=email&utm_term=0_e44feaa081-c79ac2dc65-33531037
確かに引数が少ないほうが関数と引数との関連性も明確化される
*/