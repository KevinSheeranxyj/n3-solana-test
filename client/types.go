package client

import (
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type ClaimRentalFeeEvent struct {
	Provider   ag_solanago.PublicKey
	Controller ag_solanago.PublicKey
	Amount     uint64
}

func (obj ClaimRentalFeeEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `Controller` param:
	err = encoder.Encode(obj.Controller)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ClaimRentalFeeEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `Controller`:
	err = decoder.Decode(&obj.Controller)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type DeviceKValueUpdated struct {
	SpecId uint16
	Old    uint64
	New    uint64
	Admin  ag_solanago.PublicKey
}

func (obj DeviceKValueUpdated) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `SpecId` param:
	err = encoder.Encode(obj.SpecId)
	if err != nil {
		return err
	}
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DeviceKValueUpdated) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `SpecId`:
	err = decoder.Decode(&obj.SpecId)
	if err != nil {
		return err
	}
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

type DeviceStakedEvent struct {
	Provider ag_solanago.PublicKey
	DeviceId uint64
	SpecId   uint64
	Amount   uint64
}

func (obj DeviceStakedEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
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
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DeviceStakedEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
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
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type DeviceState struct {
	State              uint16
	SpecId             uint16
	StakingCoefficient uint64
	Kvalue             uint64
}

func (obj DeviceState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `State` param:
	err = encoder.Encode(obj.State)
	if err != nil {
		return err
	}
	// Serialize `SpecId` param:
	err = encoder.Encode(obj.SpecId)
	if err != nil {
		return err
	}
	// Serialize `StakingCoefficient` param:
	err = encoder.Encode(obj.StakingCoefficient)
	if err != nil {
		return err
	}
	// Serialize `Kvalue` param:
	err = encoder.Encode(obj.Kvalue)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DeviceState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `State`:
	err = decoder.Decode(&obj.State)
	if err != nil {
		return err
	}
	// Deserialize `SpecId`:
	err = decoder.Decode(&obj.SpecId)
	if err != nil {
		return err
	}
	// Deserialize `StakingCoefficient`:
	err = decoder.Decode(&obj.StakingCoefficient)
	if err != nil {
		return err
	}
	// Deserialize `Kvalue`:
	err = decoder.Decode(&obj.Kvalue)
	if err != nil {
		return err
	}
	return nil
}

type DeviceUnstakeEvent struct {
	Provider               ag_solanago.PublicKey
	DeviceId               uint64
	Amount                 uint64
	ProviderVestingInfoKey ag_solanago.PublicKey
}

func (obj DeviceUnstakeEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `DeviceId` param:
	err = encoder.Encode(obj.DeviceId)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	// Serialize `ProviderVestingInfoKey` param:
	err = encoder.Encode(obj.ProviderVestingInfoKey)
	if err != nil {
		return err
	}
	return nil
}

func (obj *DeviceUnstakeEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `DeviceId`:
	err = decoder.Decode(&obj.DeviceId)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	// Deserialize `ProviderVestingInfoKey`:
	err = decoder.Decode(&obj.ProviderVestingInfoKey)
	if err != nil {
		return err
	}
	return nil
}

type PayRentalEvent struct {
	Tenant ag_solanago.PublicKey
	Amount uint64
}

func (obj PayRentalEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Tenant` param:
	err = encoder.Encode(obj.Tenant)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *PayRentalEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Tenant`:
	err = decoder.Decode(&obj.Tenant)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type Policy struct {
	Decimals           uint8
	RewardLockedTime   uint64
	StakingCoefficient uint64
	KValues            []uint64
}

func (obj Policy) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Decimals` param:
	err = encoder.Encode(obj.Decimals)
	if err != nil {
		return err
	}
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
	// Serialize `KValues` param:
	err = encoder.Encode(obj.KValues)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Policy) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Decimals`:
	err = decoder.Decode(&obj.Decimals)
	if err != nil {
		return err
	}
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
	// Deserialize `KValues`:
	err = decoder.Decode(&obj.KValues)
	if err != nil {
		return err
	}
	return nil
}

type ProviderControllerChangedEvent struct {
	Provider      ag_solanago.PublicKey
	Action        string
	NewController ag_solanago.PublicKey
	Operator      ag_solanago.PublicKey
	OldController ag_solanago.PublicKey
}

func (obj ProviderControllerChangedEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `Action` param:
	err = encoder.Encode(obj.Action)
	if err != nil {
		return err
	}
	// Serialize `NewController` param:
	err = encoder.Encode(obj.NewController)
	if err != nil {
		return err
	}
	// Serialize `Operator` param:
	err = encoder.Encode(obj.Operator)
	if err != nil {
		return err
	}
	// Serialize `OldController` param:
	err = encoder.Encode(obj.OldController)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProviderControllerChangedEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `Action`:
	err = decoder.Decode(&obj.Action)
	if err != nil {
		return err
	}
	// Deserialize `NewController`:
	err = decoder.Decode(&obj.NewController)
	if err != nil {
		return err
	}
	// Deserialize `Operator`:
	err = decoder.Decode(&obj.Operator)
	if err != nil {
		return err
	}
	// Deserialize `OldController`:
	err = decoder.Decode(&obj.OldController)
	if err != nil {
		return err
	}
	return nil
}

type ProviderStakeInfo struct {
	ExtraControllers [2]ag_solanago.PublicKey
	Devices          []DeviceState
}

func (obj ProviderStakeInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `ExtraControllers` param:
	err = encoder.Encode(obj.ExtraControllers)
	if err != nil {
		return err
	}
	// Serialize `Devices` param:
	err = encoder.Encode(obj.Devices)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProviderStakeInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `ExtraControllers`:
	err = decoder.Decode(&obj.ExtraControllers)
	if err != nil {
		return err
	}
	// Deserialize `Devices`:
	err = decoder.Decode(&obj.Devices)
	if err != nil {
		return err
	}
	return nil
}

type ProviderVestingInfo struct {
	EndIdx         uint8
	LastReleaseDay uint16
	ReleasedAmount uint64
	Schedules      []Schedule
}

func (obj ProviderVestingInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `EndIdx` param:
	err = encoder.Encode(obj.EndIdx)
	if err != nil {
		return err
	}
	// Serialize `LastReleaseDay` param:
	err = encoder.Encode(obj.LastReleaseDay)
	if err != nil {
		return err
	}
	// Serialize `ReleasedAmount` param:
	err = encoder.Encode(obj.ReleasedAmount)
	if err != nil {
		return err
	}
	// Serialize `Schedules` param:
	err = encoder.Encode(obj.Schedules)
	if err != nil {
		return err
	}
	return nil
}

func (obj *ProviderVestingInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `EndIdx`:
	err = decoder.Decode(&obj.EndIdx)
	if err != nil {
		return err
	}
	// Deserialize `LastReleaseDay`:
	err = decoder.Decode(&obj.LastReleaseDay)
	if err != nil {
		return err
	}
	// Deserialize `ReleasedAmount`:
	err = decoder.Decode(&obj.ReleasedAmount)
	if err != nil {
		return err
	}
	// Deserialize `Schedules`:
	err = decoder.Decode(&obj.Schedules)
	if err != nil {
		return err
	}
	return nil
}

type RewardClaimedEvent struct {
	Provider ag_solanago.PublicKey
	Amount   uint64
}

func (obj RewardClaimedEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RewardClaimedEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type RewardLockedTimeUpdated struct {
	Old   uint64
	New   uint64
	Admin ag_solanago.PublicKey
}

func (obj RewardLockedTimeUpdated) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

func (obj *RewardLockedTimeUpdated) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

type Schedule struct {
	Day    uint16
	Amount uint64
}

func (obj Schedule) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Day` param:
	err = encoder.Encode(obj.Day)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *Schedule) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Day`:
	err = decoder.Decode(&obj.Day)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type StakingCoefficientUpdated struct {
	Old   uint64
	New   uint64
	Admin ag_solanago.PublicKey
}

func (obj StakingCoefficientUpdated) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Old` param:
	err = encoder.Encode(obj.Old)
	if err != nil {
		return err
	}
	// Serialize `New` param:
	err = encoder.Encode(obj.New)
	if err != nil {
		return err
	}
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

func (obj *StakingCoefficientUpdated) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Old`:
	err = decoder.Decode(&obj.Old)
	if err != nil {
		return err
	}
	// Deserialize `New`:
	err = decoder.Decode(&obj.New)
	if err != nil {
		return err
	}
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	return nil
}

type SupernodeState struct {
	Admin  ag_solanago.PublicKey
	Token  ag_solanago.PublicKey
	Policy Policy
}

func (obj SupernodeState) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Admin` param:
	err = encoder.Encode(obj.Admin)
	if err != nil {
		return err
	}
	// Serialize `Token` param:
	err = encoder.Encode(obj.Token)
	if err != nil {
		return err
	}
	// Serialize `Policy` param:
	err = encoder.Encode(obj.Policy)
	if err != nil {
		return err
	}
	return nil
}

func (obj *SupernodeState) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Admin`:
	err = decoder.Decode(&obj.Admin)
	if err != nil {
		return err
	}
	// Deserialize `Token`:
	err = decoder.Decode(&obj.Token)
	if err != nil {
		return err
	}
	// Deserialize `Policy`:
	err = decoder.Decode(&obj.Policy)
	if err != nil {
		return err
	}
	return nil
}

type TenantInfo struct {
	Funds     uint64
	Withdrawn uint64
}

func (obj TenantInfo) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Funds` param:
	err = encoder.Encode(obj.Funds)
	if err != nil {
		return err
	}
	// Serialize `Withdrawn` param:
	err = encoder.Encode(obj.Withdrawn)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TenantInfo) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Funds`:
	err = decoder.Decode(&obj.Funds)
	if err != nil {
		return err
	}
	// Deserialize `Withdrawn`:
	err = decoder.Decode(&obj.Withdrawn)
	if err != nil {
		return err
	}
	return nil
}

