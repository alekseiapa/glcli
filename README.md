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
- Shows project which depend on a certain package (Golang: `go.mod`, PHP: `composer.json`)
- Lint `.gtilab-ci.yml`

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
  ci          Manage gitlab ci
  config      Init or list glcli CLI options
  depend      Shows project which depend on a certain package
  mr          Manage merge requests
  own         Manage own resources
  version     Print version number of glcli

Flags:
  -h, --help   help for glcli

Use "glcli [command] --help" for more information about a command.
```

### Initial config

```shell script
$ glcli config init                                                                                                    master ↓ 1 ↑ 1 ✚ 1 
Gitlab Host [https://gitlab.com]: <INPUT>
Create a token here: https://gitlab.com/profile/personal_access_tokens
Gitlab Token (scope: api) [None]: <INPUT>

Config saved to /Users/<USER>/.config/glcli.yaml
```

### Show current repository merge requests
```shell script
$ glcli mr list
  MRID   TITLE                        URL                                                                   
  1      Catch your first Pokémon     https://example.com/pokemon/trainer/merge_requests/1  
  2      To become a Pokémon Master   https://example.com/pokemon/trainer/merge_requests/2  
```

### Show all merge requests that assigned to you
```shell script
$ glcli own mr list --review
  PID    MRID   PROJECT           TITLE                        URL                                                                   
  4255   1      pokemon/trainer   Catch your first Pokémon     https://example.com/pokemon/trainer/merge_requests/1  
  4255   2      pokemon/trainer   To become a Pokémon Master   https://example.com/pokemon/trainer/merge_requests/2  
```

### Open merge requests page in browser
```shell script
$ glcli mr open <MR-ID>
```

### Show `pokemon/eevee` [composer](https://getcomposer.org/) package which project depend on
```shell script
$ glcli depend php pokemon/eevee --group pokemon
  PROJECT    VERSION   BRANCH    URL
  vaporeon   v0.1.2    master    https://example.com/pokemon/vaporeon
  jolteon    v1.2.0    staging   https://example.com/pokemon/jolteon
  flareon    v3.0.0    staging   https://example.com/pokemon/flareon
```

### Show `example.com/pokemon/eevee` [go modules](https://github.com/golang/go/wiki/Modules) which project depend on
```shell script
$ glcli depend go example.com/pokemon/eevee --group pokemon
  PROJECT    VERSION   BRANCH    URL
  vaporeon   v0.1.2    master    https://example.com/pokemon/vaporeon
  jolteon    v1.2.0    staging   https://example.com/pokemon/jolteon
  flareon    v3.0.0    staging   https://example.com/pokemon/flareon
```

### Check `.gitlab-ci.yml` is valid

```shell script
$ glcli ci lint .gitlab-ci.yml
Valid!
```
