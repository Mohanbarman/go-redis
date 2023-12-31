# go-redis
A Redis clone built using go uses [RESP](https://redis.io/docs/reference/protocol-spec/) protocol.

## Features
1. Set and get the key value.
2. Set ttl of the key, which will expire after the defined ttl.
3. Persistence using AOF(Append only file).
4. Hash map data structure support `HSET` and `HGET` commands.
5. Accept concurrent connections using goroutines.

## Supported commands
- PING    
- GET     
- SET     
- HGET    
- HSET    
- HGETALL 
- TTL     
- PEXPIREAT

## Setup
```bash
$ docker-compose up
```
This will start the application on port `6380` and it can be connected using standard `redis-cli` redis client
