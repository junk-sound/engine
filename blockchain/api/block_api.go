/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package api

import (
	"errors"

	"github.com/it-chain/engine/blockchain"
)

var ErrGetLastCommitedBlock = errors.New("Error in getting last commited block")
var ErrCreateProposedBlock = errors.New("Error in creating proposed block")
var ErrFailBlockTypeCasting = errors.New("Error failed type casting block")

type syncService struct {
	queryService
	peerService
}

func (s syncService) isSynced(peer blockchain.Peer, lastBlock blockchain.DefaultBlock) (blockchain.IsSynced, error) {

	// 쿼리 서비로 standard block get 후 last block 과 standard block 비

	return nil, nil
}

func (s syncService) isProcessing() bool  {
	return true
}



type storageService struct {
	queryService queryService
}

func (s storageService) Construct(peer blockchain.Peer) error {
	//lastblock get
	//쿼리 서비스로 standard block get
	//lastblock이 standard block 과 같은 높이가 될 때까지 add, commit

	return nil
}

type poolService struct {
}

func (p poolService) AddBlockToPool(block blockchain.Block) error   {

}



type consensusService struct {
}

func (c consensusService) requestConsensus(block blockchain.Block)  {
	
}

func (c consensusService) OnConsensus()  {

}

func (c consensusService) OffConsensus()  {

}

func (c consensusService)isProcessing() bool {
	return true
}

func (c consensusService)AddToConsensusPool()  {

}

type queryService struct {
}

type peerService struct {
}

type eventService struct {
}

type BlockApi struct {
	publisherId       string
	blockQueryService blockchain.BlockQueryService
	storageService storageService
	poolservice	poolService
	syncService 	syncService
	consensusService consensusService
	eventService eventService
}

func NewBlockApi(publisherId string, blockQueryService blockchain.BlockQueryService) (BlockApi, error) {
	return BlockApi{
		publisherId:         publisherId,
	}, nil
}

func (bApi BlockApi) SyncWithPeer(peer blockchain.Peer) error  {

	// 싱크 서비스로 싱크 됐는지 확인
	// 싱크가 덜 됐다면 스토리지 서비스로 construct
	return nil
}

func (bApi BlockApi) Synchronize() error {
	// 싱크 서비스로 peer get
	// 스토리지 서비스로 last block get
	// 싱크 서비스로 싱크 됐는지 확인.
	// 싱크가 덜 됐다면 스토리지 서비스로 construct

	return nil
}
//// TODO: Check 과정에서 임의의 노드에게서 받은 blockchain 정보로 동기화 되었는지 확인한다.
//func (bApi BlockApi) SyncedCheck(block blockchain.Block) error {
//	return nil
//}

// 받은 block을 block pool에 추가한다.
func (bApi BlockApi) AddBlockToPool(block blockchain.Block) error {
	return nil
}

func (bApi BlockApi) isReadyToAddBlockFromPool() {
	//poolService와 blockService 사용
	//그냥syncedCheck를 해도 될 수도 있음.
}


func (bApi BlockApi) SaveBlockFromPool(height blockchain.BlockHeight) error {
	return nil
}

//func (bApi BlockApi) SyncIsProgressing() blockchain.ProgressState {
//	return blockchain.DONE
//}

func (bApi BlockApi) CreateBlock(txList []blockchain.Transaction) (blockchain.DefaultBlock, error) {

	lastBlock, err := bApi.blockQueryService.GetLastCommitedBlock()
	if err != nil {
		return blockchain.DefaultBlock{}, ErrGetLastCommitedBlock
	}

	prevSeal := lastBlock.GetSeal()
	height := lastBlock.GetHeight() + 1
	defaultTxList := blockchain.GetBackTxType(txList)
	creator := bApi.publisherId

	block, err := blockchain.CreateProposedBlock(prevSeal, height, defaultTxList, []byte(creator))

	if err != nil {
		return blockchain.DefaultBlock{}, ErrCreateProposedBlock
	}

	defaultBlock, ok := block.(*blockchain.DefaultBlock)
	if !ok {
		return blockchain.DefaultBlock{}, ErrFailBlockTypeCasting
	}

	return *defaultBlock, nil
}

func (bApi BlockApi) CreateGenesisBlock(GenesisConfPath string) error {

	GenesisBlock, err := blockchain.CreateGenesisBlock(GenesisConfPath)

	if err != nil {
		return ErrCreateGenesisBlock
	}

	err = blockchain.CommitBlock(GenesisBlock)

	if err != nil {
		return ErrCommitBlock
	}

	return nil
}
