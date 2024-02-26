package types

import "fmt"

const WithdrawStatusNil = WithdrawStatus_WITHDRAW_STATUS_UNSPECIFIED

// WithdrawStatusFromString turns a string into a WithdrawStatus
func WithdrawStatusFromString(str string) (WithdrawStatus, error) {
	num, ok := WithdrawStatus_value[str]
	if !ok {
		return WithdrawStatusNil, fmt.Errorf("'%s' is not a valid withdraw status", str)
	}
	return WithdrawStatus(num), nil
}
