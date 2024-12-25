package ai

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Ramcache/git-helper/config"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
)

// GetAccessToken получает временный токен доступа для GigaChat
func GetAccessToken(cfg *config.Config) (string, error) {
	authKey := base64.StdEncoding.EncodeToString([]byte(cfg.ClientID + ":" + cfg.ClientSecret))
	data := "scope=" + cfg.Scope

	req, err := http.NewRequest("POST", cfg.AuthURL, bytes.NewBufferString(data))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Basic "+authKey)
	req.Header.Set("RqUID", uuid.New().String())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("неожиданный статус ответа: %s, тело ответа: %s", resp.Status, string(bodyBytes))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	token, ok := result["access_token"].(string)
	if !ok {
		return "", fmt.Errorf("не удалось получить токен доступа из ответа")
	}

	return token, nil
}
