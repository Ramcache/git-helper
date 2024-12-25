package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Ramcache/git-helper/internal/commit"
	"github.com/Ramcache/git-helper/internal/git"
)

func main() {
	repoPath := "R:\\ProjectsGo\\InstaSpace" // Path to your Git repository

	// Scan changes in the repository
	changes, err := git.ScanChanges(repoPath)
	if err != nil {
		log.Fatalf("Error scanning repository: %v", err)
	}

	// Generate a commit message
	message, err := commit.GenerateCommitMessage(changes)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	fmt.Println("Generated Commit Message:")
	fmt.Println(message)

	// Commit and push (optional)
	if len(changes) > 0 {
		if err := os.WriteFile("commit_message.txt", []byte(message), 0644); err != nil {
			log.Fatalf("Failed to save commit message: %v", err)
		}

		fmt.Println("Commit message saved to commit_message.txt.")
		fmt.Println("You can now use `git commit -F commit_message.txt` to commit.")
	}
}
