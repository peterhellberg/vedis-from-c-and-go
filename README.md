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

### Contents of the data.vdb

``` sh
$ hexdump -C data.vdb
00000000  53 79 6d 69 73 63 56 65  64 69 73 ca 1d b6 34 44  |SymiscVedis...4D|
00000010  c9 9a b4 00 00 02 00 00  00 10 00 00 04 68 61 73  |.............has|
00000020  68 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |h...............|
00000030  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00001000  de 67 1c ef 03 f5 23 f5  00 00 00 00 00 00 00 00  |.g....#.........|
00001010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 01  |................|
00001020  00 00 00 00 00 00 00 00  00 00 00 01 00 00 00 00  |................|
00001030  00 00 00 00 00 00 00 00  00 00 00 02 00 00 00 00  |................|
00001040  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00002000  00 0c 00 75 00 00 00 00  00 00 00 00 0b 88 60 ba  |...u..........`.|
00002010  00 00 00 03 00 00 00 00  00 00 00 03 00 2c 00 00  |.............,..|
00002020  00 00 00 00 00 00 62 61  72 31 32 33 00 02 b6 1d  |......bar123....|
00002030  00 00 00 01 00 00 00 00  00 00 00 03 00 4d 00 00  |.............M..|
00002040  00 00 00 00 00 00 78 31  32 33 41 54 41 0b 88 73  |......x123ATA..s|
00002050  89 00 00 00 03 00 00 00  00 00 00 00 0b 00 00 00  |................|
00002060  00 00 00 00 00 00 00 66  6f 6f 44 41 54 41 20 46  |.......fooDATA F|
00002070  52 4f 4d 20 43 00 00 0f  8b 00 00 00 00 00 00 00  |ROM C...........|
00002080  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00003000
```
