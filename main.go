package main

import "github.com/ken1kasap/learning-wire/di"

func main() {
	e := di.InitializeEvent()
	e.Start()
}
