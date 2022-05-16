// Author: Roland Lux
// Date: 16.05.2022
package main

import (
	"flag"
	"fmt"
	"os"
)

// version number
var version string = "0.1"

func main() {

	// make usage use this instead
	os.Args[0] = "SmartDarts"

	versionFlag := flag.Bool("v", false, "show Version #")
	helpFlag := flag.Bool("h", false, "get some help")
	flag.Parse()

	// show version information
	if *versionFlag {
		fmt.Println(version)
		return
	}
	// show usage
	if *helpFlag {
		flag.Usage()
		return
	}

	// main
	fmt.Println("Start Darting!!!")

}
