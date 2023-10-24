package commands

import (
	"time"

	"github.com/Mohanbarman/redis-clone/resp"
)

func ttl(args []resp.Value, options Options) resp.Value {
	if len(args) < 1 {
		return resp.Value{Typ: "error", Str: "ERR syntax error"}
	}

	key := args[0].Bulk

	SETsMu.RLock()
	defer SETsMu.RUnlock()

	value, ok := SETs[key]
	if !ok {
		return resp.Value{Typ: "null"}
	}

	var ttl int64 = -1
	if !value.expiresAt.IsZero() {
		ttl = value.expiresAt.Sub(time.Now()).Milliseconds()
	}

	if ttl < -2 {
		ttl = -2
	}

	return resp.Value{Typ: "integer", Num: ttl}
}
