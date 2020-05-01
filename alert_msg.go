package main

import (
	"bytes"
	"encoding/json"
	"github.com/prometheus/alertmanager/template"
	"k8s.io/klog"
	"net/http"
	gotemplate "text/template"
)

func alertMsg(alert template.Alert) (*SendMsg, error) {
	msg := &SendMsg{
		Msgtype: "markdown",
	}
	var doc bytes.Buffer
	t, err := gotemplate.New("alert").Parse(templ)
	if err != nil {
		klog.Errorf("Webhook: initial go template error: %v", err.Error())
		return nil, err
	}
	if err := t.Execute(&doc, alert); err != nil {
		klog.Errorf("Webhook: go template execute error: %v", err.Error())
		return nil, err
	}
	msg.Markdown = &MsgContent{Content: doc.String()}

	return msg, nil
}

func sendToWechatWork(msg interface{}, webhookURL string) error {
	jsonBytes, err := json.Marshal(msg)
	if err != nil {
		klog.Errorf("sendToWechatWork: json marshal error: %v", err.Error())
		return err
	}

	reader := bytes.NewReader(jsonBytes)
	url := wechatWorkURL
	request, err := http.NewRequest("POST", url, reader)
	if err != nil {
		klog.Errorf("sendToWechatWork: http new request error: %v", err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	_, err = client.Do(request)
	if err != nil {
		klog.Errorf("sendToWechatWork: http post request error: %v", err.Error())
		return err
	}

	return nil
}
