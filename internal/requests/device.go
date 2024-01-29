package requests

import "go.mongodb.org/mongo-driver/bson/primitive"

type PairDeviceRequest struct {
	Name         string   `json:"name" validate:"required,min=8,max=32"`
	Mac          string   `json:"mac" validate:"required,mac"`
	Ip           string   `json:"ip" validate:"required,ip4_addr"`
	Capabilities []string `json:"capabilities" validate:"required"`
}

type UnPairDeviceRequest struct {
	Mac string `json:"mac" validate:"required,mac"`
}
type SetupDeviceRequest struct {
	// CapabilityIDs []primitive.ObjectID `json:"capability_ids"`
	ZoneID primitive.ObjectID `json:"zone_id" validate:"required,mongodb"`
}
