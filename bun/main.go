package main

import (
	"github.com/adyatlov/bun/bun/cmd"
	_ "github.com/adyatlov/bun/import"
)

const printProgress = false
const printLong = false

func main() {
	cmd.Execute()
}
