package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
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
type AccessFlowTag struct {
	totallen uint32
	bid      uint32 //业务bid
	site_id  uint32 //站点集合
	uid      []byte
	uidtype  uint8
	trace_id []byte
}

func (access *AccessFlowTag) TagEnCode() error {
	str := "hzh"
	access.uid = []byte(str)

	access.bid = 22222
	access.site_id = 55555
	access.uidtype = byte('3')
	str = "tracechage"
	access.trace_id = []byte(str)
	access.totallen = uint32(unsafe.Sizeof(access.bid))*5 + uint32(len(access.uid)) +
		uint32(unsafe.Sizeof(access.uidtype)) +
		uint32(len(access.trace_id))
	fmt.Println("access len:", access.totallen)

	return errors.New("sdf")
}
func (access *AccessFlowTag) TagWrite(buf *bytes.Buffer) error {

	if err := binary.Write(buf, binary.BigEndian, access.totallen); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, access.bid); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, access.site_id); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, access.uid); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, access.uidtype); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, access.trace_id); err != nil {
		return err
	}
	return errors.New("sdf")
}

type AdposInfo struct {
	totallen     uint32
	adposid      uint32
	siteid       uint32
	version      uint32
	sub_version  uint32
	file_version uint32
	adidnum      uint32

	relfect_adpos uint32
}

func (info *AdposInfo) InfoEnCode() error {
	info.totallen = 36
	info.adposid = 22222
	info.siteid = 55555
	info.version = 104
	info.sub_version = 1
	info.file_version = 0
	info.adidnum = 1
	info.relfect_adpos = 1
	fmt.Println("access len:", info.totallen)

	return errors.New("sdf")
}
func (info *AdposInfo) InfoWrite(buf *bytes.Buffer) error {
	if err := binary.Write(buf, binary.BigEndian, info.totallen); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.adposid); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.siteid); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.version); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.sub_version); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.file_version); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.adidnum); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, info.relfect_adpos); err != nil {
		return err
	}
	fmt.Println("access len:", info.totallen)

	return errors.New("sdf")
}
func (c *ClientCodec) GetReqbuf() ([]byte, error) {
	// 开始打包
	buf := bytes.NewBuffer(make([]byte, 0, 1024*1024*5))
	//	str := reqHead{recommending: "testst"}
	tag := AccessFlowTag{}
	tag.TagEnCode()
	info := AdposInfo{}
	info.InfoEnCode()
	//	strbyte := []byte(str.recommending)
	msgmethod := uint32(102)
	//magicnum := uint16(256)
	msgtype := uint8(1)

	// msgtype := uint8(1)
	uuid := uint32(12324)
	sessionid := uint32(33312324)
	msgid := uint32(6666)
	var pbbuf []byte
	str := "hell rank2.0"
	pbbuf = []byte(str)
	//totalLen := uint32(unsafe.Sizeof(msgtype)) + uint32(unsafe.Sizeof(magicnum)) +
	//	uint32(unsafe.Sizeof(msgmethod)) + uint32(unsafe.Sizeof(uuid))
	totalLen := uint32(unsafe.Sizeof(msgtype)) +
		uint32(unsafe.Sizeof(msgmethod)) + uint32(unsafe.Sizeof(uuid)) + uint32(unsafe.Sizeof(sessionid)) +
		uint32(unsafe.Sizeof(msgid)) + tag.totallen + info.totallen + uint32(len(pbbuf))
	fmt.Println("totallen:", totalLen)

	if err := binary.Write(buf, binary.BigEndian, totalLen); err != nil {
		return nil, err
	}
	//if err := binary.Write(buf, binary.BigEndian, magicnum); err != nil {
	//	return nil, err
	//}
	if err := binary.Write(buf, binary.BigEndian, msgtype); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, msgmethod); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, uuid); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, sessionid); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, msgid); err != nil {
		return nil, err
	}
	if err := tag.TagWrite(buf); err != nil {
		return nil, err
	}
	if err := info.InfoWrite(buf); err != nil {
		return nil, err
	}
	if err := binary.Write(buf, binary.BigEndian, pbbuf); err != nil {
		return nil, err
	}
	fmt.Println("buf len:", buf.Len())
	return buf.Bytes(), nil
}
