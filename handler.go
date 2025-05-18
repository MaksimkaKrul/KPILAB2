package lab2

import (
	"io"
	"io/ioutil"
	"strconv"
)

type ComputeHandler struct {
	Input  io.Reader
	Output io.Writer
}

func (ch *ComputeHandler) Compute() error {
	data, err := ioutil.ReadAll(ch.Input)
	if err != nil {
		return err
	}

	expression := string(data)
	result, err := EvaluatePrefix(expression)
	if err != nil {
		return err
	}

	_, err = ch.Output.Write([]byte(strconv.Itoa(result)))
	return err
}
