
<p align="center">
  <img width="256" height="256" src="./assets/gkill-logo.png" />
</p>

<h1 align="center">gkill - GOLang Kill Process</h1>

<p align="center">
  <a href="https://github.com/waldirborbajr/gkill/actions/workflows/ci-cd.yaml">
    <img alt="tests" src="https://github.com/waldirborbajr/gkill/actions/workflows/ci-cd.yaml/badge.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/waldirborbajr/gkill">
    <img alt="goreport" src="https://goreportcard.com/badge/github.com/waldirborbajr/gkill" />
  </a>
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
</p>

**BETA:** This project is in active development. Please check out the issues and contribute if you're interested in helping out.

## gkill
`tl;dr:` `gkill`, a.k.a GOlang Symbolic Link (symlink), is an open-source software built-in with the main aim of being a personal alternative to **GNU Stow**.

As `GNU Stow`, `gkill` is a symlink farm manager which takes distinct packages of software and/or data located in separate directories on the filesystem, and makes them appear to be installed in the same place.

With `gkill` it is eeasy to track and manage configuration files in the user's home directory, especially when coupled with version control systems.

## How to install

### Homebrew

To install gkill, run the following [homebrew](https://brew.sh/) command:

```sh
brew install waldirborbajr/gkill/gkill
```

### Go

Alternatively, you can install gkill using Go's go install command:

```sh
go install github.com/waldirborbajr/gkill@latest
```

This will download and install the latest version of gkill. Make sure that your Go environment is properly set up.

**Note:** Do you want this on another package manager? [Create an issue](https://github.com/waldirborbajr/gkill/issues/new) and let me know!

## How to use

The main goal of `gkill` is to be as simple as that, `easy peasy lemon squeezy`, with few commands and straight to the target.

```sh
# To create a link to $HOMR
gkill l

# To force overwrite existing link : **TODO** not implemented
gkill f -f

# To remove (kill) all symblinks : **TODO** not implemented
gkill k

# To remove a specific symblinks : **TODO** not implemented
gkill r symlink-name


# To print all symlink created : **TODO** not implemented
gkill p
```

## .gkill-ignore`

You can add files/directories to ignore list, so when execute `gkill` the content will no be linked.

```sh
touch .gkill-ignore
```

## Contributing to gkill

If you are interested in contributing to `gkill`, we would love to have your help! You can start by checking out the [ open issues ](https://github.com/waldirborbajr/gkill/issues) on our GitHub repository to see if there is anything you can help with. You can also suggest new features or feel free to create a new feature by opening a new issue.

To contribute code to `gkill`, you will need to fork the repository and create a new branch for your changes. Once you have made your changes, you can submit a pull request for them to be reviewed and merged into the main codebase.

## Contributors

<a href="https://github.com/waldirborbajr/gkill/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=waldirborbajr/gkill" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
