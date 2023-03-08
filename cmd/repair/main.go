package main

import (
	"repair/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
