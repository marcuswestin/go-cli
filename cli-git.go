package cli

import "os"

func GetGitBranch() (branchName string) {
	branchName, _ = Cmd("git", "rev-parse", "--abbrev-ref", "HEAD")
	return
}

func GetGitRoot() (rootDir string) {
	rootDir, _ = Cmd("git", "rev-parse", "--show-toplevel")
	return
}
