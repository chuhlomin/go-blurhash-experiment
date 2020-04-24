package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/bbrks/go-blurhash"
	log "github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
)

var (
	xComponentsPtr = flag.Int("x", 2, "Number of X components from 1 (more blur, shorter hash, faster) to 9 (less blur, longer hash, slower)")
	yComponentsPtr = flag.Int("y", 2, "Number of Y components from 1 (more blur, shorter hash, faster) to 9 (less blur, longer hash, slower)")
	silentPtr      = flag.Bool("silent", false, "Only return blurhash(-s) to standard output")
	debugPtr       = flag.Bool("debug", false, "Enables debug mode")
)

func main() {
	flag.Parse()

	filenames := flag.Args()

	if *debugPtr {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
	} else {
		log.Setup(log.Msec, log.LevelBraces)
	}
	if *silentPtr {
		log.Setup(log.Out(ioutil.Discard))
	}

	log.Printf("DEBUG xComponents: %d, yComponents: %d, filenames: %v", *xComponentsPtr, *yComponentsPtr, filenames)

	log.Printf("INFO Starting...")

	if err := run(filenames); err != nil {
		log.Fatalf("ERROR %v", err)
	}

	log.Printf("INFO Finished.")
}

func run(filenames []string) error {
	if len(filenames) == 0 {
		reader := bufio.NewReader(os.Stdin)
		hash, err := getHashForReader(reader)
		if err != nil {
			return errors.Wrap(err, "Blurhash failed for image in standard input")
		}

		log.Printf("INFO Blurhash for image from standard input: %s", hash)
		if *silentPtr {
			fmt.Println(hash)
		}
		return nil
	}

	for _, filename := range filenames {
		hash, err := getHashForFile(filename)
		if err != nil {
			log.Printf("ERROR Blurhash failed for %s: %v", filename, err)
			continue
		}

		log.Printf("INFO Blurhash for %s is %s", filename, hash)
		if *silentPtr {
			fmt.Println(hash)
		}
	}

	return nil
}

func getHashForFile(filename string) (string, error) {
	reader, err := os.Open(filename)
	if err != nil {
		return "", errors.Wrapf(err, "failed to open file")
	}
	defer reader.Close()

	return getHashForReader(reader)
}

func getHashForReader(reader io.Reader) (string, error) {
	m, _, err := image.Decode(reader)
	if err != nil {
		return "", err
	}

	return blurhash.Encode(*xComponentsPtr, *yComponentsPtr, m)
}
