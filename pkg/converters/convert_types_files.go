package converters

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
)

func FileHeaderToBytes(file *multipart.FileHeader) ([]byte, error) {
	fileConverted, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("fail on converted fileHeader: %w", err)
	}
	defer fileConverted.Close()

	buff := bytes.NewBuffer(nil)
	if _, err := io.Copy(buff, fileConverted); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
