package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			updateZendeskTicketsCount()
			time.Sleep(10 * time.Second)
		}
	}()
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func updateZendeskTicketsCount() {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", "https://"+os.Getenv("ZENDESK_DOMAIN")+".zendesk.com/api/v2/tickets.json", nil)
	req.Header.Add("Authorization", "Basic "+basicAuth(os.Getenv("ZENDESK_USER"), os.Getenv("ZENDESK_PASSWORD")))
	resp, _ := client.Do(req)

	bodyText, _ := ioutil.ReadAll(resp.Body)

	data := make(map[string]interface{})
	err := json.Unmarshal(bodyText, &data)

	if err != nil {
		return
	}

	cnt, ok := data["count"].(float64)

	if !ok {
		return
	}
	ticketsCount.Set(float64(cnt))
}

var (
	ticketsCount = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "zendesk_tickets_count",
		Help: "Count of zendesk tickets",
	})
)

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":9803", nil)
}
