package main

import (
	"log"

	"github.com/go_tour/tour/cmd"
	//"github.com/spf13/cobra/cobra/cmd"
)

func main() {
	err := cmd.Excute()
	if err != nil {
		log.Fatal("cmd.Excute err:%v", err)
	}
}
