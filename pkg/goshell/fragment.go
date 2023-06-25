package goshell

type Fragment struct {
	Outputs []Output `protobuf:"bytes,4,opt,name=Output,proto3" json:"Output,omitempty"`
}
