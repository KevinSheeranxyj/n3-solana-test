package reward

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"github.com/test-go/testify/assert"
	"log"
	sol_client "n3-solana-test/client"
	"testing"
)

func Test_Reward(t *testing.T) {

	ctx := context.Background()
	client := rpc.New(rpcURL)

	// Load admin wallet (private key)
	//adminWallet, err := solana.PrivateKeyFromSolanaKeygenFile("/Users/kevinsheeran/Developer/depinx-repo/n3-solana-smart-contract/target/deploy/supernode-keypair.json") // Replace with your keypair path
	adminWallet, err := solana.PrivateKeyFromSolanaKeygenFile("/Users/kevinsheeran/.config/solana/id.json") // Replace with your keypair path
	if err != nil {
		log.Fatalf("failed to generate private key: %v", err)
	}

	// Derive the PDA for the Supernode
	supernodePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode")}, programID)
	if err != nil {
		log.Fatalf("Failed to derive supernode PDA: %v", err)
	}

	//supernodePDA, _, err := sol_client.NewInitRewardAccountInstructionBuilder().FindSupernodeAddress()
	//if err != nil {
	//	log.Fatalf("failed to find supernode reward account address: %v", err)
	//}

	// Derive the PDA for the Supernode Reward Account
	supernodeRewardPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode_reward_account")}, programID)
	if err != nil {
		log.Fatalf("Failed to derive supernode reward account PDA: %v", err)
	}

	// Define Token Mint Account (Replace with actual mint address)
	tokenMint := solana.MustPublicKeyFromBase58("4vSVTKJtE1Fmr5z9HgDuh4d7yw93PUQXXVUDp23vXtG5")

	// Solana System Programs
	tokenProgram := solana.TokenProgramID
	systemProgram := solana.SystemProgramID
	associatedTokenProgram := solana.SPLAssociatedTokenAccountProgramID

	log.Printf("supernodePDA %v", supernodePDA)
	log.Printf("supernodeRewardPDA %v", supernodePDA)
	log.Printf("tokenMint %v", tokenMint)

	log.Printf("adminWallet %v", adminWallet.PublicKey())

	// Create the instruction
	instruction := sol_client.NewInitRewardAccountInstruction(
		supernodePDA,
		supernodeRewardPDA,
		tokenMint,
		adminWallet.PublicKey(),
		tokenProgram,
		systemProgram,
		associatedTokenProgram,
	)

	accountMeta := instruction.GetAccounts()
	for _, acc := range accountMeta {
		log.Printf("accountMeta: %v", acc.PublicKey)
	}

	tokenAccount := instruction.GetTokenAccount()
	log.Printf("tokenAccount %v", tokenAccount)

	adminAccount := instruction.GetAdminAccount()
	log.Printf("adminAccount %v", adminAccount)

	// Validate instruction before building the transaction
	if err := instruction.Validate(); err != nil {
		log.Fatalf("Instruction validation failed: %v", err)
	}

	assert.NoError(t, err, "Failed to execute transaction")

	// Get the latest blockhash for the transaction
	recentBlockhash, err := client.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("Failed to fetch blockhash: %v", err)
	}

	log.Printf("recentBlockhash %v", recentBlockhash.Value)

	// Create a new transaction
	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction.Build()},
		recentBlockhash.Value.Blockhash,
		solana.TransactionPayer(adminWallet.PublicKey()),
	)
	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err)
	}

	// Sign the transaction
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(adminWallet.PublicKey()) {
			return &adminWallet
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Transaction signing failed: %v", err)
	}

	//Send the transaction to Solana blockchain
	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	fmt.Printf("Transaction sent! Signature: %s\n", sig)

}

//func createRewardInstruction() *sol_client.InitRewardAccount {
//
//	builder := sol_client.NewInitRewardAccountInstructionBuilder()
//	supernodePDA := builder.MustFindSupernodeAddress()
//	builder.SetSupernodeAccount(supernodePDA)
//	supernodeStakePDA := builder.MustFindSupernodeRewardAccountAddress()
//	builder.SetSupernodeRewardAccountAccount(supernodeStakePDA)
//
//	builder.SetTokenAccount()
//	builder.SetAdminAccount()
//	builder.SetTokenProgramAccount()
//	builder.SetSystemProgramAccount()
//	return builder
//}

func NewTokenAccount() {

}
