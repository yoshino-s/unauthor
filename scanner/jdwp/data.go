package jdwp

import (
	"encoding/binary"
	"io"
)

type DataReader struct {
	io.Reader
}

func (r *DataReader) ReadString() (string, error) {
	len, err := r.ReadUint32()
	if err != nil {
		return "", err
	}

	buffer := make([]byte, len)
	if _, err := r.Read(buffer); err != nil {
		return "", err
	}
	return string(buffer), nil
}

func (r *DataReader) ReadUint32() (uint32, error) {
	var data uint32
	if err := binary.Read(r, binary.BigEndian, &data); err != nil {
		return 0, err
	}
	return data, nil
}

type JdwpVersionResponse struct {
	Description string `json:"description"`
	JdwpMajor   uint32 `json:"jdwp_major"`
	JdwpMinor   uint32 `json:"jdwp_minor"`
	VmVersion   string `json:"vm_version"`
	VmName      string `json:"vm_name"`
}

func (c *JdwpVersionResponse) Unmarshal(_reader io.Reader) error {
	var err error
	reader := &DataReader{_reader}
	if c.Description, err = reader.ReadString(); err != nil {
		return err
	}
	if c.JdwpMajor, err = reader.ReadUint32(); err != nil {
		return err
	}
	if c.JdwpMinor, err = reader.ReadUint32(); err != nil {
		return err
	}
	if c.VmVersion, err = reader.ReadString(); err != nil {
		return err
	}

	return nil
}
