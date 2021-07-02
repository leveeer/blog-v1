package enum

type ChatTypeEnum uint8

const (
	OnlineCount ChatTypeEnum = iota + 1
	HistoryRecord
	SendMessage
	RecallMessage
	VoiceMessage
	HeartBeat
)

var (
	ChatTypeDesc = map[ChatTypeEnum]string{
		OnlineCount:   "在线人数",
		HistoryRecord: "历史记录",
		SendMessage:   "发送消息",
		RecallMessage: "撤回消息",
		VoiceMessage:  "语音消息",
		HeartBeat:     "心跳消息",
	}

	ChatType = map[ChatTypeEnum]uint8{
		OnlineCount:   1,
		HistoryRecord: 2,
		SendMessage:   3,
		RecallMessage: 4,
		VoiceMessage:  5,
		HeartBeat:     6,
	}
)

func (receiver ChatTypeEnum) GetChatTypeDesc() string {
	return ChatTypeDesc[receiver]
}

func (receiver ChatTypeEnum) GetChatType() uint8 {
	return ChatType[receiver]
}
