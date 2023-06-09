package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type HeaderHandlers func(*Curl)

func FormHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
}

func JsonHeader() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Content-Type", "application/json")
	}
}

func UserAgent() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36")
	}
}

func Accept() HeaderHandlers {
	return func(c *Curl) {
		c.request.Header.Add("Accept", "*/*")
	}
}

func SetHeader(key, val string) HeaderHandlers {
	return func(c *Curl) {
		if c.debug {
			fmt.Println()
			fmt.Println("===> header：", key, val)
			fmt.Println()
		}
		c.request.Header.Add(key, val)
	}
}

func Authorization(token string) HeaderHandlers {
	return func(c *Curl) {
		if c.debug {
			fmt.Println()
			fmt.Println("===> token：", token)
			fmt.Println()
		}
		c.request.Header.Add("Authorization", token)
	}
}

func (c *Curl) JsonData(data interface{}) (*Curl, error) {
	bs, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	if c.debug {
		fmt.Println()
		fmt.Println("===> 请求报文：", string(bs))
		fmt.Println()
	}
	c.data = bytes.NewReader(bs)
	return c, nil
}

func (c *Curl) QueryData(data map[string]string) *Curl {
	c.data = strings.NewReader(HttpBuildQuery(data))
	return c
}

func (c *Curl) Data(data *bytes.Buffer) *Curl {
	c.data = data
	return c
}

func HttpBuildQuery(params map[string]string) string {
	v := make(url.Values)
	for key, val := range params {
		v.Set(key, val)
	}
	return v.Encode()
}
