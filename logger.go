package logger

import (
	"fmt"
)

var Version string = "0.1.0"

func Logger(str string) {
	fmt.Println("[LOG] - ", str)
}
