package duobo

import "testing"

func TestMultiClient(t *testing.T) {
	UDPSenderMulti()
}
func TestMultiServer(t *testing.T) {
	UDPReceiverMulti()
}
