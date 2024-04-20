package invites_fetch

import (
	qrcodes "invitr.io.com/services/qr-codes/endpoints/qr-codes"
)

type QRMgrMock struct{}

func NewQRMgrMock() *QRMgrMock {
	return &QRMgrMock{}
}

func (mgr *QRMgrMock) GenerateQRCode() (*qrcodes.QRCode, error) {
	return &qrcodes.QRCode{
		Width:  100,
		Height: 100,
		Pixels: []int{
			0,
		},
	}, nil
}

func (mgr *QRMgrMock) BytesToQR(b []byte) qrcodes.QRCode {
	return qrcodes.QRCode{}
}

func (mgr *QRMgrMock) QrToBytes(qr qrcodes.QRCode) []byte {
	return nil
}
