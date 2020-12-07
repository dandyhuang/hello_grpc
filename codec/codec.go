package codec

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"unsafe"
)

const TrpcMagic_TRPC_MAGIC_VALUE = 0x21
const (
	// 起始魔数
	stx = uint16(TrpcMagic_TRPC_MAGIC_VALUE)
	// 帧头 2 bytes stx + 1 byte type + 1 byte status + 4 bytes total len
	// + 2 bytes pb header len + 2 bytes stream id + 4 bytes reserved
	frameHeadLen = uint16(16)

	DyeingKey   = "trpc-dyeing-key" // 染色key
	UserIP      = "trpc-user-ip"    // 客户端最上游ip
	EnvTransfer = "trpc-env"        // 透传环境数据
)

// ClientCodec rpc客户端编解码
type ClientCodec struct {
	RequestID uint32 //全局唯一request id
}

//
type reqHead struct {
	recommending string
}

func (c *ClientCodec) GetReqbuf() ([]byte, error) {
	//	str := reqHead{recommending: "testst"}

	//	strbyte := []byte(str.recommending)
	body := uint32(1022)
	magicnum := uint16(256)
	msgtype := byte('1')
	totalLen := uint32(unsafe.Sizeof(msgtype)) + uint32(unsafe.Sizeof(magicnum)) + uint32(unsafe.Sizeof(body))
	fmt.Println("totallen:", totalLen)

	// 开始打包
	buf := bytes.NewBuffer(make([]byte, 0, totalLen))
	if err := binary.Write(buf, binary.BigEndian, totalLen); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, magicnum); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, msgtype); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, body); err != nil {
		return nil, err
	}
	fmt.Println("buf len:", buf.Len())
	return buf.Bytes(), nil
}
