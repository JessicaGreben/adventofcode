# adventofcode

Advent of Code solutions

https://adventofcode.com

- 2019 https://adventofcode.com/2019
- 2021 https://adventofcode.com/2021
- 2022 https://adventofcode.com/2022
- 2023 https://adventofcode.com/2023
- 2024 https://adventofcode.com/2024

## Setup

To setup a directory for a new problem, copy the template directory with this `make` command:

```
$ make copy year=2023 day=01
```

This creates a new directory at the path `<year>/<day>/<partX>` for part 1 and part 2 of the problem.

## Tests

Run unit tests for a problem on a specific day and part

```
$ go test ./2024/02/part1 -v
```

Get test coverage info for a problem on a day and part
```
$ go test -coverprofile cover.out ./2024/02/part2/
$ go tool cover -html=cover.out
```
