package trading212

import "time"

const (
	// localTimeFmt set time format to utc+0
	localTimeFmt = "2006-01-02 15:04:05 +0000 UTC"
)

// Statistics is the Statistics of the bill file.
type Statistics struct {
	UserID          string    `json:"user_id,omitempty"`
	Username        string    `json:"username,omitempty"`
	ParsedItems     int       `json:"parsed_items,omitempty"`
	Start           time.Time `json:"start,omitempty"`
	End             time.Time `json:"end,omitempty"`
	TotalInRecords  int       `json:"total_in_records,omitempty"`
	TotalInMoney    float64   `json:"total_in_money,omitempty"`
	TotalOutRecords int       `json:"total_out_records,omitempty"`
	TotalOutMoney   float64   `json:"total_out_money,omitempty"`
}

// Order is a single order.
type Order struct {
	Action           string
	Time             time.Time `json:"payTime,omitempty"` // 交易时间
	Money            float64   `json:"money,omitempty"`   // 金额
	Currency         string
	Notes            string
	ID               string
	Peer             string `json:"peer,omitempty"` // merchant name
	MerchantCategory string
}

