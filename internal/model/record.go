package model

type Record struct {
	Timestamp int64 // unique primary key
	URL       string
	Method    string
	IP        string
}
