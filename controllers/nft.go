package controllers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
	"tulip/backend/models"

	gv "tulip/backend/vsys"

	"github.com/gin-gonic/gin"
	ipfs "github.com/ipfs/go-ipfs-api"
)

func GetNFTs(c *gin.Context) {
	nfts, err := models.GetNFTs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// TODO: is c.Abort necessary in non middleware?
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, nfts)
}

func GetNFTById(c *gin.Context) {
	id := c.Param("id")
	nft, err := models.GetNFTById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, nft)
}

func GetNFTsByCollection(c *gin.Context) {
	name, berr := c.Params.Get("name")
	if !berr {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "params wrong"})
		// c.Abort()
		return
	}

	nfts, err := models.GetNFTsByCollection(name)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "nfts not found"})
		// c.Abort()
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nfts})
}

func GetNFTMintHistory(c *gin.Context) {
	//	TODO: ahh too lazy
}

type textForm struct {
	File        *multipart.File `form:"File" binding:"required"`
	Description string          `form:"Description" binding:"required"`
	MetaInfo    string          `form:"MetaInfo" binding:"required"`
	Collection  string          `form:"Collection" binding:"required"`
}

func MintFileNFT(c *gin.Context) {
	VsysHost := os.Getenv("VSYS_HOST")
	VsysSeed := os.Getenv("VSYS_SEED")

	api := gv.NewNodeAPI(VsysHost)
	ch := gv.NewChain(api, gv.TEST_NET)
	wal, _ := gv.NewWalletFromSeedStr(VsysSeed)
	acnt, _ := wal.GetAccount(ch, 0)

	// nc, _ := gv.RegisterNFTCtrt(acnt, "")
	// fmt.Println(nc.CtrtId)

	// test ctrtId
	ctrtId := "CEyXxA6joy1PN5nxnn55uvC9Ny7AZQ3EZvX"
	nc, err := gv.NewNFTCtrt(ctrtId, ch)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}

	f, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	sh := ipfs.NewShell("http://127.0.0.1:5001")

	metainfo := f.Value["MetaInfo"][0]
	metacid, err := sh.Add(strings.NewReader(metainfo))
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	File, err := (f.File["File"][0]).Open()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	nftcid, err := sh.Add(File)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("%T, %v\n", nftcid, nftcid)

	ret, err := nc.Issue(acnt, metacid, nftcid)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	txId := string(ret.Id)

	fmt.Println(txId)

	time.Sleep(6 * time.Second)

	res, err := api.GetTxInfo(string(txId))

	// TODO: error handling
	if err != nil {
		fmt.Println(err.Error())
		// c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	}
	fmt.Println(res)

	// TODO: tokId api
	ctrtInfo, err := api.GetCtrtInfo(ctrtId)
	fmt.Println(ctrtInfo)

	m := models.NFT{}

	m.TokenIdx = 0
	m.TokenId = "asdfgh"
	m.Price = 0
	m.Owner = string(acnt.Addr.B58Str())
	m.TokenMeta = metainfo
	m.TokenDescription = f.Value["Description"][0]
	m.IssueTransactionId = txId

	_, err = m.InsertNFT()
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"metacId": metacid, "nftcId": nftcid, "txId": txId})

}

type TextForm struct {
	Text        string `form:"Text" binding:"required"`
	Description string `form:"Description" binding:"required"`
	MetaInfo    string `form:"MetaInfo" binding:"required"`
	Collection  string `form:"Collection" binding:"required"`
}

func MintTextNFT(c *gin.Context) {
	// VsysHost := os.Getenv("VSYS_HOST")
	// VsysSeed := os.Getenv("VSYS_SEED")

	// api := gv.NewNodeAPI(VsysHost)
	// ch := gv.NewChain(api, gv.TEST_NET)
	// wal, _ := gv.NewWalletFromSeedStr(VsysSeed)
	// acnt, _ := wal.GetAccount(ch, 0)

	// // nc, _ := gv.RegisterNFTCtrt(acnt, "")
	// // fmt.Println(nc.CtrtId)

	// // test ctrtId
	// ctrtId := "CEyXxA6joy1PN5nxnn55uvC9Ny7AZQ3EZvX"
	// nc, err := gv.NewNFTCtrt(ctrtId, ch)
	// if err != nil {
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// }

	// f, err := c.MultipartForm()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }

	// sh := ipfs.NewShell("http://127.0.0.1:5001")

	// metainfo := f.Value["MetaInfo"][0]
	// metacid, err := sh.Add(strings.NewReader(metainfo))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// 	return
	// }

	// text := f.Value["Text"][0]
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// }

	// nftcid, err := sh.Add(strings.NewReader(text))
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// 	return
	// }

	// fmt.Printf("%T, %v\n", nftcid, nftcid)

	// ret, err := nc.Issue(acnt, metacid, nftcid)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// }
	// txId := string(ret.Id)

	// fmt.Println(txId)

	// time.Sleep(6 * time.Second)

	// res, err := api.GetTxInfo(string(txId))

	// // TODO: error handling
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	// c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
	// }
	// fmt.Println(res)

	// // TODO: tokId api
	// ctrtInfo, err := api.GetCtrtInfo(ctrtId)
	// fmt.Println(ctrtInfo)

	// m := models.NFT{}

	// m.TokenIdx = 0
	// m.TokenId = string(rune(rand.Int()))
	// m.Price = 0
	// m.Owner = string(acnt.Addr.B58Str())
	// m.TokenMeta = metainfo
	// m.TokenDescription = f.Value["Description"][0]
	// m.IssueTransactionId = txId
	// m.NFTcid = nftcid
	// m.Metacid = metacid
	// m.CollectionBelong = f.Value["Collection"][0]

	// _, err = m.InsertNFT()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	c.JSON(http.StatusOK, gin.H{"metacId": "QmUJVPH7CXmDXf1gs5zYswVfFMizeZ64G6Ky9PmGZWvtra", "nftcId": "QmcBfzyyY5zb2VMykPKv8FY9CfzGUfgZPhNTFeUjU6KPqM", "txId": "97DrMriBWNFTfYSCXwgbK4eYwJg91FeQT2m6tyhRH7cQ"})

}
