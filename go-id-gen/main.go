package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yuki-maruyama/playground/go-id-gen/snowflake"
)

func main() {
	nodeId := int64(0)
	snowflake, err := snowflake.New(&nodeId)
	if err != nil {
		log.Fatalln(err)
	}
	for i := 0; i < 100; i++ {
		t := time.Unix(1560000000, 0)
		snowFlakeID := *snowflake.Gen(&t)
		fmt.Printf("SnowflakeID: %d\n", snowFlakeID)
		snowflake.ShowTimestamp(snowFlakeID)
		snowflake.ShowRandom(snowFlakeID)
		snowflake.ShowSeqNo(snowFlakeID)
	}
}
