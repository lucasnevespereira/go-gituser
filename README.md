<img src="inline-logo.png" alt="logo" width="200" />

## Overview

User of multiple git accounts, in order to meet my need to switch regularly between these accounts (student, professional, personal), I developed an open source cli (a command line interface) in Golang.

This program helps switch between different git user accounts easily.

![](demo.gif)

It automates the following commands:

```
git config --global user.name "yourUsername"
```

```
git config --global user.email "yourEmail"
```

#### Modes

There is currently 3 modes in this script:

- "work" : for a work related git account.
- "school" : for a school related git account.
- "personal" : for a personal related git account.

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
<i>Other OS coming soon..</i>

<!--
[Install for MacOS](MACOS_PATH.md) </br>
Other OS - coming soon.. -->

The goal of adding the program to the <b>PATH</b> is to be able to call `gituser` globally in your machine.

To add the program to your path on <b>MacOS</b> do the following:

Go to the directory where you keep the program

<small>Example: </small>

```
cd ~/projects/go-gituser/
```

If you have not build the program build it

```
go build -o gituser
```

Edit your `.bash_profile` or `.zshrc` if you use zsh

```
nano ~/.bash_profile
```

Add the path of your project and the path of your `data/config.json`

<small>Example: </small>

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

Examples:

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

The flag `--help` will print some information about the program.

`gituser --help`

The flag `--info` that will print some information about the accounts.

`gituser --info`

<hr>

## License

This project is under [MIT LICENSE](LICENSE)
