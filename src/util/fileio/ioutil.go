package fileio

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// DiscardReader implements io.Reader interface, does nothing
type DiscardReader struct{}

func (d DiscardReader) Read(b []byte) (n int, err error) { return }

// ReadFileToLines reads in a file line by line and returns an array of strings,
// trims spaces at begin and end of each line if trimSpaces flag is set
func ReadFileToLines(filename string, trimSpaces bool) ([]string, error) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if trimSpaces {
			text = strings.TrimSpace(text)
		}
		lines = append(lines, text)
	}
	return lines, scanner.Err()
}

// GetPDDLFiles returns all files in a given folder that contain the word
// `problem` or `domain`. An error is returned if a file is encountered that
// does not contain either domain or problem in its filename OR if the number of
// domain files does not match the number of domain files.
func GetPDDLFiles(folder string) (domfiles, probfiles []os.FileInfo, err error) {
	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		if strings.HasPrefix(f.Name(), ".") {
			continue
		}
		if strings.Contains(strings.ToLower(f.Name()), "domain") {
			domfiles = append(domfiles, f)
		} else if strings.Contains(strings.ToLower(f.Name()), "problem") {
			probfiles = append(probfiles, f)
		} else {
			err = fmt.Errorf("file: %s cannot decide type (problem|domain)\n", f.Name())
		}
	}
	if len(domfiles) != len(probfiles) {
		err = fmt.Errorf("#dom files != #prob files: %v vs %v\n", domfiles, probfiles)
	}
	return
}
