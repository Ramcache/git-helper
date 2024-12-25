package commit

import (
	"fmt"

	"github.com/Ramcache/git-helper/internal/ai"
	"github.com/Ramcache/git-helper/internal/git"
)

// GenerateCommitMessage generates a commit message using AI
func GenerateCommitMessage(changes []git.Change) (string, error) {
	if len(changes) == 0 {
		return "No changes detected", nil
	}

	// Prepare changes for AI input
	var details []string
	for _, change := range changes {
		details = append(details, fmt.Sprintf("%s: %s", change.Change, change.File))
	}

	// Use AI summarizer for a concise message
	message, err := ai.SummarizeChanges(details)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	return message, nil
}
