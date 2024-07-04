package proxychecker

import (
	"net/http"
	"net/url"
)

func ProxyChecker(ipPort string) string {
	proxyURL, err := url.Parse("http://" + ipPort)
	if err != nil {
		return "0"
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		},
	}

	targetURL := "http://google.com"

	resp, err := client.Get(targetURL)
	if err != nil {
		return "0"
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return "1"
	} else {
		return "0"
	}
}
