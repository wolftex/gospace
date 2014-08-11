// hello.go (from Summerfield)

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	who := "GoWorld!"

	if len(os.Args) > 1 { /* os.Args[0] is "hello" */

		who = strings.Join(os.Args[1:], " ")

	}

	fmt.Println("Hello", who)
}
