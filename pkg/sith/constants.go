package sith

const (
	// FetchBranchPrefix represents the prefix before the branch name when using GetRef API
	FetchBranchPrefix = "heads/%s"

	// CreateBranchPrefix represents the prefix before the branch name when using CreateRef API
	CreateBranchPrefix = "refs/head/%s"
)
