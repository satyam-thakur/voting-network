package main

import (
    "crypto/sha256"
    "encoding/json"
    "fmt"

    "github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VotingContract provides functions for managing voting on mychannel
type VotingContract struct {
    contractapi.Contract
}

// BallotRecord represents the structure for storing ballot information on mychannel
type BallotRecord struct {
    Ballot         string `json:"ballot"`
    VotingTokenHash string `json:"votingTokenHash"` // Updated field name
    ID             string `json:"id"`
}

// VotingTokenRecord represents the structure for storing voting token information on mychannel1
type VotingTokenRecord struct {
    VcmsTokenHash string `json:"vcmsTokenhash"` // Updated field name
    ID            string `json:"id"`
    Salt          string `json:"salt"` // Added Salt field
	VCMSdSignature  string `json:"vcmsdSignature"`
}

// TransactionType1 represents a transaction structure for type1 transactions
type TransactionType1 struct {
    VcmsTokenHash  string `json:"vcmsTokenHash"`
    ID             string `json:"id"`
    // Index1        string `json:"index1"` // VCMSdSignature
	VCMSdSignature  string `json:"vcmsdSignature"`
}



// CastVote takes voter input and processes it
func (c *VotingContract) CastVote(ctx contractapi.TransactionContextInterface, ballot string, votingTokenhash string) error {
    // Step 1: Verify the voting token on channel2
    votingTokenRecord, err := c.GetVotingTokenFromChannel2(ctx, votingTokenhash)
    if err != nil {
        return fmt.Errorf("failed to verify voting token: %v", err)
    }

    // Step 2: Construct Index1 using VCMSdSignature
    Index1 := "00" + votingTokenRecord.ID // Assuming constant ID for second part of Type
    // Step 3: Construct Index2
    Index2 := "01" + votingTokenRecord.ID

    // Step 4: Check if Index2 already exists in channel1
    exists, err := c.idExistsInChannel1(ctx, Index2)
    if err != nil {
        return err
    }
    if exists {
        return fmt.Errorf("a vote with ID %s has already been cast", Index2)
    }

    // Step 5: Hash the voting token with salt
    hashedVotingToken := c.hashWithSalt(votingTokenhash, votingTokenRecord.Salt)

    // Step 6: Create and store the ballot record
    ballotRecord := BallotRecord{
        Ballot:         ballot,
        VotingTokenHash: hashedVotingToken, // Store the hashed value
        ID:             Index2,
    }

    ballotJSON, err := json.Marshal(ballotRecord)
    if err != nil {
        return fmt.Errorf("failed to marshal ballot record: %v", err)
    }

    err = ctx.GetStub().PutState(Index2, ballotJSON)
    if err != nil {
        return fmt.Errorf("failed to store ballot: %v", err)
    }

	// Step 7: Create and store TransactionType1 record
	transactionType1 := TransactionType1{
		VcmsTokenHash:  hashedVotingToken,
		ID:             Index1,
		VCMSdSignature: votingTokenRecord.VCMSdSignature,
	}

	transactionJSON, err := json.Marshal(transactionType1)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction type1 record: %v", err)
	}

	err = ctx.GetStub().PutState(Index1, transactionJSON)
	if err != nil {
		return fmt.Errorf("failed to store transaction type1 record: %v", err)
	}

	return nil
}

