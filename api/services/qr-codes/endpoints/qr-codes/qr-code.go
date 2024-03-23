package qrcodes

type QRCode struct {
	Width  int   `json:"width"`
	Height int   `json:"height"`
	Pixels []int `json:"pixels"`
}
