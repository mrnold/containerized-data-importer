package ovirtclient

import (
	"encoding/binary"
	"io"
)

func extractQCOWParameters(fileSize uint64, reader readSeekCloser) (
	ImageFormat,
	error,
) {
	format := ImageFormatCow
	qcowSize := fileSize
	header := make([]byte, qcowHeaderSize)

	_, err := io.ReadAtLeast(reader, header, qcowHeaderSize)
	if err != nil {
		return "", wrap(err, EBadArgument, "failed to read QCOW header")
	}

	isQCOW := string(header[0:len(qcowMagicBytes)]) == qcowMagicBytes
	if !isQCOW {
		format = ImageFormatRaw
	} else {
		// See https://people.gnome.org/~markmc/qcow-image-format.html
		qcowSize = binary.BigEndian.Uint64(header[qcowSizeStartByte : qcowSizeStartByte+8])
	}
	if qcowSize <= 0 {
		return format, newError(EBadArgument, "expected positive image size, got %d instead", qcowSize)
	}
	return format, err
}
