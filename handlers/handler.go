package handlers

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/tqtcloud/alertmanager-dispose/impl"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	linuxDownErr = "主机linux无法连接"
	linuxTcpErr  = "警告：tcp连接数大于10000"
)

//type GenerateStruct struct {
//	Annotations  *Annotations `json:"annotations,omitempty"`
//	EndsAt       string       `json:"endsAt,omitempty"`
//	StartsAt     string       `json:"startsAt,omitempty"`
//	GeneratorURL string       `json:"generatorURL,omitempty"`
//	Labels       *Labels      `json:"labels,omitempty"`
//}
//type Annotations struct {
//	Description string `json:"description,omitempty"`
//	Summary     string `json:"summary,omitempty"`
//}
//type Labels struct {
//	Alertname string `json:"alertname,omitempty"`
//	Cluster   string `json:"cluster,omitempty"`
//	Env       string `json:"env,omitempty"`
//	Instance  string `json:"instance,omitempty"`
//	Job       string `json:"job,omitempty"`
//	Severity  string `json:"severity,omitempty"`
//	Team      string `json:"team,omitempty"`
//	User      string `json:"user,omitempty"`
//}
type GenerateStruct struct {
	Receiver          string             `json:"receiver,omitempty"`
	Status            string             `json:"status,omitempty"`
	Alerts            []Alerts           `json:"alerts,omitempty"`
	GroupLabels       *GroupLabels       `json:"groupLabels,omitempty"`
	CommonLabels      *CommonLabels      `json:"commonLabels,omitempty"`
	CommonAnnotations *CommonAnnotations `json:"commonAnnotations,omitempty"`
	ExternalURL       string             `json:"externalURL,omitempty"`
	Version           string             `json:"version,omitempty"`
	GroupKey          string             `json:"groupKey,omitempty"`
	TruncatedAlerts   int                `json:"truncatedAlerts,omitempty"`
}

type Labels struct {
	Alertname string `json:"alertname,omitempty"`
	Cluster   string `json:"cluster,omitempty"`
	Env       string `json:"env,omitempty"`
	Instance  string `json:"instance,omitempty"`
	Job       string `json:"job,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Team      string `json:"team,omitempty"`
	User      string `json:"user,omitempty"`
}

type Annotations struct {
	Description string `json:"description,omitempty"`
	Summary     string `json:"summary,omitempty"`
}

type Alerts struct {
	Status       string      `json:"status,omitempty"`
	Labels       Labels      `json:"labels,omitempty"`
	Annotations  Annotations `json:"annotations,omitempty"`
	StartsAt     string      `json:"startsAt,omitempty"`
	EndsAt       string      `json:"endsAt,omitempty"`
	GeneratorURL string      `json:"generatorURL,omitempty"`
	Fingerprint  string      `json:"fingerprint,omitempty"`
}

type GroupLabels struct {
	Alertname string `json:"alertname,omitempty"`
}

type CommonLabels struct {
	Alertname string `json:"alertname,omitempty"`
	Cluster   string `json:"cluster,omitempty"`
	Env       string `json:"env,omitempty"`
	Instance  string `json:"instance,omitempty"`
	Job       string `json:"job,omitempty"`
	Severity  string `json:"severity,omitempty"`
	Team      string `json:"team,omitempty"`
	User      string `json:"user,omitempty"`
}

type CommonAnnotations struct {
	Description string `json:"description,omitempty"`
	Summary     string `json:"summary,omitempty"`
}
type AlertHandler struct{}

func (mh *AlertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	var text GenerateStruct
	if err != nil {
		fmt.Printf("read body err, %v\n", err)
		return
	}
	fmt.Println("接受到告警：", string(body))
	err = json.Unmarshal(body, &text)
	if err != nil {
		fmt.Printf("json Unmarshal, %v\n", err)
		return
	}
	//fmt.Println(time.Now())
	//fmt.Println(text[0])
	//fmt.Printf("%#v\n", text[0].Labels.Instance)
	//fmt.Println(text[0].Annotations.Description)
	//fmt.Println(text.Alerts[0].Annotations.Summary)
	//
	//fmt.Println(text[0].Labels.Cluster)

	// 遍历所有的告警条目，找到匹配条目
	os.Open("/data")
	for i := 0; i < len(text.Alerts); i++ {
		switch strings.ReplaceAll(text.Alerts[i].Annotations.Summary, " ", "") {
		case linuxDownErr:
			//if strings.ReplaceAll(text.Alerts[i].Annotations.Summary, " ", "") == linuxDownErr && text.Alerts[i].Status == "firing" {
			//	strIP := strings.Split(text.Alerts[i].Labels.Instance, ":")
			//	log.Infof("告警规则命中，触发自愈操作,操作主机：%s", strIP[0])
			//
			//	if err := impl.ConnectHost("root", "centos", strIP[0], 22); err != nil {
			//		//if err := impl.ConnectHost("root", "Bs+425tBa%", "192.168.196.95", 22); err != nil {
			//		log.Errorf("连接远端服务器失败：%#v \n", err)
			//	} else {
			//		log.Info("自愈成功")
			//		time.Sleep(60 * time.Second)
			//	}
			//}
		case linuxTcpErr:
			if strings.ReplaceAll(text.Alerts[i].Annotations.Summary, " ", "") == linuxTcpErr && text.Alerts[i].Status == "firing" {
				strIP := strings.Split(text.Alerts[i].Labels.Instance, ":")
				log.Infof("告警规则命中: %s ，触发自愈操作,操作主机：%s", linuxTcpErr, strIP[0])

				if err := impl.ConnectHost("root", "Bs+425tBa%", strIP[0], 22); err != nil {
					//if err := impl.ConnectHost("root", "Bs+425tBa%", "192.168.196.95", 22); err != nil {
					log.Errorf("连接远端服务器失败：%#v \n", err)
				} else {
					log.Info("自愈成功")
					time.Sleep(60 * time.Second)
				}
			}
		}
	}
	log.Info("告警恢复：", text.Alerts)
	w.WriteHeader(http.StatusOK)
}
