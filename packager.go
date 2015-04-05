// Type of packages
package omg8583

type pkg struct {
	Length int
	Name   string
	Type   string
}

var t = map[int]pkg{
	0: {
		Length: 4,
		Name:   "Message Type Indicator",
		Type:   "n",
	},
	1: {
		Length: 16,
		Name:   "Bit map (b 128 if secondary is present and b 192 if tertiary is present)",
		Type:   "b",
	},
	2: {
		Length: 19,
		Name:   "Primary account number (PAN)",
		Type:   "n",
	},
	3: {
		Length: 6,
		Name:   "Processing code",
		Type:   "n",
	},
	4: {
		Length: 12,
		Name:   "Amount, transaction",
		Type:   "n",
	},
	5: {
		Length: 12,
		Name:   "Amount, settlement",
		Type:   "n",
	},
	6: {
		Length: 12,
		Name:   "Amount, cardholder billing",
		Type:   "n",
	},
	7: {
		Length: 10,
		Name:   "Transmission date & time",
		Type:   "n",
	},
	8: {
		Length: 8,
		Name:   "Amount, cardholder billing fee",
		Type:   "n",
	},
	9: {
		Length: 8,
		Name:   "Conversion rate, settlement",
		Type:   "n",
	},
	10: {
		Length: 8,
		Name:   "Conversion rate, cardholder billing",
		Type:   "n",
	},
	11: {
		Length: 6,
		Name:   "System trace audit number",
		Type:   "n",
	},
	12: {
		Length: 6,
		Name:   "Time, local transaction (hhmmss)",
		Type:   "n",
	},
	13: {
		Length: 4,
		Name:   "Date, local transaction (MMDD)",
		Type:   "n",
	},
	14: {
		Length: 4,
		Name:   "Date, expiration",
		Type:   "n",
	},
	15: {
		Length: 4,
		Name:   "Date, settlement",
		Type:   "n",
	},
	16: {
		Length: 4,
		Name:   "Date, conversion",
		Type:   "n",
	},
	17: {
		Length: 4,
		Name:   "Date, capture",
		Type:   "n",
	},
	18: {
		Length: 4,
		Name:   "Merchant type",
		Type:   "n",
	},
	19: {
		Length: 3,
		Name:   "Acquiring institution country code",
		Type:   "n",
	},
	20: {
		Length: 3,
		Name:   "PAN extended, country code",
		Type:   "n",
	},
	21: {
		Length: 3,
		Name:   "Forwarding institution. country code",
		Type:   "n",
	},
	22: {
		Length: 3,
		Name:   "Point of service entry mode",
		Type:   "n",
	},
	23: {
		Length: 3,
		Name:   "Application PAN sequence number",
		Type:   "n",
	},
	24: {
		Length: 3,
		Name:   "Function code (ISO 8583:1993)/Network International identifier (NII)",
		Type:   "n",
	},
	25: {
		Length: 2,
		Name:   "Point of service condition code",
		Type:   "n",
	},
	26: {
		Length: 2,
		Name:   "Point of service capture code",
		Type:   "n",
	},
	27: {
		Length: 1,
		Name:   "Authorizing identification response length",
		Type:   "n",
	},
	28: {
		Length: 8,
		Name:   "Amount, transaction fee",
		Type:   "x+n",
	},
	29: {
		Length: 8,
		Name:   "Amount, settlement fee",
		Type:   "x+n",
	},
	30: {
		Length: 8,
		Name:   "Amount, transaction processing fee",
		Type:   "x+n",
	},
	31: {
		Length: 8,
		Name:   "Amount, settlement processing fee",
		Type:   "x+n",
	},
	32: {
		Length: 11,
		Name:   "Acquiring institution identification code",
		Type:   "n",
	},
	33: {
		Length: 11,
		Name:   "Forwarding institution identification code",
		Type:   "n",
	},
	34: {
		Length: 28,
		Name:   "Primary account number, extended",
		Type:   "ns",
	},
	35: {
		Length: 37,
		Name:   "Track 2 data",
		Type:   "z",
	},
	36: {
		Length: 104,
		Name:   "Track 3 data",
		Type:   "n",
	},
	37: {
		Length: 12,
		Name:   "Retrieval reference number",
		Type:   "an",
	},
	38: {
		Length: 6,
		Name:   "Authorization identification response",
		Type:   "an",
	},
	39: {
		Length: 2,
		Name:   "Response code",
		Type:   "an",
	},
	40: {
		Length: 3,
		Name:   "Service restriction code",
		Type:   "an",
	},
	41: {
		Length: 8,
		Name:   "Card acceptor terminal identification",
		Type:   "ans",
	},
	42: {
		Length: 15,
		Name:   "Card acceptor identification code",
		Type:   "ans",
	},
	43: {
		Length: 40,
		Name:   "Card acceptor name/location (1-23 address 24-36 city 37-38 state 39-40 country)",
		Type:   "ans",
	},
	44: {
		Length: 25, // ..
		Name:   "Additional response data",
		Type:   "an",
	},
	45: {
		Length: 76, // ..
		Name:   "Track 1 data",
		Type:   "an",
	},
	46: {
		Length: 999, // ..
		Name:   "Additional data - ISO",
		Type:   "an",
	},
	47: {
		Length: 999, // ...
		Name:   "Additional data - national",
		Type:   "an",
	},
	48: {
		Length: 999, // ...
		Name:   "Additional data - private",
		Type:   "an",
	},
	49: {
		Length: 3,
		Name:   "Currency code, transaction",
		Type:   "a or n 3",
	},
	50: {
		Length: 3,
		Name:   "Currency code, settlement",
		Type:   "a or n 3",
	},
	51: {
		Length: 3,
		Name:   "Currency code, cardholder billing",
		Type:   "a or n 3",
	},
	52: {
		Length: 64,
		Name:   "Personal identification number data",
		Type:   "b",
	},
	53: {
		Length: 16,
		Name:   "Security related control information",
		Type:   "n",
	},
	54: {
		Length: 120,
		Name:   "Additional amounts",
		Type:   "an",
	},
	55: {
		Length: 999, // ...
		Name:   "Reserved ISO",
		Type:   "ans",
	},
	56: {
		Length: 999, // ...
		Name:   "Reserved ISO",
		Type:   "ans",
	},
	57: {
		Length: 999, // ...
		Name:   "Reserved national",
		Type:   "ans",
	},
	58: {
		Length: 999, // ...
		Name:   "Reserved national",
		Type:   "ans",
	},
	59: {
		Length: 999, // ...
		Name:   "Reserved national",
		Type:   "ans",
	},
	60: {
		Length: 999, // ...
		Name:   "Reserved national",
		// Type:   "ans",
		Type: "lllan",
	},
	61: {
		Length: 999, // ...
		Name:   "Reserved private",
		Type:   "ans",
	},
	62: {
		Length: 999, // ...
		Name:   "Reserved private",
		Type:   "ans",
	},
	63: {
		Length: 999, // ...
		Name:   "Reserved private",
		Type:   "ans",
	},
	64: {
		Length: 16,
		Name:   "Message authentication code (MAC)",
		Type:   "b",
	},
	65: {
		Length: 1,
		Name:   "Bitmap, extended",
		Type:   "b",
	},
	66: {
		Length: 1,
		Name:   "Settlement code",
		Type:   "n",
	},
	67: {
		Length: 2,
		Name:   "Extended payment code",
		Type:   "n",
	},
	68: {
		Length: 3,
		Name:   "Receiving institution country code",
		Type:   "n",
	},
	69: {
		Length: 3,
		Name:   "Settlement institution country code",
		Type:   "n",
	},
	70: {
		Length: 3,
		Name:   "Network management information code",
		Type:   "n",
	},
	71: {
		Length: 4,
		Name:   "Message number",
		Type:   "n",
	},
	72: {
		Length: 4,
		Name:   "Message number, last",
		Type:   "n",
	},
	73: {
		Length: 6,
		Name:   "Date, action (YYMMDD)",
		Type:   "n",
	},
	74: {
		Length: 10,
		Name:   "Credits, number",
		Type:   "n",
	},
	75: {
		Length: 10,
		Name:   "Credits, reversal number",
		Type:   "n",
	},
	76: {
		Length: 10,
		Name:   "Debits, number",
		Type:   "n",
	},
	77: {
		Length: 10,
		Name:   "Debits, reversal number",
		Type:   "n",
	},
	78: {
		Length: 10,
		Name:   "Transfer number",
		Type:   "n",
	},
	79: {
		Length: 10,
		Name:   "Transfer, reversal number",
		Type:   "n",
	},
	80: {
		Length: 10,
		Name:   "Inquiries number",
		Type:   "n",
	},
	81: {
		Length: 10,
		Name:   "Authorizations, number",
		Type:   "n",
	},
	82: {
		Length: 12,
		Name:   "Credits, processing fee amount",
		Type:   "n",
	},
	83: {
		Length: 12,
		Name:   "Credits, transaction fee amount",
		Type:   "n",
	},
	84: {
		Length: 12,
		Name:   "Debits, processing fee amount",
		Type:   "n",
	},
	85: {
		Length: 12,
		Name:   "Debits, transaction fee amount",
		Type:   "n",
	},
	86: {
		Length: 16,
		Name:   "Credits, amount",
		Type:   "n",
	},
	87: {
		Length: 16,
		Name:   "Credits, reversal amount",
		Type:   "n",
	},
	88: {
		Length: 16,
		Name:   "Debits, amount",
		Type:   "n",
	},
	89: {
		Length: 16,
		Name:   "Debits, reversal amount",
		Type:   "n",
	},
	90: {
		Length: 42,
		Name:   "Original data elements",
		Type:   "n",
	},
	91: {
		Length: 1,
		Name:   "File update code",
		Type:   "an",
	},
	92: {
		Length: 2,
		Name:   "File security code",
		Type:   "an",
	},
	93: {
		Length: 5,
		Name:   "Response indicator",
		Type:   "an",
	},
	94: {
		Length: 7,
		Name:   "Service indicator",
		Type:   "an",
	},
	95: {
		Length: 42,
		Name:   "Replacement amounts",
		Type:   "an",
	},
	96: {
		Length: 64,
		Name:   "Message security code",
		Type:   "b",
	},
	97: {
		Length: 16,
		Name:   "Amount, net settlement",
		Type:   "x+n",
	},
	98: {
		Length: 25,
		Name:   "Payee",
		Type:   "ans",
	},
	99: {
		Length: 11, // ..
		Name:   "Settlement institution identification code",
		Type:   "n",
	},
	100: {
		Length: 11, // ..
		Name:   "Receiving institution identification code",
		Type:   "n",
	},
	101: {
		Length: 17, // ..
		Name:   "File name",
		Type:   "ans",
	},
	102: {
		Length: 28, // ..
		Name:   "Account identification 1",
		Type:   "ans",
	},
	103: {
		Length: 28, // ..
		Name:   "Account identification 2",
		Type:   "ans",
	},
	104: {
		Length: 100, // ...
		Name:   "Transaction description",
		Type:   "ans",
	},
	105: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	106: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	107: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	108: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	109: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	110: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	111: {
		Length: 999, // ...
		Name:   "Reserved for ISO use",
		Type:   "ans",
	},
	112: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	113: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	114: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	115: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	116: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	117: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	118: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	119: {
		Length: 999, // ...
		Name:   "Reserved for national use",
		Type:   "ans",
	},
	120: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	121: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	122: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	123: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	124: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	125: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	126: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	127: {
		Length: 999, // ...
		Name:   "Reserved for private use",
		Type:   "ans",
	},
	128: {
		Length: 64,
		Name:   "Message authentication code",
		Type:   "b",
	},
}

