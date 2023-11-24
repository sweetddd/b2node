package bitcoin

import (
	"encoding/json"

	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
)

// Descriptor @see https://developer.bitcoin.org/reference/rpc/importdescriptors.html
type Descriptor struct {
	Desc      string      `json:"desc"`
	Active    *bool       `json:"active,omitempty"`
	Range     interface{} `json:"range,omitempty"`
	NextIndex *int        `json:"next_index,omitempty"`
	Timestamp interface{} `json:"timestamp"`
	Internal  *bool       `json:"internal,omitempty"`
	Label     *string     `json:"label,omitempty"`
}

// ImportDescriptorsCmd @see https://developer.bitcoin.org/reference/rpc/importdescriptors.html
type ImportDescriptorsCmd struct {
	Descriptors []Descriptor `json:""`
}

// NewImportDescriptorsCmd creates a new instance of the importdescriptors command.
func NewImportDescriptorsCmd(descriptors []Descriptor) *ImportDescriptorsCmd {
	return &ImportDescriptorsCmd{
		Descriptors: descriptors,
	}
}

// ImportDescriptorsResult @see ImportDescriptorsResult
type ImportDescriptorsResultElement struct {
	Success  bool              `json:"success"`
	Warnings []string          `json:"warnings,omitempty"`
	Error    *btcjson.RPCError `json:"error,omitempty"`
}

type ImportDescriptorsResult []ImportDescriptorsResultElement

type FutureImportDescriptorsResult chan *rpcclient.Response

// Receive waits for the response promised by the future and returns the result of the importdescriptors command.
func (r FutureImportDescriptorsResult) Receive() (*ImportDescriptorsResult, error) {
	res, err := rpcclient.ReceiveFuture(r)
	if err != nil {
		return nil, err
	}

	var importDescriptors ImportDescriptorsResult
	err = json.Unmarshal(res, &importDescriptors)
	if err != nil {
		return nil, err
	}

	return &importDescriptors, nil
}

// ImportDescriptorsAsync @see ImportDescriptorsAsync
func ImportDescriptorsAsync(c *rpcclient.Client, descriptors []Descriptor) FutureImportDescriptorsResult {
	cmd := &ImportDescriptorsCmd{
		Descriptors: descriptors,
	}
	return c.SendCmd(cmd)
}

// ImportDescriptors @see https://developer.bitcoin.org/reference/rpc/importdescriptors.html
func ImportDescriptors(c *rpcclient.Client, descriptors []Descriptor) (*ImportDescriptorsResult, error) {
	return ImportDescriptorsAsync(c, descriptors).Receive()
}

// register
func init() {
	flags := btcjson.UsageFlag(0)
	btcjson.MustRegisterCmd("importdescriptors", (*ImportDescriptorsCmd)(nil), flags)
}
