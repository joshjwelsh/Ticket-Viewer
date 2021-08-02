package main

// device handles user input

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type ReadDevice struct {
	io.Reader
	Input string
}

func CreateDevice(stdin io.Reader) ReadDevice {
	return ReadDevice{
		Reader: stdin,
	}
}

func (r *ReadDevice) GetInput() error {
	var err error
	reader := bufio.NewReader(r.Reader)
	if r.Reader == os.Stdin {
		text, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("Read device error: %v with io.Reader %v", err, r.Reader)
		}
		r.Input = text

	} else {

		buf := new(strings.Builder)
		_, err := io.Copy(buf, reader)
		if err != nil {
			return fmt.Errorf("Read device error: %v with io.Reader %v", err, r.Reader)
		}
		r.Input = buf.String()
	}

	return err

}
