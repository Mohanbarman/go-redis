package commands

import (
	"strconv"
	"time"

	"github.com/Mohanbarman/go-redis/resp"
)

func pexpireat(args []resp.Value, options Options) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	msInt, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		return resp.Value{Typ: "error", Str: "ERR syntax error"}
	}

	SETsMu.RLock()
	val, ok := SETs[key]
	if !ok {
		return resp.Value{Typ: "integer", Num: 0}
	}
	val.expiresAt = time.Unix(0, msInt*int64(time.Millisecond))
	SETsMu.RUnlock()

	return resp.Value{Typ: "string", Str: args[0].Bulk}
}
