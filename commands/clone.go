package commands

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/github/hub/github"
	"github.com/github/hub/utils"
)

var cmdClone = &Command{
	Run:          clone,
	GitExtension: true,
	Usage:        "clone [-p] [<OPTIONS>] [<USER>/]<REPOSITORY> [<DESTINATION>]",
	Long: `Clone a repository from GitHub.

## Options:
	-p
		(Deprecated) Clone private repositories over SSH.

	[<USER>/]<REPOSITORY>
		<USER> defaults to your own GitHub username.

	<DESTINATION>
		Directory name to clone into (default: <REPOSITORY>).

## Protocol used for cloning

The 'git:' protocol will be used for cloning public repositories, while the SSH
protocol will be used for private repositories and those that you have push
access to. Alternatively, hub can be configured to use HTTPS protocol for
everything. See "HTTPS instead of git protocol" and "HUB_PROTOCOL" of hub(1).

## Examples:
		$ hub clone rtomayko/ronn
		> git clone git://github.com/rtomayko/ronn.git

## See also:

hub-fork(1), hub(1), git-clone(1)
`,
}

func init() {
	CmdRunner.Use(cmdClone)
}

func clone(command *Command, args *Args) {
	if !args.IsParamsEmpty() {
		transformCloneArgs(args)
	}
}

func transformCloneArgs(args *Args) {
	isSSH := parseClonePrivateFlag(args)

	// git help clone | grep -e '^ \+-.\+<'
	p := utils.NewArgsParser()
	p.RegisterValue("--branch", "-b")
	p.RegisterValue("--depth")
	p.RegisterValue("--reference")
	if args.Command == "submodule" {
		p.RegisterValue("--name")
	} else {
		p.RegisterValue("--config", "-c")
		p.RegisterValue("--jobs", "-j")
		p.RegisterValue("--origin", "-o")
		p.RegisterValue("--reference-if-able")
		p.RegisterValue("--separate-git-dir")
		p.RegisterValue("--shallow-exclude")
		p.RegisterValue("--shallow-since")
		p.RegisterValue("--template")
		p.RegisterValue("--upload-pack", "-u")
	}
	p.Parse(args.Params)

	nameWithOwnerRegexp := regexp.MustCompile(NameWithOwnerRe)
	for _, i := range p.PositionalIndices {
		a := args.Params[i]
		if nameWithOwnerRegexp.MatchString(a) && !isCloneable(a) {
			url := getCloneUrl(a, isSSH, args.Command != "submodule")
			args.ReplaceParam(i, url)
		}
		break
	}
}

func parseClonePrivateFlag(args *Args) bool {
	if i := args.IndexOfParam("-p"); i != -1 {
		args.RemoveParam(i)
		return true
	}

	return false
}

func getCloneUrl(nameWithOwner string, isSSH, allowSSH bool) string {
	name := nameWithOwner
	owner := ""
	if strings.Contains(name, "/") {
		split := strings.SplitN(name, "/", 2)
		owner = split[0]
		name = split[1]
	}

	var host *github.Host
	if owner == "" {
		config := github.CurrentConfig()
		h, err := config.DefaultHost()
		if err != nil {
			utils.Check(github.FormatError("cloning repository", err))
		}

		host = h
		owner = host.User
	}

	var hostStr string
	if host != nil {
		hostStr = host.Host
	}

	expectWiki := strings.HasSuffix(name, ".wiki")
	if expectWiki {
		name = strings.TrimSuffix(name, ".wiki")
	}

	project := github.NewProject(owner, name, hostStr)
	gh := github.NewClient(project.Host)
	repo, err := gh.Repository(project)
	if err != nil {
		if strings.Contains(err.Error(), "HTTP 404") {
			err = fmt.Errorf("Error: repository %s/%s doesn't exist", project.Owner, project.Name)
		}
		utils.Check(err)
	}

	owner = repo.Owner.Login
	name = repo.Name
	if expectWiki {
		if !repo.HasWiki {
			utils.Check(fmt.Errorf("Error: %s/%s doesn't have a wiki", owner, name))
		} else {
			name = name + ".wiki"
		}
	}

	if !isSSH &&
		allowSSH &&
		!github.IsHttpsProtocol() {
		isSSH = repo.Private || repo.Permissions.Push
	}

	return project.GitURL(name, owner, isSSH)
}
