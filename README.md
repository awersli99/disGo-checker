# disGo checker

DisGo Checker is a command line tool written in Go that was created to quickly and efficiently check the validity of thousands of discord authentication tokens.

## Installation

Install Go here: https://golang.org/doc/install

Use 'go get' to install the discord go package

```bash
go get github.com/bwmarrin/discordgo
```
Then use 'go build' to create an executable.

```bash
go build main.go
```

## Usage

```bash
./{executable name} -t 8 -o tokens.txt -i unchecked.txt -b broken.txt
```

