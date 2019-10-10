package sms

import (
	"github.com/buger/jsonparser"
	"github.com/henrylee2cn/goutil"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var RequestError = errors.New("request error")

type LexinClient struct {
	User     string
	Password string
	SignTile string
}

func NewLexinClient(user, password, signTitle string) *LexinClient {
	if !strings.HasPrefix(signTitle, "【") {
		signTitle = "【" + signTitle
	}
	if !strings.HasSuffix(signTitle, "】") {
		signTitle = signTitle + "】"
	}
	c := &LexinClient{
		User:     user,
		SignTile: signTitle,
	}
	s := goutil.Md5([]byte(password))
	s = strings.ToUpper(s)
	c.Password = s

	return c
}
func (c *LexinClient) Send(phones string, content string) error {
	bizId := time.Now().Format("20160102150405")
	content = content + c.SignTile
	params := url.Values{
		"accName":  {c.User},
		"accPwd":   {c.Password},
		"aimcodes": {phones},
		"content":  {content},
		"bizId":    {bizId},
		"dataType": {"json"},
	}

	resp, err := http.PostForm("http://sdk.lx198.com/sdk/send", params)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return RequestError
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	replyCode, err := jsonparser.GetInt(bytes, "replyCode")
	if err != nil {
		return err
	}
	if replyCode != 1 {
		replyMsg, err := jsonparser.GetString(bytes, "replyMsg")
		if err != nil {
			return err
		}
		return errors.New(replyMsg)
	}
	return nil
}
