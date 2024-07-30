package sticky

import (
	"bytes"
	"encoding/binary"
	"io"
)

// 编译编码器
type Encoder struct {
	w io.Writer
}

func NewEncoder(w io.Writer) *Encoder {

	return &Encoder{
		w: w,
	}
}

func (enc *Encoder) Encode(message string) error {
	// 1. 定义长度
	l := int32(len(message))

	buf := new(bytes.Buffer)
	// 写入长度
	// 转为二进制  写入
	if err := binary.Write(buf, binary.BigEndian, l); err != nil {
		return err
	}

	// 写入主体
	// if err := binary.Write(buf, binary.BigEndian, []byte(message)); err != nil { //这个写法 使用这个长度需要新的二进制长度
	// return err
	// }

	// 写入消息内容（以字节形式）
	if _, err := buf.Write([]byte(message)); err != nil {
		return err
	}

	// io.writer 发送
	_, err := enc.w.Write(buf.Bytes())
	return err
}

// 解码

// 编译编码器
type Decoder struct {
	r io.Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		r: r,
	}
}

func (dec *Decoder) Decode(message *string) error {
	// 先读取长度
	header := make([]byte, 4)
	_, err := io.ReadFull(dec.r, header)
	if err != nil {
		return err
	}

	// 转为int
	var length int32
	buf := bytes.NewBuffer(header)
	err = binary.Read(buf, binary.BigEndian, &length)
	if err != nil {
		return err
	}

	body := make([]byte, length)
	if _, err := io.ReadFull(dec.r, body); err != nil {
		return err
	}

	*message = string(body)

	// 读取
	return nil
}
