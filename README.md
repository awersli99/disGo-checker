# DisGo Checker

DisGo Checker is a command line tool for checking the validity of discord tokens.

## Installation

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
./disgo-checker -t 8 -o tokens.txt -i unchecked.txt -b broken.txt
```

