package alpha_stake

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"n3-solana-test/client"
	"testing"
)

var (
	// Define your variables
	ProgramID    = solana.MustPublicKeyFromBase58("549dKMjWhEy5GeeR9uaQmhgL9nZafPMBpv9cG1wXdzjy")
	RPC_ENDPOINT = "https://api.devnet.solana.com"                                                     // Or your preferred endpoint
	provider, _  = solana.PrivateKeyFromSolanaKeygenFile("/Users/kevinsheeran/.config/solana/id.json") // Generate or load your provider keypair
	admin, _     = solana.PrivateKeyFromSolanaKeygenFile("/Users/kevinsheeran/.config/solana/id.json") // Generate or load your admin keypair
	//providerTokenAccount solana.PublicKey                  // Set your provider token account
	//mint                 solana.PublicKey                  // Set your token mint
	mint                 = solana.MustPublicKeyFromBase58("4vSVTKJtE1Fmr5z9HgDuh4d7yw93PUQXXVUDp23vXtG5") // Your token mint address
	providerTokenAccount = solana.MustPublicKeyFromBase58("9R8bHte4xYauxEm3pGrAzniWcfKMuCRcbGzqrDvf4zzV")
	connection           = rpc.New(RPC_ENDPOINT)
)

func createStakeDeviceTx() (*solana.Transaction, error) {
	// Find or set PDAs
	supernodePDA, _, err := client.NewStakeDeviceInstructionBuilder().FindSupernodeAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to find supernode PDA: %v", err)
	}
	fmt.Println(supernodePDA.String())

	supernodeStakePDA, _, err := client.NewStakeDeviceInstructionBuilder().FindSupernodeStakeAccountAddress()
	if err != nil {
		return nil, fmt.Errorf("failed to find supernode stake PDA: %v", err)
	}
	fmt.Println(supernodeStakePDA.String())

	providerStakeInfoPDA, _, err := client.NewStakeDeviceInstructionBuilder().FindProviderStakeInfoAddress(provider.PublicKey())
	if err != nil {
		return nil, fmt.Errorf("failed to find provider stake info PDA: %v", err)
	}
	fmt.Println(providerStakeInfoPDA.String())

	// Build the instruction
	stakeInst := client.NewStakeDeviceInstruction(
		// Parameters
		1, // device_id (equivalent to BN(1))
		1, // spec_id (equivalent to BN(1))
		// Accounts
		supernodePDA,           // supernode
		supernodeStakePDA,      // supernode_stake_account
		providerStakeInfoPDA,   // provider_stake_info
		providerTokenAccount,   // provider_token_account
		mint,                   // token
		provider.PublicKey(),   // provider
		provider.PublicKey(),   // controller
		admin.PublicKey(),      // admin
		solana.TokenProgramID,  // token_program
		solana.SystemProgramID, // system_program
		solana.SPLAssociatedTokenAccountProgramID, // associated_token_program
	)

	// Get recent blockhash
	recent, err := connection.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent blockhash: %v", err)
	}

	// Build the transaction
	instruction, err := stakeInst.ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("failed to validate and build instruction: %v", err)
	}

	tx, err := solana.NewTransaction(
		[]solana.Instruction{instruction},
		recent.Value.Blockhash,
		solana.TransactionPayer(provider.PublicKey()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}

	return tx, nil
}

func doubleSignAndSend(tx *solana.Transaction, user1, user2 solana.PrivateKey) (solana.Signature, error) {
	// Get latest blockhash
	recent, err := connection.GetLatestBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to get recent blockhash: %v", err)
	}

	// Update transaction with recent blockhash
	tx.Message.RecentBlockhash = recent.Value.Blockhash

	// Sign with both keys
	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		if key.Equals(user1.PublicKey()) {
			return &user1
		}
		if key.Equals(user2.PublicKey()) {
			return &user2
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to sign transaction: %v", err)
	}

	// Send transaction
	sig, err := connection.SendTransactionWithOpts(
		context.Background(),
		tx,
		rpc.TransactionOpts{
			SkipPreflight:       false,
			PreflightCommitment: rpc.CommitmentConfirmed,
		},
	)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("failed to send transaction: %v", err)
	}

	// Wait for confirmation
	_, err = connection.GetSignatureStatuses(context.Background(), true, []solana.Signature{sig}...)
	if err != nil {
		return solana.Signature{}, fmt.Errorf("transaction confirmation failed: %v", err)
	}

	return sig, nil
}

func executeStakeDevice() (solana.Signature, error) {
	// Create the transaction
	tx, err := createStakeDeviceTx()
	if err != nil {
		return solana.Signature{}, err
	}

	// Sign and send with both signatures
	sig, err := doubleSignAndSend(tx, provider, admin)
	if err != nil {
		return solana.Signature{}, err
	}

	return sig, nil
}

func TestStakeDevice(t *testing.T) {
	sig, err := executeStakeDevice()
	if err != nil {
		fmt.Printf("Error executing stake device: %v\n", err)
		return
	}
	fmt.Printf("Transaction successful with signature: %s\n", sig.String())
}
