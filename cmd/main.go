package main

import (
	"fmt"
	"github.com/Ramcache/git-helper/config"
	"github.com/joho/godotenv"
	"log"

	"github.com/Ramcache/git-helper/ai"
	"github.com/Ramcache/git-helper/commit"
	"github.com/Ramcache/git-helper/git"
)

func main() {
	cfg := config.LoadConfig()
	// Определение флага для генерации
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Ошибка при загрузке файла .env: %v", err)
	}
	// Получение токена доступа
	token, err := ai.GetAccessToken(cfg)
	if err != nil {
		log.Fatalf("Error getting access token: %v\n", err)
	}

	repoPath := "."
	// Получение изменений в репозитории
	changes, err := git.GetGitDiff(repoPath)
	if err != nil {
		log.Fatalf("Error getting repository changes: %v\n", err)
	}

	// Генерация сообщения коммита
	commitMessage, err := commit.GenerateCommitMessage(token, changes)
	if err != nil {
		log.Fatalf("Error generating commit message: %v\n", err)
	}

	// Вывод только текста коммита
	fmt.Print(commitMessage)
}
