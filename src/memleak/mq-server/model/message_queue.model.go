package model

import "time"

type MessageQueueModel struct {
	ID          string    // Assuming UUID is represented as a string
	OwnerSystem string
	Topic       string
	Data     string // Assuming LZ4 codec is handled externally
	CreatedAt   time.Time
	CreatedBy   string // Assuming FixedString(36) is handled externally
	Remarks     string // Assuming FixedString(255) is handled externally
	StatusID    uint8
	Sign        int8
}