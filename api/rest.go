package api

import (
	"net/http"

	"blockchain.com/m/internal"
	"github.com/gin-gonic/gin"
)

var blockchain = internal.NewBlockChain()

func getChain(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, blockchain)
}

func mineBlock(c *gin.Context) {
	prevBlock := blockchain.GetPreviousBlock()
	prevHash := blockchain.Hash(prevBlock)
	proof := blockchain.ProofOfWork()
	block := blockchain.CreateBlock(proof, prevHash)
	c.IndentedJSON(http.StatusOK, block)
}

func isValid(c *gin.Context) {
	valid := blockchain.IsChainValid()
	c.IndentedJSON(http.StatusOK, valid)
}

func CreateRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/get_chain", getChain)
	router.POST("/mine_block", mineBlock)
	router.GET("/is_valid", isValid)
	return router
}
