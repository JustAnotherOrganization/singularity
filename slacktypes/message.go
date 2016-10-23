package slacktypes

type Message struct {
	Type      string `json:"type"`
	Channel   string `json:"channel"`
	User      string `json:"user"`
	Text      string `json:"text"`
	TimeStamp string `json:"ts"`
	SubType   string `json:"subtype"`
}
