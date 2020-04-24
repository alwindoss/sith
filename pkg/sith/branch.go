package sith

//go:generate mockgen -destination=../../mocks/mock_branch.go -package=mocks github.com/alwindoss/sith/pkg/sith GitService

import (
	"context"

	"github.com/google/go-github/github"
)

// CreateBranchInput is the input to create a branch
type CreateBranchInput struct {
}

// GitService gets
type GitService interface {
	CreateBlob(ctx context.Context, owner string, repo string, blob *github.Blob) (*github.Blob, *github.Response, error)
	CreateCommit(ctx context.Context, owner string, repo string, commit *github.Commit) (*github.Commit, *github.Response, error)
	CreateRef(ctx context.Context, owner string, repo string, ref *github.Reference) (*github.Reference, *github.Response, error)
	CreateTag(ctx context.Context, owner string, repo string, tag *github.Tag) (*github.Tag, *github.Response, error)
	CreateTree(ctx context.Context, owner string, repo string, baseTree string, entries []github.TreeEntry) (*github.Tree, *github.Response, error)
	DeleteRef(ctx context.Context, owner string, repo string, ref string) (*github.Response, error)
	GetBlob(ctx context.Context, owner string, repo string, sha string) (*github.Blob, *github.Response, error)
	GetBlobRaw(ctx context.Context, owner, repo, sha string) ([]byte, *github.Response, error)
	GetCommit(ctx context.Context, owner string, repo string, sha string) (*github.Commit, *github.Response, error)
	GetRef(ctx context.Context, owner string, repo string, ref string) (*github.Reference, *github.Response, error)
	GetRefs(ctx context.Context, owner string, repo string, ref string) ([]*github.Reference, *github.Response, error)
	GetTag(ctx context.Context, owner string, repo string, sha string) (*github.Tag, *github.Response, error)
	GetTree(ctx context.Context, owner string, repo string, sha string, recursive bool) (*github.Tree, *github.Response, error)
	ListRefs(ctx context.Context, owner, repo string, opt *github.ReferenceListOptions) ([]*github.Reference, *github.Response, error)
	UpdateRef(ctx context.Context, owner string, repo string, ref *github.Reference, force bool) (*github.Reference, *github.Response, error)
}
