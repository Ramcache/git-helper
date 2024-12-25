package git

import (
	"fmt"

	git "github.com/go-git/go-git/v5"
)

// Change represents a file change
type Change struct {
	File   string
	Change string // Added, Modified, Deleted
}

// ScanChanges scans a Git repository for changes
func ScanChanges(repoPath string) ([]Change, error) {
	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open repo: %w", err)
	}

	wt, err := repo.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to get worktree: %w", err)
	}

	status, err := wt.Status()
	if err != nil {
		return nil, fmt.Errorf("failed to get repository status: %w", err)
	}

	var changes []Change
	for file, change := range status {
		changes = append(changes, Change{
			File:   file,
			Change: fmt.Sprintf("%s", change.Worktree), // Преобразуем Worktree в строку
		})
	}

	return changes, nil
}
