package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UpdateKValue is the `update_k_value` instruction.
type UpdateKValue struct {
	SpecId *uint16
	Val    *uint64

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

// NewUpdateKValueInstructionBuilder creates a new `UpdateKValue` instruction builder.
func NewUpdateKValueInstructionBuilder() *UpdateKValue {
	nd := &UpdateKValue{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 5),
	}
	nd.AccountMetaSlice[2] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[3] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[4] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetSpecId sets the "spec_id" parameter.
func (inst *UpdateKValue) SetSpecId(spec_id uint16) *UpdateKValue {
	inst.SpecId = &spec_id
	return inst
}

// SetVal sets the "val" parameter.
func (inst *UpdateKValue) SetVal(val uint64) *UpdateKValue {
	inst.Val = &val
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *UpdateKValue) SetSupernodeAccount(supernode ag_solanago.PublicKey) *UpdateKValue {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *UpdateKValue) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UpdateKValue) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *UpdateKValue) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *UpdateKValue) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *UpdateKValue) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *UpdateKValue) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetAdminAccount sets the "admin" account.
func (inst *UpdateKValue) SetAdminAccount(admin ag_solanago.PublicKey) *UpdateKValue {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UpdateKValue) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *UpdateKValue) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UpdateKValue {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *UpdateKValue) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UpdateKValue) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UpdateKValue {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UpdateKValue) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *UpdateKValue) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *UpdateKValue {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *UpdateKValue) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

func (inst UpdateKValue) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UpdateKValue,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UpdateKValue) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UpdateKValue) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.SpecId == nil {
			return errors.New("SpecId parameter is not set")
		}
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

func (inst *UpdateKValue) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UpdateKValue")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" SpecId", *inst.SpecId))
						paramsBranch.Child(ag_format.Param("    Val", *inst.Val))
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

func (obj UpdateKValue) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SpecId` param:
	err = encoder.Encode(obj.SpecId)
	if err != nil {
		return err
	}
	// Serialize `Val` param:
	err = encoder.Encode(obj.Val)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UpdateKValue) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SpecId`:
	err = decoder.Decode(&obj.SpecId)
	if err != nil {
		return err
	}
	// Deserialize `Val`:
	err = decoder.Decode(&obj.Val)
	if err != nil {
		return err
	}
	return nil
}

// NewUpdateKValueInstruction declares a new UpdateKValue instruction with the provided parameters and accounts.
func NewUpdateKValueInstruction(
	// Parameters:
	spec_id uint16,
	val uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *UpdateKValue {
	return NewUpdateKValueInstructionBuilder().
		SetSpecId(spec_id).
		SetVal(val).
		SetSupernodeAccount(supernode).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
