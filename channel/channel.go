package channel

import (
	_"fmt"
	_"sync"
	"fmt"
	//"time"
)


func Count(ch chan int) {
	ch <- 1
	fmt.Println("Counting\n")
}

func ChannelFunc() {
	chs := make([] chan int, 10)
	
	for i:=0; i<10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i])
	}
	
	for id , ch := range(chs) {
		 <- ch
		fmt.Println("id:%d\n", id)
	}
	//time.Sleep(3000000000)
}