type TokenReleasedEvent struct {
	Controller ag_solanago.PublicKey
	Provider   ag_solanago.PublicKey
	Amount     uint64
}

func (obj TokenReleasedEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Controller` param:
	err = encoder.Encode(obj.Controller)
	if err != nil {
		return err
	}
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *TokenReleasedEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Controller`:
	err = decoder.Decode(&obj.Controller)
	if err != nil {
		return err
	}
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type VestingScheduledEvent struct {
	Provider ag_solanago.PublicKey
	Day      uint16
	Amount   uint64
}

func (obj VestingScheduledEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Provider` param:
	err = encoder.Encode(obj.Provider)
	if err != nil {
		return err
	}
	// Serialize `Day` param:
	err = encoder.Encode(obj.Day)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *VestingScheduledEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Provider`:
	err = decoder.Decode(&obj.Provider)
	if err != nil {
		return err
	}
	// Deserialize `Day`:
	err = decoder.Decode(&obj.Day)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

type WithdrawEvent struct {
	Tenant ag_solanago.PublicKey
	Amount uint64
}

func (obj WithdrawEvent) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Tenant` param:
	err = encoder.Encode(obj.Tenant)
	if err != nil {
		return err
	}
	// Serialize `Amount` param:
	err = encoder.Encode(obj.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (obj *WithdrawEvent) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Tenant`:
	err = decoder.Decode(&obj.Tenant)
	if err != nil {
		return err
	}
	// Deserialize `Amount`:
	err = decoder.Decode(&obj.Amount)
	if err != nil {
		return err
	}
	return nil
}
