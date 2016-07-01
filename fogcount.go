package main

import (
	"fmt"
	"os"

	"github.com/ccirello/gunning-fog/fogcount"
)

func main() {
	fmt.Println(fogcount.Analyze(os.Stdin))
}
