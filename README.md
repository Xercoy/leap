# LEAP

A clone of hop in Go, made for linux. Keeps a hidden file in your home directory called .leap to keep track of all of the places that you can hop to.

# Usage

Create places to leap to:
```
leap new ./ home
```

List all of the places you can leap to:
```
leap list
```

Leap to somewhere via the alias you created:
```
leap home
```

# Notes

The package contains an init function that will attempt to create a hidden file in a User's home directory called `.leap`

