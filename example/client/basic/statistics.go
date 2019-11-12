package basic

//easyjson:json
type Statistics struct {
	Platform string        `json:"plf"`
	Message  ReportMessage `json:"rms"`
}

//easyjson:json
type ReportMessage struct {
	MsgType    int      `json:"mty"`
	BodyType   string   `json:"bt"`
	Time       int64    `json:"tm"`
	TTopics    []string `json:"ttp"`
	TOpIDs     []string `json:"tid"`
	FromOpID   string   `json:"fid"`
	FromTopic  string   `json:"ftp"`
	MsgID      string   `json:"mid"`
	Online     int      `json:"onl"`
	Dev        int      `json:"dev"`
	PacketType int      `json:"pty"`
}
