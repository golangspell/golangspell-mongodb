# golangspell-mongodb
Golang Spell MongoDB is the spell designed for adding the MongoDB infrastructure and application logic to your Golangspell Microservices

## Golang Spell
The Golang Spell MongoDB contains the specific commands (and the respective templates) for building MongoDB backed projects in the platform [Golang Spell](https://github.com/golangspell/golangspell).

![alt text](https://golangspell.com/golangspell/blob/master/img/gopher_spell.png?raw=true)

## Test and coverage

Run the tests

```sh 
TESTRUN=true go test ./... -coverprofile=cover.out

go tool cover -html=cover.out
```

Install [golangci-lint](https://github.com/golangci/golangci-lint#install) and run lint:

```sh
golangci-lint run
```

## Install
To install the golangspell-mongodb spell use the command
```sh
golangspell addspell github.com/golangspell/golangspell-mongodb golangspell-mongodb
```

## Update
To update the golangspell-core version use the command
```sh
golangspell updatespell github.com/golangspell/golangspell-mongodb golangspell-mongodb
```

## After installing, the Golangspell MongoDB commands will be available in the platform help:
```sh
golangspell help
```
