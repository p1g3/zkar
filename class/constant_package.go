package class

import (
	"encoding/binary"
	"fmt"
	"github.com/phith0n/zkar/commons"
)

type ConstantPackage struct {
	NameIndex uint16
}

func (c *ConstantPackage) ToBytes() []byte {
	var bs = []byte{CONSTANT_PACKAGE_INFO}
	bs = append(bs, commons.NumberToBytes(c.NameIndex)...)
	return bs
}

func readConstantPackage(stream *commons.Stream) (*ConstantPackage, error) {
	_, _ = stream.ReadN(1)
	bs, err := stream.ReadN(2)
	if err != nil {
		return nil, fmt.Errorf("read constant package failed, no enough data in the stream")
	}

	return &ConstantPackage{
		NameIndex: binary.BigEndian.Uint16(bs),
	}, nil
}
