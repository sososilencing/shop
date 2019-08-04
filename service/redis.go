package service

import (
	 "github.com/gomodule/redigo/redis"
)

func getRedis()  {
	redis.Dial("tcp", "127.0.0.1:6379")
}


