# go-blurhash-experiment

This repository is an attempt to create CLI for [go-blurhash](https://github.com/bbrks/go-blurhash) library, mainly to try it out.

It contains source code for 2 binaries: encoder and decoder.

To build binaries locally run:

```bash
make build-encoder build-decoder
```

## Encoder

Encoder takes an image(-s) na generates "blurhash" for it.

```bash
./encoder -silent images/*.jpg
AKA].{xZ0LRk
AsGc7J%M_Noe

# or encoder can read image from standard input
cat images/photo_1.jpg | ./encoder -silent 
AKA].{xZ0LRk
```

See help:

```bash
./encoder --help
Usage of ./encoder:
  -debug
        Enables debug mode
  -silent
        Only return blurhash(-s) to standard output
  -x int
        Number of X components from 1 (more blur, shorter hash, faster) to 9 (less blur, longer hash, slower) (default 2)
  -y int
        Number of Y components from 1 (more blur, shorter hash, faster) to 9 (less blur, longer hash, slower) (default 2)
```

## Decoder

Decoder takes "blurhash" and creates file with blurred image in there.

```bash
./decoder -silent -o output.jpg -p 5 -w 640 -h 960 "AKA].{xZ0LRk"
```

Or you can pipe one to another:

```
cat images/photo_1.jpg | ./encoder -silent -x 6 -y 9 | ./decoder -silent -o output.jpg -p 1 -w 640 -h 960
```

```bash
Usage of ./decoder:
  -debug
        Enables debug mode
  -h int
        Result blur image height (default 150)
  -o string
        Destination filename
  -p int
        Punch (default 2)
  -silent
        Only return errors to standard output
  -w int
        Result blur image width (default 200)
```