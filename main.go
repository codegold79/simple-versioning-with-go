package main

import (
	// Change the import path to be v1 or v2 and you can see nothing else in the code needs to change to upgrade versions.
	"github.com/codegold79/simple-versioning-with-go/greeting/v2"
)

func main() {
	greeting.Hello()
}
