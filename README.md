# golangspell-mongodb
Golang Spell MongoDB is the spell designed for adding the MongoDB infrastructure and application logic to your Golangspell Microservices

## Golang Spell
The Golang Spell MongoDB contains the specific commands (and the respective templates) for building MongoDB backed projects in the platform [Golang Spell](https://github.com/golangspell/golangspell).

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
This action should be done only once. After doing it, the golangspell-mongodb Spell is installed in your computer and ready for being used

## Update
To update the golangspell-core version use the command
```sh
golangspell updatespell github.com/golangspell/golangspell-mongodb golangspell-mongodb
```

## After installing, the Golangspell MongoDB commands will be available in the platform help:
```sh
golangspell help
```

## Before executing the Golangspell MongoDB commands (mongodbinit, for example), it is necessary to be inside the root directory of an application initialized with the command:
```sh
golangspell init [module] [appname]
```
More details can be found with the command:

```sh
golangspell help init
```

Find more insights on how to use **golangspell** here:

* [Golangspell Wiki](https://github.com/golangspell/golangspell/wiki)
* [Golangspell Article - Medium](https://medium.com/golangspell-go-fast-small-and-productive/go-fast-small-and-productive-with-golangspell-be193c65a382)
* [Golangspell Article - HackerNoon](https://hackernoon.com/go-fast-small-and-productive-with-golangspell-d0163zqk)

<p align="center">
    <img src="https://github.com/golangspell/golangspell/blob/master/img/gopher_spell.png" width="350" alt="Golang Spell logo"/>
</p>
