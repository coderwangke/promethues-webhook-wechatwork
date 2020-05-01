package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
)

const templ = `Promethues Alert:
>状态:<font color=\"comment\">{{.Status}}</font>
>开始于:<font color=\"comment\">{{.StartsAt}}</font>
>Labels:
{{ range $key, $value := .Labels }}
	{{ $key }}:{{ $value }}
{{end}}
>Annotations:
{{ range $key, $value := .Annotations }}
	{{ $key }}:{{ $value }}
{{end}}
>详情:[点击查看]({{.GeneratorURL}})`

// Parameter
var wechatWorkURL string

// Msg to wechat work
type SendMsg struct {
	Msgtype  string      `json:"msgtype"`
	Markdown interface{} `json:"markdown"`
}

type MsgContent struct {
	Content string `json:"content"`
}

func init() {
	flag.StringVar(&wechatWorkURL, "url", "", "url for wechat work robot.")
}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Write(response)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, http.StatusOK, "OK!")
}

func webhook(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data := template.Data{}

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		responseWithJSON(w, http.StatusBadRequest, err.Error())
		return
	}

	for _, alert := range data.Alerts {
		klog.Infof("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)

		msg, err := alertMsg(alert)
		if err != nil {
			responseWithJSON(w, http.StatusInternalServerError, "failed")
		}

		if err := sendToWechatWork(msg, wechatWorkURL); err != nil {
			responseWithJSON(w, http.StatusInternalServerError, "failed")
		}
	}

	responseWithJSON(w, http.StatusOK, "success")
}

func main() {
	flag.Parse()
	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthz)
	router.HandleFunc("/webhook", webhook)

	listenAddress := ":8080"
	klog.Fatal(http.ListenAndServe(listenAddress, router))
}
