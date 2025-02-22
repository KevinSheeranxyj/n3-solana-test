package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// ReplaceExtraController is the `replace_extra_controller` instruction.
type ReplaceExtraController struct {

	// [0] = [] supernode
	//
	// [1] = [WRITE] provider_stake_info
	//
	// [2] = [] provider
	//
	// [3] = [SIGNER] operator
	//
	// [4] = [] old_controller
	//
	// [5] = [SIGNER] admin
	//
	// [6] = [] new_controller
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewReplaceExtraControllerInstructionBuilder creates a new `ReplaceExtraController` instruction builder.
func NewReplaceExtraControllerInstructionBuilder() *ReplaceExtraController {
	nd := &ReplaceExtraController{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 7),
	}
	return nd
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *ReplaceExtraController) SetSupernodeAccount(supernode ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode)
	return inst
}

func (inst *ReplaceExtraController) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *ReplaceExtraController) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *ReplaceExtraController) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *ReplaceExtraController) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *ReplaceExtraController) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *ReplaceExtraController) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *ReplaceExtraController) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(providerStakeInfo).WRITE()
	return inst
}

func (inst *ReplaceExtraController) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *ReplaceExtraController) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *ReplaceExtraController) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *ReplaceExtraController) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *ReplaceExtraController) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *ReplaceExtraController) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderAccount sets the "provider" account.
func (inst *ReplaceExtraController) SetProviderAccount(provider ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *ReplaceExtraController) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOperatorAccount sets the "operator" account.
func (inst *ReplaceExtraController) SetOperatorAccount(operator ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(operator).SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
func (inst *ReplaceExtraController) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetOldControllerAccount sets the "old_controller" account.
func (inst *ReplaceExtraController) SetOldControllerAccount(oldController ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(oldController)
	return inst
}

// GetOldControllerAccount gets the "old_controller" account.
func (inst *ReplaceExtraController) GetOldControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAdminAccount sets the "admin" account.
func (inst *ReplaceExtraController) SetAdminAccount(admin ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *ReplaceExtraController) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetNewControllerAccount sets the "new_controller" account.
func (inst *ReplaceExtraController) SetNewControllerAccount(newController ag_solanago.PublicKey) *ReplaceExtraController {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(newController)
	return inst
}

// GetNewControllerAccount gets the "new_controller" account.
func (inst *ReplaceExtraController) GetNewControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

func (inst ReplaceExtraController) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_ReplaceExtraController,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst ReplaceExtraController) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *ReplaceExtraController) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.ProviderStakeInfo is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.Provider is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Operator is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.OldController is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.NewController is not set")
		}
	}
	return nil
}

func (inst *ReplaceExtraController) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("ReplaceExtraController")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=7]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("provider_stake_info", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           provider", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           operator", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("     old_controller", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("              admin", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("     new_controller", inst.AccountMetaSlice.Get(6)))
					})
				})
		})
}

func (obj ReplaceExtraController) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *ReplaceExtraController) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewReplaceExtraControllerInstruction declares a new ReplaceExtraController instruction with the provided parameters and accounts.
func NewReplaceExtraControllerInstruction(
	// Accounts:
	supernode ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	operator ag_solanago.PublicKey,
	oldController ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	newController ag_solanago.PublicKey) *ReplaceExtraController {
	return NewReplaceExtraControllerInstructionBuilder().
		SetSupernodeAccount(supernode).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderAccount(provider).
		SetOperatorAccount(operator).
		SetOldControllerAccount(oldController).
		SetAdminAccount(admin).
		SetNewControllerAccount(newController)
}
