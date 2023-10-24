package commands

import (
	"github.com/Mohanbarman/redis-clone/resp"
)

func ping(args []resp.Value, options Options) resp.Value {
	if len(args) == 0 {
		return resp.Value{Typ: "string", Str: "PONG"}
	}
	return resp.Value{Typ: "string", Str: args[0].Bulk}
}
