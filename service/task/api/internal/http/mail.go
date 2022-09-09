/**
 * @Author: Sun
 * @Description:
 * @File:  mail
 * @Version: 1.0.0
 * @Date: 2022/9/9 16:20
 */

package http

import (
	"bytes"
	"cron/service/task/api/internal/global"
	"cron/service/task/api/internal/message"
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"mime"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// GetToken 获取企业微信Access_Token
func GetToken(CorPid, CorpSecret string) {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s", CorPid, CorpSecret)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("get access token err:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var respBody message.ResponseAccessToken
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		fmt.Println("unmarshaler json err:", err)
		return
	}
	global.AccessToken = respBody.AccessToken
	time.Sleep(2 * time.Hour)
	//return respBody.AccessToken, nil
}

// SendMessage  给多人同时发送单条消息
func SendMessage(accessToken string, toUser interface{}, content interface{}, agentID int64) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token=%s&debug=1", accessToken)

	//fmt.Println(content)
	req := map[string]interface{}{
		"touser":  toUser,
		"msgtype": "text",
		"agentid": agentID,
		"text": map[string]interface{}{
			"content": content,
		},
	}
	// fmt.Println(accessToken)
	byte, err := json.Marshal(req)
	if err != nil {
		fmt.Println("json marshal err:", err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(byte))
	if err != nil {
		fmt.Println("send message err:", err)
		return err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	return nil
}

func SendMail(host, username, password, mailTo, serverName, subject, filePath, fileName string) error {
	fmt.Println(host, username, password, mailTo, serverName, subject, filePath, fileName)
	mail := gomail.NewMessage()
	mail.SetHeader("From", mail.FormatAddress(username, serverName))
	mail.SetHeader("To", mailTo)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/html", "")
	if len(filePath) > 0 {
		mail.Attach(filePath,
			gomail.Rename(fileName),
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": {
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", fileName)),
				},
			}),
		)
	}

	hp := strings.Split(host, ":")
	port, _ := strconv.Atoi(hp[1])
	dialer := gomail.NewDialer(hp[0], port, username, password)
	return dialer.DialAndSend(mail)
}