// PostVoting handles post-voting token input and appends it to the blockchain
func (c *VotingContract) PostVoting(ctx contractapi.TransactionContextInterface, votingTokenHash string, votingToken string) error {
    // Step 1: Verify the voting token on channel1 (voting6)
    votingTokenRecord, err := c.GetVotingTokenFromChannel2(ctx, votingTokenHash)
    if err != nil {
        return fmt.Errorf("failed to verify voting token from channel1: %v", err)
    }

    // Step 2: Construct Index3 with prefix "10"
    Index3 := "10" + votingTokenRecord.ID

    // Step 3: Check if this Index3 already exists in current blockchain (mychannel)
    exists, err := c.idExistsInChannel1(ctx, Index3)
    if err != nil {
        return err
    }
    
    if exists {
        return fmt.Errorf("a VCMS hash token record with ID %s already exists", Index3)
    }

    // Step 4: Hash VcmsTokenHash and VotingTokenHash with salt
    hashedVcmsToken := c.hashWithSalt(votingTokenRecord.VcmsTokenHash, votingTokenRecord.Salt)
    // hashedVotingToken := c.hashWithSalt(votingTokenHash, votingTokenRecord.Salt)

    // Step 5: Create a new record for post-voting
    postVotingRecord := struct {
        HashedVcmsToken   string `json:"hashedVcmsToken"`
        VotingToken string `json:"votingToken"`
        ID                string `json:"id"`
        // Salt              string `json:"salt"` // Include salt in post-voting record if needed
    }{
        HashedVcmsToken:   hashedVcmsToken,
        VotingToken: 	votingToken,
        ID:                Index3,
        // Salt:              votingTokenRecord.Salt, // Use salt from retrieved record
    }

    postVotingJSON, err := json.Marshal(postVotingRecord)
    if err != nil {
        return fmt.Errorf("failed to marshal post-voting record: %v", err)
    }

    // Store the post-voting record in the world state
    err = ctx.GetStub().PutState(Index3, postVotingJSON)
    if err != nil {
        return fmt.Errorf("failed to store post-voting record: %v", err)
    }

    return nil
}

// hashWithSalt computes SHA-256 hash of a given input concatenated with salt
func (c *VotingContract) hashWithSalt(input string, salt string) string {
   data := input + salt
   hash := sha256.Sum256([]byte(data))
   return fmt.Sprintf("%x", hash[:]) // Return hex representation of the hash
}

// GetPostVotingRecord retrieves a post-voting record by ID
func (c *VotingContract) GetPostVotingRecord(ctx contractapi.TransactionContextInterface, id string) (interface{}, error) {
   postVotingJSON, err := ctx.GetStub().GetState(id)
   if err != nil {
       return nil, fmt.Errorf("failed to read post-voting record: %v", err)
   }
   if postVotingJSON == nil {
       return nil, fmt.Errorf("the post-voting record with ID %s does not exist", id)
   }

   var postVotingRecord struct {
       HashedVcmsToken   string `json:"hashedVcmsToken"`
       VotingToken string `json:"VotingToken"`
       ID                string `json:"id"`
    //    Salt              string `json:"salt"`
   }
   err = json.Unmarshal(postVotingJSON, &postVotingRecord)
   if err != nil {
       return nil, fmt.Errorf("failed to unmarshal post-voting record: %v", err)
   }

   return postVotingRecord, nil
}

// GetVotingTokenFromChannel2 retrieves and verifies the voting token from channel2
func (c *VotingContract) GetVotingTokenFromChannel2(ctx contractapi.TransactionContextInterface, votingTokenhash string) (*VotingTokenRecord, error) {
   // Retrieve the voting token from channel2
   channel2Stub := ctx.GetStub().InvokeChaincode("voting6", [][]byte{
       []byte("GetVotingTokenRecord"),
       []byte(votingTokenhash),
   }, "mychannel1")

   if channel2Stub.Status != 200 {
       return nil, fmt.Errorf("failed to retrieve voting token from channel2: %s", channel2Stub.Message)
   }

   var record VotingTokenRecord
   err := json.Unmarshal(channel2Stub.Payload, &record)
   if err != nil {
       return nil, fmt.Errorf("failed to unmarshal voting token record: %v", err)
   }

   return &record, nil
}

// idExistsInChannel1 checks if an ID already exists in channel1
func (c *VotingContract) idExistsInChannel1(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
   ballotJSON, err := ctx.GetStub().GetState(id)
   if err != nil {
       return false, fmt.Errorf("failed to read from world state: %v", err)
   }
   return ballotJSON != nil, nil
}

// GetBallot retrieves a ballot by ID
func (c *VotingContract) GetBallot(ctx contractapi.TransactionContextInterface, id string) (*BallotRecord, error) {
   ballotJSON, err := ctx.GetStub().GetState(id)
   if err != nil {
       return nil, fmt.Errorf("failed to read ballot record: %v", err)
   }
   if ballotJSON == nil {
       return nil, fmt.Errorf("the ballot with ID %s does not exist", id)
   }

   var ballot BallotRecord
   err = json.Unmarshal(ballotJSON, &ballot)
   if err != nil {
       return nil, fmt.Errorf("failed to unmarshal ballot record: %v", err)
   }

   return &ballot, nil
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