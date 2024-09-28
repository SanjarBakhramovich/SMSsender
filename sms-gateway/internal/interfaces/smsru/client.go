package smsru

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type SmsRuClient struct {
    apiID string
}

func NewSmsRuClient(apiID string) *SmsRuClient {
    return &SmsRuClient{apiID: apiID}
}

func (client *SmsRuClient) SendSMS(phoneNumber string, message string) error {
    smsURL := "https://sms.ru/sms/send"
    params := map[string]string{
        "api_id":  client.apiID,
        "to":      phoneNumber,
        "msg":     message,
        "json":    "1", // Просим вернуть ответ в формате JSON
    }

    jsonData, err := json.Marshal(params)
    if err != nil {
        return fmt.Errorf("error marshaling json: %w", err)
    }

    resp, err := http.Post(smsURL, "application/json", bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("error sending request to sms.ru: %w", err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return fmt.Errorf("error reading response from sms.ru: %w", err)
    }

    log.Printf("Response from sms.ru: %s", string(body))
    return nil
}
