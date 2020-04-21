package models

import (
	"time"
)

//Payment entity database object
type Payment struct {
	ID         string     `json:"id,omitempty" csv:"id,omitempty"`
	TrxnRef    string     `json:"trxnRef,omitempty" csv:"trxnRef,omitempty"`
	SenderID   string     `json:"senderId,omitempty" csv:"senderId,omitempty"`
	ReceiverID string     `json:"receiverId,omitempty" csv:"receiverId,omitempty"`
	Amount     float32    `json:"amount,omitempty" csv:"amount,omitempty"`
	Currency   string     `json:"currency,omitempty" csv:"currency,omitempty"`
	Narration  string     `json:"narration,omitempty" csv:"narration,omitempty"`
	CreatedOn  *time.Time `json:"createdon,omitempty" csv:"createdon,omitempty"`
}
