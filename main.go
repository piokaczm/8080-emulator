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
		disassemble(*dFlag)
	}
}

func disassemble(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf(err.Error())
	}

	err = disassembler.Decode(data)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
