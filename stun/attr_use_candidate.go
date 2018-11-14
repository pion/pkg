package stun

// UseCandidate has no field struct
type UseCandidate struct {
}

// Pack message user Candidate
func (u *UseCandidate) Pack(message *Message) error {
	message.AddAttribute(AttrUseCandidate, []byte{})
	return nil
}

// Unpack message with Candidate
func (u *UseCandidate) Unpack(message *Message, rawAttribute *RawAttribute) error {
	return nil
}
