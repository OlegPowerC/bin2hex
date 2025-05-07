package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	hfname := flag.String("hex", "", "filename with hex data")
	bname := flag.String("bin", "", "filename with binary data")
	flag.Parse()

	if len(*hfname) < 3 {
		fmt.Println("please provide hex filename")
		return
	}
	if len(*bname) < 3 {
		fmt.Println("please provide binary filename")
		return
	}
	hexstring := ""

	_, err := os.Stat(*bname)
	if os.IsNotExist(err) {
		fmt.Println("bin file not found")
		return
	}
	f, fopenerr := os.Open(*bname)
	if fopenerr != nil {
		panic(any(fopenerr))
	}
	defer f.Close()
	lb, _ := io.ReadAll(f)
	hexstringrep := ""
	hexstringrep = hex.EncodeToString(lb)
	every16counter := 0
	for rbl := 0; rbl < len(hexstringrep); rbl++ {
		hexstring += string(hexstringrep[rbl])
		every16counter++
		if every16counter >= 32 {
			every16counter = 0
			hexstring += "\r\n"
		}
	}
	fmt.Println("Write file:", *hfname)
	werr := os.WriteFile(*hfname, []byte(hexstring), 0644)
	if werr != nil {
		panic(any(werr))
	}
}
