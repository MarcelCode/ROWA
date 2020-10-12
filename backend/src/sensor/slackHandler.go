package sensor

import (
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "time"
)

type SlackRequestBody struct {
    Text string `json:"text"`
}
var lastNotification time.Time
func CheckReadings(timestamp time.Time, waterLevel float64){
	if(waterLevel > 15){
        if(timestamp.Sub(lastNotification).Hours() > 24) {
            SendSlackNotification("https://hooks.slack.com/services/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","Water level low!")
            lastNotification = timestamp
        }
    }

}

func SendSlackNotification(webhookUrl string, msg string) error {

    slackBody, _ := json.Marshal(SlackRequestBody{Text: msg})
    req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
    if err != nil {
        return err
    }

    req.Header.Add("Content-Type", "application/json")

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return err
    }

    buf := new(bytes.Buffer)
    buf.ReadFrom(resp.Body)
    if buf.String() != "ok" {
        return errors.New("Non-ok response returned from Slack")
    }
    return nil
}