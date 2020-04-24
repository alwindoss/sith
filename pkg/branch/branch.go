package branch

import (
	"context"
	"log"

	"github.com/alwindoss/sith/internal/branch"
	"github.com/alwindoss/sith/pkg/sith"
)

// CreateBranch creates a branch
func CreateBranch(ctx context.Context, gitService sith.GitService) {
	// gitService.CreateRef(ctx, "", "", nil)

	log.Printf("CreateBranch called")
	branch.GetBranchDetails(ctx, gitService, "alwindoss", "conduit", "master")
}
