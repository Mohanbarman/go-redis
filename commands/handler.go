package commands

import (
	"sync"
	"time"

	"github.com/Mohanbarman/go-redis/resp"
)

var Handlers = map[string]func([]resp.Value, Options) resp.Value{
	"PING":      ping,
	"GET":       get,
	"SET":       set,
	"HGET":      hget,
	"HSET":      hset,
	"HGETALL":   hgetall,
	"TTL":       ttl,
	"PEXPIREAT": pexpireat,
}

type SetValue struct {
	value     string
	expiresAt time.Time
}

var SETs = map[string]*SetValue{}
var SETsMu = sync.RWMutex{}
var HSETs = map[string]map[string]string{}
var HSETsMu = sync.RWMutex{}
