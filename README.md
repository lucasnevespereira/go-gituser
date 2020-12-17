<img src="assets/inline-logo.png" alt="logo" width="200" />

## Overview

User of multiple git accounts, in order to meet my need to switch regularly between these accounts (student, professional, personal), I developed an open source cli (a command line interface) in Golang.

This program helps switch between different git user accounts easily.

![](assets/demo.gif)

It automates the following commands:

```
git config --global user.name "yourUsername"
```

```
git config --global user.email "yourEmail"
```

#### Modes

There is currently 3 modes in this script:

- 💻 <b>work</b> : for a work related git account.
- 📚 <b>school</b> : for a school related git account.
- 🏠 <b>personal</b> : for a personal related git account.

## How to install

### Prerequesites

At this moment to use this program you need to have `go`.

Install with Homebrew

```
brew update
```

```
brew install go
```

or visit https://golang.org/doc/install

### Add your git data

To add your respective accounts, you need to fill out the `data/config.json` file.

```
{
  "personalUsername": "enterYourUsernameHere",
  "personalEmail": "enterYourEmailHere",
  "schoolUsername": "enterYourUsernameHere",
  "schoolEmail": "enterYourEmailHere",
  "workUsername": "enterYourUsernameHere",
  "workEmail": "enterYourEmailHere"
}

```

### Install Globally

Set up the program to run globally on your machine (MacOs). <br>
<em><i>Other OS coming soon..</i></em>

<!--
[Install for MacOS](MACOS_PATH.md) </br>
Other OS - coming soon.. -->

The goal of adding the program to the <b>PATH</b> is to be able to call `gituser` globally in your machine.

To add the program to your path on <b>MacOS</b> do the following:

Go to the directory where you keep the program

<em>Example: </em>

```
cd ~/projects/go-gituser/
```

You can compile your program now or later, by running :

```
go build -o gituser
```

Edit your `.bash_profile` or `.zshrc` if you use zsh

```
nano ~/.bash_profile
```

Add the path of <b>this project</b> (go-gituser) and the path of `data/config.json`

<em>Example: </em>

```
# GitUser program
export PATH=~/projects/go-gituser/:$PATH
export PATH_TO_GITUSER_CONFIG=~/projects/go-gituser/data/config.json
```

Save the file and exit.

Reopen a terminal window or source your bash_profile

```
source ~/.bash_profile
```

And now you can call `gituser` globally 😀

## Usage

<i>Attention: </i> Make sure you've entered your information in `config.json` before compile program

Compile by running `go build -o gituser`

Call executable file with mode

```
gituser <mode>
```

<em>Examples: </em>

```
gituser work
```

```
gituser school
```

```
gituser personal
```

#### Flags

The flag `--help` or `-help` is a default flag in go that prints existing flags.

The flag `--manual` or `-manual` will print some information about the program.

The flag `--info` or `-info` that will print some information about the accounts.

```
gituser <flag>
```

<hr>

## License

This project is under [MIT LICENSE](LICENSE)
