// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package client

import (
	"fmt"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
)

type ProviderStakeInfoAccount struct {
	ExtraControllers [2]ag_solanago.PublicKey
	Devices          []DeviceState
}

var ProviderStakeInfoAccountDiscriminator = [8]byte{200, 104, 62, 31, 28, 110, 31, 45}

func (obj ProviderStakeInfoAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ProviderStakeInfoAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
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

func (obj *ProviderStakeInfoAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ProviderStakeInfoAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[200 104 62 31 28 110 31 45]",
				fmt.Sprint(discriminator[:]))
		}
	}
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

type ProviderVestingInfoAccount struct {
	EndIdx         uint8
	LastReleaseDay uint16
	ReleasedAmount uint64
	Schedules      []Schedule
}

var ProviderVestingInfoAccountDiscriminator = [8]byte{138, 109, 208, 248, 196, 12, 164, 112}

func (obj ProviderVestingInfoAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(ProviderVestingInfoAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
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

func (obj *ProviderVestingInfoAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(ProviderVestingInfoAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[138 109 208 248 196 12 164 112]",
				fmt.Sprint(discriminator[:]))
		}
	}
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

type SupernodeStateAccount struct {
	Admin  ag_solanago.PublicKey
	Token  ag_solanago.PublicKey
	Policy Policy
}

var SupernodeStateAccountDiscriminator = [8]byte{2, 22, 141, 62, 77, 123, 126, 67}

func (obj SupernodeStateAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(SupernodeStateAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
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

func (obj *SupernodeStateAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(SupernodeStateAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[2 22 141 62 77 123 126 67]",
				fmt.Sprint(discriminator[:]))
		}
	}
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

type TenantInfoAccount struct {
	Funds     uint64
	Withdrawn uint64
}

var TenantInfoAccountDiscriminator = [8]byte{239, 62, 8, 238, 217, 205, 200, 193}

func (obj TenantInfoAccount) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Write account discriminator:
	err = encoder.WriteBytes(TenantInfoAccountDiscriminator[:], false)
	if err != nil {
		return err
	}
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

func (obj *TenantInfoAccount) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Read and check account discriminator:
	{
		discriminator, err := decoder.ReadTypeID()
		if err != nil {
			return err
		}
		if !discriminator.Equal(TenantInfoAccountDiscriminator[:]) {
			return fmt.Errorf(
				"wrong discriminator: wanted %s, got %s",
				"[239 62 8 238 217 205 200 193]",
				fmt.Sprint(discriminator[:]))
		}
	}
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
