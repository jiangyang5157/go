package stuff

import "fmt"

func Chan1() {
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		go func() {
			ch <- i
		}()
	}

	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
	fmt.Println(<-ch) // 5
}

func Chan2() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	fmt.Println(<-ch) // 0
	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3
	fmt.Println(<-ch) // 4
}

func Chan3() {
	c := fanIn(boring("Joe"), boring("Ann"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're boring, I'm leaving.")
}

//Returns receive-only channel of strings
func boring(msg string) <-chan string {
	c := make(chan string)
	//launch the goroutine from inside the function
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}
