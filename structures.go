package model_go

import (
	"net/http"
	"time"

	"google.golang.org/appengine/datastore"
)

type ModelConnector interface {
	Open(r *http.Request, config string) ModelDatabase
	/**
	Coordinator
	*/

}

type ModelDatabase interface {
	/*
		Coordinator
	*/
	Coord_MasterKeyExists() bool
	Coord_InsertKey(key *MasterKey)
	Coord_GetKey() *MasterKey

	Coord_Insert_ExternalNode(node *NodeIdentification)
	Coord_GetRandomNodeIdentification(limit int) []NodeIdentification

	Coord_InsertBlockRequest(blreq *BlockRequest)
	Coord_InsertBlockRequestForInstance(transport *BlockRequestForInstance)
	Coord_GetLocalBlockRequestForInstance(transport *BlockRequestForInstance) *BlockRequestForInstance
	Coord_UpdateBlockRequestForInstance(transport *BlockRequestForInstance)

	Coord_SetGenesisBlocks(genesis *GenesisBlocksTransport)
	Coord_GetGenesisBlocks() *GenesisBlocksTransport

	/*
		Nodes
	*/
	GetNodeId() *NodeIdentification
	IsRegisteredNodeID() bool
	RegisteredNodeID(us *NodeIdentification)
	SetupSetEndPointIfNull(endpoint string) bool
	SetupNodeRegistrationInDeployment(t *Transaction)

	/*
		Apps
	*/
	AppIdExists(r *http.Request, id string) bool

	/*
		Ledger
	*/
	NodeExists(r *http.Request) bool
	GetCoordinatorKey(r *http.Request) *Transaction
	UserSignExist(r *http.Request, id string) *Transaction
	UserPayloadExist(r *http.Request, id string) *Transaction

	/*
		Transactions
	*/
	InsertPairVerificationTransaction(r *http.Request, t *Transaction) *Transaction
	InsertTransaction(r *http.Request, t *Transaction) *Transaction
	GetParentTransaction(r *http.Request, transactionID string) *Transaction
	/*
		Contracts
	*/
	InsertSignRequest(r *http.Request, t *Transaction) *Transaction
	GetSignRequest(r *http.Request, id int64) *Transaction

	/*
		Blocks
	*/

	InsertBlockRequest(blrq *BlockRequest)
	InsertBlockRequestForInstance(transport *BlockRequestForInstance)
	UpdateBlockRequestForInstance(transport *BlockRequestForInstance)
	GetBlockRequestForInstanceFromHash(hash string) *BlockRequestForInstance

	InsertBlock(block *Block)
	CountBlocks() int
	GetLastBlocks(size int) []Block

	BlockTransactionsCursor(sign string) StorageCursor
	NextTransactionSign(cursor StorageCursor) ([]byte, bool)
	BlockTransactionsCursorClose(mod StorageCursor)

	BlockExists(BlockSign string) bool

	/*
		Query
	*/
	GetBlocksByOffset(size, offset int) []Block
	GetBlockTransactions(blockid string, size, offset int, metadata bool) []Transaction
	GetGroupTransactions(blockid string, size, offset int, metadata bool) []Transaction

	/*
		Query Explorer
	*/
	GetApplicationTransactions(appid, from, to string, metadata bool) []Transaction
	GetApplicationTransaction(appid, sign string, metadata bool) *Transaction
	GetApplicationGroupTransactions(appid, groupsign string, metadata bool) []Transaction
}

/**
Structures for Coordinator
*/

type MasterKey struct {
	MasterPublicKey    []byte
	URL                string
	CoordinatorPublic  []byte
	CoordinatorPrivate []byte
}

type BlockRequest struct {
	IdVal        int64
	TimeCreation int64
	TimeValidity int64
	SeedTrCount  int
}

type BlockRequestForInstance struct {
	NodeIdVal int64

	BlockRequest int64
	Nonce        int64
	Endpoint     string
	NodePubK     string

	Content string
	Sign    string

	BlockSignProposal      string
	BlockSignProposalSign  string //signed by node
	BlockSignProposalCount int
	BlockSign              string
}

type BlockRequestTransport struct {
	Request     *BlockRequest
	ForInstance *BlockRequestForInstance
}

/**
Structures for nodes
*/
type NodeIdentification struct {
	Creation   int64  `json:",omitempty"`
	PublicKey  string `json:",omitempty"`
	PrivateKey string `json:",omitempty"`
	Endpoint   string `json:",omitempty"`
	Myself     bool
}

type Block struct {
	Creation                  time.Time
	TransactionsCount         int
	NextBlockTransactionsUsed int
	Hash                      string
	Sign                      string
	BlockTime                 time.Time
}

type Transaction struct {
	IdVal int64

	BlockSign string `json:",omitempty"`

	//Control
	//OriginatorURl string

	//Signed
	InsertMoment int64  `json:",omitempty"`
	Sign         string `json:",omitempty"`
	Signer       string `json:",omitempty"`

	Hash     string `json:",omitempty"`
	Content  string `datastore:",noindex" json:",omitempty"`
	Creation int64  `json:",omitempty"`

	FromNode NodeIdentification `json:",omitempty"`
	ToNode   NodeIdentification `json:",omitempty"`

	Payload string `datastore:",noindex" json:",omitempty"`
	Parent  string `json:",omitempty"`
	//ParentBlock int64 //TODO AGREGARIN SINGLE TRANSACTIONS, Y CONTRACT CREATION
	AppID       string   `json:",omitempty"`
	SignerKinds []string `json:",omitempty"`
	SignKind    string   `json:",omitempty"`
	Callback    string   `json:",omitempty"`

	External bool `json:",omitempty"`
}

type StorageCursor struct {
	GAE *datastore.Iterator
}
