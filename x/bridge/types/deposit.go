package types

import "fmt"

const (
	DepositStatusNil = DepositStatus_DEPOSIT_STATUS_UNSPECIFIED
	CoinTypeNil      = CoinType_COIN_TYPE_UNSPECIFIED
)

// DepositStatusFromString turns a string into a DepositStatus
func DepositStatusFromString(str string) (DepositStatus, error) {
	num, ok := DepositStatus_value[str]
	if !ok {
		return DepositStatusNil, fmt.Errorf("'%s' is not a valid Deposit status", str)
	}
	return DepositStatus(num), nil
}

// CoinTypeFromString turns a string into a CoinType
func CoinTypeFromString(str string) (CoinType, error) {
	num, ok := CoinType_value[str]
	if !ok {
		return CoinTypeNil, fmt.Errorf("'%s' is not a valid coin type", str)
	}
	return CoinType(num), nil
}
