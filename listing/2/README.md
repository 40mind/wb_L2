Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и порядок их вызовов.

package main

import (
"fmt"
)

func test() (x int) {
defer func() {
x++
}()
x = 1
return
}


func anotherTest() int {
var x int
defer func() {
x++
}()
x = 1
return x
}


func main() {
fmt.Println(test())
fmt.Println(anotherTest())
}



2 1 defer выводится после завершения фукнции, во втором случае сперва выполнится
return x(=1), а потом defer func(){x++}()