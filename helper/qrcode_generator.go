package helper

import (
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

func QRCodeGenerator(token string) ([]byte, error) {

	qr, err := qrcode.Encode(token, qrcode.Low, 256)
	if err != nil {
		return []byte{}, errors.Wrap(err, " Error while encode qrcode")
	}

	return qr, nil
}
