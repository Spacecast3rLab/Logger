package Logger

import (
	"fmt"
)

var Version string = "0.1.0"

func Log(str string) {
	fmt.Println("[LOG] - ", str)
}
