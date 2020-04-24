package branch_test

import (
	"context"
	"fmt"

	"github.com/alwindoss/sith/pkg/branch"
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
		branch.CreateBranch(ctx, client.Git)
	})
})
