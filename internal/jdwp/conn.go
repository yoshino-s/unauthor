package jdwp

import (
	"bytes"
	"fmt"
	"net"
)

const (
	handShake = "JDWP-Handshake"
)

type JdwpConn struct {
	net.Conn
	id uint32
}

func NewJdwpConn(conn net.Conn) *JdwpConn {
	return &JdwpConn{conn, 0x01}
}

func (c *JdwpConn) readPacket() (*Packet, error) {
	pkt := &Packet{}
	if err := pkt.Unmarshal(c.Conn); err != nil {
		return nil, err
	}
	if pkt.Flags != 0x80 {
		return nil, fmt.Errorf("invalid packet flags")
	}
	return pkt, nil
}

func (c *JdwpConn) writePacket(pkt *Packet) error {
	_, err := c.Conn.Write(pkt.Marshal())
	return err
}

func (c *JdwpConn) Handshake() error {
	if _, err := c.Conn.Write([]byte(handShake)); err != nil {
		return err
	}
	buffer := make([]byte, len(handShake))
	if _, err := c.Conn.Read(buffer); err != nil {
		return err
	}
	if string(buffer) != handShake {
		return fmt.Errorf("invalid handshake response")
	}
	return nil
}

func (c *JdwpConn) Version() (*JdwpVersionResponse, error) {
	pkt := &Packet{
		ID:     c.id,
		CmdSet: 0x01,
		Cmd:    0x01,
	}
	c.id += 2
	if err := c.writePacket(pkt); err != nil {
		return nil, err
	}
	pkt, err := c.readPacket()
	if err != nil {
		return nil, err
	}

	resp := &JdwpVersionResponse{}

	if err := resp.Unmarshal(bytes.NewReader(pkt.Data)); err != nil {
		return nil, err
	}
	return resp, nil
}
