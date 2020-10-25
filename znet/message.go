package znet

type Message struct {
	ID      uint32 // 消息的ID
	DataLen uint32 // 消息的长度
	Data    []byte // 消息的内容
}

// 获取消息的ID
func (m *Message) GetMsgID() uint32 {
	return m.ID
}

// 获取消息长度
func (m *Message) GetMsgLen() uint32 {
	return m.DataLen
}

// 获取消息的内容

func (m *Message) GetData() []byte {
	return m.Data
}

// 设置消息id
func (m *Message) SetMsgID(ID uint32) {
	m.ID = ID
}

// 设置消息长度
func (m *Message) SetDataLen(Len uint32) {
	m.DataLen = Len
}

// 设置消息内容
func (m *Message) SetData(Data []byte) {
	m.Data = Data
}
