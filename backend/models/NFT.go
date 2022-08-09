package models

import (
	"errors"
)

// TODO: add json hints
// type NFT struct {
// 	TokenId            string   `gorm:"primaryKey; size: 41;not null" json:"token_id"`
// 	ContractId         string   `gorm:"size: 35; not null" json:"contract_id"`
// 	Contract           Contract `json:"-"`
// 	TokenIdx           int64    `gorm:"not null" json:"token_idx"`
// 	Owner              string   `gorm:"size: 35; not null" json:"owner"`
// 	IssueTransactionId string   `gorm:"size: 44;not null" json:"issue_transaction_id"`
// 	TokenDescription   string   `gorm:"size: 500" json:"token_description"` // TODO: What size to choose?
// 	OnMarketplace      bool     `gorm:"not null; default: true" json:"on_marketplace"`
// }

// for testing currently
type NFT struct {
	TokenId            string `gorm:"primaryKey; size: 41;not null" json:"token_id"`
	TokenIdx           int64  `gorm:"not null" json:"token_idx"`
	Owner              string `gorm:"size: 35; not null" json:"owner"`
	IssueTransactionId string `gorm:"size: 44;not null" json:"issue_transaction_id"`
	TokenDescription   string `json:"token_description"` // TODO: What size to choose?
	TokenMeta          string `json:"token_meta"`        // TODO: size
	Price              int64  `gorm:"not null" json:"token_price"`
	OnSale             bool   `gorm:"not null; default: true" json:"on_sale"`
	CollectionBelong   string `gorm:"not null" json:"collection_belong"`
	NFTcid             string `gorm:"size: 44;not null" json:"nft_cid"`
	Metacid            string `gorm:"size: 44;not null" json:"meta_cid"`
}

func (n *NFT) InsertNFT() (*NFT, error) {
	err := Db.Create(&n).Error
	if err != nil {
		return &NFT{}, err
	}
	return n, nil
}

func GetNFTsByCollection(collectionName string) ([]NFT, error) {
	var nfts []NFT

	if err := Db.First(&nfts, "collection_belong = ?", collectionName).Error; err != nil {
		return nfts, errors.New("nfts not found")
	}

	return nfts, nil
}

func GetNFTs() ([]NFT, error) {
	var nfts []NFT
	if err := Db.Find(&nfts).Error; err != nil {
		return nfts, err
	}
	return nfts, nil
}

func GetNFTById(tokenId string) (NFT, error) {
	var n NFT
	if err := Db.First(&n, tokenId).Error; err != nil {
		return n, errors.New("nft not found")
	}
	return n, nil
}
