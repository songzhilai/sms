package server

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Luxurioust/excelize"
)

//开始发送短信
func (s *Server) startSms() {
	filename := fmt.Sprintf("conf/%s", s.opts.ExcelName)
	xlsx, err := excelize.OpenFile(filename)
	if err != nil {
		s.logger.Error(err)
		os.Exit(0)
	}
	sheetname := strings.Split(s.opts.ExcelSheetName, ",")
	if len(sheetname) <= 0 {
		s.logger.Error("sheet读取失败")
		os.Exit(0)
	}
	var rowss []PhoneNameGroup
	for _, sv := range sheetname {
		rows := xlsx.GetRows(sv)
		for k, v := range rows {
			if k == 0 {
				if v[0] != "姓名" && v[1] != "电话" {
					s.logger.Error("sheet:" + sv + "的格式不对，不在短信发送范围内")
					break
				}
			} else {
				if v[0] != "" && v[1] != "" {
					var row PhoneNameGroup
					row.Name = v[0]
					row.Phone = v[1]
					rowss = append(rowss, row)
				}
			}

		}
	}
	fmt.Println(rowss)
	if len(rowss) <= 0 {
		s.logger.Info("没有需要发送的短信")
		os.Exit(0)
	}
	if len(rowss) > 100 {
		s.logger.Info("要发的人太多了，超过100个了")
	}
	s.SendSms(rowss)
	os.Exit(0)

}

//SendSms 发送短信
func (s *Server) SendSms(rows []PhoneNameGroup) {
	var smsinfo SMSHTTPInfo
	smsinfo.TplId = SMSTPLID
	nowtime := time.Now().Unix()
	smsinfo.Time = int(nowtime)
	smsinfo.Sign = "退休通知"
	randnum := GenerateRandnum()
	httpurl := fmt.Sprintf("https://yun.tim.qq.com/v5/tlssmssvr/sendsms?sdkappid=%s&random=%d", APPID, randnum)
	for _, v := range rows {
		var trow TelGroup
		trow.Mobile = v.Phone
		trow.Nationcode = "86"
		smsinfo.Sig = getSig(int(nowtime), randnum, v.Phone)
		smsinfo.Params = []string{v.Name, s.opts.SmsContent}
		smsinfo.Tel = trow
		res, sta := s.HTTPRequest("POST", httpurl, smsinfo)
		if sta < 0 {
			s.logger.Error(v.Name + " 发送失败")
			continue
		}
		var result SMSResult
		err := json.Unmarshal(res, &result)
		if err != nil {
			s.logger.Error(v.Name+" 发送失败", err, res)
			continue
		}
		if result.Result != 0 {
			s.logger.Error(v.Name+" 发送失败", result.Errmsg)
		} else {
			s.logger.Info(v.Name + " :" + v.Phone + " 发送成功")
		}
	}
	time.Sleep(1 * time.Second)
}

func getSig(nowtime, randnumint int, mobile string) string {
	s := fmt.Sprintf("appkey=%s&random=%d&time=%d&mobile=%s", APPKEY, randnumint, nowtime, mobile)
	fmt.Println(s)
	h := sha256.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// GenerateRandnum 生成随机数
func GenerateRandnum() int {
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(10000000)
	return randNum
}
