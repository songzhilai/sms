package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

//HTTPRequest 发送http请求
func (s *Server) HTTPRequest(method string, url string, data interface{}) ([]byte, int) {
	var datas *bytes.Buffer
	rejson, err := json.Marshal(data)
	if err != nil {
		s.logger.Error(err)
		return []byte{}, -4
	}
	datas = bytes.NewBuffer(rejson)
	req, err := http.NewRequest(method, url, datas)
	// req.Header.Set("content-type", "application/json")
	// ba64authstr := base64.StdEncoding.EncodeToString([]byte(libs.ConfigValue("FrpsUser") + ":" + libs.ConfigValue("FrpsPasswd")))
	// req.Header.Set("Authorization", "Basic "+ba64authstr)
	if err != nil {
		s.logger.Error(err)
		return []byte{}, -1
	}
	Client := &http.Client{Timeout: 5 * time.Second}
	resp, err := Client.Do(req)
	if err != nil {
		s.logger.Error(err)
		return []byte{}, -2
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Error(err)
		return []byte{}, -3
	}

	return body, 1
}
