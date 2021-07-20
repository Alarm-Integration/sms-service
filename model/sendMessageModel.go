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

type SendRequestDto struct {
	Message SendMessageDto `json:"message"`
}

type SendMessageDto struct {
	To   string `json:"to"`
	From string `json:"from"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type TopicMessageDto struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Sender    string   `json:"sender"`
	TraceID   string   `json:"traceId"`
	UserID    int      `json:"userId"`
	Addresses []string `json:"addresses"`
}

type SendMessageResponseDto struct {
	GroupId       string
	MessageId     string
	AccountId     string
	StatusMessage string
	StatusCode    string
	To            string
	From          string
	Type          string
	Country       string
}

type SendMessageFailResponseDto struct {
	ErrorCode    string
	ErrorMessage string
}
