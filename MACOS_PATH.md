## MacOS PATH

The goal of adding the program to the <b>PATH</b> is to be able to call `gituser` globally in your machine

So to add the program to your path on <b>MacOS</b> do the following.

Go to the directory where you keep the program

```
cd Users/lucas/projects/go-gituser/
```

If you have not build the program build it

```
go build -o temp
```

Edit your `.bash_profile` or `.zshrc` if you use zsh

```
nano ~/.bash_profile
```

Add the following to this file

```
# GitUser program
export PATH=Users/lucas/projects/go-gituser/:$PATH
```

Save the file and exit.

Reopen a terminal window or source your bash_profile

```
source ~/.bash_profile
```

And now you can call `gituser` globally 😀

<hr>

#### Tips

The command `~/` is equivalent to `Users/lucas`, so you can use it instead

`cd ~/`
