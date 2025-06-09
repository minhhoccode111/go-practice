package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	n, _ := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(2))) // sub 2 because we have to add 2 later to be greater than 1
	return new(big.Int).Add(n, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	gBig := big.NewInt(g)
	return new(big.Int).Exp(gBig, private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
