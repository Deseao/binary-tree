package main

import (
	"fmt"
	//"golang.org/x/tour/tree"

	"github.com/Deseao/binary-tree/internal/tree"
)

//Stolen solution from https://stackoverflow.com/a/24580091/1162540
func main() {
	big := tree.New(1)
	big.Insert(19)
	big.Insert(12)
	big.Insert(2)
	big.Insert(-3)
	fmt.Println("Big tree: ", big.Print())
	fmt.Println("Same: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Different: ", Same(tree.New(1), tree.New(2)))
}

func Walk(t *tree.Tree, ch chan int) {
	var walk func(t *tree.Tree, ch chan int)
	walk = func(t *tree.Tree, ch chan int) {
		if t == nil {
			return
		}
		walk(t.Left, ch)
		ch <- t.Value
		walk(t.Right, ch)
	}
	walk(t, ch)
	close(ch)
}

func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}
