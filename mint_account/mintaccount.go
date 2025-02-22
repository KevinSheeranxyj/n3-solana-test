package mint_account

import (
	"github.com/gagliardetto/solana-go"
	"log"
)

func creat() {
	tokenMintKeypair := solana.NewWallet()
	adminWallet := solana.NewWallet()

	tx, err := solana.NewTransaction(
		[]solana.Instruction{createMintInstruction},
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

	// Send the transaction
	sig, err := client.SendTransaction(ctx, tx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	log.Printf("Created new token mint! Signature: %s\n", sig)
}
