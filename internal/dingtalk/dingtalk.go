package dingtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

/*
[{
	"scopeId": 2,
	"name": "growing-segmentation-pid:15149@seg3",
	"id0": 47,
	"id1": 0,
	"alarmMessage": "Response time of service instance growing-segmentation-pid:15149@seg3 is more than 1000ms in 2 minutes of last 10 minutes",
	"startTime": 1568888544862
}, {
	"scopeId": 2,
	"name": "growing-segmentation-pid:11847@seg2",
	"id0": 46,
	"id1": 0,
	"alarmMessage": "Response time of service instance growing-segmentation-pid:11847@seg2 is more than 1000ms in 2 minutes of last 10 minutes",
	"startTime": 1568888544862
}]
*/
type message struct {
	ScopeId int
	Name string
	Id0	int
	Id1 int
	AlarmMessage	string
	StartTime	int
}


// Dingtalk 发送钉钉消息体
func Dingtalk(data []byte) error {
	var m []message
	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println(err.Error())
	}
	contents, alertSummary := createContent(m)
	bodys := strings.NewReader(contents)
	token := viper.GetString("dingtalk.p3")
	resp, err := http.Post(
		fmt.Sprintf("https://oapi.dingtalk.com/robot/send?access_token=%s", token), "application/json", bodys)
	if err != nil {
		return err
	}
	log.Println(resp.StatusCode, alertSummary)
	return nil
}


/*
状态: notify

等级: P1

告警: Skywalking
  growing-segmentation-pid:6494@seg1  id: 44  time: 1568945304861
  growing-segmentation-pid:6908@seg0  id: 43  time: 1568945304861


Item values:

0  Response time of service instance growing-segmentation-pid:6494@seg1 is more than 1000ms in 2 minutes of last 10 minutes
1  Response time of service instance growing-segmentation-pid:6908@seg0 is more than 1000ms in 2 minutes of last 10 minutes


故障修复:
*/
func createContent(message []message) (string, string) {
	grade := "P1"
	description	:=	""

	var alertname bytes.Buffer
	var alertSummary bytes.Buffer
	for i, alert := range message {
		alertname.WriteString(fmt.Sprintf("\t%s\tid: %d\ttime: %d\n", alert.Name, alert.Id0, alert.StartTime))
		alertSummary.WriteString(fmt.Sprintf("%b\t%s\n", i, alert.AlarmMessage))
	}

	contents := fmt.Sprintf(`状态: notify\n\n等级: %s\n\n告警: Skywalking\n%s\n\nItem values:\n\n%s\n\n故障修复: %s`,
		grade, alertname.String(), alertSummary.String(), description)
	data := fmt.Sprintf(`{
        "msgtype": "text",
            "text": {
            "content": "%s",
        },
        "at": {
            "isAtAll": "",
        },
    }`, contents)
	return data, alertSummary.String()
}
