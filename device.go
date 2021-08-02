package main

// device handles user input

import (
	"bufio"
	"fmt"
	"io"
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

	reader := bufio.NewReader(r.Reader)

	text, err := reader.ReadString('\n')

	if err != nil {
		return fmt.Errorf("Read device error: %v with io.Reader %v", err, r.Reader)
	}

	r.Input = text
	return nil

}
