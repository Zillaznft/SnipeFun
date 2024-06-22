package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"GoSnipeFun/config"
)

func sendNotification(notification Notification) error {
	discordMessage := DiscordMessage{
		Content: strings.ReplaceAll(notification.Message, "``", ""),
	}

	messageBytes, err := json.Marshal(discordMessage)
	if err != nil {
		return fmt.Errorf("error marshaling message: %w", err)
	}

	req, err := http.NewRequest("POST", getWebhookByType(notification), bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("received non-204 response: %d", resp.StatusCode)
	}

	return nil
}

func getWebhookByType(notification Notification) string {
	switch notification.Type {
	case infoNotification:
		return config.NotifWebhookURL
	case eventNotification:
		return config.EventWebhookURL
	case errorNotification:
		return config.ErrorWebhookURL
	default:
		return ""
	}
}

func getSymbol(flag bool) string {
	if flag {
		return "🟢"
	}
	return "🔴"
}
