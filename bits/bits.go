package bits

type Bits uint8

func Set(b, flag Bits) Bits       { return b | flag }
func Unset(b, flag Bits) Bits     { return b &^ flag }
func (b Bits) Has(flag Bits) bool { return b&flag != 0 }
