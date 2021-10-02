## gopass

gopass is a small password generation ulility for the command line, written in go.

usage:
```
SYNOPSIS
gopass [options]

OPTIONS
	--all			Use all character types when generating a password
	--len			Number of characters in the password
	-l 			Use letters in password
	-n 			Use numbers in password
	-s 			Use symbols in password
	--use-upper-case	Include upper case letters in password
```

## build
to build gopass, glone the repo
```
git clone https://github.com/theory-of-everything/gopass
```
update/install dependencies
```
go mod tidy
```
build (spits out binary `gopass` in project root)
```
make build
```
