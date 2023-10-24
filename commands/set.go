package commands

import (
	"github.com/Mohanbarman/go-redis/resp"
)

func set(args []resp.Value, options Options) resp.Value {
	if len(args) < 2 {
		return resp.Value{Typ: "error", Str: "ERR wrong number of arguments for 'set' command"}
	}

	key := args[0].Bulk
	value := args[1].Bulk

	_value := &SetValue{value: value}
  if (!options.ExpiresAt.IsZero()) {
    _value.expiresAt = options.ExpiresAt
  }

	SETsMu.Lock()
	SETs[key] = _value
	SETsMu.Unlock()

	return resp.Value{Typ: "string", Str: "OK"}
}
