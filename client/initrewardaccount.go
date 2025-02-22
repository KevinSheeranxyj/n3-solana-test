package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// InitRewardAccount is the `init_reward_account` instruction.
type InitRewardAccount struct {

	// [0] = [WRITE] supernode
	//
	// [1] = [WRITE] supernode_reward_account
	//
	// [2] = [] token
	//
	// [3] = [WRITE, SIGNER] admin
	//
	// [4] = [] token_program
	//
	// [5] = [] system_program
	//
	// [6] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitRewardAccountInstructionBuilder creates a new `InitRewardAccount` instruction builder.
func NewInitRewardAccountInstructionBuilder() *InitRewardAccount {
	nd := &InitRewardAccount{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[5] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[6] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *InitRewardAccount) SetSupernodeAccount(supernode ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *InitRewardAccount) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeAddressWithBumpSeed calculates Supernode account address with given seeds and a known bump seed.
func (inst *InitRewardAccount) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *InitRewardAccount) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *InitRewardAccount) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *InitRewardAccount) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *InitRewardAccount) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeRewardAccountAccount sets the "supernode_reward_account" account.
func (inst *InitRewardAccount) SetSupernodeRewardAccountAccount(supernodeRewardAccount ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeRewardAccount).WRITE()
	return inst
}

func (inst *InitRewardAccount) findFindSupernodeRewardAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode_reward_account
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65), byte(0x5f), byte(0x72), byte(0x65), byte(0x77), byte(0x61), byte(0x72), byte(0x64), byte(0x5f), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeRewardAccountAddressWithBumpSeed calculates SupernodeRewardAccount account address with given seeds and a known bump seed.
func (inst *InitRewardAccount) FindSupernodeRewardAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeRewardAccountAddress(bumpSeed)
	return
}

func (inst *InitRewardAccount) MustFindSupernodeRewardAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRewardAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeRewardAccountAddress finds SupernodeRewardAccount account address with given seeds.
func (inst *InitRewardAccount) FindSupernodeRewardAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeRewardAccountAddress(0)
	return
}

func (inst *InitRewardAccount) MustFindSupernodeRewardAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRewardAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeRewardAccountAccount gets the "supernode_reward_account" account.
func (inst *InitRewardAccount) GetSupernodeRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenAccount sets the "token" account.
func (inst *InitRewardAccount) SetTokenAccount(token ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *InitRewardAccount) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetAdminAccount sets the "admin" account.
func (inst *InitRewardAccount) SetAdminAccount(admin ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *InitRewardAccount) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *InitRewardAccount) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *InitRewardAccount) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *InitRewardAccount) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *InitRewardAccount) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *InitRewardAccount) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *InitRewardAccount {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *InitRewardAccount) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst InitRewardAccount) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_InitRewardAccount,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst InitRewardAccount) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *InitRewardAccount) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SupernodeRewardAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *InitRewardAccount) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("InitRewardAccount")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       supernode_reward_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                   token", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj InitRewardAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *InitRewardAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewInitRewardAccountInstruction declares a new InitRewardAccount instruction with the provided parameters and accounts.
func NewInitRewardAccountInstruction(
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeRewardAccount ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *InitRewardAccount {
	return NewInitRewardAccountInstructionBuilder().
		SetSupernodeAccount(supernode).
		SetSupernodeRewardAccountAccount(supernodeRewardAccount).
		SetTokenAccount(token).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
