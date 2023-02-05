package main

import (
	"context"
	"juejin/boot"
)

func main() {

	boot.ViperSetup()
	boot.LoggerSetup()
	boot.MysqlDBSetup()
	boot.RedisSetup()
	boot.SnowFlakeSetup()
	boot.ExecuteCron(context.Background())
	boot.ServerSetup()
}
