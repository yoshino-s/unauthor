package main

import "github.com/yoshino-s/unauthor/cmd"

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
