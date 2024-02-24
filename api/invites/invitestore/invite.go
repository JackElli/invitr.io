package invitestore

import qrcodes "invitio.com/qr-codes/endpoints/qr-codes"

type Invite struct {
	Id         string `json:"id"`
	Organiser  string `json:"organiser"`
	Location   string `json:"location"`
	Date       string `json:"date"`
	Passphrase string `json:"passphrase"`
}

type InviteDB struct {
	Invite
	QRCode string `json:"qr_code,omitempty"`
}

type InviteJSON struct {
	Invite
	QRCode qrcodes.QRCode `json:"qr_code,omitempty"`
}
