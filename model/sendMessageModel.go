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

type SendMessageResponseDto struct {
	Count          Count
	CountForCharge CountForCharge
	Balance        BalanceForGroup
	Point          Point
	App            App
	SdkVersion     string
	OsPlatform     string
	Log            []map[string]interface{}
	Status         string
	DateSent       string
	DateCompleted  string
	IsRefunded     bool
	FlagUpdated    bool
	AccountId      string
	ApiVersion     string
	GroupId        string
	Price          map[string]Price
	DateCreated    string
	DateUpdated    string
	Id             string `json:"_id"`
}

type Count struct {
	Total             int
	SentTotal         int
	SentFailed        int
	SentSuccess       int
	SentPending       int
	SentReplacement   int
	Refund            int
	RegisteredFailed  int
	RegisteredSuccess int
}

type CountForCharge struct {
	SMS map[string]int
	LMS map[string]int
	MMS map[string]int
	ATA map[string]int
	CTA map[string]int
	CTI map[string]int
}
type BalanceForGroup struct {
	Requested   float32
	Replacement float32
	Refund      float32
	Sum         float32
}
type Point struct {
	Requested   int
	Replacement int
	Refund      int
	Sum         int
}

type App struct {
	Profit  Profit
	AppId   string
	Version string
}

type Profit struct {
	SMS float32
	LMS float32
	MMS float32
	ATA float32
	CTA float32
	CTI float32
}

type Price struct {
	SMS float32
	LMS float32
	MMS float32
	ATA float32
	CTA float32
	CTI float32
}
