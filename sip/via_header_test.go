package sip

import "testing"

var viaTarget = "Via SIP/2.0/UDP 192.168.1.171:5060;rport;branch=z9hG4bK.CKk7LZ6Pi9lAINvmVDK35B9LFTu4cSwN"

func TestViaHeaders(t *testing.T) {
	params := NewParams()
	params.Add("rport", nil)
	params.Add("branch", &String{Value: "z9hG4bK.CKk7LZ6Pi9lAINvmVDK35B9LFTu4cSwN"})

	via := &ViaHeader{
		{
			ProtocolName:    "SIP",
			ProtocolVersion: "2.0",
			Transport:       "UDP",
			Host:            "192.168.1.171",
			Port:            NewPort(5060),
			Params:          params,
		},
	}
	t.Log("got : ", via.String())
	if via.String() != viaTarget {
		t.Errorf("want %v, got %v", viaTarget, via)
	}
	t.Log("Ok")
}
