package invites_fetch

import (
	"encoding/json"
	"net/http"

	qrcodes "invitr.io.com/services/qr-codes/endpoints/qr-codes"
)

const QR_CODE_URL = "http://qr-codes:3201/qr-code"

type IQRMgr interface {
	GenerateQRCode() (*qrcodes.QRCode, error)
	BytesToQR(b []byte) qrcodes.QRCode
	QrToBytes(qr qrcodes.QRCode) []byte
}

type QRMgr struct{}

func NewQRMgr() *QRMgr {
	return &QRMgr{}
}

// GenerateQRCode fetches a QR code from the QR code microservice
func (mgr *QRMgr) GenerateQRCode() (*qrcodes.QRCode, error) {
	resp, err := http.Get(QR_CODE_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var qrCode qrcodes.QRCode
	json.NewDecoder(resp.Body).Decode(&qrCode)

	return &qrCode, nil
}

// BytesToQR returns a QR code type based on the bytes
// provided
func (mgr *QRMgr) BytesToQR(b []byte) qrcodes.QRCode {
	var qrcode qrcodes.QRCode
	json.Unmarshal(b, &qrcode)

	return qrcode
}

// QrToBytes returns a byte array based on QR code type
// given
func (mgr *QRMgr) QrToBytes(qr qrcodes.QRCode) []byte {
	qrcodeBytes, _ := json.Marshal(qr)
	return qrcodeBytes
}
