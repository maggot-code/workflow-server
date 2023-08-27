/*
 * @FilePath: /workflow-server/internal/data/wechat.go
 * @Author: maggot-code
 * @Date: 2023-08-14 20:04:55
 * @LastEditors: maggot-code
 * @LastEditTime: 2023-08-16 17:34:50
 * @Description:
 */
package data

import (
	"fmt"
	"net/url"
)

var code2Session = "https://api.weixin.qq.com/sns/jscode2session"

func buildCodeToSession(ur *userRepo, code string) (string, error) {
	query := url.Values{}
	query.Add("grant_type", "authorization_code")
	query.Add("appid", ur.conf.Wechat.Appid)
	query.Add("secret", ur.conf.Wechat.Secret)
	query.Add("js_code", code)

	uri, err := url.Parse(code2Session)
	if err != nil {
		return "", fmt.Errorf("wechat: code2Session url parse fail; %w", err)
	}

	uri.RawQuery = query.Encode()
	return uri.String(), nil
}
