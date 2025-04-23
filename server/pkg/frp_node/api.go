package frp_node

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type api struct{}

var Api *api

func init() {
	Api = &api{}
}

func (*api) Kill(url, key, runId string) error {
	t := time.Now().Unix()

	v := map[string]string{
		"runId": runId,
		"t":     fmt.Sprintf("%d", t),
	}
	sign := GenerateSign(v, key)

	req := KillReq{
		RunId: runId,
		T:     t,
		Sign:  sign,
	}

	reqBody, err := json.Marshal(req)
	if err != nil {
		return errors.New("服务端出错")
	}

	request, err := http.NewRequest("POST",
		fmt.Sprintf("%s%s", url, "/api/kill"),
		bytes.NewBuffer(reqBody))
	if err != nil {
		return errors.New("服务端出错")
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return errors.New("服务端出错")
	}
	defer response.Body.Close()

	var res Res

	readAll, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("服务端出错")
	}
	err = json.Unmarshal(readAll, &res)
	if err != nil {
		return errors.New("服务端出错")
	}

	if res.Code != 200 {
		return errors.New(res.Msg)
	}

	return nil
}

func (*api) Flow() error {
	return nil
}
