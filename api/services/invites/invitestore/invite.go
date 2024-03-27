package invitestore

import qrcodes "invitr.io.com/services/qr-codes/endpoints/qr-codes"

type Invite struct {
	Id         string   `json:"id"`
	Title      string   `json:"title"`
	Organiser  string   `json:"organiser"`
	Location   string   `json:"location"`
	Notes      string   `json:"notes"`
	Date       string   `json:"date"`
	Passphrase string   `json:"passphrase"`
	Invitees   []Person `json:"invitees"`
}

type Person struct {
	Name    string `json:"name"`
	IsGoing *int8  `json:"is_going"`
}

type InviteDB struct {
	Invite
	QRCode string `json:"qr_code,omitempty"`
}

type InviteJSON struct {
	Invite
	QRCode qrcodes.QRCode `json:"qr_code,omitempty"`
}
