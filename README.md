# Gitlab CLI 🌟


`glcli` is a command line tool that makes working with GitLab easier.

---

## ⚙️ Features

- Easy configuration of GitLab settings.
- Display all merge requests created by or assigned to you.
- Quickly open pages in your default browser.
- Identify projects dependent on a certain package (Golang: `go.mod`, PHP: `composer.json`).
- Validate `.gitlab-ci.yml` files with ease.

---

## 🚀 Installation

Clone the repository and use the Makefile provided:

```bash
git clone https://github.com/alekseiapa/glcli.git
cd glcli
make install
```

---

## 📘 Usage Instructions

```
A CLI tool for GitLab management.

Usage:
  glcli [command]

Available Commands:
  ci          Manage GitLab CI
  clone       Clone a repository from GitLab
  config      Configure or display CLI options
  depend      Show projects relying on a particular package
  mr          Manage merge requests
  own         Manage your own resources
  project     Manage GitLab projects
  version     Display version number of glcli

Flags:
  -h, --help   Help information for glcli

Use "glcli [command] --help" for more information about a command.
```

### Initial Configuration 🛠️

To start using `glcli`, run:

```bash
glcli config init
```

The command will prompt you for GitLab host and token information.

### Clone a Repository 📂

Clone a repository from GitLab with ease:

```bash
glcli clone [REPO]
```

### Working with Merge Requests ⚙️

- **Create a Merge Request**
  Push your local branch to the remote and start a merge request:

  ```bash
  glcli mr create
  ```

- **List Merge Requests**
  Display current repository merge requests:

```bash
glcli mr list
```

- **Open a Merge Request**
  Open a merge request page in the default web browser:

  ```bash
  glcli mr open <MR-ID>
  ```

### Analyze Dependencies 🔍

- **Show PHP Composer Dependencies**
  See projects dependent on a specific PHP package:

```bash
glcli depend php package/name --group=group
```

- **Show Go Module Dependencies**
  Display Go projects that depend on a specific module:

```bash
glcli depend go module/path --group=group
  ```

### Validate GitLab CI Config ✅

Check if `.gitlab-ci.yml` files are valid:

```bash
glcli ci lint .gitlab-ci.yml
```

---