# LEAP

Bookmark your directories and `cd` to them instantly. Keeps a hidden file in your home directory named .leap to keep track of all of the places that you can leap to.

This is not 100% in Go due to the limitation that a subprocess is unable to change its parent's working directory. Therefore, if you run the program as originally intended (executing `cd` or using os.Chdir()) from a terminal, the subprocess that spawns and runs the program will change its working directory and exit, returning to the parent process whose working directory remains unaffected.

The workaround? Create a bash function that will call the binary and parse the results. For now, the bash function passes arguments to the program, which in turn will return a single string as an argument to `cd`.

# Installation

Add the bash function to your `.bashrc` or `.zshrc` file. As you can see, it's calling the binary from the workspace.
```
funcion leap() {
	cd $($GOPATH/bin/leap "$@")
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

# TODO

- Modify the bash function to parse a result. This way we can implement `leap list`, `leap remove`, and print out error messages.
- Invalidate addition of a Place with a duplicate alias
- Implement `leap list` which will list all of the places you can leap to.
