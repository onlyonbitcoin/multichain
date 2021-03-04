package bitcoinsv

import (
	"github.com/btcsuite/btcd/chaincfg"
)

func init() {
	if err := chaincfg.Register(&MainNetParams); err != nil {
		panic(err)
	}
	if err := chaincfg.Register(&RegressionNetParams); err != nil {
		panic(err)
	}
}

var MainNetParams = chaincfg.Params{
	Name: "mainnet",
	Net:  0xe3e1f3e8,

	// Address encoding magics
	PubKeyHashAddrID: 10,
	ScriptHashAddrID: 15,
	PrivateKeyID:     128,

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x04, 0x88, 0xB2, 0x1E}, // starts with xpub
	HDPrivateKeyID: [4]byte{0x04, 0x88, 0xAD, 0xE4}, // starts with xprv

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. Bitcoinsv does not actually support this, but we do not want to
	// collide with real addresses, so we specify it.
	Bech32HRPSegwit: "bsv",
}

var RegressionNetParams = chaincfg.Params{
	Name: "regtest",

	// Dogecoin has 0xdab5bffa as RegTest (same as Bitcoin's RegTest).
	// Setting it to an arbitrary value (leet_hex(dogecoin)), so that we can
	// register the regtest network.
	Net: 0xf4e5f3f4,

	// Address encoding magics
	PubKeyHashAddrID: 111,
	ScriptHashAddrID: 196,
	PrivateKeyID:     239,

	// BIP32 hierarchical deterministic extended key magics
	HDPublicKeyID:  [4]byte{0x04, 0x35, 0x87, 0xCF}, // starts with xpub
	HDPrivateKeyID: [4]byte{0x04, 0x35, 0x83, 0x94}, // starts with xprv

	// Human-readable part for Bech32 encoded segwit addresses, as defined in
	// BIP 173. Dogecoin does not actually support this, but we do not want to
	// collide with real addresses, so we specify it.
	Bech32HRPSegwit: "dogert",
}
