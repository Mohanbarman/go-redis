package commands

import (
	"errors"
	"strconv"
	"time"

	"github.com/Mohanbarman/go-redis/resp"
)

type Options struct {
	TTLMs     int64
	ExpiresAt time.Time
}

func ParseOptions(command string, args []resp.Value) (Options, error) {
	switch command {
	case "SET":
		return parseSetOptions(args)
	default:
		return Options{}, nil
	}
}

func parseSetOptions(args []resp.Value) (options Options, err error) {
	i := 2
	for i < len(args) {
		if args[i].Bulk == "PX" || args[i].Bulk == "EX" {
			if i+1 >= len(args) {
				return Options{}, errors.New("ERR syntax error")
			}
			ms, err := strconv.ParseInt(args[i+1].Bulk, 10, 64)
			if err != nil {
				return Options{}, errors.New("ERR syntax error")
			}
			if args[i].Bulk == "EX" {
				ms *= 1000
			}
			options.ExpiresAt = time.Now().Add(time.Millisecond * time.Duration(ms))
			i += 2
			continue
		}
	}

	return options, nil
}
