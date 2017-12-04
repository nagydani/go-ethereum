// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package build

import (
	"flag"
	"fmt"
	"os"
<<<<<<< HEAD
=======
	"strings"
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
)

var (
	// These flags override values in build env.
	GitCommitFlag   = flag.String("git-commit", "", `Overrides git commit hash embedded into executables`)
	GitBranchFlag   = flag.String("git-branch", "", `Overrides git branch being built`)
	GitTagFlag      = flag.String("git-tag", "", `Overrides git tag being built`)
	BuildnumFlag    = flag.String("buildnum", "", `Overrides CI build number`)
	PullRequestFlag = flag.Bool("pull-request", false, `Overrides pull request status of the build`)
<<<<<<< HEAD
=======
	CronJobFlag     = flag.Bool("cron-job", false, `Overrides cron job status of the build`)
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
)

// Environment contains metadata provided by the build environment.
type Environment struct {
	Name                string // name of the environment
	Repo                string // name of GitHub repo
	Commit, Branch, Tag string // Git info
	Buildnum            string
	IsPullRequest       bool
<<<<<<< HEAD
=======
	IsCronJob           bool
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
}

func (env Environment) String() string {
	return fmt.Sprintf("%s env (commit:%s branch:%s tag:%s buildnum:%s pr:%t)",
		env.Name, env.Commit, env.Branch, env.Tag, env.Buildnum, env.IsPullRequest)
}

// Env returns metadata about the current CI environment, falling back to LocalEnv
// if not running on CI.
func Env() Environment {
	switch {
	case os.Getenv("CI") == "true" && os.Getenv("TRAVIS") == "true":
		return Environment{
			Name:          "travis",
			Repo:          os.Getenv("TRAVIS_REPO_SLUG"),
			Commit:        os.Getenv("TRAVIS_COMMIT"),
			Branch:        os.Getenv("TRAVIS_BRANCH"),
			Tag:           os.Getenv("TRAVIS_TAG"),
			Buildnum:      os.Getenv("TRAVIS_BUILD_NUMBER"),
			IsPullRequest: os.Getenv("TRAVIS_PULL_REQUEST") != "false",
<<<<<<< HEAD
=======
			IsCronJob:     os.Getenv("TRAVIS_EVENT_TYPE") == "cron",
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
		}
	case os.Getenv("CI") == "True" && os.Getenv("APPVEYOR") == "True":
		return Environment{
			Name:          "appveyor",
			Repo:          os.Getenv("APPVEYOR_REPO_NAME"),
			Commit:        os.Getenv("APPVEYOR_REPO_COMMIT"),
			Branch:        os.Getenv("APPVEYOR_REPO_BRANCH"),
			Tag:           os.Getenv("APPVEYOR_REPO_TAG_NAME"),
			Buildnum:      os.Getenv("APPVEYOR_BUILD_NUMBER"),
			IsPullRequest: os.Getenv("APPVEYOR_PULL_REQUEST_NUMBER") != "",
<<<<<<< HEAD
=======
			IsCronJob:     os.Getenv("APPVEYOR_SCHEDULED_BUILD") == "True",
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
		}
	default:
		return LocalEnv()
	}
}

// LocalEnv returns build environment metadata gathered from git.
func LocalEnv() Environment {
	env := applyEnvFlags(Environment{Name: "local", Repo: "ethereum/go-ethereum"})
<<<<<<< HEAD
	if _, err := os.Stat(".git"); err != nil {
		return env
	}
	if env.Commit == "" {
		env.Commit = RunGit("rev-parse", "HEAD")
	}
	if env.Branch == "" {
		if b := RunGit("rev-parse", "--abbrev-ref", "HEAD"); b != "HEAD" {
			env.Branch = b
		}
	}
	// Note that we don't get the current git tag. It would slow down
	// builds and isn't used by anything.
	return env
}

=======

	head := readGitFile("HEAD")
	if splits := strings.Split(head, " "); len(splits) == 2 {
		head = splits[1]
	} else {
		return env
	}
	if env.Commit == "" {
		env.Commit = readGitFile(head)
	}
	if env.Branch == "" {
		if head != "HEAD" {
			env.Branch = strings.TrimLeft(head, "refs/heads/")
		}
	}
	if info, err := os.Stat(".git/objects"); err == nil && info.IsDir() && env.Tag == "" {
		env.Tag = firstLine(RunGit("tag", "-l", "--points-at", "HEAD"))
	}
	return env
}

func firstLine(s string) string {
	return strings.Split(s, "\n")[0]
}

>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
func applyEnvFlags(env Environment) Environment {
	if !flag.Parsed() {
		panic("you need to call flag.Parse before Env or LocalEnv")
	}
	if *GitCommitFlag != "" {
		env.Commit = *GitCommitFlag
	}
	if *GitBranchFlag != "" {
		env.Branch = *GitBranchFlag
	}
	if *GitTagFlag != "" {
		env.Tag = *GitTagFlag
	}
	if *BuildnumFlag != "" {
		env.Buildnum = *BuildnumFlag
	}
	if *PullRequestFlag {
		env.IsPullRequest = true
	}
<<<<<<< HEAD
=======
	if *CronJobFlag {
		env.IsCronJob = true
	}
>>>>>>> 1d06e41f04d75c31334c455063e9ec7b4136bf23
	return env
}
