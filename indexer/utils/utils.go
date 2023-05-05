package utils

import (
	"fmt"
	"os"
)

func ErrorPrinter(err string) {
	fmt.Println("error: " + err)
	os.Exit(1)
}
