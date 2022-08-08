package tools

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
)

func UrlConstructor(addr string, r *http.Request) (string, map[string]string) {
	q := r.URL.Query()

	u := fmt.Sprintf("%s?date=%s&start_date=%s&end_date=%s&count=%s&thumbs=%s&api_key=%s",
		addr, q.Get("date"), q.Get("start_date"), q.Get("end_date"), q.Get("count"), q.Get("thumbs"), os.Getenv("API_KEY"))

	return u, map[string]string{
		"hd":         q.Get("hd"),
		"store":      q.Get("store"),
		"date":       q.Get("date"),
		"start_date": q.Get("start_date"),
		"end_date":   q.Get("end_date"),
		"img_url":    q.Get("img_url"),
	}
}

func ParseTime(str string) (string, error) {
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		return "", err
	}

	return t.Format("2006-01-02"), nil
}

func NewUUID() string {
	return uuid.New().String()
}
