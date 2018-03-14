package main

import (
	"flag"
	"io/ioutil"
	"log"

	"github.com/piokaczm/8080-emulator/disassembler"
)

func main() {
	dFlag := flag.String("d", "", "use this flag to disassemble provided file")
	flag.Parse()

	if len(*dFlag) > 0 {
		data, err := ioutil.ReadFile(*dFlag)
		if err != nil {
			log.Fatalf(err.Error())
		}

		disassembler.Decode(data)
	}
}
