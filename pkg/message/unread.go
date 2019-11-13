package message

import "encoding/json"

// Unread ...
type Unread struct {
	Topics map[string]bool `json:"topics"`
}

// Encode encode
func (u *Unread) Encode() ([]byte, error) {
	body, err := json.Marshal(u)
	return body, err
}

// Decode decode
func (u *Unread) Decode(body []byte) error {
	err := json.Unmarshal(body, u)
	return err
}

// NewUnread ///
func NewUnread() *Unread {
	return &Unread{
		Topics: make(map[string]bool),
	}
}
