package aof

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sync"
	"time"

	"github.com/Mohanbarman/redis-clone/commands"
	"github.com/Mohanbarman/redis-clone/resp"
)

type Aof struct {
	file  *os.File
	rd    *bufio.Reader
	mutex sync.Mutex
}

func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: f,
		rd:   bufio.NewReader(f),
	}

	go func() {
		for {
			aof.mutex.Lock()
			aof.file.Sync()
			aof.mutex.Unlock()
			time.Sleep(time.Second)
		}
	}()

	return aof, err
}

func (aof *Aof) Close() error {
	aof.mutex.Lock()
	defer aof.mutex.Unlock()
	return aof.file.Close()
}

func (aof *Aof) Write(value resp.Value, options commands.Options) error {
	aof.mutex.Lock()
	defer aof.mutex.Unlock()

	_, err := aof.file.Write(value.Marshal())
	if err != nil {
		return err
	}

  // setting the expiry of key in aof file
	if !options.ExpiresAt.IsZero() && value.Array[0].Bulk == "SET" {
    expireAtCommand := resp.Value{
      Typ: "array",
      Array: []resp.Value{
        {
          Typ: "bulk",
          Bulk: "PEXPIREAT",
        },
        {
          Typ: "bulk",
          Bulk: value.Array[1].Bulk,
        },
        {
          Typ: "bulk",
          Bulk: fmt.Sprint(options.ExpiresAt.UnixMilli()),
        },
      },
    }
    aof.file.Write(expireAtCommand.Marshal())
	}

	return nil
}

func (aof *Aof) Read(callback func(resp.Value)) error {
	aof.mutex.Lock()
	defer aof.mutex.Unlock()

	resp := resp.NewResp(aof.file)

	for {
		value, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		callback(value)
	}

	return nil
}
