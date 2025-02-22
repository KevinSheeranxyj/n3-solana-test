package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Releasable is the `releasable` instruction.
type Releasable struct {

	// [0] = [WRITE] provider_vesting_info
	//
	// [1] = [] provider_stake_info
	//
	// [2] = [] provider
	//
	// [3] = [WRITE, SIGNER] controller
	//
	// [4] = [] token_program
	//
	// [5] = [] system_program
	//
	// [6] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReleasableInstructionBuilder creates a new `Releasable` instruction builder.
func NewReleasableInstructionBuilder() *Releasable {
	nd := &Releasable{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[5] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[6] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetProviderVestingInfoAccount sets the "provider_vesting_info" account.
func (inst *Releasable) SetProviderVestingInfoAccount(providerVestingInfo ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(providerVestingInfo).WRITE()
	return inst
}

func (inst *Releasable) findFindProviderVestingInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Releasable) FindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	return
}

func (inst *Releasable) MustFindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderVestingInfoAddress finds ProviderVestingInfo account address with given seeds.
func (inst *Releasable) FindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderVestingInfoAddress(provider, 0)
	return
}

func (inst *Releasable) MustFindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderVestingInfoAccount gets the "provider_vesting_info" account.
func (inst *Releasable) GetProviderVestingInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *Releasable) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(providerStakeInfo)
	return inst
}

func (inst *Releasable) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Releasable) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *Releasable) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *Releasable) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *Releasable) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *Releasable) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderAccount sets the "provider" account.
func (inst *Releasable) SetProviderAccount(provider ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *Releasable) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetControllerAccount sets the "controller" account.
func (inst *Releasable) SetControllerAccount(controller ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(controller).WRITE().SIGNER()
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *Releasable) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Releasable) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Releasable) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Releasable) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Releasable) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *Releasable) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *Releasable {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *Releasable) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst Releasable) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Releasable,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Releasable) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Releasable) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.ProviderVestingInfo is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ProviderStakeInfo is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Provider is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Controller is not set")
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

func (inst *Releasable) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Releasable")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("   provider_vesting_info", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("     provider_stake_info", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                provider", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("              controller", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj Releasable) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *Releasable) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewReleasableInstruction declares a new Releasable instruction with the provided parameters and accounts.
func NewReleasableInstruction(
	// Accounts:
	providerVestingInfo ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	controller ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *Releasable {
	return NewReleasableInstructionBuilder().
		SetProviderVestingInfoAccount(providerVestingInfo).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderAccount(provider).
		SetControllerAccount(controller).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
