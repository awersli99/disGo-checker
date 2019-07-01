# discord-token-checker
A command line discord token checker wrtiten in Go.

Usage of ./discord-token-checker:
  -b string
        [OPTIONAL] The output file for broken tokens
  -i string
        [REQUIRED] Your input file containing unchecked tokens
  -o string
        [REQUIRED] Your output file name that will contain checked tokens
  -t int
        [REQUIRED] The amount of threads the app should run on (At least one) (default 1)
