package basic

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// UUIDResult uuid result
//easyjson:json
type UUIDResult struct {
	Suc bool   `json:"suc"`
	Val string `json:"val"`
}

// UUID uuid server
type UUID struct {
	GroupURL string
	MsgIDURL string
}

// GroupID get group id
func (uid *UUID) GroupID() (string, error) {
	return uid.getuuid(uid.GroupURL)
}

// MsgID get MsgID id
func (uid *UUID) MsgID() (string, error) {
	msgid, err := uid.getuuid(uid.MsgIDURL)
	if err != nil {
		fmt.Println("[ERROR] [MsgID] - get msgid failed, err message is", err)
	}
	return msgid, err
}

func (uid *UUID) getuuid(url string) (string, error) {
	ur := UUIDResult{}
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = ur.UnmarshalJSON(body)
	if err != nil {
		return "", err
	}
	return ur.Val, err
}

// NewUUID new uuid class point
func NewUUID(uidaddr string) *UUID {
	GroupURL := "http://" + uidaddr + "?type=increment&group=zeus"
	MsgIDURL := "http://" + uidaddr
	u := &UUID{GroupURL: GroupURL, MsgIDURL: MsgIDURL}
	return u
}
