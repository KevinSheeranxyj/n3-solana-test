package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// StakeDevice is the `stake_device` instruction.
type StakeDevice struct {
	DeviceId *uint64
	SpecId   *uint64

	// [0] = [] supernode
	//
	// [1] = [WRITE] supernode_stake_account
	//
	// [2] = [WRITE] provider_stake_info
	//
	// [3] = [WRITE] provider_token_account
	//
	// [4] = [] token
	//
	// [5] = [] provider
	//
	// [6] = [WRITE, SIGNER] controller
	//
	// [7] = [SIGNER] admin
	//
	// [8] = [] token_program
	//
	// [9] = [] system_program
	//
	// [10] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewStakeDeviceInstructionBuilder creates a new `StakeDevice` instruction builder.
func NewStakeDeviceInstructionBuilder() *StakeDevice {
	nd := &StakeDevice{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	nd.AccountMetaSlice[8] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[9] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[10] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetDeviceId sets the "device_id" parameter.
func (inst *StakeDevice) SetDeviceId(device_id uint64) *StakeDevice {
	inst.DeviceId = &device_id
	return inst
}

// SetSpecId sets the "spec_id" parameter.
func (inst *StakeDevice) SetSpecId(spec_id uint64) *StakeDevice {
	inst.SpecId = &spec_id
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *StakeDevice) SetSupernodeAccount(supernode ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode)
	return inst
}

func (inst *StakeDevice) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *StakeDevice) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *StakeDevice) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *StakeDevice) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *StakeDevice) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *StakeDevice) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeStakeAccountAccount sets the "supernode_stake_account" account.
func (inst *StakeDevice) SetSupernodeStakeAccountAccount(supernodeStakeAccount ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeStakeAccount).WRITE()
	return inst
}

func (inst *StakeDevice) findFindSupernodeStakeAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode_stake_account
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65), byte(0x5f), byte(0x73), byte(0x74), byte(0x61), byte(0x6b), byte(0x65), byte(0x5f), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeStakeAccountAddressWithBumpSeed calculates SupernodeStakeAccount account address with given seeds and a known bump seed.
func (inst *StakeDevice) FindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	return
}

func (inst *StakeDevice) MustFindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeStakeAccountAddress finds SupernodeStakeAccount account address with given seeds.
func (inst *StakeDevice) FindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeStakeAccountAddress(0)
	return
}

func (inst *StakeDevice) MustFindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeStakeAccountAccount gets the "supernode_stake_account" account.
func (inst *StakeDevice) GetSupernodeStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *StakeDevice) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(providerStakeInfo).WRITE()
	return inst
}

func (inst *StakeDevice) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *StakeDevice) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *StakeDevice) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *StakeDevice) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *StakeDevice) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *StakeDevice) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProviderTokenAccountAccount sets the "provider_token_account" account.
func (inst *StakeDevice) SetProviderTokenAccountAccount(providerTokenAccount ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(providerTokenAccount).WRITE()
	return inst
}

// GetProviderTokenAccountAccount gets the "provider_token_account" account.
func (inst *StakeDevice) GetProviderTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenAccount sets the "token" account.
func (inst *StakeDevice) SetTokenAccount(token ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *StakeDevice) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProviderAccount sets the "provider" account.
func (inst *StakeDevice) SetProviderAccount(provider ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *StakeDevice) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetControllerAccount sets the "controller" account.
func (inst *StakeDevice) SetControllerAccount(controller ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(controller).WRITE().SIGNER()
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *StakeDevice) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAdminAccount sets the "admin" account.
func (inst *StakeDevice) SetAdminAccount(admin ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *StakeDevice) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *StakeDevice) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *StakeDevice) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *StakeDevice) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *StakeDevice) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *StakeDevice) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *StakeDevice {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *StakeDevice) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst StakeDevice) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_StakeDevice,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst StakeDevice) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *StakeDevice) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DeviceId == nil {
			return errors.New("DeviceId parameter is not set")
		}
		if inst.SpecId == nil {
			return errors.New("SpecId parameter is not set")
		}
	}

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

func (inst *StakeDevice) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("StakeDevice")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" DeviceId", *inst.DeviceId))
						paramsBranch.Child(ag_format.Param("   SpecId", *inst.SpecId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        supernode_stake_", inst.AccountMetaSlice.Get(1)))
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

func (obj StakeDevice) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DeviceId` param:
	err = encoder.Encode(obj.DeviceId)
	if err != nil {
		return err
	}
	// Serialize `SpecId` param:
	err = encoder.Encode(obj.SpecId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *StakeDevice) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DeviceId`:
	err = decoder.Decode(&obj.DeviceId)
	if err != nil {
		return err
	}
	// Deserialize `SpecId`:
	err = decoder.Decode(&obj.SpecId)
	if err != nil {
		return err
	}
	return nil
}

// NewStakeDeviceInstruction declares a new StakeDevice instruction with the provided parameters and accounts.
func NewStakeDeviceInstruction(
	// Parameters:
	device_id uint64,
	spec_id uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeStakeAccount ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	providerTokenAccount ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	controller ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *StakeDevice {
	return NewStakeDeviceInstructionBuilder().
		SetDeviceId(device_id).
		SetSpecId(spec_id).
		SetSupernodeAccount(supernode).
		SetSupernodeStakeAccountAccount(supernodeStakeAccount).
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
