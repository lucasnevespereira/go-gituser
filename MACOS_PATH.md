## MacOS PATH

To add `./gituser` to your path on <b>MacOS</b> do the following.

Go to `Users/YourName` you can use this shortcut `~`

```
cd ~
```

Create a directory for your personal scripts

```
mkdir myScripts
```

Edit your `.bash_profile` or `.zshrc` if you use zsh

```
nano ~/.bash_profile
```

Add the following to this file

```
# For Personal Scripts
export PATH=~/myScripts:$PATH
```

Save the file and exit.

Now all you have to do is copy the `gituser` program to your `myScripts` folder.

```
cp yourFolderPath/gituser ~/myScripts/
```

Reopen a terminal window or source your bash_profile

```
source ~/.bash_profile
```

And now you can call `gituser` globally 😀
