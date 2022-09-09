/**
 * @Author: Sun
 * @Description:
 * @File:  msg_resp
 * @Version: 1.0.0
 * @Date: 2022/9/9 16:25
 */

package message

type ResponseAccessToken struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
