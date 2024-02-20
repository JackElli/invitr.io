package invitesmgr

type Invite struct {
	Id         string `json:"id"`
	Organiser  string `json:"organiser"`
	Location   string `json:"location"`
	Date       string `json:"date"`
	QRCode     QRCode `json:"qr_code,omitempty"`
	Passphrase string `json:"passphrase"`
}
