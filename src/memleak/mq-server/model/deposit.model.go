package model

import (
	"fmt"
	"time"
)

type DepositModel struct {
	DepositID   string
	AccountID   string
	Amount      float64
	DepositDate time.Time
	CreatedAt time.Time
	CreatedBy   string
}

func (d DepositModel) JoinFields() string {
	layout := "2006-01-02 15:04:05"

	// Format the DepositDate and CreatedAt fields using the custom layout
	formattedDepositDate := d.DepositDate.Format(layout)
	return fmt.Sprintf("('%s','%s',%.2f,'%s','%s')",
		d.DepositID, d.AccountID, d.Amount, formattedDepositDate, d.CreatedBy)
}
