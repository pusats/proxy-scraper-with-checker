package main

import (
	"encoding/json"
	"fmt"
	"gohere/ProxyPress/proxychecker"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

type ProxyData struct {
	Data  []ProxyInfo `json:"data"`
	Count int         `json:"count"`
}

type ProxyInfo struct {
	IpPort      string      `json:"ipPort"`
	Ip          string      `json:"ip"`
	Port        string      `json:"port"`
	Country     string      `json:"country"`
	LastChecked string      `json:"last_checked"`
	ProxyLevel  string      `json:"proxy_level"`
	Type        string      `json:"type"`
	Speed       string      `json:"speed"`
	Support     SupportInfo `json:"support"`
}

type SupportInfo struct {
	Https     int `json:"https"`
	Get       int `json:"get"`
	Post      int `json:"post"`
	Cookies   int `json:"cookies"`
	Referer   int `json:"referer"`
	UserAgent int `json:"user_agent"`
	Google    int `json:"google"`
}

func main() {
	var thrds int
	fmt.Printf("Threads: ")
	fmt.Scan(&thrds)

	wg := sync.WaitGroup{}
	wg.Add(thrds)
	for i := 0; i < thrds; i++ {
		go func() {
			defer wg.Done()
			for true {
				file, _ := os.OpenFile("list/all.txt", os.O_WRONLY|os.O_APPEND, 0666)
				file2, _ := os.OpenFile("list/good.txt", os.O_WRONLY|os.O_APPEND, 0666)
				file3, _ := os.OpenFile("list/bad.txt", os.O_WRONLY|os.O_APPEND, 0666)

				time.Sleep(time.Millisecond * 1500)
				URL := "http://pubproxy.com/api/proxy"
				proxyURL := "http://pusat:pusats1023@geo.iproyal.com:12321"
				proxy := func(_ *http.Request) (*url.URL, error) {
					return url.Parse(proxyURL)
				}

				transport := &http.Transport{Proxy: proxy}
				client := &http.Client{Transport: transport}

				req, _ := http.NewRequest("GET", URL, nil)
				resp, _ := client.Do(req)
				body, _ := io.ReadAll(resp.Body)

				var proxyData ProxyData
				json.Unmarshal(body, &proxyData)
				if len(proxyData.Data) > 0 {
					ipPort := proxyData.Data[0].IpPort
					log.Println("Proxy: ", ipPort)
					file.WriteString(ipPort + "\n")
					isWork := proxychecker.ProxyChecker(ipPort)
					if isWork == "1" {
						file2.WriteString(ipPort + "\n")
						log.Println("Good: ", ipPort)
					} else if isWork == "0" {
						file3.WriteString(ipPort + "\n")
						log.Println("Bad: ", ipPort)
					}
				} else {
					time.Sleep(time.Millisecond * 500)
				}
			}
		}()
	}
	for true {
		time.Sleep(time.Second * 3)
	}
}
