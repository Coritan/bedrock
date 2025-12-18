package main

import (
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// universalProtocol is a protocol implementation that accepts any protocol version
// and passes packets through without conversion. This is ideal for proxy servers
// that don't need to inspect or modify packet contents.
type universalProtocol struct {
	protocolID int32
}

func (p universalProtocol) ID() int32 {
	return p.protocolID
}

func (p universalProtocol) Ver() string {
	return "universal"
}

func (p universalProtocol) Packets(listener bool) packet.Pool {
	if listener {
		return packet.NewClientPool()
	}
	return packet.NewServerPool()
}

func (p universalProtocol) NewReader(r minecraft.ByteReader, shieldID int32, enableLimits bool) protocol.IO {
	return protocol.NewReader(r, shieldID, enableLimits)
}

func (p universalProtocol) NewWriter(w minecraft.ByteWriter, shieldID int32) protocol.IO {
	return protocol.NewWriter(w, shieldID)
}

func (p universalProtocol) ConvertToLatest(pk packet.Packet, _ *minecraft.Conn) []packet.Packet {
	return []packet.Packet{pk}
}

func (p universalProtocol) ConvertFromLatest(pk packet.Packet, _ *minecraft.Conn) []packet.Packet {
	return []packet.Packet{pk}
}

// generateProtocolRange creates a slice of protocols supporting versions from min to max.
// This allows the proxy to accept a wide range of Minecraft versions.
func generateProtocolRange(min, max int32) []minecraft.Protocol {
	protocols := make([]minecraft.Protocol, 0, max-min+1)
	for i := min; i <= max; i++ {
		protocols = append(protocols, universalProtocol{protocolID: i})
	}
	return protocols
}
