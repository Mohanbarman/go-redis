package main

import (
	"fmt"
	"github.com/Mohanbarman/go-redis/aof"
	"github.com/Mohanbarman/go-redis/commands"
	"github.com/Mohanbarman/go-redis/resp"
	"net"
	"os"
	"strings"
)

func main() {
	port := ":6380"
	if len(os.Args) > 1 {
		port = fmt.Sprintf(":%s", os.Args[1])
	}

	l, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Listening on port", port)

	aof, err := aof.NewAof("db.aof")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer aof.Close()

	err = aof.Read(func(v resp.Value) {
		command := strings.ToUpper(v.Array[0].Bulk)
		args := v.Array[1:]
		handler, ok := commands.Handlers[command]
		options, err := commands.ParseOptions(command, args)
		if err != nil {
			return
		}
		if !ok {
			return
		}
		handler(v.Array[1:], options)
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		defer conn.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		go func() {
			for {
				writer := resp.NewWriter(conn)
				_resp := resp.NewResp(conn)
				value, err := _resp.Read()
				if err != nil {
					fmt.Println(err)
					return
				}

				if value.Typ != "array" {
					fmt.Println("Invalid request")
					continue
				}

				if len(value.Array) == 0 {
					fmt.Println("Invalid request, array is empty")
					continue
				}

				command := strings.ToUpper(value.Array[0].Bulk)
				args := value.Array[1:]
				commandOptions, err := commands.ParseOptions(command, args)

				if err != nil {
					writer.Write(resp.Value{Typ: "error", Str: err.Error()})
					continue
				}

				handler, ok := commands.Handlers[command]
				if !ok {
					fmt.Println("Invalid command ", command)
					writer.Write(resp.Value{Typ: "string", Str: ""})
					continue
				}

				result := handler(args, commandOptions)

				if command == "SET" || command == "HSET" {
					aof.Write(value, commandOptions)
				}

				writer.Write(result)
			}
		}()
	}
}
