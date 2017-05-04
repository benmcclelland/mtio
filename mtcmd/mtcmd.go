//Example program using mtio package
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/benmcclelland/mtio"
)

func printStatus(s *mtio.MtGet) {
	fmt.Printf("%v (%v)\n", mtio.MtTypeToString(s.Type), s.Type)
	fmt.Println("Residual count:", s.ResID)
	fmt.Printf("Device registers: %x\n", s.DsReg)
	fmt.Printf("Status registers: %x\n", s.GStat)
	fmt.Println(mtio.MtStatusToString(s.GStat))
	fmt.Println("Error register:", s.ErReg)
	fmt.Println("Possibly inaccurate:")
	fmt.Println("  Current file:", s.FileNo)
	fmt.Println("  Current block number:", s.BlkNo)
}

func doOp(f *os.File, op int16, count int32) {
	log.Println("doing operation:", op, "count:", count)
	m := mtio.NewMtOp(
		mtio.WithOperation(op),
		mtio.WithCount(count))
	err := mtio.DoOp(f, m)
	if err != nil {
		log.Fatalln("Operation failed", err)
	}
}

func doStatus(f *os.File) {
	log.Println("getting status")
	s, err := mtio.GetStatus(f)
	if err != nil {
		log.Fatalln("Status failed", err)
	}
	printStatus(s)
}

func doTell(f *os.File) {
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

	f, err := os.OpenFile(os.Args[1], os.O_RDWR, 0)
	if err != nil {
		log.Fatalln("Open failed", os.Args[1], err)
	}
	defer f.Close()

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
			doOp(f, int16(i), 1)
			return
		}

		c, err := strconv.ParseInt(os.Args[4], 10, 32)
		if err != nil {
			log.Fatalln("invalid operation id", err)
		}
		doOp(f, int16(i), int32(c))

	case "status":
		doStatus(f)

	case "tell":
		doTell(f)

	default:
		log.Fatalln("invalid operation")
	}
}
