package main

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"time"
)

var data = []byte("nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好nihao,你好,大家好")

type PackConn struct {
	*tls.Conn
}

func (s *PackConn) OnData() {
	var msgType uint8
	var length uint32
	var payload []byte
	var err error
	for {
		if err = binary.Read(s, binary.LittleEndian, &msgType); err != nil {
			return
		}
		if err = binary.Read(s, binary.LittleEndian, &length); err != nil {
			return
		}
		payload = make([]byte, length)
		if err = binary.Read(s, binary.LittleEndian, &payload); err != nil {
			return
		}
		fmt.Println(msgType, length, string(payload))
	}
}
func (s *PackConn) Send(typ uint8, payload []byte) (err error) {
	b := bytes.NewBuffer(nil)
	if err = binary.Write(b, binary.LittleEndian, typ); err != nil {
		return
	}
	if err = binary.Write(b, binary.LittleEndian, uint32(len(payload))); err != nil {
		return
	}
	if err = binary.Write(b, binary.LittleEndian, payload); err != nil {
		return
	}
	if _, err = s.Write(b.Bytes()); err != nil {
		return
	}
	return
}

func main() {
	client, err := tls.Dial("tcp4", "127.0.0.1:8888", &tls.Config{InsecureSkipVerify: true})
	if err != nil {
		panic(err)
	}
	conn := PackConn{client}
	go func() {
		time.Sleep(time.Second * 3)
		for i := 0; i < 10; i++ {
			if err = conn.Send(1, data); err != nil {
				return
			}
			time.Sleep(time.Second)
		}
	}()
	conn.OnData()
	fmt.Println("conn disconnected")
}
