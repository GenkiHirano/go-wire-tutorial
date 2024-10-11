package main

import "go-wire-tutorial/di"

func main() {
	event := di.InitializeEvent()

	event.Start()
}
