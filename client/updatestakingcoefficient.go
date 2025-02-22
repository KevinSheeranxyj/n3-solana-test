package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateStakingCoefficient is the `update_staking_coefficient` instruction.
type UpdateStakingCoefficient struct {
	Val *uint64

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

// NewUpdateStakingCoefficientInstructionBuilder creates a new `UpdateStakingCoefficient` instruction builder.
func NewUpdateStakingCoefficientInstructionBuilder() *UpdateStakingCoefficient {
	nd := &UpdateStakingCoefficient{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetVal sets the "val" parameter.
func (inst *UpdateStakingCoefficient) SetVal(val uint64) *UpdateStakingCoefficient {
	inst.Val = &val
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *UpdateStakingCoefficient) SetSupernodeAccount(supernode ag_solanago.PublicKey) *UpdateStakingCoefficient {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *UpdateStakingCoefficient) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UpdateStakingCoefficient) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *UpdateStakingCoefficient) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *UpdateStakingCoefficient) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *UpdateStakingCoefficient) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *UpdateStakingCoefficient) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateStakingCoefficient) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateStakingCoefficient {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateStakingCoefficient) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *UpdateStakingCoefficient) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UpdateStakingCoefficient {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *UpdateStakingCoefficient) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UpdateStakingCoefficient) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateStakingCoefficient {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UpdateStakingCoefficient) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *UpdateStakingCoefficient) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *UpdateStakingCoefficient {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *UpdateStakingCoefficient) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateStakingCoefficient) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateStakingCoefficient,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateStakingCoefficient) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateStakingCoefficient) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Val == nil {
			return errors.New("Val parameter is not set")
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

func (inst *UpdateStakingCoefficient) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateStakingCoefficient")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Val", *inst.Val))
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

func (obj UpdateStakingCoefficient) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Val` param:
	err = encoder.Encode(obj.Val)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateStakingCoefficient) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Val`:
	err = decoder.Decode(&obj.Val)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateStakingCoefficientInstruction declares a new UpdateStakingCoefficient instruction with the provided parameters and accounts.
func NewUpdateStakingCoefficientInstruction(
	// Parameters:
	val uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *UpdateStakingCoefficient {
	return NewUpdateStakingCoefficientInstructionBuilder().
		SetVal(val).
		SetSupernodeAccount(supernode).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
