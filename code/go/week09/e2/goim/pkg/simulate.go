package pkg

import (
	"encoding/binary"
	"errors"
)

var (
	packLengthFieldSize    = 4
	headerLengthFieldSize  = 2
	protocVersionFieldSize = 2
	operationCodeFieldSize = 4
	sequenceIdFieldSize    = 4
	ErrPackInComplete      = errors.New("error: package is incomplete")
)

//Pack define package
type Pack struct {
	Length          int
	HeaderLength    int
	ProtocolVersion int
	OperationCode   int
	Seq             int
	Content         []byte
}

func NewPack(version, code, seq int, content []byte) *Pack {
	headerSize := headerSize()
	return &Pack{
		Length:          len(content) + headerSize,
		HeaderLength:    headerSize,
		ProtocolVersion: version,
		OperationCode:   code,
		Seq:             seq,
		Content:         content,
	}
}

func Encoder(pack *Pack) []byte {
	res := make([]byte, pack.Length)

	//  package length
	binary.BigEndian.PutUint32(
		res[:headerLengthStart()],
		uint32(pack.Length),
	)
	//  header length
	binary.BigEndian.PutUint16(
		res[headerLengthStart():protocolVersionStart()],
		uint16(headerSize()),
	)
	//  protocol version
	binary.BigEndian.PutUint16(
		res[protocolVersionStart():operationCodeStart()],
		uint16(pack.ProtocolVersion),
	)
	//  operation code
	binary.BigEndian.PutUint32(
		res[operationCodeStart():sequenceIdStart()],
		uint32(pack.OperationCode),
	)
	//  sequence id
	binary.BigEndian.PutUint32(
		res[sequenceIdStart():sequenceIdStart()+sequenceIdFieldSize],
		uint32(pack.Seq),
	)
	//  body
	copy(res[headerSize():], pack.Content)

	return res
}

//Decoder   request message
func Decoder(msg []byte) (*Pack, error) {
	if len(msg) < headerSize()+1 {
		return nil, ErrPackInComplete
	}
	//  package length
	packageLength := binary.BigEndian.Uint32(msg[:headerLengthStart()])
	//  header length
	headerLength := binary.BigEndian.Uint16(msg[headerLengthStart():protocolVersionStart()])
	//  protocol version
	protocolVersion := binary.BigEndian.Uint16(msg[protocolVersionStart():operationCodeStart()])
	//  operation code
	operationCode := binary.BigEndian.Uint32(msg[operationCodeStart():sequenceIdStart()])
	//  sequence id
	sequenceId := binary.BigEndian.Uint32(msg[sequenceIdStart() : sequenceIdStart()+sequenceIdFieldSize])
	//  data
	content := msg[headerSize():]
	return &Pack{
		Length:          int(packageLength),
		HeaderLength:    int(headerLength),
		ProtocolVersion: int(protocolVersion),
		OperationCode:   int(operationCode),
		Seq:             int(sequenceId),
		Content:         content,
	}, nil
}

func headerSize() int {
	return packLengthFieldSize +
		headerLengthFieldSize +
		protocVersionFieldSize +
		operationCodeFieldSize +
		sequenceIdFieldSize
}

func headerLengthStart() int {
	return packLengthFieldSize
}

func protocolVersionStart() int {
	return headerLengthStart() + headerLengthFieldSize
}

func operationCodeStart() int {
	return protocolVersionStart() + protocVersionFieldSize
}

func sequenceIdStart() int {
	return operationCodeStart() + operationCodeFieldSize
}

func PackageLengthSize() int {
	return packLengthFieldSize
}
