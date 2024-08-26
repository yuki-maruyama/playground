package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid"
)

func main() {
	uidMode := flag.String("uid-mode", "uuid", "userid mode")
	usersNum := flag.Uint64("users", 1000, "number of users")
	shardsNum := flag.Uint64("shards", 10, "number of shards")
	flag.Parse()

	shards := make([][]string, *shardsNum)

	for i := uint64(0); i < *usersNum; i++ {
		var uid string
		switch *uidMode {
		case "uuid":
			uid = makeUUIDString()
		case "seq":
			uid = makeSeqString(i)
		case "ulid":
			uid = makeULIDString()
		default:
			panic("unknown uid mode")
		}

		uidHash := sha1.Sum([]byte(uid))
		shard := new(big.Int).SetBytes(uidHash[:]).Uint64() % *shardsNum
		shards[shard] = append(shards[shard], uid)
	}

	// show user id mode and number of users and shards
	fmt.Printf("uid mode: %s, users: %d, shards: %d\n", *uidMode, *usersNum, *shardsNum)
	// for each shard length and percentage and first user
	for i, shard := range shards {
		fmt.Printf("shard %d: %d users (%.2f%%), first user: %s\n", i, len(shard), float64(len(shard))/float64(*usersNum)*100, shard[0])
	}
}

func makeUUIDString() string {
	return uuid.New().String()
}

func makeSeqString(i uint64) string {
	return "user-" + strconv.FormatUint(i, 10)
}

func makeULIDString() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	id := ulid.MustNew(ulid.Now(), r)
	return id.String()
}
