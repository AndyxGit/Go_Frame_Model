package structure

import "time"

type RequestHeaderStruct struct {
	SessionID string
	ChannelID string
	UserID    string
	Number    int
	Date      time.Time
}
