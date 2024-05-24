package tcp

import (
	"encoding/binary"
	"io"

	log "github.com/sirupsen/logrus"
)

func Encode(w io.Writer, encoded []byte) error {

	len := uint32(len(encoded))

	len += 4

	if err := binary.Write(w, binary.BigEndian, len); err != nil {
		return err
	}

	if _, err := w.Write(encoded); err != nil {
		return err
	}

	return nil
}

func Decode(r io.Reader) ([]byte, error) {
	var len uint32
	if err := binary.Read(r, binary.BigEndian, &len); err != nil {
		log.WithError(err).Error("Error reading length")
		return nil, err
	}

	len -= 4

	payload := make([]byte, len)

	_, err := io.ReadFull(r, payload)
	if err != nil {
		log.WithError(err).Error("Error reading payload")
		return nil, err
	}

	return payload, nil
}
