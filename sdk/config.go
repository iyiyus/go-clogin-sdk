package oauth

const (
	APIURL = "https://u.cccyun.cc/"
)

type Config struct {
	APIURL   string
	AppID    string
	AppKey   string
	Callback string
}

func NewConfig(appID, appKey, callback string) *Config {
	return &Config{
		APIURL:   APIURL,
		AppID:    appID,
		AppKey:   appKey,
		Callback: callback,
	}
}
