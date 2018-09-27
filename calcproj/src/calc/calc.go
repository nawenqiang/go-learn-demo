package main

import (
	"fmt"
	"os"
	"simplemath"
	"strconv"
)

var Usage = func() {
	fmt.Println("calc command [arg] ...")
	fmt.Println("the commands are:\n\tadd\n\tsqrt")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[1] {
	case "add":
		if len(args) != 4 {
			Usage()
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			Usage()
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("result:", ret)
	case "sqrt":
		if len(args) != 3 {
			Usage()
			return
		}
		v1, err := strconv.Atoi(args[2])
		if err != nil {
			Usage()
			return
		}
		ret := simplemath.Sqrt(v1)
		fmt.Println("result:", ret)

	default:
		Usage()
		return
	}

}
