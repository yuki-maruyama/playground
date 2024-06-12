package snowflake

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"time"
)

// id spec
// Snowflake ID like format
// first 1 bit: 0
// next 41 bits: milisecond level timestamp
// next 10 bits: random or node id
// next 12 bits: sequence number in millis

type snowflake struct {
	now    int64
	seq    int16
	nodeId *int64
}

func New(nodeId *int64) (*snowflake, error) {
	if nodeId != nil && *nodeId > 1024 || *nodeId < 0 {
		return nil, errors.New("node ID must be between 0 and 1023")
	}
	return &snowflake{
		now:    time.Now().UnixMilli() << 22,
		seq:    1,
		nodeId: nodeId,
	}, nil
}

var timestampBitMask = int64(0x7FFFFFFFFFC00000)
var randomIdBitMask = int64(0x3FF000)
var seqNoBitMask = int64(0xFFF)

func (s *snowflake) Gen(t *time.Time) *int64 {
	var timestamp int64
	if t != nil {
		timestamp = t.UnixMilli() << 22
	} else {
		timestamp = time.Now().UnixMilli() << 22
	}

	var nid int64
	if s.nodeId != nil {
		nid = *s.nodeId
	} else {
		rid, err := rand.Int(rand.Reader, big.NewInt(1024))
		if err != nil {
			return nil
		}
		nid = rid.Int64()
	}

	if s.now == timestamp {
		s.seq++
	} else {
		s.seq = 1
		s.now = timestamp
	}

	var id int64 = 0b0 + timestamp + (nid << 12) + (int64(s.seq) & seqNoBitMask)
	return &id
}

func (s *snowflake) ShowTimestamp(id int64) {
	timestamp := id & timestampBitMask >> 22
	fmt.Println("Timestamp: ", time.UnixMilli(timestamp))
}

func (s *snowflake) ShowRandom(id int64) {
	random := id & randomIdBitMask >> 12
	fmt.Println("Random: ", random)
}

func (s *snowflake) ShowSeqNo(id int64) {
	seqNo := id & seqNoBitMask
	fmt.Println("SeqNo: ", seqNo)
}

func (s *snowflake) CompareNewer(a int64, b int64) int64 {
	timestampA := a & timestampBitMask
	timestampB := b & timestampBitMask
	if timestampA != timestampB {
		if timestampA > timestampB {
			return a
		} else {
			return b
		}
	} else {
		seqNoA := a & seqNoBitMask
		seqNoB := b & seqNoBitMask
		if seqNoA > seqNoB {
			return a
		} else {
			return b
		}
	}
}
