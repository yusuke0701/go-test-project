package main

import "fmt"

// Field は、キャラクターを配置するフィールドを表す
type Field struct{}

// Print は、Fieldを表示する
func (f Field) Print() {
	hoge := `
--------------------|
|  
|  --------------
|  |            |
|  --------------
|
|  --------------
|  |            |
|  |            |
|  --------------
|
--------------------
`
	fmt.Printf(hoge)
}
