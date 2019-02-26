package legacy

const (
	P2KH        uint8 = 31
	P2SH              = 90
	P2KHTestnet       = 66
	P2SHTestnet       = 127
)

// Address is a structure which has a raw unpacked version of a legacy
// address.
type Address struct {
	Version uint8
	Payload []uint8
}
