package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/tqtcloud/alertmanager-dispose/handlers"
	"net/http"
)

func main() {
	//http.Handle("/api/v2/alerts", &handlers.AlertHandler{})
	http.Handle("/webhook", &handlers.AlertHandler{})
	log.Info("webhook 启动成功,启动端口在：:18083")
	if err := http.ListenAndServe(":18083", nil); err != nil {
		log.Errorf("启动失败：%#v", err)
	}
}
