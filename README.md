# LEAP

Bookmark your directories and `cd` to them instantly. Keeps a hidden file in your home directory named .leap to keep track of all of the places that you can leap to.

This is not 100% in Go due to the limitation that a subprocess is unable to change its parent's working directory. Therefore, if you run the program as originally intended (executing `cd` or using os.Chdir()) from a terminal, the subprocess that spawns and runs the program will change its working directory and exit, returning to the parent process whose working directory remains unaffected.

The workaround? Create a bash function that will call the binary and parse the results. For now, the bash function passes arguments to the program, which in turn will return a single string as an argument to `cd`.

# Installation

Add the bash function to your `.bashrc` or `.zshrc` file. As you can see, it's calling the binary from the workspace.
```
function leap() {
    local LEAP
    LEAP="$GOPATH/bin/leap"

    if [ $# -lt 1 ]; then
	$LEAP
    elif [ $1 = "add" ] || [ $1 = "rm" ] || [ $1 = "list" ] || [ $1 = "help" ]; then
	$LEAP $@
    else
        cd $($LEAP "$@")
    fi
}

```

# Usage

Create a Place, which is a combo of the directory to change to, and an alias that leap will use:
```
leap add <DIR> <ALIAS>
```

Example:
```
leap new ./ home
```

Leap to somewhere via the alias you created:
```
leap <ALIAS>
```

Example:
```
leap home
```

# Dev Notes

- The package contains an init function that will attempt to create a hidden file in a User's home directory called `.leap`.
- Until `TODO` bullet #1 is implemented, there is no erroneous or informational output. Incorrect usage will simply return "./". This is essentially to negate `cd`.

Helpful links:
http://unix.stackexchange.com/questions/259460/how-can-i-parse-a-multi-line-command-output-in-bash
http://unix.stackexchange.com/questions/78470/pass-arguments-to-function-exactly-as-is
http://zsh.sourceforge.net/Intro/intro_4.html
https://groups.google.com/forum/#!topic/golang-nuts/8o7S3fq5fN8
http://stackoverflow.com/questions/255414/why-doesnt-cd-work-in-a-bash-shell-script
http://stackoverflow.com/questions/17026290/golang-chdir-and-stay-there-on-program-termination

# TODO

- Modify the bash function to parse a result. This way we can implement `leap list`, `leap remove`, and print out error messages.
- Invalidate addition of a Place with a duplicate alias
- Implement `leap list` which will list all of the places you can leap to.
