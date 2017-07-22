//nolint
package coin

import (
	"fmt"

	abci "github.com/tendermint/abci/types"

	"github.com/tendermint/basecoin/errors"
)

var (
	errNoAccount         = fmt.Errorf("No such account")
	errInsufficientFunds = fmt.Errorf("Insufficient funds")
	errNoInputs          = fmt.Errorf("No input coins")
	errNoOutputs         = fmt.Errorf("No output coins")
	errInvalidAddress    = fmt.Errorf("Invalid address")
	errInvalidCoins      = fmt.Errorf("Invalid coins")
	errUnknownKey        = fmt.Errorf("Unknown key")

	invalidInput   = abci.CodeType_BaseInvalidInput
	invalidOutput  = abci.CodeType_BaseInvalidOutput
	unknownAddress = abci.CodeType_BaseUnknownAddress
	unknownRequest = abci.CodeType_UnknownRequest
)

// here are some generic handlers to grab classes of errors based on code
func IsInputErr(err error) bool {
	return errors.HasErrorCode(err, invalidInput)
}
func IsOutputErr(err error) bool {
	return errors.HasErrorCode(err, invalidOutput)
}
func IsAddressErr(err error) bool {
	return errors.HasErrorCode(err, unknownAddress)
}
func IsCoinErr(err error) bool {
	return err != nil && (IsInputErr(err) || IsOutputErr(err) || IsAddressErr(err))
}

func ErrNoAccount() errors.TMError {
	return errors.WithCode(errNoAccount, unknownAddress)
}

func IsNoAccountErr(err error) bool {
	return errors.IsSameError(errNoAccount, err)
}

func ErrInvalidAddress() errors.TMError {
	return errors.WithCode(errInvalidAddress, invalidInput)
}
func IsInvalidAddressErr(err error) bool {
	return errors.IsSameError(errInvalidAddress, err)
}

func ErrInvalidCoins() errors.TMError {
	return errors.WithCode(errInvalidCoins, invalidInput)
}
func IsInvalidCoinsErr(err error) bool {
	return errors.IsSameError(errInvalidCoins, err)
}

func ErrInsufficientFunds() errors.TMError {
	return errors.WithCode(errInsufficientFunds, invalidInput)
}
func IsInsufficientFundsErr(err error) bool {
	return errors.IsSameError(errInsufficientFunds, err)
}

func ErrNoInputs() errors.TMError {
	return errors.WithCode(errNoInputs, invalidInput)
}
func IsNoInputsErr(err error) bool {
	return errors.IsSameError(errNoInputs, err)
}

func ErrNoOutputs() errors.TMError {
	return errors.WithCode(errNoOutputs, invalidOutput)
}
func IsNoOutputsErr(err error) bool {
	return errors.IsSameError(errNoOutputs, err)
}

func ErrUnknownKey(mod string) errors.TMError {
	return errors.WithMessage(mod, errUnknownKey, unknownRequest)
}
func IsUnknownKeyErr(err error) bool {
	return errors.IsSameError(errUnknownKey, err)
}