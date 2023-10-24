package commands

import (
	"github.com/Mohanbarman/go-redis/resp"
)

func hset(args []resp.Value, options Options) resp.Value {
	if len(args) < 3 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hset' command"}
	}

	hkey := args[0].Bulk
	key := args[1].Bulk
	value := args[2].Bulk

	HSETsMu.Lock()
	_, ok := HSETs[hkey]
	if !ok {
		HSETs[hkey] = make(map[string]string)
	}
	HSETs[hkey][key] = value
	HSETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}
