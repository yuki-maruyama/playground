package main

import (
	"fmt"
	"log"

	"github.com/yuki-maruyama/playground/go-id-gen/snowflake"
)

func main() {
	nodeId := int64(0)
	snowflake, err := snowflake.New(&nodeId)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 2; i++ {
		snowFlakeID := snowflake.Gen()
		fmt.Printf("SnowflakeID: %d\n", snowFlakeID)
		snowflake.ShowTimestamp(snowFlakeID)
		snowflake.ShowRandom(snowFlakeID)
		snowflake.ShowSeqNo(snowFlakeID)
	}
}
