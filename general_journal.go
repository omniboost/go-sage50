package sage50

import (
	"fmt"
)

func NewCSVHeaderLine() []string {
	return []string{
		"BlgNr",
		"Date",
		"AccId",
		"Grp",
		"Orig",
		"MType",
		"Type",
		"CAcc",
		"TaxId",
		"Tidx",
		"Cidx",
		"BType",
		"Code",
		"ValNt",
		"ValTx",
		"ValFW",
		"Text",
		"Text2",
		"PrKey",
		"OpId",
		"Flags",
	}
}

type CSVBookingLine struct {
	BlgNr int
	Date
	// 	AccId : String[13]
	// The Id of a book keeping account. [Fibu-Konto]
	AccId string
	// Grp: String[1]
	// Any character to mark your bookings so your can retreave them later.
	Grp string
	// Orig : Integer
	// The origin of this [Beleg]. The following are defined:
	// 0 Finanzbuchhaltung
	// 1 Debitorenbuchhaltung
	// 2 Kreditorenbuchhaltung
	// 3 Other
	Orig int
	// MType : Integer
	// This tag allows the TAF-reader to abort processing if the mandate it currently is logged into has a different type, than what was specified here. For example if this file contains data for Debitorenbuchhaltung, it cannot be imported by Kreditorenbuchhaltung. The following types are currently defined:
	// 0 or 4 - only the INI file
	// 1 - Finanzbuchhaltung files
	// 2 - Debitorenbuchhaltung files
	// 3 - Kreditorenbuchhaltung files
	MType int
	// Type : Integer
	// The type of the account:
	// 0 Active account [Aktiv]
	// 1 Passive account [Passiv]
	// 2 Expense account [Aufwand]
	// 3 Revenue account [Ertrag]
	// 4 Debit cost centre account [Kst-Soll]
	// 5 Credit cost centre account [Kst-Haben]
	Type int
	// CAcc: String[13]
	// This is the cost type account
	CAcc string
	// TaxId : String[5]
	// The Id of the tax. [MWSt-Kürzel]
	TaxId string
	// TIdx : Integer
	// This is the index of the booking that represents the tax booking which is attached to this booking.
	Tidx int
	// CIdx : Integer
	// This is the index of the booking that represents the cost booking which is
	Cidx int
	// BType : Integer
	// Booking type: 1=cost booking, 2=tax booking
	BType int
	// Code : String[9]
	// A user definable code.
	Code string
	// ValNt : Currency
	// The net amount for this booking. [Netto-Betrag]
	ValNt Decimal
	// ValTx : Currency
	// The tax amount for this booking. [Steuer-Betrag]
	ValTx Decimal
	// ValFW : Currency The gross amount for this booking in the foreign currency specified by currency of the account AccId. [FW-Betrag]
	ValFW Decimal
	// Text : String[31]
	// The first line of the description of the account.
	Text string
	// Text2 : String[31]
	// The second line of the description of the account.
	Text2 string
	// PkKey : Long
	// The PK number of this booking.
	PrKey int
	// OpId : String[13] (Primary Key)
	// OP Nr: The unique identifier for the OP records.
	OpId string
	// Flags : Integer Auxilliary flags.
	// This value consists of the sum of one or more of the following biases:
	// 1 OP is from a creditor.
	// 2 OP has foreign currency.
	Flags int
}

func (l CSVBookingLine) Validate() []error {
	var errs []error
	return errs
}

func (l CSVBookingLine) Values() []interface{} {
	return []interface{}{
		l.BlgNr,
		l.Date,
		l.AccId,
		l.Grp,
		l.Orig,
		l.MType,
		l.Type,
		l.CAcc,
		l.TaxId,
		l.Tidx,
		l.Cidx,
		l.BType,
		l.Code,
		l.ValNt,
		l.ValTx,
		l.ValFW,
		l.Text,
		l.Text2,
		l.PrKey,
		l.OpId,
		l.Flags,
	}
}

func (l CSVBookingLine) StringValues() []string {
	values := l.Values()
	st := make([]string, len(values))
	for k, v := range values {
		st[k] = fmt.Sprint(v)
	}
	return st
}
