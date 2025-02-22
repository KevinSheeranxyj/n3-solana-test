package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// RemoveExtraController is the `remove_extra_controller` instruction.
type RemoveExtraController struct {

	// [0] = [] supernode
	//
	// [1] = [WRITE] provider_stake_info
	//
	// [2] = [] provider
	//
	// [3] = [SIGNER] operator
	//
	// [4] = [SIGNER] admin
	//
	// [5] = [] old_controller
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewRemoveExtraControllerInstructionBuilder creates a new `RemoveExtraController` instruction builder.
func NewRemoveExtraControllerInstructionBuilder() *RemoveExtraController {
	nd := &RemoveExtraController{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 6),
	}
	return nd
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *RemoveExtraController) SetSupernodeAccount(supernode ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode)
	return inst
}

func (inst *RemoveExtraController) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *RemoveExtraController) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *RemoveExtraController) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *RemoveExtraController) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *RemoveExtraController) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *RemoveExtraController) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *RemoveExtraController) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(providerStakeInfo).WRITE()
	return inst
}

func (inst *RemoveExtraController) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *RemoveExtraController) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *RemoveExtraController) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *RemoveExtraController) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *RemoveExtraController) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *RemoveExtraController) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderAccount sets the "provider" account.
func (inst *RemoveExtraController) SetProviderAccount(provider ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *RemoveExtraController) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOperatorAccount sets the "operator" account.
func (inst *RemoveExtraController) SetOperatorAccount(operator ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(operator).SIGNER()
	return inst
}

// GetOperatorAccount gets the "operator" account.
func (inst *RemoveExtraController) GetOperatorAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAdminAccount sets the "admin" account.
func (inst *RemoveExtraController) SetAdminAccount(admin ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *RemoveExtraController) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetOldControllerAccount sets the "old_controller" account.
func (inst *RemoveExtraController) SetOldControllerAccount(oldController ag_solanago.PublicKey) *RemoveExtraController {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(oldController)
	return inst
}

// GetOldControllerAccount gets the "old_controller" account.
func (inst *RemoveExtraController) GetOldControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

func (inst RemoveExtraController) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_RemoveExtraController,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RemoveExtraController) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RemoveExtraController) Validate() error {
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
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.OldController is not set")
		}
	}
	return nil
}

func (inst *RemoveExtraController) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RemoveExtraController")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=6]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("          supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("provider_stake_info", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           provider", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("           operator", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("              admin", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("     old_controller", inst.AccountMetaSlice.Get(5)))
					})
				})
		})
}

func (obj RemoveExtraController) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RemoveExtraController) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRemoveExtraControllerInstruction declares a new RemoveExtraController instruction with the provided parameters and accounts.
func NewRemoveExtraControllerInstruction(
	// Accounts:
	supernode ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	operator ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	oldController ag_solanago.PublicKey) *RemoveExtraController {
	return NewRemoveExtraControllerInstructionBuilder().
		SetSupernodeAccount(supernode).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderAccount(provider).
		SetOperatorAccount(operator).
		SetAdminAccount(admin).
		SetOldControllerAccount(oldController)
}
