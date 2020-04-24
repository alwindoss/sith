package branch_test

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/alwindoss/sith/mocks"
	"github.com/alwindoss/sith/pkg/branch"

	// internalbranch "github.com/alwindoss/sith/internal/branch"
	"github.com/alwindoss/sith/pkg/sith"
	"github.com/golang/mock/gomock"
	"github.com/google/go-github/github"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var (
	client   *github.Client
	err      error
	ctx      context.Context
	mockCtrl *gomock.Controller
)

var _ = Describe("Branch", func() {

	BeforeSuite(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		ctx = context.Background()
		client, err = sith.NewGithubClient(ctx)
		// fmt.Fprintf(GinkgoWriter, "Initializer called")
		fmt.Printf("Initializer called")
		Expect(err, BeNil)
	})

	AfterSuite(func() {
		defer mockCtrl.Finish()
	})

	BeforeEach(func() {
	})

	AfterEach(func() {

	})

	It("create branch for a single repository", func() {
		mockGitService := mocks.NewMockGitService(mockCtrl)

		returnedRefJSON := `{
			"ref": "refs/heads/master",
			"node_id": "MDM6UmVmcmVmcy9oZWFkcy9mZWF0dXJlQQ==",
			"url": "https://api.github.com/repos/octocat/Hello-World/git/refs/heads/master",
			"object": {
			  "type": "commit",
			  "sha": "aa218f56b14c9653891f9e74264a383fa43fefbd",
			  "url": "https://api.github.com/repos/octocat/Hello-World/git/commits/aa218f56b14c9653891f9e74264a383fa43fefbd"
			}
		  }`
		returnedRefObj := &github.Reference{}
		returnedRespObj := &github.Response{}
		var returnedErr error = nil
		owner := "alwindoss"
		repo := "conduit"
		err := json.Unmarshal([]byte(returnedRefJSON), returnedRefObj)
		mockGitService.EXPECT().GetRef(ctx, owner, repo, "heads/master").Return(returnedRefObj, returnedRespObj, returnedErr)
		createRefObj := &github.Reference{}
		createRefJSON := `{
			"ref": "refs/heads/newbranch",
			"node_id": "MDM6UmVmcmVmcy9oZWFkcy9mZWF0dXJlQQ==",
			"url": "https://api.github.com/repos/octocat/Hello-World/git/refs/heads/newbranch",
			"object": {
			  "type": "commit",
			  "sha": "aa218f56b14c9653891f9e74264a383fa43fefbd",
			  "url": "https://api.github.com/repos/octocat/Hello-World/git/commits/aa218f56b14c9653891f9e74264a383fa43fefbd"
			}
		  }`
		err = json.Unmarshal([]byte(createRefJSON), createRefObj)
		Expect(err, BeNil())
		mockGitService.EXPECT().CreateRef(ctx, owner, repo, gomock.Any()).Return(createRefObj, returnedRespObj, returnedErr)
		branch.CreateBranch(ctx, mockGitService, owner, repo, "master", "newbranch")
	})
})
