package main

import (
	"log"
	_ "matchers"
	"os"
	"search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
