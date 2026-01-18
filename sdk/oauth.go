package oauth

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/big"
	"net/http"
	"net/url"
)

type Oauth struct {
	apiURL   string
	appID    string
	appKey   string
	callback string
}

func NewOauth(apiURL, appID, appKey, callback string) *Oauth {
	return &Oauth{
		apiURL:   apiURL + "connect.php",
		appID:    appID,
		appKey:   appKey,
		callback: callback,
	}
}

type LoginResponse struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Type   string `json:"type"`
	URL    string `json:"url,omitempty"`
	QRCode string `json:"qrcode,omitempty"`
}

type CallbackResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	Type        string `json:"type"`
	AccessToken string `json:"access_token"`
	SocialUID   string `json:"social_uid"`
	FaceImg     string `json:"faceimg"`
	Nickname    string `json:"nickname"`
	Gender      string `json:"gender"`
	Location    string `json:"location"`
	IP          string `json:"ip"`
}

type QueryResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	Type        string `json:"type"`
	SocialUID   string `json:"social_uid"`
	AccessToken string `json:"access_token"`
	Nickname    string `json:"nickname"`
	FaceImg     string `json:"faceimg"`
	Gender      string `json:"gender"`
	Location    string `json:"location"`
	IP          string `json:"ip"`
}

func (o *Oauth) Login(loginType string) (*LoginResponse, error) {
	state := generateState()

	params := url.Values{}
	params.Set("act", "login")
	params.Set("appid", o.appID)
	params.Set("appkey", o.appKey)
	params.Set("type", loginType)
	params.Set("redirect_uri", o.callback)
	params.Set("state", state)

	loginURL := o.apiURL + "?" + params.Encode()

	resp, err := o.doRequest(loginURL)
	if err != nil {
		return nil, err
	}

	var result LoginResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (o *Oauth) Callback(code string) (*CallbackResponse, error) {
	params := url.Values{}
	params.Set("act", "callback")
	params.Set("appid", o.appID)
	params.Set("appkey", o.appKey)
	params.Set("code", code)

	tokenURL := o.apiURL + "?" + params.Encode()

	resp, err := o.doRequest(tokenURL)
	if err != nil {
		return nil, err
	}

	var result CallbackResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (o *Oauth) Query(loginType, socialUID string) (*QueryResponse, error) {
	params := url.Values{}
	params.Set("act", "query")
	params.Set("appid", o.appID)
	params.Set("appkey", o.appKey)
	params.Set("type", loginType)
	params.Set("social_uid", socialUID)

	queryURL := o.apiURL + "?" + params.Encode()

	resp, err := o.doRequest(queryURL)
	if err != nil {
		return nil, err
	}

	var result QueryResponse
	if err := json.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (o *Oauth) doRequest(requestURL string) ([]byte, error) {
	client := &http.Client{
		Timeout: 10 * 1000000000,
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}

func generateState() string {
	h := md5.New()
	h.Write([]byte(randomString()))
	return hex.EncodeToString(h.Sum(nil))
}

func randomString() string {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000000))
	return n.String()
}
