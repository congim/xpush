package basic

// CheckUsersMsg check user msg
type CheckUsersMsg struct {
	System  int    `msgp:"sys"`
	Session string `msgp:"s"`
}

// BatchCheckMsg batch check msg
type BatchCheckMsg struct {
	Msgs []*CheckUsersMsg `msgp:"msgs"`
}

// LineMsg 在线结果返回
type LineMsg struct {
	List []bool
}
