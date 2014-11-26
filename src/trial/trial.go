// hello.go (from Summerfield)

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	who := "GoWorld!"

	fmt.Println("os.Args", os.Args)

	if len(os.Args) > 1 { /* os.Args[0] is "hello" */

		who = strings.Join(os.Args[1:], " ")

	}

	fmt.Println("Hello", who)

	for day := range longWeekend {
		fmt.Println(longWeekend[day])
	}
	for i := range lowPrimes {
		fmt.Println(lowPrimes[i])
	}

}

var longWeekend = []string{"Friday", "Saturday", "Sunday", "Monday"}
var lowPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19}
