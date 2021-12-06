package shared

import "github.com/hashicorp/go-plugin"

var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "Greeting_Key",
	MagicCookieValue: "Greeting_Value",
}
