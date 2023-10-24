package commands

import (
	"github.com/Mohanbarman/go-redis/resp"
)

func hget(args []resp.Value, options Options) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hget' command"}
	}

	hkey := args[0].Bulk
	key := args[1].Bulk

	HSETsMu.RLock()
	value, ok := HSETs[hkey][key]
	HSETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: "null"}
	}

	return resp.Value{Typ: "bulk", Bulk: value}
}