// var t = map[int]pkg{
// 	0: {
// 		Length: 4,
// 		Name:   "Message Type Indicator",
// 		Type:   "n",
// 	},
// 	1: {
// 		Length: 16,
// 		Name:   "Bitmap",
// 		Type:   "b",
// 	},
// 	2: {
// 		Length: 19,
// 		Name:   "Primary Account Number",
// 		Type:   "lln",
// 	},
// 	3: {
// 		Length: 6,
// 		Name:   "Processing Code",
// 		Type:   "n",
// 	},
// 	4: {
// 		Length: 12,
// 		Name:   "Amount, Transaction",
// 		Type:   "n",
// 	},
// 	5: {
// 		Length: 12,
// 		Name:   "Amount, Settlement",
// 		Type:   "n",
// 	},
// 	6: {
// 		Length: 12,
// 		Name:   "Amount, Cardholder Billing",
// 		Type:   "n",
// 	},
// 	7: {
// 		Length: 10,
// 		Name:   "Transmission Date and Time",
// 		Type:   "n",
// 	},
// 	8: {
// 		Length: 8,
// 		Name:   "Amount, Cardholder Billing Fee",
// 		Type:   "n",
// 	},
// 	9: {
// 		Length: 8,
// 		Name:   "Conversion Rate, Settlement",
// 		Type:   "n",
// 	},
// 	10: {
// 		Length: 8,
// 		Name:   "Conversion Rate, Cardholder Billing",
// 		Type:   "n",
// 	},
// 	11: {
// 		Length: 6,
// 		Name:   "System Trace Audit Number",
// 		Type:   "n",
// 	},
// 	12: {
// 		Length: 6,
// 		Name:   "Time, Local Transaction",
// 		Type:   "n",
// 	},
// 	13: {
// 		Length: 4,
// 		Name:   "Date, Local Transaction",
// 		Type:   "n",
// 	},
// 	15: {
// 		Length: 4,
// 		Name:   "Date, Expiration", //Date, Settlement
// 		Type:   "n",
// 	},
// 	32: {
// 		Length: 11,
// 		Name:   "Acquiring Institution Ident Code",
// 		Type:   "lln",
// 	},
// 	37: {
// 		Length: 12,
// 		Name:   "Retrieval Reference Number",
// 		Type:   "an",
// 	},
// 	38: {
// 		Length: 6,
// 		Name:   "Authorization identification response",
// 		Type:   "an",
// 	},
// 	39: {
// 		Length: 2,
// 		Name:   "Response code",
// 		Type:   "an",
// 	},
// 	41: {
// 		Length: 8,
// 		Name:   "Card Acceptor Terminal Identification",
// 		Type:   "ans",
// 	},
// 	49: {
// 		Length: 3,
// 		Name:   "Currency code, transaction",
// 		Type:   "a",
// 	},
// 	54: {
// 		Length: 120,
// 		Name:   "Additional amounts",
// 		Type:   "lllan",
// 	},
// 	70: {
// 		Length: 3,
// 		Name:   "Network Management Information Code",
// 		Type:   "n",
// 	},
// }
