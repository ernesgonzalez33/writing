package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("/tmp/dat1")
	check(err)

	defer f.Close()

	for {
		_, err := f.WriteString("writing\n")
		check(err)
		fmt.Printf("Writing!\n")

		f.Sync()
	}

}
