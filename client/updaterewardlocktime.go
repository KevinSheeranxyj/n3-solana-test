package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateRewardLockTime is the `update_reward_lock_time` instruction.
type UpdateRewardLockTime struct {
	New *uint64

	// [0] = [WRITE] supernode
	//
	// [1] = [SIGNER] admin
	//
	// [2] = [] token_program
	//
	// [3] = [] system_program
	//
	// [4] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewUpdateRewardLockTimeInstructionBuilder creates a new `UpdateRewardLockTime` instruction builder.
func NewUpdateRewardLockTimeInstructionBuilder() *UpdateRewardLockTime {
	nd := &UpdateRewardLockTime{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetNew sets the "new" parameter.
func (inst *UpdateRewardLockTime) SetNew(new uint64) *UpdateRewardLockTime {
	inst.New = &new
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *UpdateRewardLockTime) SetSupernodeAccount(supernode ag_solanago.PublicKey) *UpdateRewardLockTime {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *UpdateRewardLockTime) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UpdateRewardLockTime) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *UpdateRewardLockTime) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *UpdateRewardLockTime) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *UpdateRewardLockTime) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *UpdateRewardLockTime) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateRewardLockTime) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateRewardLockTime {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateRewardLockTime) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *UpdateRewardLockTime) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UpdateRewardLockTime {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *UpdateRewardLockTime) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UpdateRewardLockTime) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateRewardLockTime {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UpdateRewardLockTime) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *UpdateRewardLockTime) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *UpdateRewardLockTime {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *UpdateRewardLockTime) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateRewardLockTime) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateRewardLockTime,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateRewardLockTime) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateRewardLockTime) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.New == nil {
			return errors.New("New parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *UpdateRewardLockTime) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateRewardLockTime")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("New", *inst.New))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=5]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(4)))
					})
				})
		})
}

func (obj UpdateRewardLockTime) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateRewardLockTime) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateRewardLockTimeInstruction declares a new UpdateRewardLockTime instruction with the provided parameters and accounts.
func NewUpdateRewardLockTimeInstruction(
	// Parameters:
	new uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *UpdateRewardLockTime {
	return NewUpdateRewardLockTimeInstructionBuilder().
		SetNew(new).
		SetSupernodeAccount(supernode).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
