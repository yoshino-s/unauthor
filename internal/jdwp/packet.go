package jdwp

import (
	"encoding/binary"
	"io"
)

type Packet struct {
	ID     uint32
	Flags  byte
	CmdSet byte
	Cmd    byte
	Data   []byte
}

func (p *Packet) Marshal() []byte {
	var pktLen uint32 = uint32(11 + len(p.Data))
	var pkt []byte = make([]byte, pktLen)
	binary.BigEndian.PutUint32(pkt[0:4], pktLen)
	binary.BigEndian.PutUint32(pkt[4:8], p.ID)
	pkt[8] = p.Flags
	pkt[9] = p.CmdSet
	pkt[10] = p.Cmd
	copy(pkt[11:], p.Data)
	return pkt
}

func (p *Packet) Unmarshal(reader io.Reader) error {
	var pktLen uint32
	if err := binary.Read(reader, binary.BigEndian, &pktLen); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &p.ID); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &p.Flags); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &p.CmdSet); err != nil {
		return err
	}
	if err := binary.Read(reader, binary.BigEndian, &p.Cmd); err != nil {
		return err
	}
	p.Data = make([]byte, pktLen-11)
	if _, err := reader.Read(p.Data); err != nil {
		return err
	}
	return nil
}
