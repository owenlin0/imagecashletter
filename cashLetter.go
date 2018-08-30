// Copyright 2018 The x9 Authors
// Use of this source code is governed by an Apache License
// license that can be found in the LICENSE file.

package x9

// CashLetter contains CashLetterHeader, CashLetterControl and Bundle records.
type CashLetter struct {
	// ID is a client defined string used as a reference to this record.
	ID string `json:"id"`
	// CashLetterHeader is an ICL Cash Letter Header Record
	CashLetterHeader *CashLetterHeader `json:"cashLetterHeader,omitempty"`
	// Bundle is an array of ICL Bundle
	Bundles []Bundle `json:"bundles,omitempty"`
	// ReturnBundles is an array of ICL Return Bundle
	//ReturnBundles      []ReturnBundle      `json:"returnBundle,omitempty"`
	// CashLetterControl is an ICL Cash Letter Control Record
	CashLetterControl CashLetterControl `json:"cashLetterControl,omitempty"`
	// Converters is composed for x9 to GoLang Converters
	converters
}

// NewCashLetter takes a CashLetterHeader and returns a CashLetter
// ToDo:  Follow up on returning a pointer when implementing tests and examples
func NewCashLetter(clh *CashLetterHeader) *CashLetter {
	cl := &CashLetter{}
	cl.SetControl(NewCashLetterControl())
	cl.SetHeader(clh)
	return cl
}

// Validate performs X9 validations and format rule checks and returns an error if not Validated
func (cl *CashLetter) Validate() error {
	return nil
}

// SetHeader appends an CashLetterHeader to the CashLetter
func (cl *CashLetter) SetHeader(cashLetterHeader *CashLetterHeader) {
	cl.CashLetterHeader = cashLetterHeader
}

// SetControl appends an CashLetterControl to the CashLetter
func (cl *CashLetter) SetControl(cashLetterControl CashLetterControl) {
	cl.CashLetterControl = cashLetterControl
}

// AddBundle appends a Bundle to the x9.File.CashLetter
func (cl *CashLetter) AddBundle(bundle Bundle) []Bundle {
	cl.Bundles = append(cl.Bundles, bundle)
	return cl.Bundles
}
