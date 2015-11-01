package main

import (
	"fmt"
	"os"
)

func main() {
	userFilename := "pic.jpg"
	userfileWritename := "picc.jpg"
	fout, err := os.Create(userfileWritename)

	if err != nil {
		fmt.Println(userfileWritename, err)
		return
	}

	fl, err := os.Open(userFilename)
	if err != nil {
		fmt.Println(userFilename, err)
		return
	}

	defer fl.Close()
	defer fout.Close()
	buf := make([]byte, 1024*13)
	for {
		n, _ := fl.Read(buf)

		if 0 == n {
			break
		}
	}
	fout.Write(buf)
	fmt.Println(buf)
	for {
		n, _ := fout.Read(buf)
		if 0 == n {
			break
		}
	}
	fmt.Println("\n")
	fmt.Println(buf)
}
