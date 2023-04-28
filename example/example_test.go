package example

import (
	"errors"
	"fmt"
	"testing"
)

func TestUseError(t *testing.T) {
	fmt.Println("========= type Err:")
	printError(ErrPlanIsEmpty)

	fmt.Println("========= type Err2:")
	printError(Err2Addr)

	fmt.Println("========= custorm type:")
	printError(errors.New("custom error"))
}

func printError(err error) {
	if err == nil {
		fmt.Println("error is Nil")
		return
	}
	fmt.Println("err.Error():", err.Error())
	fmt.Println("err Code:", Code(err))
}
