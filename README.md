# Experimenting with Vedis from C and Go

Just playing around with Vedis and refreshing my rusty C programming skills.

## Vedis?

[Vedis](http://vedis.symisc.net/) is an embeddable datastore C library
built with over 70 [commands](http://vedis.symisc.net/commands.html)
similar in concept to [Redis](http://redis.io/) but without the networking
layer since Vedis run in the same process of the host application.

## Requirements

The CGO bindings to Vedis:

``` sh
$ go get github.com/icholy/vedis
```

## Usage

``` sh
$ make
Building the examples:
go build -o bin/example-from-go example.go
gcc example.c vedis.c -o bin/example-from-c

Running C example:

 foo: DATA FROM C

Running Go example:

 foo: DATA FROM C
 bar: 123
```
