package reward

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
)

const rpcURL = "https://api.devnet.solana.com" // Change to mainnet if needed

// Replace with the actual program ID from your Solana smart contract
var programID = solana.MustPublicKeyFromBase58("549dKMjWhEy5GeeR9uaQmhgL9nZafPMBpv9cG1wXdzjy")

func main() {
	ctx := context.Background()
	client := rpc.New(rpcURL)

	// Load admin wallet (private key)
	adminWallet := solana.MustPrivateKeyFromBase58("2fs5YwA2QkpkWx5qty4VhuASrNnnxFNTUhwaBUbcCS6R") // Replace with your keypair path

	// Derive the PDA for the Supernode
	supernodePDA, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode")}, programID)
	if err != nil {
		log.Fatalf("Failed to derive supernode PDA: %v", err)
	}

	// Derive the PDA for the Supernode Reward Account
	supernodeRewardPDA, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode_reward_account")}, programID)
	if err != nil {
		log.Fatalf("Failed to derive supernode reward account PDA: %v", err)
	}

	// Define Token Mint Account (Replace with actual mint address)
	tokenMint := solana.MustPublicKeyFromBase58("TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA")

	// Solana System Programs
	tokenProgram := solana.TokenProgramID
	systemProgram := solana.SystemProgramID
	associatedTokenProgram := solana.MustPublicKeyFromBase58("ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL")

	// Create the instruction
	//instruction := sol_client.NewInitRewardAccountInstruction(
	//	supernodePDA,
	//	supernodeRewardPDA,
	//	tokenMint,
	//	adminWallet.PublicKey(),
	//	tokenProgram,
	//	systemProgram,
	//	associatedTokenProgram,
	//)
	instruction := create

	// Validate instruction before building the transaction
	if err := instruction.Validate(); err != nil {
		log.Fatalf("Instruction validation failed: %v", err)
	}

	// Get the latest blockhash for the transaction
	recentBlockhash, err := client.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("Failed to fetch blockhash: %v", err)
	}

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

	// Send the transaction to Solana blockchain
	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}

	fmt.Printf("Transaction sent! Signature: %s\n", sig)
}
