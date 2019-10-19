hub is a command line tool that wraps `git` in order to extend it with extra
features and commands that make working with GitHub easier.

This repository and its issue tracker is **not for reporting problems with
GitHub.com** web interface. If you have a problem with GitHub itself, please
[contact Support](https://github.com/contact).

Usage
-----

``` sh
$ hub clone rtomayko/tilt

# expands to:
#=> git clone git://github.com/rtomayko/tilt.git
```

hub can be safely [aliased](#aliasing) as `git` so you can type `$ git
<command>` in the shell and get all the usual `hub` features.

See [Usage documentation](https://hub.github.com/hub.1.html) for the list of all
commands and their arguments.

Hub can also be used to make shell scripts that [manually interface with the
GitHub API](https://hub.github.com/hub-api.1.html).

Installation
------------

The `hub` executable has no dependencies, but since it was designed to wrap
`git`, it's recommended to have at least **git 1.7.3** or newer.

platform | manager | command to run
---------|---------|---------------
macOS, Linux | [Homebrew](https://docs.brew.sh/Installation) | `brew install hub`
Windows | [Scoop](http://scoop.sh/) | `scoop install hub`
Windows | [Chocolatey](https://chocolatey.org/) | `choco install hub`
Fedora Linux | [DNF](https://fedoraproject.org/wiki/DNF) | `sudo dnf install hub`
Arch Linux | [pacman](https://wiki.archlinux.org/index.php/pacman) | `sudo pacman -S hub`
FreeBSD | [pkg(8)](http://man.freebsd.org/pkg/8) | `pkg install hub`
Debian | [apt(8)](https://manpages.debian.org/buster/apt/apt.8.en.html) | `sudo apt install hub`
Ubuntu | [Snap](https://snapcraft.io) | `snap install hub --classic`

#### Standalone

`hub` can be easily installed as an executable. Download the latest
[compiled binaries](https://github.com/github/hub/releases) and put it anywhere
in your executable path.

#### Source

Prerequisites for building from source are:

* `make`
* [Go 1.9+](https://golang.org/doc/install)

Clone this repository and run `make install`:

```sh
git clone \
  --config transfer.fsckobjects=false \
  --config receive.fsckobjects=false \
  --config fetch.fsckobjects=false \
  https://github.com/github/hub.git

cd hub
make install prefix=/usr/local
```

This assumes support for [Go 1.11+
modules](https://github.com/golang/go/wiki/Modules). If you are building on an
older version of Go, you will need to clone the repository into
`$GOPATH/src/github.com/github/hub`.

Aliasing
--------

Some hub features feel best when it's aliased as `git`. This is not dangerous; your
_normal git commands will all work_. hub merely adds some sugar.

`hub alias` displays instructions for the current shell. With the `-s` flag, it
outputs a script suitable for `eval`.

You should place this command in your `.bash_profile` or other startup script:

``` sh
eval "$(hub alias -s)"
```

#### PowerShell

If you're using PowerShell, you can set an alias for `hub` by placing the
following in your PowerShell profile (usually
`~/Documents/WindowsPowerShell/Microsoft.PowerShell_profile.ps1`):

``` sh
Set-Alias git hub
```

A simple way to do this is to run the following from the PowerShell prompt:

``` sh
Add-Content $PROFILE "`nSet-Alias git hub"
```

Note: You'll need to restart your PowerShell console in order for the changes to be picked up.

If your PowerShell profile doesn't exist, you can create it by running the following:

``` sh
New-Item -Type file -Force $PROFILE
```

### Shell tab-completion

hub repository contains tab-completion scripts for bash, zsh and fish.
These scripts complement existing completion scripts that ship with git.

[Installation instructions](etc)

* [hub bash completion](https://github.com/github/hub/blob/master/etc/hub.bash_completion.sh)
* [hub zsh completion](https://github.com/github/hub/blob/master/etc/hub.zsh_completion)
* [hub fish completion](https://github.com/github/hub/blob/master/etc/hub.fish_completion)

Meta
----

* Home: <https://github.com/github/hub>
* Bugs: <https://github.com/github/hub/issues>
* Authors: <https://github.com/github/hub/contributors>
