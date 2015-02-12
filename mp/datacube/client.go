// @description wechat2 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/philsong/wechat2 for the canonical source repository
// @license     https://github.com/philsong/wechat2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package datacube

import (
	"net/http"

	"github.com/philsong/wechat2/mp"
)

type Client struct {
	mp.WechatClient
}

// 创建一个新的 Client.
//  如果 HttpClient == nil 则默认用 http.DefaultClient
func NewClient(TokenServer mp.TokenServer, HttpClient *http.Client) *Client {
	if TokenServer == nil {
		panic("TokenServer == nil")
	}
	if HttpClient == nil {
		HttpClient = http.DefaultClient
	}

	return &Client{
		WechatClient: mp.WechatClient{
			TokenServer: TokenServer,
			HttpClient:  HttpClient,
		},
	}
}
