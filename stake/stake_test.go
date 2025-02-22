package stake

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	sendandconfirmtransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/test-go/testify/assert"
	"log"
	"n3-solana-test/client"
	"testing"
)

func Test_stakedevice(t *testing.T) {

	rpcClient := rpc.New(rpc.DevNet_RPC)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)

	ctx := context.TODO()

	admin, err := solana.PrivateKeyFromSolanaKeygenFile("/Users/kevinsheeran/.config/solana/id.json") // Replace with your keypair path

	supernodeAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode")}, programID)
	assert.NoError(t, err)
	supernodeStakeAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("supernode_stake_account")}, programID)
	assert.NoError(t, err)
	providerStakeInfoAccount, _, err := solana.FindProgramAddress([][]byte{[]byte("provider_stake_info")}, programID)
	assert.NoError(t, err)

	provider := solana.NewWallet()

	tokenAccount := solana.NewWallet()

	controller := solana.NewWallet()

	//providerStakeAccount, err := token.NewInitializeMintInstructionBuilder().
	//	SetDecimals(9).
	//	SetMintAuthority(admin.PublicKey()).
	//	SetMintAccount(supernodeAccount).
	//	SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).ValidateAndBuild()

	// Get the latest blockhash for the transaction
	recent, err := rpcClient.GetLatestBlockhash(ctx, rpc.CommitmentFinalized)
	if err != nil {
		log.Fatalf("Failed to fetch blockhash: %v", err)
	}

	log.Printf("recentBlockhash %v", recent.Value)

	providerStakeTokenAccount, err := token.NewInitializeMintInstructionBuilder().
		SetDecimals(9).
		SetMintAuthority(admin.PublicKey()).
		SetMintAccount(supernodeAccount).
		SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).ValidateAndBuild()
	inst, err := client.NewStakeDeviceInstructionBuilder().
		SetDeviceId(solana.LAMPORTS_PER_SOL * 2).
		SetSpecId(solana.LAMPORTS_PER_SOL * 2).
		SetSupernodeAccount(supernodeAccount).
		SetSupernodeStakeAccountAccount(supernodeStakeAccount).
		SetProviderStakeInfoAccount(providerStakeInfoAccount).
		SetProviderTokenAccountAccount(providerStakeTokenAccount.ProgramID()).
		SetTokenAccount(tokenAccount.PublicKey()).
		SetProviderAccount(provider.PublicKey()).
		SetControllerAccount(controller.PublicKey()).
		SetAdminAccount(admin.PublicKey()).
		ValidateAndBuild()
	assert.NoError(t, err)
	fmt.Printf("%v", inst)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			inst,
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(admin.PublicKey()),
	)
	assert.NoError(t, err)

	sig, err := sendandconfirmtransaction.SendAndConfirmTransaction(
		ctx,
		rpcClient,
		wsClient,
		tx,
	)
	fmt.Println(sig)
	assert.NoError(t, err)

}
