package model

type messageType int

const (
	SMS messageType = iota
	LMS
	MMS
)

func (d messageType) String() string {
	return [...]string{"SMS", "LMS", "MMS"}[d]
}

type RequestBody struct {
	Messages []SendMessageDto `json:"messages"`
}

type SendMessageDto struct {
	To   string `json:"to"`
	From string `json:"from"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type TopicMessageDto struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Sender  string   `json:"sender"`
	TraceID string   `json:"traceId"`
	UserID  int      `json:"userId"`
	Raws    []string `json:"raws"`
}
