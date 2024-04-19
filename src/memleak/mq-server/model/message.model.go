package model
import "time"

type MessageModel struct {
	CreatedAt time.Time `json:"createdAt"`
	Text      string `json:"text"`
}