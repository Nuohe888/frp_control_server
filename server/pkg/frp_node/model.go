package frp_node

type Res struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type KillReq struct {
	RunId string `json:"runId"`
	T     int64  `json:"t"`
	Sign  string `json:"sign"`
}

type FlowReq struct {
	Uuid string `json:"uuid"`
	T    int64  `json:"t"`
	Sign string `json:"sign"`
}
