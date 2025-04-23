package request

type Auth struct {
	Node  string `json:"node"`
	Uuid  string `json:"uuid"`
	T     int64  `json:"t"`
	RunId string `json:"runId"`
	Sign  string `json:"sign"`
}

type Auth2 struct {
	Node string `json:"node"`
	Uuid string `json:"uuid"`
	Port string `json:"port"`
	Type string `json:"type"`
	T    int64  `json:"t"`
	Sign string `json:"sign"`
}

type Speed struct {
	RunId string `json:"runId"`
	T     int64  `json:"t"`
	Sign  string `json:"sign"`
}
