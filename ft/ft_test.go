package ft

import (
	"context"
	"fmt"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	sendandconfirmtransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/test-go/testify/assert"
	"testing"
	"time"
)

func TestFT(t *testing.T) {
	rpcClient := rpc.New(rpc.DevNet_RPC)
	wsClient, err := ws.Connect(context.Background(), rpc.DevNet_WS)
	assert.NoError(t, err)

	ctx := context.TODO()

	to := solana.MustPublicKeyFromBase58("9B5XszUGdMaxCZ7uSQhPzdks5ZQSmWxrmzCSvtJ6Ns6b")
	fmt.Println(to)
	from := solana.NewWallet()
	fmt.Println(from.PublicKey())

	// Airdrop 5 SOL to the new account:
	out, err := rpcClient.RequestAirdrop(
		ctx,
		from.PublicKey(),
		solana.LAMPORTS_PER_SOL*2,
		rpc.CommitmentFinalized,
	)
	fmt.Println(out)
	assert.NoError(t, err)

	time.Sleep(30 * time.Second)

	recent, err := rpcClient.GetRecentBlockhash(ctx, rpc.CommitmentFinalized)
	assert.NoError(t, err)

	FTAccount := solana.NewWallet()
	createInst, err := system.NewCreateAccountInstruction(0, 82, solana.TokenProgramID, from.PublicKey(), FTAccount.PublicKey()).ValidateAndBuild()
	assert.NoError(t, err)

	mintInst, err := token.NewInitializeMintInstructionBuilder().
		SetDecimals(9).
		SetMintAuthority(from.PublicKey()).
		SetMintAccount(FTAccount.PublicKey()).
		SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).ValidateAndBuild()
	assert.NoError(t, err)

	tx, err := solana.NewTransaction(
		[]solana.Instruction{
			createInst,
			mintInst,
		},
		recent.Value.Blockhash,
		solana.TransactionPayer(from.PublicKey()),
	)
	assert.NoError(t, err)

	_, err = tx.Sign(
		func(key solana.PublicKey) *solana.PrivateKey {
			if from.PublicKey().Equals(key) {
				return &from.PrivateKey
			}
			if FTAccount.PublicKey().Equals(key) {
				return &FTAccount.PrivateKey
			}
			return nil
		},
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
