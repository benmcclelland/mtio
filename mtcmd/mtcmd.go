//Example program using mtio package
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"syscall"

	"github.com/benmcclelland/mtio"
)

func doOp(file string, op int16, count int32) {
	f, err := os.OpenFile(file, os.O_RDWR, 0)
	if err != nil {
		log.Fatalln("Open failed", os.Args[1], err)
	}
	defer f.Close()

	log.Println("doing operation:", op, "count:", count)
	m := mtio.NewMtOp(
		mtio.WithOperation(op),
		mtio.WithCount(count))
	err = mtio.DoOp(f, m)
	if err != nil {
		log.Fatalln("Operation failed", err)
	}
}

func doStatus(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		log.Fatalln("Open failed", os.Args[1], err)
	}
	defer f.Close()

	log.Println("getting status")
	s, err := mtio.GetStatus(f)
	if err != nil {
		log.Fatalln("Status failed", err)
	}
	fmt.Println(s)
}

func doTell(file string) {
	f, err := os.OpenFile(file, os.O_RDONLY|syscall.O_NONBLOCK, 0)
	if err != nil {
		log.Fatalln("Open failed", os.Args[1], err)
	}
	defer f.Close()

	log.Println("getting position")
	p, err := mtio.GetPos(f)
	if err != nil {
		log.Fatalln("Get position failed", err)
	}
	fmt.Println("Positon:", p.BlkNo)
}

func main() {
	if len(os.Args) < 3 {
		log.Println("Usage:", os.Args[0],
			"<device path> op|status|tell [<operation id>] [<operation count>]")
		log.Fatalln("Invalid arguments")
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	switch os.Args[2] {
	case "op":
		if len(os.Args) < 4 {
			log.Fatalln("operation id not specified")
		}

		i, err := strconv.ParseInt(os.Args[3], 10, 16)
		if err != nil {
			log.Fatalln("invalid operation", err)
		}

		if len(os.Args) < 5 {
			doOp(os.Args[1], int16(i), 1)
			return
		}

		c, err := strconv.ParseInt(os.Args[4], 10, 32)
		if err != nil {
			log.Fatalln("invalid operation id", err)
		}
		doOp(os.Args[1], int16(i), int32(c))

	case "status":
		doStatus(os.Args[1])

	case "tell":
		doTell(os.Args[1])

	default:
		log.Fatalln("invalid operation")
	}
}
