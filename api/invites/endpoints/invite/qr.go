package invite

import (
	"encoding/json"
	"net/http"

	qrcodes "invitr.io.com/qr-codes/endpoints/qr-codes"
)

const (
	QR_CODE_URL = "http://qr-codes:3201/qr-code"
)

// GenerateQRCode fetches a QR code from the QR code microservice
func GenerateQRCode() (*qrcodes.QRCode, error) {
	resp, err := http.Get(QR_CODE_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var qrCode qrcodes.QRCode
	json.NewDecoder(resp.Body).Decode(&qrCode)

	return &qrCode, nil
}

// bytesToQR returns a QR code type based on the bytes
// provided
func bytesToQR(b []byte) qrcodes.QRCode {
	var qrcode qrcodes.QRCode
	json.Unmarshal(b, &qrcode)

	return qrcode
}

// qrToBytes returns a byte array based on QR code type
// given
func qrToBytes(qr qrcodes.QRCode) []byte {
	qrcodeBytes, _ := json.Marshal(qr)
	return qrcodeBytes
}
