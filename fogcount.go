package main // import "cirello.io/gunning-fog"

import (
	"fmt"
	"os"

	"cirello.io/gunning-fog/fogcount"
)

func main() {
	fmt.Println(fogcount.Analyze(os.Stdin))
}
