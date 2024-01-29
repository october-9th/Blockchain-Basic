package internal

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Proof        int
	PreviousHash string
}

type Chain struct {
	Blocks []Block
}

func NewBlockChain() *Chain {
	chain := Chain{[]Block{}}
	chain.CreateBlock(1, "0")
	return &chain
}

func (c *Chain) CreateBlock(proof int, prevHash string) *Block {
	block := Block{
		Index:        len(c.Blocks) + 1,
		Timestamp:    strconv.FormatInt(time.Now().Unix(), 10),
		Proof:        proof,
		PreviousHash: prevHash,
	}
	c.Blocks = append(c.Blocks, block)
	return &block
}

func (c *Chain) GetPreviousBlock() *Block {
	if len(c.Blocks) > 1 {
		return &c.Blocks[len(c.Blocks)-1]
	}
	return nil
}

func (c *Chain) ProofOfWork() int {
	nonce := 1
	foundProof := false
	for !foundProof {
		h := sha256.New()
		str := strconv.Itoa(nonce)
		h.Write([]byte(str))
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:4] == "0000" {
			foundProof = true
		} else {
			nonce += 1
		}
	}
	return nonce
}

func (c *Chain) Hash(block *Block) string {
	h := sha256.New()
	b, _ := json.Marshal(block)
	h.Write(b)
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func (c *Chain) IsChainValid() bool {
	prevBlock := c.Blocks[0]
	blockIndex := 1
	for blockIndex < len(c.Blocks) {
		block := c.Blocks[blockIndex]
		if block.PreviousHash != c.Hash(&prevBlock) {
			return false
		}
		h := sha256.New()
		str := strconv.Itoa(block.Proof)
		h.Write([]byte(str))
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:4] != "0000" {
			return false
		}
		prevBlock = block
		blockIndex += 1
	}
	return true
}
