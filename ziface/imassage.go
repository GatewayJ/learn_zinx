package ziface

type IMessage interface {
	// 获取消息的ID
	GetMsgID() uint32
	// 获取消息长度
	GetMsgLen() uint32
	// 获取消息的内容
	GetData() []byte

	// 设置消息id
	SetMsgID(ID uint32)
	// 设置消息长度
	SetDataLen(Len uint32)
	// 设置消息内容
	SetData(Data []byte)
}
