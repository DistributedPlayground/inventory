package model

type InventoryRequest struct {
	// The ID of the inventory item
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}
