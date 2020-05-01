package main

import (
	"github.com/prometheus/alertmanager/template"
	"testing"
	"time"
)

func TestAlertMsg(t *testing.T) {
	alert := template.Alert{
		Status:   "resolved",
		StartsAt: time.Now(),
		Labels: map[string]string{
			"alertName": "FlinkTaskManagerMissing",
		},
		Annotations: map[string]string{
			"summary": "Fewer Flink TaskManagers than expected are running.",
		},
		GeneratorURL: "www.example.com",
	}

	msg, _ := alertMsg(alert)
	t.Logf("Webhook Type: %s, Content: %v\n", msg.Msgtype, msg.Markdown)
}
