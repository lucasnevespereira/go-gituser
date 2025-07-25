
<div align="center">
  <img src="assets/inline-logo.png" alt="logo" width="400" style="padding: 20px;" />
  <h3><em>Switch between git accounts easily</em></h3>
  <p>
    <a href="https://github.com/lucasnevespereira/go-gituser/releases/latest">
      <img src="https://img.shields.io/github/v/release/lucasnevespereira/go-gituser?style=flat&logo=github" alt="Latest Release" />
    </a>
    <a href="https://github.com/lucasnevespereira/go-gituser/actions/workflows/release.yml">
      <img src="https://github.com/lucasnevespereira/go-gituser/actions/workflows/release.yml/badge.svg" alt="Build Status" />
    </a>
    <a href="https://github.com/lucasnevespereira/go-gituser/stargazers">
      <img src="https://img.shields.io/github/stars/lucasnevespereira/go-gituser?style=flat&logo=github" alt="Stars" />
    </a>
    <a href="https://github.com/lucasnevespereira/go-gituser/network/members">
      <img src="https://img.shields.io/github/forks/lucasnevespereira/go-gituser?style=flat&logo=github" alt="Forks" />
    </a>
    <a href="https://github.com/lucasnevespereira/go-gituser/issues">
      <img src="https://img.shields.io/github/issues/lucasnevespereira/go-gituser?style=flat&logo=github" alt="Issues" />
    </a>
    <a href="https://github.com/lucasnevespereira/go-gituser/graphs/contributors">
      <img src="https://img.shields.io/github/contributors/lucasnevespereira/go-gituser?style=flat&logo=github" alt="Contributors" />
    </a>
    <a href="https://github.com/sponsors/lucasnevespereira">
      <img src="https://img.shields.io/badge/Sponsor-GitHub-333333?style=flat&logo=github&logoColor=white" alt="Sponsor" />
    </a>
    <a href="LICENSE">
      <img src="https://img.shields.io/badge/License-MIT-green.svg?style=flat" alt="MIT License" />
    </a>
  </p>
</div>


## Table of Contents

- [Overview](#overview)
- [What It Automates](#what-it-automates)
- [Account Modes](#account-modes)
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Commands](#commands)
- [Advanced Features](#advanced-features)
- [Contributing](#contributing)
- [License](#license)

## Overview

As a user of multiple git accounts, I needed to switch regularly between my student, professional, and personal profiles. Manually updating configs, loading SSH keys, and keeping track of identities quickly became tedious.

So I built this open source CLI tool to streamline the whole process.

GitUser helps you switch between different git accounts effortlessly. It automates all the necessary configuration commands: username, email, GPG signing, SSH key loading, so you can focus on coding instead of fiddling with your setup.

Whether you're pushing to school projects, personal repos, or work-related codebases, GitUser makes sure you're always using the right identity with a single command.


<div align="center">
<img src="assets/demo.png" alt="demo" width="600" heigth="400" style="border-radius: 20px; padding: 10px;" />
</div>


## What It Automates

Instead of manually running these commands every time you switch projects:

```bash
git config --global user.name "yourUsername"
git config --global user.email "yourEmail"
git config --global user.signingkey "yourSigningKeyID"
git config --global commit.gpgsign true
ssh-add ~/.ssh/your_ssh_key
```

Just run `gituser work` ✨

## Account Modes

There is currently 3 modes available:

- 💻 <b>work</b> : for a work related git account.
- 📚 <b>school</b> : for a school related git account.
- 🏠 <b>personal</b> : for a personal related git account.


## Features

- **🔄 Instant Account Switching** - Switch between work, school, and personal accounts
- **🔧 Complete Git Configuration** - Manages username, email, and GPG signing
- **🗝️ SSH Key Management** - Automatically loads the correct SSH key for each account
- **🎯 Interactive Setup** - Guided wizard to configure all your accounts
- **🛡️ Secure** - Each account uses its own SSH key for isolation
- **🚀 Zero Configuration After Setup** - One command switches everything



## Installation

### Homebrew (Recommended)

```
brew tap lucasnevespereira/homebrew-tools
```

```
brew install --cask gituser
```

or

```
brew install --cask lucasnevespereira/homebrew-tools/gituser
```

### Manual Installation

Make sur your bin path is in your `$PATH`, you can check in your `.zshrc` or `.bash` file.

_e.g_
```shell
export PATH="$HOME/bin:$PATH"
```

Run the following command from the root of the project:

```
make install
```
<em>This will build gituser and move it to your `$HOME/bin`</em>

## Quick Start

Setup your accounts:

```bash
gituser setup
```

Switch between accounts:

```bash
gituser work      # Switch to work account
gituser personal  # Switch to personal account
gituser school    # Switch to school account
```

Check current account:

```bash
gituser now
```

## Commands


| Command | Description |
|---------|-------------|
| **Account Management** | |
| `gituser setup` | Interactive setup for all accounts (username, email, GPG, SSH) |
| `gituser work` | Switch to work account |
| `gituser personal` | Switch to personal account |
| `gituser school` | Switch to school account |
| **Information** | |
| `gituser now` | Show current active account |
| `gituser info` | Display all configured accounts |
| **SSH Management** | |
| `gituser ssh list` | List SSH keys currently loaded |
| `gituser ssh discover` | Find existing SSH keys on your system |
| `gituser ssh test` | Test SSH connections to GitHub/GitLab |
| `gituser ssh guide` | Show SSH setup guide |
| **Help** | |
| `gituser help` | Show help information |
| `gituser manual` | Show detailed manual |
| `gituser quickstart` | Show quick start guide |


## Advanced Features

### SSH Key Management

GitUser automatically handles SSH keys for each account:

- **Auto-discovery** - Finds existing SSH keys on your system
- **Key generation** - Helps create new SSH keys during setup
- **Automatic switching** - Loads the correct SSH key when switching accounts
- **Connection testing** - Verifies SSH setup works with GitHub/GitLab

### GPG Signing

Configure different GPG keys for each account to enable signed commits with proper identity verification.


## Contributing

If you want to contribute to this project please read the [Contribution Guide](CONTRIBUTING.md).

<hr>

## License

This project is under [MIT LICENSE](LICENSE)

<br />
<div align="center">
  <strong>⭐ If GitUser helps you, please consider giving it a star!</strong>
</div>
