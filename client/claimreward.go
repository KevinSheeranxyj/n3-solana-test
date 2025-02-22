// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ClaimReward is the `claim_reward` instruction.
type ClaimReward struct {
	Amount *uint64

	// [0] = [WRITE] supernode
	//
	// [1] = [WRITE] supernode_reward_account
	//
	// [2] = [] provider_stake_info
	//
	// [3] = [WRITE] provider_token_account
	//
	// [4] = [] token
	//
	// [5] = [] provider
	//
	// [6] = [WRITE, SIGNER] controller
	//
	// [7] = [WRITE, SIGNER] admin
	//
	// [8] = [] token_program
	//
	// [9] = [] system_program
	//
	// [10] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewClaimRewardInstructionBuilder creates a new `ClaimReward` instruction builder.
func NewClaimRewardInstructionBuilder() *ClaimReward {
	nd := &ClaimReward{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	nd.AccountMetaSlice[8] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[9] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[10] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *ClaimReward) SetAmount(amount uint64) *ClaimReward {
	inst.Amount = &amount
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *ClaimReward) SetSupernodeAccount(supernode ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *ClaimReward) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *ClaimReward) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *ClaimReward) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *ClaimReward) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *ClaimReward) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *ClaimReward) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeRewardAccountAccount sets the "supernode_reward_account" account.
func (inst *ClaimReward) SetSupernodeRewardAccountAccount(supernodeRewardAccount ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeRewardAccount).WRITE()
	return inst
}

func (inst *ClaimReward) findFindSupernodeRewardAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *ClaimReward) FindSupernodeRewardAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeRewardAccountAddress(bumpSeed)
	return
}

func (inst *ClaimReward) MustFindSupernodeRewardAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRewardAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeRewardAccountAddress finds SupernodeRewardAccount account address with given seeds.
func (inst *ClaimReward) FindSupernodeRewardAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeRewardAccountAddress(0)
	return
}

func (inst *ClaimReward) MustFindSupernodeRewardAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRewardAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeRewardAccountAccount gets the "supernode_reward_account" account.
func (inst *ClaimReward) GetSupernodeRewardAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *ClaimReward) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(providerStakeInfo)
	return inst
}

func (inst *ClaimReward) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: provider_stake_info
	seeds = append(seeds, []byte{byte(0x70), byte(0x72), byte(0x6f), byte(0x76), byte(0x69), byte(0x64), byte(0x65), byte(0x72), byte(0x5f), byte(0x73), byte(0x74), byte(0x61), byte(0x6b), byte(0x65), byte(0x5f), byte(0x69), byte(0x6e), byte(0x66), byte(0x6f)})
	// path: provider
	seeds = append(seeds, provider.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindProviderStakeInfoAddressWithBumpSeed calculates ProviderStakeInfo account address with given seeds and a known bump seed.
func (inst *ClaimReward) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *ClaimReward) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *ClaimReward) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *ClaimReward) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *ClaimReward) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProviderTokenAccountAccount sets the "provider_token_account" account.
func (inst *ClaimReward) SetProviderTokenAccountAccount(providerTokenAccount ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(providerTokenAccount).WRITE()
	return inst
}

// GetProviderTokenAccountAccount gets the "provider_token_account" account.
func (inst *ClaimReward) GetProviderTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenAccount sets the "token" account.
func (inst *ClaimReward) SetTokenAccount(token ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *ClaimReward) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProviderAccount sets the "provider" account.
func (inst *ClaimReward) SetProviderAccount(provider ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *ClaimReward) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetControllerAccount sets the "controller" account.
func (inst *ClaimReward) SetControllerAccount(controller ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(controller).WRITE().SIGNER()
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *ClaimReward) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAdminAccount sets the "admin" account.
func (inst *ClaimReward) SetAdminAccount(admin ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *ClaimReward) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *ClaimReward) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *ClaimReward) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *ClaimReward) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *ClaimReward) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *ClaimReward) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *ClaimReward {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *ClaimReward) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst ClaimReward) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ClaimReward,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ClaimReward) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ClaimReward) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SupernodeRewardAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ProviderStakeInfo is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ProviderTokenAccount is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Provider is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *ClaimReward) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ClaimReward")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       supernode_reward_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     provider_stake_info", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("         provider_token_", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                   token", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                provider", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              controller", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(10)))
					})
				})
		})
}

func (obj ClaimReward) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *ClaimReward) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewClaimRewardInstruction declares a new ClaimReward instruction with the provided parameters and accounts.
func NewClaimRewardInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeRewardAccount ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	providerTokenAccount ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	controller ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *ClaimReward {
	return NewClaimRewardInstructionBuilder().
		SetAmount(amount).
		SetSupernodeAccount(supernode).
		SetSupernodeRewardAccountAccount(supernodeRewardAccount).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderTokenAccountAccount(providerTokenAccount).
		SetTokenAccount(token).
		SetProviderAccount(provider).
		SetControllerAccount(controller).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
