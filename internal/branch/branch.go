package branch

import (
	"context"
	"fmt"
	"log"

	"github.com/alwindoss/sith/pkg/sith"
	"github.com/google/go-github/github"
)

// GetBranchDetails fetches details of the branch
func GetBranchDetails(ctx context.Context, gitService sith.GitService, owner, repo, branchName string) (*github.Reference, error) {
	// When fetching the branch details using GetRef the reference should be of the form 'heads/<branch-name>'
	referenceQualifiedName := fmt.Sprintf(sith.FetchBranchPrefix, branchName)
	log.Printf("Fully qualified branch name: %s", referenceQualifiedName)
	ref, _, err := gitService.GetRef(ctx, owner, repo, referenceQualifiedName)
	if err != nil {
		log.Printf("Error when fetching the references: %v", err)
		return nil, err
	}
	log.Printf("Ref.SHA: %s\n", ref.GetObject().GetSHA())
	return ref, nil
}
