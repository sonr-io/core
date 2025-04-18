package types

import (
	"github.com/sonr-io/coins"
)

// CoinType represents different cryptocurrency types
type CoinType int

const (
	// Main coins
	Aptos CoinType = iota
	Avax
	Bitcoin
	Ethereum
	Filecoin
	Helium
	Polkadot
	Solana
	Sui
	Tezos

	// Cosmos ecosystem coins
	Atom
	Axelar
	Cronos
	Evmos
	Iris
	Juno
	Kava
	Kujira
	Okc
	Osmo
	Secret
	Sei
	Sonr
	Stargaze
	Tia
)

// FormatAddress returns the string representation of the address
func FormatAddress(c CoinType, pubKey []byte) (string, error) {
	switch c {
	case Aptos:
		return coins.AptosAddressFromPublicKey(pubKey)
	case Avax:
		return coins.AvaxAddressFromPublicKey(pubKey)
	case Bitcoin:
		return coins.BitcoinAddressFromPublicKey(pubKey)
	case Ethereum:
		return coins.EthereumAddressFromPublicKey(pubKey)
	case Filecoin:
		return coins.FilecoinAddressFromPublicKey(pubKey)
	case Helium:
		return coins.HeliumAddressFromPublicKey(pubKey)
	case Polkadot:
		return coins.PolkadotAddressFromPublicKey(pubKey)
	case Solana:
		return coins.SolanaAddressFromPublicKey(pubKey)
	case Sui:
		return coins.SuiAddressFromPublicKey(pubKey)
	case Tezos:
		return coins.TezosAddressFromPublicKey(pubKey)
	case Atom:
		return coins.AtomAddressFromPublicKey(pubKey)
	case Axelar:
		return coins.AxelarAddressFromPublicKey(pubKey)
	case Cronos:
		return coins.CronosAddressFromPublicKey(pubKey)
	case Evmos:
		return coins.EvmosAddressFromPublicKey(pubKey)
	case Iris:
		return coins.IrisAddressFromPublicKey(pubKey)
	case Juno:
		return coins.JunoAddressFromPublicKey(pubKey)
	case Kava:
		return coins.KavaAddressFromPublicKey(pubKey)
	case Kujira:
		return coins.KujiraAddressFromPublicKey(pubKey)
	case Okc:
		return coins.OkcAddressFromPublicKey(pubKey)
	case Osmo:
		return coins.OsmoAddressFromPublicKey(pubKey)
	case Secret:
		return coins.SecretAddressFromPublicKey(pubKey)
	case Sei:
		return coins.SeiAddressFromPublicKey(pubKey)
	case Sonr:
		return coins.SonrAddressFromPublicKey(pubKey)
	case Stargaze:
		return coins.StargazeAddressFromPublicKey(pubKey)
	case Tia:
		return coins.TiaAddressFromPublicKey(pubKey)
	default:
		return "", nil
	}
}

// String returns the string representation of the coin type
func (c CoinType) String() string {
	return [...]string{
		"aptos",
		"avax",
		"bitcoin",
		"ethereum",
		"filecoin",
		"helium",
		"polkadot",
		"solana",
		"sui",
		"tezos",
		"atom",
		"axelar",
		"cronos",
		"evmos",
		"iris",
		"juno",
		"kava",
		"kujira",
		"okc",
		"osmo",
		"secret",
		"sei",
		"sonr",
		"stargaze",
		"tia",
	}[c]
}

func CoinList() []CoinType {
	return []CoinType{
		Aptos,
		Avax,
		Bitcoin,
		Ethereum,
		Filecoin,
		Helium,
		Polkadot,
		Solana,
		Sui,
		Tezos,
		Atom,
		Axelar,
		Cronos,
		Evmos,
		Iris,
		Juno,
		Kava,
		Kujira,
		Okc,
		Osmo,
		Secret,
		Sei,
		Sonr,
		Stargaze,
		Tia,
	}
}
