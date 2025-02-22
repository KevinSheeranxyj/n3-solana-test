package client

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// WithdrawRentalFee is the `withdraw_rental_fee` instruction.
type WithdrawRentalFee struct {
	Amount *uint64

	// [0] = [WRITE] supernode
	//
	// [1] = [WRITE] supernode_rental_account
	//
	// [2] = [WRITE] tenant_token_account
	//
	// [3] = [WRITE] tenant_info
	//
	// [4] = [] token
	//
	// [5] = [WRITE, SIGNER] tenant
	//
	// [6] = [SIGNER] admin
	//
	// [7] = [] token_program
	//
	// [8] = [] system_program
	//
	// [9] = [] associated_token_program
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewWithdrawRentalFeeInstructionBuilder creates a new `WithdrawRentalFee` instruction builder.
func NewWithdrawRentalFeeInstructionBuilder() *WithdrawRentalFee {
	nd := &WithdrawRentalFee{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 10),
	}
	nd.AccountMetaSlice[7] = ag_solanago.Meta(Addresses["TokenkegQfeZyiNwAJbNbGKPFXCWuBvf9Ss623VQ5DA"])
	nd.AccountMetaSlice[8] = ag_solanago.Meta(Addresses["11111111111111111111111111111111"])
	nd.AccountMetaSlice[9] = ag_solanago.Meta(Addresses["ATokenGPvbdGVxr1b2hvZbsiqW5xWH25efTNsLJA8knL"])
	return nd
}

// SetAmount sets the "amount" parameter.
func (inst *WithdrawRentalFee) SetAmount(amount uint64) *WithdrawRentalFee {
	inst.Amount = &amount
	return inst
}

// SetSupernodeAccount sets the "supernode" account.
func (inst *WithdrawRentalFee) SetSupernodeAccount(supernode ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(supernode).WRITE()
	return inst
}

func (inst *WithdrawRentalFee) findFindSupernodeAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *WithdrawRentalFee) FindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeAddress(bumpSeed)
	return
}

func (inst *WithdrawRentalFee) MustFindSupernodeAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeAddress finds Supernode account address with given seeds.
func (inst *WithdrawRentalFee) FindSupernodeAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeAddress(0)
	return
}

func (inst *WithdrawRentalFee) MustFindSupernodeAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeAccount gets the "supernode" account.
func (inst *WithdrawRentalFee) GetSupernodeAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetSupernodeRentalAccountAccount sets the "supernode_rental_account" account.
func (inst *WithdrawRentalFee) SetSupernodeRentalAccountAccount(supernodeRentalAccount ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(supernodeRentalAccount).WRITE()
	return inst
}

func (inst *WithdrawRentalFee) findFindSupernodeRentalAccountAddress(knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
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
func (inst *WithdrawRentalFee) FindSupernodeRentalAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindSupernodeRentalAccountAddress(bumpSeed)
	return
}

