package netstring

import (
	"bytes"
	"io"
	"strconv"
)

var lengthTokenDelim = byte(':')
var dataTokenDelim = byte(',')

func Decode(in []byte) ([][]byte, error) {

	var out [][]byte

	var buf = bytes.NewBuffer(in)

	for {

		if lengthToken, err := buf.ReadBytes(lengthTokenDelim); err != nil {

			if err == io.EOF {
				break
			} else {
				return nil, err
			}

		} else if length, err := strconv.ParseInt(string(lengthToken[:len(lengthToken)-1]), 10, 32); err != nil {
			return nil, err
		} else if dataToken, err := buf.ReadBytes(dataTokenDelim); err != nil {
			return nil, err
		} else {
			out = append(out, dataToken[0:length])
		}

	}

	return out, nil

}

func Encode(in ...[]byte) ([]byte, error) {

	var buf bytes.Buffer

	for _, e := range in {

		if _, err := buf.WriteString(strconv.FormatInt(int64(len(e)), 10)); err != nil {
			return nil, err
		} else if err := buf.WriteByte(lengthTokenDelim); err != nil {
			return nil, err
		} else if _, err := buf.Write(e); err != nil {
			return nil, err
		} else if err := buf.WriteByte(dataTokenDelim); err != nil {
			return nil, err
		}

	}

	return buf.Bytes(), nil

}
