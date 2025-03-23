package errors

import (
	"fmt"
	"os"
)

func HandleError(err error) {
	fmt.Println(err.Error())
	os.Exit(1)
}

