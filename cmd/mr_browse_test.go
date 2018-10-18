package cmd

import (
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_mrBrowseWithParameter(t *testing.T) {
	oldBrowse := browse
	defer func() { browse = oldBrowse }()

	browse = func(url string) error {
		require.Equal(t, "https://gitlab.com/zaquestion/test/merge_requests/1", url)
		return nil
	}

	mrBrowseCmd.Run(nil, []string{"1"})
}

func Test_mrBrowseCurrent(t *testing.T) {
	t.Parallel()
	repo := copyTestRepo(t)
	git := exec.Command("git", "checkout", "mrtest")
	git.Dir = repo
	b, err := git.CombinedOutput()
	if err != nil {
		t.Log(string(b))
		t.Fatal(err)
	}

	oldBrowse := browse
	defer func() { browse = oldBrowse }()

	browse = func(url string) error {
		require.Equal(t, "https://gitlab.com/zaquestion/test/merge_requests/1", url)
		return nil
	}

	mrBrowseCmd.Run(nil, nil)
}