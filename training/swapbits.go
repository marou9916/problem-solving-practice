package training

func SwapBits(octet byte) byte {
	return octet>>4 | octet<<4
}
