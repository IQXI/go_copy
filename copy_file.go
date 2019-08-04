package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var from, to string
var offset, limit int

func init() {
	flag.StringVar(&from, "from", "", "path to folder copy --from")
	flag.StringVar(&to, "to", "", "path to folder copy --to")

	flag.IntVar(&offset, "offset", 0, "offset in folder --from")
	flag.IntVar(&limit, "limit", 0, "limit of files count")

}

func Copy(from, to string, offset, limit int) error {
	file, err := os.OpenFile(from, os.O_RDONLY, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Read file is not found!\n")
		} else {
			log.Print(err)
		}
		return err

	}

	stat, err := file.Stat()
	if err != nil {
		log.Print(err)
		return err
	}

	if int(stat.Size()) <= offset {
		err = errors.New("Offset bigger than size of file!\n")
		return err
	}

	//if int(stat.Size()) < offset+limit {
	//	errors.New("Offset + limit bigger than size of file!\n")
	//}

	fileWrite, err := os.OpenFile(to, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Write file is not found!\n")
		} else {
			log.Print(err)
		}
		return err
	}

	defer file.Close()
	defer fileWrite.Close()

	_, err = file.Seek(int64(offset), io.SeekStart)
	if err != nil {
		log.Printf("Offset bigger than size of file!\n")
		return err
	}

	//# иммем на входе 640 байт, нужно прочитать их по 256
	rw_limit := 256
	buf := make([]byte, rw_limit)
	bufOffset := 0
	writed := 0
	for bufOffset < limit {

		readed, err := file.Read(buf)

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Can't read: %v\n", err)
			return err
		}
		if limit-bufOffset > rw_limit {
			writed, err = fileWrite.Write(buf[:readed])
		} else {
			readed = limit - bufOffset
			writed, err = fileWrite.Write(buf[:readed])
		}

		if err != nil {
			log.Printf("Cannot write: %v\n", err)
			return err
		}
		bufOffset += readed
		log.Printf("Readed %v bytes Writed %v bytes Total %v:%v bytes", readed, writed, bufOffset, limit)

		if readed < rw_limit {
			break
		}

	}

	return nil

}

func main() {
	flag.Parse()
	fmt.Printf("%v %v %v %v\n", from, to, offset, limit)
	//err := Copy("from\\1.txt", "to\\6.txt", 0, 400)
	err := Copy(from, to, offset, limit)
	if err != nil {
		log.Fatal(err)
	}
}
