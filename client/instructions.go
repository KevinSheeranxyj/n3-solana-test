package client

import (
	"bytes"
	"fmt"
	ag_spew "github.com/davecgh/go-spew/spew"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_text "github.com/gagliardetto/solana-go/text"
	ag_treeout "github.com/gagliardetto/treeout"
)

var ProgramID ag_solanago.PublicKey

func SetProgramID(PublicKey ag_solanago.PublicKey) {
	ProgramID = PublicKey
	ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
}

const ProgramName = "Dummy"

func init() {
	if !ProgramID.IsZero() {
		ag_solanago.RegisterInstructionDecoder(ProgramID, registryDecodeInstruction)
	}
}

var (
	Instruction_AddExtraController = ag_binary.TypeID([8]byte{212, 217, 0, 188, 6, 34, 34, 97})

	Instruction_ClaimRentalFee = ag_binary.TypeID([8]byte{182, 236, 192, 130, 204, 223, 128, 117})

	Instruction_ClaimReward = ag_binary.TypeID([8]byte{149, 95, 181, 242, 94, 90, 158, 162})

	Instruction_InitRewardAccount = ag_binary.TypeID([8]byte{206, 38, 8, 106, 162, 1, 134, 90})

	Instruction_Initialize = ag_binary.TypeID([8]byte{175, 175, 109, 31, 13, 152, 155, 237})

	Instruction_PayRentalFee = ag_binary.TypeID([8]byte{45, 10, 39, 144, 246, 2, 157, 94})

	Instruction_Releasable = ag_binary.TypeID([8]byte{115, 143, 23, 169, 117, 191, 173, 152})

	Instruction_Release = ag_binary.TypeID([8]byte{253, 249, 15, 206, 28, 127, 193, 241})

	Instruction_RemoveExtraController = ag_binary.TypeID([8]byte{180, 173, 240, 243, 97, 70, 236, 52})

	Instruction_ReplaceExtraController = ag_binary.TypeID([8]byte{96, 25, 235, 129, 158, 174, 203, 179})

	Instruction_StakeDevice = ag_binary.TypeID([8]byte{92, 226, 123, 60, 82, 196, 27, 73})

	Instruction_UnstakeDevice = ag_binary.TypeID([8]byte{86, 80, 60, 38, 72, 185, 250, 61})

	Instruction_UpdateKValue = ag_binary.TypeID([8]byte{213, 26, 29, 112, 239, 146, 241, 16})

	Instruction_UpdateRewardLockTime = ag_binary.TypeID([8]byte{143, 163, 14, 177, 108, 6, 205, 81})

	Instruction_UpdateStakingCoefficient = ag_binary.TypeID([8]byte{0, 178, 197, 201, 209, 217, 39, 245})

	Instruction_WithdrawRentalFee = ag_binary.TypeID([8]byte{185, 67, 135, 4, 155, 6, 77, 156})
)

// InstructionIDToName returns the name of the instruction given its ID.
func InstructionIDToName(id ag_binary.TypeID) string {
	switch id {
	case Instruction_AddExtraController:
		return "AddExtraController"
	case Instruction_ClaimRentalFee:
		return "ClaimRentalFee"
	case Instruction_ClaimReward:
		return "ClaimReward"
	case Instruction_InitRewardAccount:
		return "InitRewardAccount"
	case Instruction_Initialize:
		return "Initialize"
	case Instruction_PayRentalFee:
		return "PayRentalFee"
	case Instruction_Releasable:
		return "Releasable"
	case Instruction_Release:
		return "Release"
	case Instruction_RemoveExtraController:
		return "RemoveExtraController"
	case Instruction_ReplaceExtraController:
		return "ReplaceExtraController"
	case Instruction_StakeDevice:
		return "StakeDevice"
	case Instruction_UnstakeDevice:
		return "UnstakeDevice"
	case Instruction_UpdateKValue:
		return "UpdateKValue"
	case Instruction_UpdateRewardLockTime:
		return "UpdateRewardLockTime"
	case Instruction_UpdateStakingCoefficient:
		return "UpdateStakingCoefficient"
	case Instruction_WithdrawRentalFee:
		return "WithdrawRentalFee"
	default:
		return ""
	}
}

type Instruction struct {
	ag_binary.BaseVariant
}

func (inst *Instruction) EncodeToTree(parent ag_treeout.Branches) {
	if enToTree, ok := inst.Impl.(ag_text.EncodableToTree); ok {
		enToTree.EncodeToTree(parent)
	} else {
		parent.Child(ag_spew.Sdump(inst))
	}
}

