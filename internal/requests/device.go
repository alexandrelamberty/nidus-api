package requests

type PairDeviceRequest struct {
	Name string `json:"name"`
	Mac  string `json:"mac"`
	Ip   string `json:"ip"`
}
