package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Release is the `release` instruction.
type Release struct {

	// [0] = [] supernode
	//
	// [1] = [WRITE] supernode_stake_account
	//
	// [2] = [] provider_stake_info
	//
	// [3] = [WRITE] provider_vesting_info
	//
	// [4] = [WRITE] provider_token_account
	//
	// [5] = [] token
	//
	// [6] = [] provider
	//
	// [7] = [WRITE, SIGNER] controller
	//
	// [8] = [SIGNER] admin
	//
	// [9] = [] token_program
	//
	// [10] = [] system_program
	//
	// [11] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReleaseInstructionBuilder creates a new `Release` instruction builder.
func NewReleaseInstructionBuilder() *Release {
	nd := &Release{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 12),
	}
	nd.AccountMetaSlice[9] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[10] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[11] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *Release) SetSupernodeAccount(supernode ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode)
	return inst
}

func (inst *Release) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Release) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *Release) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *Release) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *Release) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *Release) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeStakeAccountAccount sets the "supernode_stake_account" account.
func (inst *Release) SetSupernodeStakeAccountAccount(supernodeStakeAccount ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeStakeAccount).WRITE()
	return inst
}

func (inst *Release) findFindSupernodeStakeAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode_token_account
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65), byte(0x5f), byte(0x74), byte(0x6f), byte(0x6b), byte(0x65), byte(0x6e), byte(0x5f), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeStakeAccountAddressWithBumpSeed calculates SupernodeStakeAccount account address with given seeds and a known bump seed.
func (inst *Release) FindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	return
}

func (inst *Release) MustFindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeStakeAccountAddress finds SupernodeStakeAccount account address with given seeds.
func (inst *Release) FindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeStakeAccountAddress(0)
	return
}

func (inst *Release) MustFindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeStakeAccountAccount gets the "supernode_stake_account" account.
func (inst *Release) GetSupernodeStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *Release) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(providerStakeInfo)
	return inst
}

func (inst *Release) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Release) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *Release) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *Release) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *Release) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *Release) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProviderVestingInfoAccount sets the "provider_vesting_info" account.
func (inst *Release) SetProviderVestingInfoAccount(providerVestingInfo ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(providerVestingInfo).WRITE()
	return inst
}

func (inst *Release) findFindProviderVestingInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: provider_vesting_info
	seeds = append(seeds, []byte{byte(0x70), byte(0x72), byte(0x6f), byte(0x76), byte(0x69), byte(0x64), byte(0x65), byte(0x72), byte(0x5f), byte(0x76), byte(0x65), byte(0x73), byte(0x74), byte(0x69), byte(0x6e), byte(0x67), byte(0x5f), byte(0x69), byte(0x6e), byte(0x66), byte(0x6f)})
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

// FindProviderVestingInfoAddressWithBumpSeed calculates ProviderVestingInfo account address with given seeds and a known bump seed.
func (inst *Release) FindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	return
}

func (inst *Release) MustFindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderVestingInfoAddress finds ProviderVestingInfo account address with given seeds.
func (inst *Release) FindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderVestingInfoAddress(provider, 0)
	return
}

func (inst *Release) MustFindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderVestingInfoAccount gets the "provider_vesting_info" account.
func (inst *Release) GetProviderVestingInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProviderTokenAccountAccount sets the "provider_token_account" account.
func (inst *Release) SetProviderTokenAccountAccount(providerTokenAccount ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(providerTokenAccount).WRITE()
	return inst
}

// GetProviderTokenAccountAccount gets the "provider_token_account" account.
func (inst *Release) GetProviderTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTokenAccount sets the "token" account.
func (inst *Release) SetTokenAccount(token ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *Release) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetProviderAccount sets the "provider" account.
func (inst *Release) SetProviderAccount(provider ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *Release) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetControllerAccount sets the "controller" account.
func (inst *Release) SetControllerAccount(controller ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(controller).WRITE().SIGNER()
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *Release) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAdminAccount sets the "admin" account.
func (inst *Release) SetAdminAccount(admin ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *Release) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Release) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Release) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Release) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Release) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *Release) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *Release {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *Release) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

func (inst Release) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Release,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Release) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Release) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SupernodeStakeAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.ProviderStakeInfo is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ProviderVestingInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ProviderTokenAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Provider is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Controller is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *Release) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Release")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=12]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        supernode_stake_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("     provider_stake_info", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("   provider_vesting_info", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("         provider_token_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   token", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                provider", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("              controller", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(11)))
					})
				})
		})
}

func (obj Release) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Release) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewReleaseInstruction declares a new Release instruction with the provided parameters and accounts.
func NewReleaseInstruction(
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeStakeAccount ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	providerVestingInfo ag_solanago.PublicKey,
	providerTokenAccount ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	controller ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *Release {
	return NewReleaseInstructionBuilder().
		SetSupernodeAccount(supernode).
		SetSupernodeStakeAccountAccount(supernodeStakeAccount).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderVestingInfoAccount(providerVestingInfo).
		SetProviderTokenAccountAccount(providerTokenAccount).
		SetTokenAccount(token).
		SetProviderAccount(provider).
		SetControllerAccount(controller).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