var InstructionImplDef = ag_binary.NewVariantDefinition(
	ag_binary.AnchorTypeIDEncoding,
	[]ag_binary.VariantType{
		{
			Name: "add_extra_controller", Type: (*AddExtraController)(nil),
		},
		{
			Name: "claim_rental_fee", Type: (*ClaimRentalFee)(nil),
		},
		{
			Name: "claim_reward", Type: (*ClaimReward)(nil),
		},
		{
			Name: "init_reward_account", Type: (*InitRewardAccount)(nil),
		},
		{
			Name: "initialize", Type: (*Initialize)(nil),
		},
		{
			Name: "pay_rental_fee", Type: (*PayRentalFee)(nil),
		},
		{
			Name: "releasable", Type: (*Releasable)(nil),
		},
		{
			Name: "release", Type: (*Release)(nil),
		},
		{
			Name: "remove_extra_controller", Type: (*RemoveExtraController)(nil),
		},
		{
			Name: "replace_extra_controller", Type: (*ReplaceExtraController)(nil),
		},
		{
			Name: "stake_device", Type: (*StakeDevice)(nil),
		},
		{
			Name: "unstake_device", Type: (*UnstakeDevice)(nil),
		},
		{
			Name: "update_k_value", Type: (*UpdateKValue)(nil),
		},
		{
			Name: "update_reward_lock_time", Type: (*UpdateRewardLockTime)(nil),
		},
		{
			Name: "update_staking_coefficient", Type: (*UpdateStakingCoefficient)(nil),
		},
		{
			Name: "withdraw_rental_fee", Type: (*WithdrawRentalFee)(nil),
		},
	},
)

func (inst *Instruction) ProgramID() ag_solanago.PublicKey {
	return ProgramID
}

func (inst *Instruction) Accounts() (out []*ag_solanago.AccountMeta) {
	return inst.Impl.(ag_solanago.AccountsGettable).GetAccounts()
}

func (inst *Instruction) Data() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := ag_binary.NewBorshEncoder(buf).Encode(inst); err != nil {
		return nil, fmt.Errorf("unable to encode instruction: %w", err)
	}
	return buf.Bytes(), nil
}

func (inst *Instruction) TextEncode(encoder *ag_text.Encoder, option *ag_text.Option) error {
	return encoder.Encode(inst.Impl, option)
}

func (inst *Instruction) UnmarshalWithDecoder(decoder *ag_binary.Decoder) error {
	return inst.BaseVariant.UnmarshalBinaryVariant(decoder, InstructionImplDef)
}

func (inst *Instruction) MarshalWithEncoder(encoder *ag_binary.Encoder) error {
	err := encoder.WriteBytes(inst.TypeID.Bytes(), false)
	if err != nil {
		return fmt.Errorf("unable to write variant type: %w", err)
	}
	return encoder.Encode(inst.Impl)
}

func registryDecodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (interface{}, error) {
	inst, err := decodeInstruction(accounts, data)
	if err != nil {
		return nil, err
	}
	return inst, nil
}

func decodeInstruction(accounts []*ag_solanago.AccountMeta, data []byte) (*Instruction, error) {
	inst := new(Instruction)
	if err := ag_binary.NewBorshDecoder(data).Decode(inst); err != nil {
		return nil, fmt.Errorf("unable to decode instruction: %w", err)
	}
	if v, ok := inst.Impl.(ag_solanago.AccountsSettable); ok {
		err := v.SetAccounts(accounts)
		if err != nil {
			return nil, fmt.Errorf("unable to set accounts for instruction: %w", err)
		}
	}
	return inst, nil
}

func DecodeInstructions(message *ag_solanago.Message) (instructions []*Instruction, err error) {
	for _, ins := range message.Instructions {
		var programID ag_solanago.PublicKey
		if programID, err = message.Program(ins.ProgramIDIndex); err != nil {
			return
		}
		if !programID.Equals(ProgramID) {
			continue
		}
		var accounts []*ag_solanago.AccountMeta
		if accounts, err = ins.ResolveInstructionAccounts(message); err != nil {
			return
		}
		var insDecoded *Instruction
		if insDecoded, err = decodeInstruction(accounts, ins.Data); err != nil {
			return
		}
		instructions = append(instructions, insDecoded)
	}
	return
}
