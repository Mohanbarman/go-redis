package commands

import (
	"github.com/Mohanbarman/redis-clone/resp"
)

func hgetall(args []resp.Value, options Options) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'hgetall' command"}
	}

	key := args[0].Bulk

	HSETsMu.Lock()
	defer HSETsMu.Unlock()

	result := resp.Value{Typ: "array"}
	hmap, ok := HSETs[key]
	if !ok {
		return resp.Value{Typ: "null"}
	}

	for k := range hmap {
		result.Array = append(
			result.Array,
			resp.Value{Typ: "bulk", Bulk: k}, resp.Value{Typ: "bulk", Bulk: hmap[k]})
	}

	return result
}
