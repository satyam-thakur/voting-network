package main

import (
    "crypto/sha256"
    "encoding/hex"
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VotingContract provides functions for managing voting tokens
type VotingContract struct {
    contractapi.Contract
}

// VotingTokenRecord represents the structure for storing voting token information
type VotingTokenRecord struct {
    VcmsToken string `json:"vcmsTokenhash"`
    ID        string `json:"id"`
    Salt      string `json:"salt"`
    VCMSdSignature  string `json:"vcmsdSignature"`
}

// VcmsVotingToken stores the received hash and generates a unique ID
func (c *VotingContract) VcmsVotingToken(ctx contractapi.TransactionContextInterface, vcmsToken string, vcmsdSignature string) error {
    exists, err := c.TokenExists(ctx, vcmsToken)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("the voting token %s already exists", vcmsToken)
    }

    // Generate deterministic key for ID
    deterministicKey, err := c.generateDeterministicKey(ctx, vcmsToken)
    if err != nil {
        return err
    }

    // Generate deterministic salt
    salt := c.generateDeterministicSalt(vcmsToken, deterministicKey)

    record := VotingTokenRecord{
        VcmsToken: vcmsToken,
        ID:        deterministicKey,
        Salt:      salt,
        VCMSdSignature: vcmsdSignature,
    }

    recordJSON, err := json.Marshal(record)
    if err != nil {
        return fmt.Errorf("failed to marshal voting token record: %v", err)
    }

    err = ctx.GetStub().PutState(vcmsToken, recordJSON)
    if err != nil {
        return fmt.Errorf("failed to store voting token: %v", err)
    }

    return nil
}

// generateDeterministicKey generates a deterministic key
func (c *VotingContract) generateDeterministicKey(ctx contractapi.TransactionContextInterface, vcmsToken string) (string, error) {
    txID := ctx.GetStub().GetTxID()
    data := txID + vcmsToken

    hash := sha256.Sum256([]byte(data))
    deterministicKey := hex.EncodeToString(hash[:])[:4]

    return deterministicKey, nil
}

// generateDeterministicSalt generates a deterministic salt based on vcmsToken and ID
func (c *VotingContract) generateDeterministicSalt(vcmsToken string, id string) string {
    data := vcmsToken + id
    hash := sha256.Sum256([]byte(data))
    return hex.EncodeToString(hash[:])[:5]
}

// TokenExists checks if a token exists
func (c *VotingContract) TokenExists(ctx contractapi.TransactionContextInterface, vcmsToken string) (bool, error) {
    recordJSON, err := ctx.GetStub().GetState(vcmsToken)
    if err != nil {
        return false, fmt.Errorf("failed to read from world state: %v", err)
    }
    return recordJSON != nil, nil
}

// getVotingTokenRecord retrieves a voting token record
func (c *VotingContract) GetVotingTokenRecord(ctx contractapi.TransactionContextInterface, vcmsToken string) (*VotingTokenRecord, error) {
    recordJSON, err := ctx.GetStub().GetState(vcmsToken)
    if err != nil {
        return nil, fmt.Errorf("failed to read voting token record: %v", err)
    }
    if recordJSON == nil {
        return nil, fmt.Errorf("the voting token %s does not exist", vcmsToken)
    }

    var record VotingTokenRecord
    err = json.Unmarshal(recordJSON, &record)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal voting token record: %v", err)
    }

    return &record, nil
}

func main() {
    chaincode, err := contractapi.NewChaincode(&VotingContract{})
    if err != nil {
        fmt.Printf("Error creating voting chaincode: %v", err)
        return
    }

    if err := chaincode.Start(); err != nil {
        fmt.Printf("Error starting voting chaincode: %v", err)
    }
}