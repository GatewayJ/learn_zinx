package znet

import (
	"bytes"
	"encoding/binary"
	"zinx/ziface"
)

type DataPack struct {
}

func NewDataPack() *DataPack {
	return &DataPack{}
}

func (pd *DataPack) GetHeadLen() uint32 {
	return 8
}

func (pd *DataPack) Pack(msg ziface.IMessage) ([]byte, error) {

	dataByte := bytes.NewBuffer([]byte{})
	if err := binary.Write(dataByte, binary.LittleEndian, msg.GetMsgLen()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataByte, binary.LittleEndian, msg.GetMsgID()); err != nil {
		return nil, err
	}
	if err := binary.Write(dataByte, binary.LittleEndian, msg.GetData()); err != nil {
		return nil, err
	}
	return dataByte.Bytes(), nil
}

func (pd *DataPack) Unpack(binaryData []byte) (ziface.IMessage, error) {

	msg := &Message{}
	dataBuff := bytes.NewReader(binaryData)
	if err := binary.Read(dataBuff, binary.LittleEndian, msg.DataLen); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, msg.ID); err != nil {
		return nil, err
	}
	if err := binary.Read(dataBuff, binary.LittleEndian, msg.Data); err != nil {
		return nil, err
	}
	return msg, nil
}
