package ai

import (
	"fmt"
	"strings"
)

// SummarizeChanges generates a concise summary for commit messages
func SummarizeChanges(changes []string) (string, error) {
	// Simulate AI summarization by combining changes intelligently
	if len(changes) == 0 {
		return "", fmt.Errorf("no changes to summarize")
	}

	var added, modified, deleted []string
	for _, change := range changes {
		if strings.HasPrefix(change, "Added") {
			added = append(added, change)
		} else if strings.HasPrefix(change, "Modified") {
			modified = append(modified, change)
		} else if strings.HasPrefix(change, "Deleted") {
			deleted = append(deleted, change)
		}
	}

	var summary []string
	if len(added) > 0 {
		summary = append(summary, fmt.Sprintf("Added %d files", len(added)))
	}
	if len(modified) > 0 {
		summary = append(summary, fmt.Sprintf("Modified %d files", len(modified)))
	}
	if len(deleted) > 0 {
		summary = append(summary, fmt.Sprintf("Deleted %d files", len(deleted)))
	}

	return strings.Join(summary, ", "), nil
}
