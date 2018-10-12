package server

type TelGroup struct {
	Mobile     string `json:"mobile"`
	Nationcode string `json:"nationcode"`
}

type PhoneNameGroup struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

//SMSResult 短信发送结果
type SMSResult struct {
	Result int    `json:"result"`
	Errmsg string `json:"errmsg"`
	Ext    string `json:"ext"`
	Fee    int    `json:"fee"`
	Sid    string `json:"sid"`
}

//SMSHTTPInfo 短信发送参数
type SMSHTTPInfo struct {
	Ext    string   `json:"ext"`
	Extend string   `json:"extend"`
	Params []string `json:"params"`
	Sig    string   `json:"sig"`
	Sign   string   `json:"sign"`
	Tel    TelGroup `json:"tel"`
	Time   int      `json:"time"`
	TplId  int      `json:"tpl_id"`
}
