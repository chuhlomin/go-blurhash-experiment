package main

import (
	"bufio"
	"flag"
	"image/jpeg"
	_ "image/png"
	"io/ioutil"
	"os"

	"github.com/bbrks/go-blurhash"
	log "github.com/go-pkgz/lgr"
	"github.com/pkg/errors"
)

var (
	outPtr    = flag.String("o", "", "Destination filename")
	widthPtr  = flag.Int("w", 200, "Result blur image width")
	heightPtr = flag.Int("h", 150, "Result blur image height")
	punchPtr  = flag.Int("p", 1, "Punch")
	silentPtr = flag.Bool("silent", false, "Only return errors to standard output")
	debugPtr  = flag.Bool("debug", false, "Enables debug mode")
)

func main() {
	flag.Parse()

	args := flag.Args()

	if *debugPtr {
		log.Setup(log.Debug, log.CallerFile, log.CallerFunc, log.Msec, log.LevelBraces)
	} else {
		log.Setup(log.Msec, log.LevelBraces)
	}
	if *silentPtr {
		log.Setup(log.Out(ioutil.Discard))
	}

	log.Printf("DEBUG Output: %v [%dx%d], pinch: %d", *outPtr, *widthPtr, *heightPtr, punchPtr)

	log.Printf("INFO Starting...")

	if err := run(*outPtr, args); err != nil {
		log.Fatalf("ERROR %v", err)
	}

	log.Printf("INFO Finished.")
}

func run(out string, args []string) error {
	if out == "" {
		return errors.New("Required flag \"o\" was not set. Add \"--help\" to see more information about flags.")
	}

	if len(args) == 0 {
		reader := bufio.NewReader(os.Stdin)
		line, _, err := reader.ReadLine()
		if err != nil {
			return errors.Wrap(err, "failed read blurhash from standard input")
		}

		return writeImageToFile(out, string(line))
	}

	for _, arg := range args {
		err := writeImageToFile(out, arg)
		if err != nil {
			log.Printf("ERROR failed to decode hash %s: %v", arg, err)
			continue
		}
	}

	return nil
}

func writeImageToFile(filename, hash string) error {
	writer, err := os.Create(filename)
	if err != nil {
		return errors.Wrapf(err, "failed to open file")
	}
	defer writer.Close()

	m, err := blurhash.Decode(hash, *widthPtr, *heightPtr, *punchPtr)
	if err != nil {
		return errors.Wrap(err, "failed to decode blurhash")
	}

	if err := jpeg.Encode(writer, m, &jpeg.Options{Quality: 90}); err != nil {
		return errors.Wrap(err, "failed to encode image to JPEG")
	}

	return nil
}
