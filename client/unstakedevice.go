package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// UnstakeDevice is the `unstake_device` instruction.
type UnstakeDevice struct {
	DeviceId *uint64

	// [0] = [] supernode
	//
	// [1] = [WRITE] supernode_stake_account
	//
	// [2] = [WRITE] supernode_vesting_account
	//
	// [3] = [WRITE] provider_stake_info
	//
	// [4] = [WRITE] provider_vesting_info
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

// NewUnstakeDeviceInstructionBuilder creates a new `UnstakeDevice` instruction builder.
func NewUnstakeDeviceInstructionBuilder() *UnstakeDevice {
	nd := &UnstakeDevice{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 11),
	}
	nd.AccountMetaSlice[8] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[9] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[10] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetDeviceId sets the "device_id" parameter.
func (inst *UnstakeDevice) SetDeviceId(device_id uint64) *UnstakeDevice {
	inst.DeviceId = &device_id
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *UnstakeDevice) SetSupernodeAccount(supernode ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode)
	return inst
}

func (inst *UnstakeDevice) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UnstakeDevice) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *UnstakeDevice) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *UnstakeDevice) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeStakeAccountAccount sets the "supernode_stake_account" account.
func (inst *UnstakeDevice) SetSupernodeStakeAccountAccount(supernodeStakeAccount ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeStakeAccount).WRITE()
	return inst
}

func (inst *UnstakeDevice) findFindSupernodeStakeAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UnstakeDevice) FindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeStakeAccountAddress finds SupernodeStakeAccount account address with given seeds.
func (inst *UnstakeDevice) FindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeStakeAccountAddress(0)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeStakeAccountAccount gets the "supernode_stake_account" account.
func (inst *UnstakeDevice) GetSupernodeStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSupernodeVestingAccountAccount sets the "supernode_vesting_account" account.
func (inst *UnstakeDevice) SetSupernodeVestingAccountAccount(supernodeVestingAccount ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(supernodeVestingAccount).WRITE()
	return inst
}

func (inst *UnstakeDevice) findFindSupernodeVestingAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode_vesting_account
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65), byte(0x5f), byte(0x76), byte(0x65), byte(0x73), byte(0x74), byte(0x69), byte(0x6e), byte(0x67), byte(0x5f), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeVestingAccountAddressWithBumpSeed calculates SupernodeVestingAccount account address with given seeds and a known bump seed.
func (inst *UnstakeDevice) FindSupernodeVestingAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeVestingAccountAddress(bumpSeed)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeVestingAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeVestingAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeVestingAccountAddress finds SupernodeVestingAccount account address with given seeds.
func (inst *UnstakeDevice) FindSupernodeVestingAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeVestingAccountAddress(0)
	return
}

func (inst *UnstakeDevice) MustFindSupernodeVestingAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeVestingAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeVestingAccountAccount gets the "supernode_vesting_account" account.
func (inst *UnstakeDevice) GetSupernodeVestingAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetProviderStakeInfoAccount sets the "provider_stake_info" account.
func (inst *UnstakeDevice) SetProviderStakeInfoAccount(providerStakeInfo ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(providerStakeInfo).WRITE()
	return inst
}

func (inst *UnstakeDevice) findFindProviderStakeInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UnstakeDevice) FindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	return
}

func (inst *UnstakeDevice) MustFindProviderStakeInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderStakeInfoAddress finds ProviderStakeInfo account address with given seeds.
func (inst *UnstakeDevice) FindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderStakeInfoAddress(provider, 0)
	return
}

func (inst *UnstakeDevice) MustFindProviderStakeInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderStakeInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderStakeInfoAccount gets the "provider_stake_info" account.
func (inst *UnstakeDevice) GetProviderStakeInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetProviderVestingInfoAccount sets the "provider_vesting_info" account.
func (inst *UnstakeDevice) SetProviderVestingInfoAccount(providerVestingInfo ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(providerVestingInfo).WRITE()
	return inst
}

func (inst *UnstakeDevice) findFindProviderVestingInfoAddress(provider ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *UnstakeDevice) FindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	return
}

func (inst *UnstakeDevice) MustFindProviderVestingInfoAddressWithBumpSeed(provider ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindProviderVestingInfoAddress finds ProviderVestingInfo account address with given seeds.
func (inst *UnstakeDevice) FindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindProviderVestingInfoAddress(provider, 0)
	return
}

func (inst *UnstakeDevice) MustFindProviderVestingInfoAddress(provider ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindProviderVestingInfoAddress(provider, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetProviderVestingInfoAccount gets the "provider_vesting_info" account.
func (inst *UnstakeDevice) GetProviderVestingInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetProviderAccount sets the "provider" account.
func (inst *UnstakeDevice) SetProviderAccount(provider ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(provider)
	return inst
}

// GetProviderAccount gets the "provider" account.
func (inst *UnstakeDevice) GetProviderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetControllerAccount sets the "controller" account.
func (inst *UnstakeDevice) SetControllerAccount(controller ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(controller).WRITE().SIGNER()
	return inst
}

// GetControllerAccount gets the "controller" account.
func (inst *UnstakeDevice) GetControllerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetAdminAccount sets the "admin" account.
func (inst *UnstakeDevice) SetAdminAccount(admin ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *UnstakeDevice) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *UnstakeDevice) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *UnstakeDevice) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *UnstakeDevice) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *UnstakeDevice) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *UnstakeDevice) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *UnstakeDevice {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *UnstakeDevice) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

func (inst UnstakeDevice) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_UnstakeDevice,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst UnstakeDevice) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *UnstakeDevice) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.DeviceId == nil {
			return errors.New("DeviceId parameter is not set")
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
			return errors.New("accounts.SupernodeVestingAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.ProviderStakeInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.ProviderVestingInfo is not set")
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

func (inst *UnstakeDevice) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("UnstakeDevice")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param(" DeviceId", *inst.DeviceId))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=11]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        supernode_stake_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      supernode_vesting_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("     provider_stake_info", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("   provider_vesting_info", inst.AccountMetaSlice.Get(4)))
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

func (obj UnstakeDevice) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `DeviceId` param:
	err = encoder.Encode(obj.DeviceId)
	if err != nil {
		return err
	}
	return nil
}
func (obj *UnstakeDevice) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `DeviceId`:
	err = decoder.Decode(&obj.DeviceId)
	if err != nil {
		return err
	}
	return nil
}

// NewUnstakeDeviceInstruction declares a new UnstakeDevice instruction with the provided parameters and accounts.
func NewUnstakeDeviceInstruction(
	// Parameters:
	device_id uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeStakeAccount ag_solanago.PublicKey,
	supernodeVestingAccount ag_solanago.PublicKey,
	providerStakeInfo ag_solanago.PublicKey,
	providerVestingInfo ag_solanago.PublicKey,
	provider ag_solanago.PublicKey,
	controller ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *UnstakeDevice {
	return NewUnstakeDeviceInstructionBuilder().
		SetDeviceId(device_id).
		SetSupernodeAccount(supernode).
		SetSupernodeStakeAccountAccount(supernodeStakeAccount).
		SetSupernodeVestingAccountAccount(supernodeVestingAccount).
		SetProviderStakeInfoAccount(providerStakeInfo).
		SetProviderVestingInfoAccount(providerVestingInfo).
		SetProviderAccount(provider).
		SetControllerAccount(controller).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
