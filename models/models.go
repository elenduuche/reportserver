package models

import (
	"time"
)

//Payment entity database object
type Payment struct {
	ID         string     `json:"id,omitempty"`
	TrxnRef    string     `json:"trxnRef,omitempty"`
	SenderID   string     `json:"senderId,omitempty"`
	ReceiverID string     `json:"receiverId,omitempty"`
	Amount     float32    `json:"amount,omitempty"`
	Currency   string     `json:"currency,omitempty"`
	Narration  string     `json:"narration,omitempty"`
	CreatedOn  *time.Time `json:"createdon,omitempty"`
}
