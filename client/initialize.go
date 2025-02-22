package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Initialize is the `initialize` instruction.
type Initialize struct {
	RewardLockedTime   *uint64
	StakingCoefficient *uint64

	// [0] = [WRITE] supernode
	//
	// [1] = [WRITE] supernode_stake_account
	//
	// [2] = [WRITE] supernode_vesting_account
	//
	// [3] = [] token
	//
	// [4] = [WRITE] supernode_rental_account
	//
	// [5] = [WRITE, SIGNER] admin
	//
	// [6] = [] system_program
	//
	// [7] = [] token_program
	//
	// [8] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewInitializeInstructionBuilder creates a new `Initialize` instruction builder.
func NewInitializeInstructionBuilder() *Initialize {
	nd := &Initialize{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	nd.AccountMetaSlice[6] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[7] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[8] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetRewardLockedTime sets the "reward_locked_time" parameter.
func (inst *Initialize) SetRewardLockedTime(reward_locked_time uint64) *Initialize {
	inst.RewardLockedTime = &reward_locked_time
	return inst
}

// SetStakingCoefficient sets the "staking_coefficient" parameter.
func (inst *Initialize) SetStakingCoefficient(staking_coefficient uint64) *Initialize {
	inst.StakingCoefficient = &staking_coefficient
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *Initialize) SetSupernodeAccount(supernode ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *Initialize) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Initialize) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *Initialize) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *Initialize) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *Initialize) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *Initialize) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeStakeAccountAccount sets the "supernode_stake_account" account.
func (inst *Initialize) SetSupernodeStakeAccountAccount(supernodeStakeAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeStakeAccount).WRITE()
	return inst
}

func (inst *Initialize) findFindSupernodeStakeAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Initialize) FindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	return
}

func (inst *Initialize) MustFindSupernodeStakeAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeStakeAccountAddress finds SupernodeStakeAccount account address with given seeds.
func (inst *Initialize) FindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeStakeAccountAddress(0)
	return
}

func (inst *Initialize) MustFindSupernodeStakeAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeStakeAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeStakeAccountAccount gets the "supernode_stake_account" account.
func (inst *Initialize) GetSupernodeStakeAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetSupernodeVestingAccountAccount sets the "supernode_vesting_account" account.
func (inst *Initialize) SetSupernodeVestingAccountAccount(supernodeVestingAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(supernodeVestingAccount).WRITE()
	return inst
}

func (inst *Initialize) findFindSupernodeVestingAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *Initialize) FindSupernodeVestingAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeVestingAccountAddress(bumpSeed)
	return
}

func (inst *Initialize) MustFindSupernodeVestingAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeVestingAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeVestingAccountAddress finds SupernodeVestingAccount account address with given seeds.
func (inst *Initialize) FindSupernodeVestingAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeVestingAccountAddress(0)
	return
}

func (inst *Initialize) MustFindSupernodeVestingAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeVestingAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeVestingAccountAccount gets the "supernode_vesting_account" account.
func (inst *Initialize) GetSupernodeVestingAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTokenAccount sets the "token" account.
func (inst *Initialize) SetTokenAccount(token ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *Initialize) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetSupernodeRentalAccountAccount sets the "supernode_rental_account" account.
func (inst *Initialize) SetSupernodeRentalAccountAccount(supernodeRentalAccount ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(supernodeRentalAccount).WRITE()
	return inst
}

func (inst *Initialize) findFindSupernodeRentalAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: supernode_rental_account
	seeds = append(seeds, []byte{byte(0x73), byte(0x75), byte(0x70), byte(0x65), byte(0x72), byte(0x6e), byte(0x6f), byte(0x64), byte(0x65), byte(0x5f), byte(0x72), byte(0x65), byte(0x6e), byte(0x74), byte(0x61), byte(0x6c), byte(0x5f), byte(0x61), byte(0x63), byte(0x63), byte(0x6f), byte(0x75), byte(0x6e), byte(0x74)})

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindSupernodeRentalAccountAddressWithBumpSeed calculates SupernodeRentalAccount account address with given seeds and a known bump seed.
func (inst *Initialize) FindSupernodeRentalAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeRentalAccountAddress(bumpSeed)
	return
}

func (inst *Initialize) MustFindSupernodeRentalAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRentalAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeRentalAccountAddress finds SupernodeRentalAccount account address with given seeds.
func (inst *Initialize) FindSupernodeRentalAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeRentalAccountAddress(0)
	return
}

func (inst *Initialize) MustFindSupernodeRentalAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRentalAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeRentalAccountAccount gets the "supernode_rental_account" account.
func (inst *Initialize) GetSupernodeRentalAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetAdminAccount sets the "admin" account.
func (inst *Initialize) SetAdminAccount(admin ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(admin).WRITE().SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *Initialize) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *Initialize) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *Initialize) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *Initialize) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *Initialize) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *Initialize) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *Initialize {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *Initialize) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

func (inst Initialize) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_Initialize,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst Initialize) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *Initialize) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.RewardLockedTime == nil {
			return errors.New("RewardLockedTime parameter is not set")
		}
		if inst.StakingCoefficient == nil {
			return errors.New("StakingCoefficient parameter is not set")
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
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SupernodeRentalAccount is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *Initialize) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("Initialize")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=2]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("   RewardLockedTime", *inst.RewardLockedTime))
						paramsBranch.Child(ag_format.Param(" StakingCoefficient", *inst.StakingCoefficient))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("        supernode_stake_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("      supernode_vesting_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                   token", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("       supernode_rental_", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(8)))
					})
				})
		})
}

func (obj Initialize) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `RewardLockedTime` param:
	err = encoder.Encode(obj.RewardLockedTime)
	if err != nil {
		return err
	}
	// Serialize `StakingCoefficient` param:
	err = encoder.Encode(obj.StakingCoefficient)
	if err != nil {
		return err
	}
	return nil
}
func (obj *Initialize) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `RewardLockedTime`:
	err = decoder.Decode(&obj.RewardLockedTime)
	if err != nil {
		return err
	}
	// Deserialize `StakingCoefficient`:
	err = decoder.Decode(&obj.StakingCoefficient)
	if err != nil {
		return err
	}
	return nil
}

// NewInitializeInstruction declares a new Initialize instruction with the provided parameters and accounts.
func NewInitializeInstruction(
	// Parameters:
	reward_locked_time uint64,
	staking_coefficient uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeStakeAccount ag_solanago.PublicKey,
	supernodeVestingAccount ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	supernodeRentalAccount ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *Initialize {
	return NewInitializeInstructionBuilder().
		SetRewardLockedTime(reward_locked_time).
		SetStakingCoefficient(staking_coefficient).
		SetSupernodeAccount(supernode).
		SetSupernodeStakeAccountAccount(supernodeStakeAccount).
		SetSupernodeVestingAccountAccount(supernodeVestingAccount).
		SetTokenAccount(token).
		SetSupernodeRentalAccountAccount(supernodeRentalAccount).
		SetAdminAccount(admin).
		SetSystemProgramAccount(systemProgram).
		SetTokenProgramAccount(tokenProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
