package main

import (
	"fmt"
	"sync"
)

// func code.error() {
// 	for i := 0; i < 10; i++ { Здесь код неисправный не все числа загружены, как пример -
// 		go fmt.Println(i + 1) Сейчас я это исправлю WaitGroup. SYNC
// 	} отсутсвие синхронизации.
// }

func codeWait() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ { // счетчик for до 10
		wg.Add(1) // ! начал! 1 задача на каждую иттерацию = 10 задач.

		go func(i int) {
			fmt.Println(i) // передали счетчик цикла в горутину(anon.func)
			wg.Done()      // ! Завершил задачу !
		}(i)
	}

	wg.Wait() // блокировка основной горутины и ожидание
	fmt.Println("Завершение задачи... Выход из программы.")
}

func main() {
	codeWait() // запустил функцию
}

/* result = 10
4
5
0
3
6
9
8
7
2
Завершение задачи... Выход из программы. */
