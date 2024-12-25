package git

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetGitDiff retrieves the changes in the repository using `git diff`
func GetGitDiff(repoPath string) (string, error) {
	cmd := exec.Command("git", "diff")
	cmd.Dir = repoPath // Указываем директорию репозитория

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении git diff: %v", err)
	}

	diff := strings.TrimSpace(string(output))

	// Проверка на пустой вывод
	if diff == "" {
		return "", fmt.Errorf("в репозитории нет изменений (git diff вернул пустой результат)")
	}

	return diff, nil
}
