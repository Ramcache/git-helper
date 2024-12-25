package commit

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	chatURL = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	model   = "GigaChat" // Выберите подходящую модель
)

func filterChanges(changes string) string {
	var filteredLines []string
	lines := strings.Split(changes, "\n")
	inImportBlock := false

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		// Игнорируем строки package
		if strings.HasPrefix(trimmedLine, "package") {
			continue
		}

		// Обрабатываем блок import
		if strings.HasPrefix(trimmedLine, "import (") {
			inImportBlock = true
			continue
		}
		if inImportBlock {
			// Проверяем конец блока import
			if strings.HasPrefix(trimmedLine, ")") {
				inImportBlock = false
			}
			continue
		}
		if strings.HasPrefix(trimmedLine, "import") {
			continue
		}

		// Добавляем все остальные строки
		filteredLines = append(filteredLines, line)
	}

	return strings.Join(filteredLines, "\n")
}

func GenerateCommitMessage(token, changes string) (string, error) {
	if changes == "" {
		return "", fmt.Errorf("изменения пусты, невозможно сгенерировать сообщение коммита")
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": "На основе следующих изменений в коде сгенерируй максимально краткое и лаконичное сообщение для коммита. Сообщение должно содержать не более 50 символов:\n" + filterChanges(changes)},
		},
	})
	if err != nil {
		return "", fmt.Errorf("ошибка при создании JSON-запроса: %v", err)
	}

	req, err := http.NewRequest("POST", chatURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("ошибка при создании HTTP-запроса: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("неудачный статус ответа: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("ошибка при парсинге ответа: %v", err)
	}

	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if choice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := choice["message"].(map[string]interface{}); ok {
				if content, ok := message["content"].(string); ok {
					return content, nil
				}
			}
		}
	}

	return "", fmt.Errorf("не удалось получить сообщение коммита из ответа")
}
