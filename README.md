# Golab

[![Build Status](https://travis-ci.com/cloudingcity/glcli.svg?branch=master)](https://travis-ci.com/cloudingcity/glcli)
[![codecov](https://codecov.io/gh/cloudingcity/glcli/branch/master/graph/badge.svg)](https://codecov.io/gh/cloudingcity/glcli)
[![Go Report Card](https://goreportcard.com/badge/github.com/alekseiapa/glcli)](https://goreportcard.com/report/github.com/alekseiapa/glcli)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](http://godoc.org/github.com/alekseiapa/glcli)

`glcli` is a command line tool that make working with GitLab easier.

## Features

- Easy to configure gitlab settings
- Show all merge requests that created by you or assigned to you
- Open page in default browser

## Installation

### Homebrew

```shell script
brew install cloudingcity/tap/glcli
```

### Source

```shell script
git clone git@github.com:cloudingcity/glcli.git
cd glcli
make install
```

## Command Usage

```
A CLI tool for gitlab

Usage:
  glcli [command]

Available Commands:
  config      Init or list glcli CLI options
  mr          Manage merge requests
  own         Manage own resources
  version     Print version number of glcli

Flags:
  -h, --help   help for glcli

Use "glcli [command] --help" for more information about a command.
```

### Initial config

```shell script
glcli config init                                                                                                    master ↓ 1 ↑ 1 ✚ 1 
```

### Show current repository merge requests
```shell script
glcli mr list
```

### Show all merge requests that assigned to you
```shell script
glcli own mr list --review
```

### Open merge requests page in browser
```shell script
glcli mr open <MR-ID>
```