func (inst *WithdrawRentalFee) MustFindSupernodeRentalAccountAddressWithBumpSeed(bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRentalAccountAddress(bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindSupernodeRentalAccountAddress finds SupernodeRentalAccount account address with given seeds.
func (inst *WithdrawRentalFee) FindSupernodeRentalAccountAddress() (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindSupernodeRentalAccountAddress(0)
	return
}

func (inst *WithdrawRentalFee) MustFindSupernodeRentalAccountAddress() (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindSupernodeRentalAccountAddress(0)
	if err != nil {
		panic(err)
	}
	return
}

// GetSupernodeRentalAccountAccount gets the "supernode_rental_account" account.
func (inst *WithdrawRentalFee) GetSupernodeRentalAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetTenantTokenAccountAccount sets the "tenant_token_account" account.
func (inst *WithdrawRentalFee) SetTenantTokenAccountAccount(tenantTokenAccount ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(tenantTokenAccount).WRITE()
	return inst
}

// GetTenantTokenAccountAccount gets the "tenant_token_account" account.
func (inst *WithdrawRentalFee) GetTenantTokenAccountAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetTenantInfoAccount sets the "tenant_info" account.
func (inst *WithdrawRentalFee) SetTenantInfoAccount(tenantInfo ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(tenantInfo).WRITE()
	return inst
}

func (inst *WithdrawRentalFee) findFindTenantInfoAddress(tenant ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	var seeds [][]byte
	// const: tenant_info
	seeds = append(seeds, []byte{byte(0x74), byte(0x65), byte(0x6e), byte(0x61), byte(0x6e), byte(0x74), byte(0x5f), byte(0x69), byte(0x6e), byte(0x66), byte(0x6f)})
	// path: tenant
	seeds = append(seeds, tenant.Bytes())

	if knownBumpSeed != 0 {
		seeds = append(seeds, []byte{byte(bumpSeed)})
		pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
	} else {
		pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
	}
	return
}

// FindTenantInfoAddressWithBumpSeed calculates TenantInfo account address with given seeds and a known bump seed.
func (inst *WithdrawRentalFee) FindTenantInfoAddressWithBumpSeed(tenant ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey, err error) {
	pda, _, err = inst.findFindTenantInfoAddress(tenant, bumpSeed)
	return
}

func (inst *WithdrawRentalFee) MustFindTenantInfoAddressWithBumpSeed(tenant ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindTenantInfoAddress(tenant, bumpSeed)
	if err != nil {
		panic(err)
	}
	return
}

// FindTenantInfoAddress finds TenantInfo account address with given seeds.
func (inst *WithdrawRentalFee) FindTenantInfoAddress(tenant ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
	pda, bumpSeed, err = inst.findFindTenantInfoAddress(tenant, 0)
	return
}

func (inst *WithdrawRentalFee) MustFindTenantInfoAddress(tenant ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
	pda, _, err := inst.findFindTenantInfoAddress(tenant, 0)
	if err != nil {
		panic(err)
	}
	return
}

// GetTenantInfoAccount gets the "tenant_info" account.
func (inst *WithdrawRentalFee) GetTenantInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetTokenAccount sets the "token" account.
func (inst *WithdrawRentalFee) SetTokenAccount(token ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(token)
	return inst
}

// GetTokenAccount gets the "token" account.
func (inst *WithdrawRentalFee) GetTokenAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetTenantAccount sets the "tenant" account.
func (inst *WithdrawRentalFee) SetTenantAccount(tenant ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(tenant).WRITE().SIGNER()
	return inst
}

// GetTenantAccount gets the "tenant" account.
func (inst *WithdrawRentalFee) GetTenantAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetAdminAccount sets the "admin" account.
func (inst *WithdrawRentalFee) SetAdminAccount(admin ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(admin).SIGNER()
	return inst
}

// GetAdminAccount gets the "admin" account.
func (inst *WithdrawRentalFee) GetAdminAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetTokenProgramAccount sets the "token_program" account.
func (inst *WithdrawRentalFee) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "token_program" account.
func (inst *WithdrawRentalFee) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetSystemProgramAccount sets the "system_program" account.
func (inst *WithdrawRentalFee) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "system_program" account.
func (inst *WithdrawRentalFee) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetAssociatedTokenProgramAccount sets the "associated_token_program" account.
func (inst *WithdrawRentalFee) SetAssociatedTokenProgramAccount(associatedTokenProgram ag_solanago.PublicKey) *WithdrawRentalFee {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(associatedTokenProgram)
	return inst
}

// GetAssociatedTokenProgramAccount gets the "associated_token_program" account.
func (inst *WithdrawRentalFee) GetAssociatedTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

func (inst WithdrawRentalFee) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: Instruction_WithdrawRentalFee,
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst WithdrawRentalFee) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *WithdrawRentalFee) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Amount == nil {
			return errors.New("Amount parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.Supernode is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SupernodeRentalAccount is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.TenantTokenAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.TenantInfo is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Token is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.Tenant is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.Admin is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.AssociatedTokenProgram is not set")
		}
	}
	return nil
}

func (inst *WithdrawRentalFee) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("WithdrawRentalFee")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Amount", *inst.Amount))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=10]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("               supernode", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("       supernode_rental_", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("           tenant_token_", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("             tenant_info", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                   token", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                  tenant", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("                   admin", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("           token_program", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("          system_program", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("associated_token_program", inst.AccountMetaSlice.Get(9)))
					})
				})
		})
}

func (obj WithdrawRentalFee) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
func (obj *WithdrawRentalFee) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

// NewWithdrawRentalFeeInstruction declares a new WithdrawRentalFee instruction with the provided parameters and accounts.
func NewWithdrawRentalFeeInstruction(
	// Parameters:
	amount uint64,
	// Accounts:
	supernode ag_solanago.PublicKey,
	supernodeRentalAccount ag_solanago.PublicKey,
	tenantTokenAccount ag_solanago.PublicKey,
	tenantInfo ag_solanago.PublicKey,
	token ag_solanago.PublicKey,
	tenant ag_solanago.PublicKey,
	admin ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	associatedTokenProgram ag_solanago.PublicKey) *WithdrawRentalFee {
	return NewWithdrawRentalFeeInstructionBuilder().
		SetAmount(amount).
		SetSupernodeAccount(supernode).
		SetSupernodeRentalAccountAccount(supernodeRentalAccount).
		SetTenantTokenAccountAccount(tenantTokenAccount).
		SetTenantInfoAccount(tenantInfo).
		SetTokenAccount(token).
		SetTenantAccount(tenant).
		SetAdminAccount(admin).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetAssociatedTokenProgramAccount(associatedTokenProgram)
}
