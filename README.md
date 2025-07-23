
<div align="center">
    <img src="assets/inline-logo.png" alt="logo" width="320" />
    <h1>GitUser CLI</h1>
    <em>Switch between git accounts easily</em>
<p>
    <a href="https://github.com/lucasnevespereira/go-gituser/releases/latest"><img alt="GitHub release" src="https://img.shields.io/github/v/release/lucasnevespereira/go-gituser.svg?logo=github&style=flat-square"></a>
    <a href="https://github.com/lucasnevespereira/go-gituser/actions/workflows/release.yml"><img alt="GitHub release" src="https://github.com/lucasnevespereira/go-gituser/actions/workflows/release.yml/badge.svg"></a>
    <a href="https://github.com/lucasnevespereira/go-gituser">
          <img alt="Stars" src="https://img.shields.io/github/stars/lucasnevespereira/go-gituser?style=flat-square&logo=github">
    </a>
    <img src="https://img.shields.io/github/issues/lucasnevespereira/go-gituser" alt="Issues">
    <a href="https://github.com/lucasnevespereira/go-gituser">
          <img alt="Forks" src="https://img.shields.io/github/forks/lucasnevespereira/go-gituser?style=flat-square&logo=github">
    </a>
    <img src="https://img.shields.io/github/contributors/lucasnevespereira/go-gituser?style=plastic" alt="Contributors">
    <a href="LICENSE">
      <img src="https://img.shields.io/badge/License-MIT-green.svg" alt="MIT License">
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

As a user of multiple git accounts, I needed to switch regularly between my student, professional, and personal accounts. To solve this, I build this open source CLI tool.

GitUser helps you switch between different git accounts effortlessly, automating all the configuration commands.

![](assets/demo.gif)

## What It Automates

```bash
git config --global user.name "yourUsername"
git config --global user.email "yourEmail"
git config --global user.signingkey "yourSigningKeyID"
git config --global commit.gpgsign true
ssh-add ~/.ssh/your_ssh_key
```

Just run: `gituser work` ✨

#### Account Modes

There is currently 3 modes available:

- 💻 <b>work</b> : for a work related git account.
- 📚 <b>school</b> : for a school related git account.
- 🏠 <b>personal</b> : for a personal related git account.


#### Features

- **🔄 Instant Account Switching** - Switch between work, school, and personal accounts
- **🔧 Complete Git Configuration** - Manages username, email, and GPG signing
- **🗝️ SSH Key Management** - Automatically loads the correct SSH key for each account
- **🎯 Interactive Setup** - Guided wizard to configure all your accounts
- **🛡️ Secure** - Each account uses its own SSH key for isolation
- **🚀 Zero Configuration After Setup** - One command switches everything



## Installation

### Homebrew (Recommended)

```
brew tap lucasnevespereira/tools
```

```
brew install gituser
```

or
```
brew install lucasnevespereira/tools/gituser
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

Now you can call `gituser` globally 😀

### Add your git account data

Run the following command :

```
gituser setup
```

<em>This command will help you setup your different git accounts. </em>

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

### Commands


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


### Advanced Features

### SSH Key Management

GitUser automatically handles SSH keys for each account:

- **Auto-discovery** - Finds existing SSH keys on your system
- **Key generation** - Helps create new SSH keys during setup
- **Automatic switching** - Loads the correct SSH key when switching accounts
- **Connection testing** - Verifies SSH setup works with GitHub/GitLab

### GPG Signing

Configure different GPG keys for each account to enable signed commits with proper identity verification.

When you run `gituser setup`, you can:
- Use existing SSH keys
- Generate new SSH keys with guided setup
- Configure different keys for each account
- Skip SSH setup if you prefer manual management

## Contributing

If you want to contribute to this project please read the [Contribution Guide](CONTRIBUTING.md).

<hr>

## License

This project is under [MIT LICENSE](LICENSE)

<div align="center">
  <strong>⭐ If GitUser helps you, please consider giving it a star!</strong>
</div>
