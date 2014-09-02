// stacker.go (Summerfield)

package main

import (
	"fmt"
	"stacker/stack"
	"strings"
)

func main() {

	var haystack stack.Stack

	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(08.15)

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}

	var aStack stack.Stack
	aStack.Push("Aarvark")
	aStack.Push(5)
	aStack.Push(19)
	x, err := aStack.Top()
	fmt.Println(x)
	aStack.Push(-6e-4)
	aStack.Push("Baker")
	aStack.Push(-3)
	aStack.Push("Cake")
	aStack.Push("Dancer")
	x, err = aStack.Top()
	fmt.Println("Top():=", x, "error_return=", err)
	aStack.Push(11.7)
	fmt.Println("stack is empty", aStack.IsEmpty())
	fmt.Printf("Len() == %d Cap == %d\n", aStack.Len(), aStack.Cap())

	difference := aStack.Cap() - aStack.Len()
	for i := 0; i < difference; i++ {
		aStack.Push(strings.Repeat("*", difference-1))
	}
	fmt.Printf("Len() == %d Cap == %d\n", aStack.Len(), aStack.Cap())

	for aStack.Len() > 0 {
		x, _ = aStack.Pop()
		fmt.Printf("%T %v\n", x, x)
	}
	fmt.Println("stack is empty", aStack.IsEmpty())
	x, err = aStack.Pop()
	fmt.Println("Pop():=", x, "error_return=", err)
	x, err = aStack.Top()
	fmt.Println("Top():=", x, "error_return=", err)

}
