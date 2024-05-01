# Programming jokes generator

Programming jokes generator command line tool

## Installation

```bash
go install github.com/nguyen-duc-loc/joke-cli@latest
```

## Usage

```
$ joke-cli -h
A command line tool provides programming jokes built with love by nguyenducloc in Go
https://github.com/nguyen-duc-loc/joke-cli

Usage:
    joke-cli [command]

Available Commands:
    generate    Generate proramming jokes
    help        Help about any command

Flags:
    -h, --help   help for joke-cli

Use "joke-cli [command] --help" for more information about a command.
```

```
$ joke-cli generate -h
Generate one or more proramming jokes. For example:
joke-cli generate -n 5

Usage:
    joke-cli generate [flags]

Flags:
    -h, --help           help for generate
    -n, --number uint8   Number of joke (1 - 10) (default 1)
```

## Example

```
$ joke-cli generate -n 3
[+] Debugging is like being the detective in a crime movie where you're also the murderer at the same time.

[+] I've got a really good UDP joke to tell you but I don’t know if you'll get it.

[+] Two SQL tables sit at the bar. A query approaches and asks "Can I join you?"
```

## Authors

- [@nguyenducloc](https://www.linkedin.com/in/nguyenducloc404/)
