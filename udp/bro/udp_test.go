package bro

import "testing"

func TestMultiClient(t *testing.T) {
	UDPSenderBroadcast()
}
func TestMultiServer(t *testing.T) {
	UDPReceiverBroadcast()
}
