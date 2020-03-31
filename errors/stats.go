package main

import "os"

func main() {
	//xx := &os.PathError{}

	miao, err := os.Stat("dare")

	if err, ok := err.(*os.PathError); ok {
		println(miao,err)
	}

}
