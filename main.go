package main

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	"log"
	sol_client "n3-solana-test/client"
)

const BILLION = 1000000000

var (
	client *rpc.Client
	ctx    context.Context

	admin       solana.PrivateKey
	fakeAdmin   solana.PrivateKey
	provider    solana.PrivateKey
	controller1 solana.PrivateKey
	controller2 solana.PrivateKey
	controller3 solana.PrivateKey
	user        solana.PrivateKey
	mint        solana.PublicKey

	providerTokenAccount    solana.PublicKey
	userTokenAccount        solana.PublicKey
	controllerTokenAccount1 solana.PublicKey
	controllerTokenAccount2 solana.PublicKey
	controllerTokenAccount3 solana.PublicKey

	supernodeProgramID = solana.MustPublicKeyFromBase58("B8YWYgxzsxGDuua6qsXZAvxL3huy5qy9AtL6AEmAVCCM")
)

func setupConnection() {
	ctx = context.Background()
	client = rpc.New(rpc.DevNet_RPC) // Change to Mainnet if needed
	fmt.Println("Connected to Solana DevNet")
}

func createAccount() solana.PrivateKey {
	account := solana.NewWallet()
	airdrop(account.PublicKey())
	return account.PrivateKey
}

func airdrop(pubKey solana.PublicKey) {
	sig, err := client.RequestAirdrop(ctx, pubKey, 10000*BILLION, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("Airdrop failed: %v", err)
	}
	fmt.Printf("Airdrop Tx: %s\n", sig)
}

func createTokenAccount(owner solana.PrivateKey, mint solana.PublicKey) solana.PublicKey {
	account, _, err := token.GetOrCreateAssociatedTokenAccount(ctx, client, owner.PublicKey(), mint, owner.PublicKey(), solana.ProgramToken)
	if err != nil {
		log.Fatalf("Failed to create token account: %v", err)
	}
	fmt.Printf("Created token account: %s\n", account)
	return account
}

func createMint(owner solana.PrivateKey) solana.PublicKey {
	mint, err := token.CreateMint(ctx, client, owner.PublicKey(), nil, 9, owner, solana.ProgramToken)
	if err != nil {
		log.Fatalf("Mint creation failed: %v", err)
	}
	fmt.Printf("Mint created: %s\n", mint)
	return mint
}

func initializeSupernode() {
	inst := NewInitializeInstruction(
		90,                // reward_locked_time
		10*BILLION,        // staking_coefficient
		mint,              // token mint
		admin.PublicKey(), // admin
	)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{inst.Build()},
		solana.TransactionPayer(admin.PublicKey()),
	)
	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err)
	}

	tx.Sign(admin)

	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}
	fmt.Printf("Supernode Initialized: %s\n", sig)
}

func stakeDevice(deviceID, specID uint64) {
	inst := sol_client.NewStakeDeviceInstruction(
		deviceID, specID,
		provider.PublicKey(),
		provider.PublicKey(),
		providerTokenAccount,
		mint,
		admin.PublicKey(),
	)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{inst.Build()},
		solana.TransactionPayer(provider.PublicKey()),
	)
	if err != nil {
		log.Fatalf("Failed to create transaction: %v", err)
	}

	tx.Sign(provider, admin)

	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	}
	fmt.Printf("Staked Device: %s\n", sig)
}
