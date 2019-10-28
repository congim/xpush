package message

type UnRead struct {
	Topics map[string]int64 `json:"topics"`
}

func NewUnRead() *UnRead {
	return &UnRead{
		Topics: make(map[string]int64),
	}
}
