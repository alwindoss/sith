package branch

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/alwindoss/sith/internal/branch"
	"github.com/alwindoss/sith/pkg/sith"
	"github.com/google/go-github/github"
)

// CreateBranch creates a branch
func CreateBranch(ctx context.Context, gitService sith.GitService, owner, repo, sourceBranch, destBranchName string) {
	// gitService.CreateRef(ctx, "", "", nil)

	log.Printf("CreateBranch called")
	log.Printf("SourceBranch: %s", sourceBranch)
	log.Printf("DestBranch: %s", destBranchName)
	sourceRef, err := branch.GetBranchDetails(ctx, gitService, owner, repo, sourceBranch)
	if err != nil {
		log.Printf("unable to fetch the source branch reference %v", err)
		os.Exit(-1)
	}
	newFullyQualifiedBranchName := fmt.Sprintf(sith.CreateBranchPrefix, destBranchName)
	destObj := &github.GitObject{
		SHA: sourceRef.Object.SHA,
	}
	destRef := &github.Reference{
		Ref:    &newFullyQualifiedBranchName,
		Object: destObj,
	}
	_, _, err = gitService.CreateRef(ctx, owner, repo, destRef)
	if err != nil {
		panic(err)
	}
}
