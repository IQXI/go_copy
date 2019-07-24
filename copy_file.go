package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var from string
var to string
var offset int
var limit int

func init() {
	flag.StringVar(&from, "from", "", "path to folder copy --from")
	flag.StringVar(&to, "to", "", "path to folder copy --to")

	flag.IntVar(&offset, "offset", 0, "offset in folder --from")
	flag.IntVar(&limit, "limit", 0, "limit of files count")

}

func Copy(from, to string, offset, limit int) {
	file, err := os.OpenFile(from, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Read file is not found!\n")
			return
		}
	}
	fileWrite, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Write file is not found!\n")
			return
		}
	}

	defer file.Close()
	defer fileWrite.Close()

	_, err = file.Seek(int64(offset), io.SeekStart)
	if err != nil {
		log.Printf("Offset bigger than size of file!\n")
		return
	}

	buf := make([]byte, limit)
	bufOffset := 0
	for bufOffset < limit {
		read, err := file.Read(buf[bufOffset:])
		bufOffset += read
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Cannot read: %v\n", err)
			return
		}
	}
	if bufOffset == 0 {
		log.Printf("Offset is bigger than size of file\n")
		return
	}
	fmt.Printf("Read %v bytes\n", bufOffset)

	written, err := fileWrite.Write(buf[:bufOffset])
	if err != nil {
		log.Printf("Cannot write: %v\n", err)
		return
	}
	fmt.Printf("Write %v bytes\n", written)

}

func main() {
	flag.Parse()
	fmt.Printf("%v %v %v %v\n", from, to, offset, limit)
	//Copy("from\\1.txt", "to\\6.txt", 480, 3)
	Copy(from, to, offset, limit)
}
