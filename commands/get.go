package commands

import (
	"time"
	"github.com/Mohanbarman/redis-clone/resp"
)

func get(args []resp.Value, options Options) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'get' command"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	value, ok := SETs[key]
	SETsMu.RUnlock()

	if !ok {
		return resp.Value{Typ: "null"}
	}

  if value.expiresAt.Sub(time.Now()) < 0 {
    delete(SETs, key)
    return resp.Value{Typ: "null"}
  }
  
	return resp.Value{Typ: "bulk", Bulk: value.value}
}
