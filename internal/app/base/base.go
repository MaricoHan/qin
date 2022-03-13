package base

import (
	"bytes"
	"encoding/binary"
)

func Int2Hex(num int64) []byte {
	buff := new(bytes.Buffer)
	binary.Write(buff, binary.BigEndian, num)
	return buff.Bytes()
}
