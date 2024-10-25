package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type Asset struct {
	DealerID    string  `json:"dealerID"`
	MSISDN      string  `json:"msisdn"`
	MPIN        string  `json:"mpin"`
	Balance     float64 `json:"balance"`
	Status      string  `json:"status"`
	TransAmount float64 `json:"transAmount"`
	TransType   string  `json:"transType"`
	Remarks     string  `json:"remarks"`
}

type AssetManagementContract struct {
	contractapi.Contract
}

func (c *AssetManagementContract) CreateAsset(ctx contractapi.TransactionContextInterface, dealerID, msisdn, mpin string, balance float64, status, remarks string) error {
	asset := Asset{
		DealerID:    dealerID,
		MSISDN:      msisdn,
		MPIN:        mpin,
		Balance:     balance,
		Status:      status,
		TransAmount: 0,
		TransType:   "",
		Remarks:     remarks,
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerID, assetJSON)
}

func (c *AssetManagementContract) UpdateAsset(ctx contractapi.TransactionContextInterface, dealerID string, balance float64, status, remarks string) error {
	assetJSON, err := ctx.GetStub().GetState(dealerID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return fmt.Errorf("asset %s does not exist", dealerID)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return err
	}

	asset.Balance = balance
	asset.Status = status
	asset.Remarks = remarks

	updatedAssetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(dealerID, updatedAssetJSON)
}

func (c *AssetManagementContract) QueryAsset(ctx contractapi.TransactionContextInterface, dealerID string) (*Asset, error) {
	assetJSON, err := ctx.GetStub().GetState(dealerID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if assetJSON == nil {
		return nil, fmt.Errorf("asset %s does not exist", dealerID)
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (c *AssetManagementContract) GetTransactionHistory(ctx contractapi.TransactionContextInterface, dealerID string) ([]*contractapi.HistoryQueryResult, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(dealerID)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var history []*contractapi.HistoryQueryResult
	for resultsIterator.HasNext() {
		modification, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		history = append(history, &modification)
	}

	return history, nil
}

func main() {
	chaincode, err := contractapi.NewChaincode(new(AssetManagementContract))
	if err != nil {
		fmt.Printf("Error create asset management chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting asset management chaincode: %s", err.Error())
	}
}

