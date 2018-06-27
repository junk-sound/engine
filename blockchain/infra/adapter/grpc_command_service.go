package adapter

import (
	"errors"

	"github.com/it-chain/it-chain-Engine/blockchain"
	"github.com/it-chain/it-chain-Engine/common"
	"github.com/it-chain/it-chain-Engine/p2p"
	"github.com/it-chain/midgard"
	"github.com/rs/xid"
)

//kind of error
var ErrEmptyNodeId = errors.New("empty nodeid proposed")
var ErrEmptyBlockSeal = errors.New("empty block seal")

// ToDo: 구현.(gitId:junk-sound)
type Publish func(exchange string, topic string, data interface{}) (err error)

type GrpcCommandService struct {
	publish Publish // midgard.client.Publish
}

func NewGrpcCommandService(publish Publish) *GrpcCommandService {
	return &GrpcCommandService{
		publish: publish,
	}
}

func (gcs *GrpcCommandService) RequestBlock(peerId p2p.PeerId, height uint64) error {

	if peerId.Id == "" {
		return ErrEmptyNodeId
	}

	body := blockchain.BlockRequestMessage{
		Height: height,
	}

	deliverCommand, err := createGrpcCommand("BlockRequestProtocol", body)
	if err != nil {
		return err
	}

	deliverCommand.Recipients = append(deliverCommand.Recipients, peerId.ToString())

	return gcs.publish("Command", "message.deliver", deliverCommand)
}

func (gcs *GrpcCommandService) ResponseBlock(peerId p2p.PeerId, block blockchain.Block) error {

	if peerId.Id == "" {
		return ErrEmptyNodeId
	}

	if block.GetSeal() == nil {
		return ErrEmptyBlockSeal
	}

	body := block

	deliverCommand, err := createGrpcCommand("BlockResponseProtocol", body)
	if err != nil {
		return err
	}

	deliverCommand.Recipients = append(deliverCommand.Recipients, peerId.ToString())

	return gcs.publish("Command", "message.deliver", deliverCommand)
}

func createGrpcCommand(protocol string, body interface{}) (blockchain.GrpcCommand, error) {

	data, err := common.Serialize(body)
	if err != nil {
		return blockchain.GrpcCommand{}, err
	}

	return blockchain.GrpcCommand{
		CommandModel: midgard.CommandModel{
			ID: xid.New().String(),
		},
		Recipients: make([]string, 0),
		Body:       data,
		Protocol:   protocol,
	}, err
}
