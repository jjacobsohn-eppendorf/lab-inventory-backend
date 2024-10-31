package entities

type Device struct {
	Id           string `json:"id"`
	Location     string `json:"location"`
	Type         string `json:"type"`
	DeviceHealth string `json:"device_health"`
	LastUsed     string `json:"last_used"`
	Price        string `json:"price"`
	Color        string `json:"color"`
}
