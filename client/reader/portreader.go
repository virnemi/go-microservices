package reader

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"

	"log"
	"os"
	"regexp"

	"../../models"
)

// Reader of the Ports in the provided file.
type PortReader struct {
	Reader *bufio.Reader
	File   *os.File
}

// PortReader's constructor
func NewPortReader() *PortReader {
	portReader := &PortReader{}
	return portReader
}

// Initializes the reader. Returns an error if there is any problem with the file
// (non-existent file or permissions issues, etc).
func (portReader *PortReader) Init(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	portReader.File = file
	portReader.Reader = bufio.NewReader(portReader.File)
	return nil
}

// Scanner-like implementation. Readns a portion of the file until
// it can reproduce a Port object. It returns an error if the reader is not ready
// or the Unmarshal of the content doesn't occour or EOF is reached.
func (portReader *PortReader) NextPort() (models.Port, error) {
	if portReader.Reader == nil {
		return nil, errors.New("Reader not started")
	}
	var portKeyFound = regexp.MustCompile(`("[A-Z]{5}"\: \{)`)
	var portEndFound = regexp.MustCompile(`(  \}(\,|)\n)`)
	content := ""
	startPortStorage := false
	for {
		fb, errf := portReader.Reader.ReadBytes('\n')
		line := string(fb)

		if !startPortStorage && portKeyFound.MatchString(line) {
			startPortStorage = true
		}

		if startPortStorage {
			content += line
		}

		if startPortStorage && portEndFound.MatchString(line) {
			if content[len(content)-2:len(content)-1] == "," {
				content = content[:len(content)-2]
			}

			content = "{\n" + content + "\n}"
			//log.Printf("Port found: %s", content)

			var portFound models.Port
			err := json.Unmarshal([]byte(content), &portFound)
			if err != nil {
				return nil, err
			}
			return portFound, nil
		}

		if errf == io.EOF {
			log.Printf("Content read but ignored: %s\n", content)
			return nil, errf
		}
	}
}

// Closes the reader, removing the file from the disk.
func (portReader *PortReader) Close() {
	if portReader.File != nil {
		filename := portReader.File.Name()
		portReader.File.Close()
		errDel := os.Remove(filename)
		if errDel != nil {
			log.Printf("Error removing temp JSON file: %s\n", filename)
			return
		}
		log.Printf("Temp JSON file removed: %s\n", filename)
	}
}
