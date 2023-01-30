// The AGPLv3 License (AGPLv3)

// Copyright (c) 2022 ZHAO Zhenhua <zhao.zhenhua@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package crypto

import (
	gocrypto "crypto"
	"hash"
)

type Key interface {
	//the key's raw byte
	Bytes() ([]byte, error)
	//PrivateKey returns true is this is a asymmetric private key or symmetric security key
	PrivateKey() bool
	//symmetric returns true if this key is symmetric, otherwise false
	Symmetric() bool
	//if this is a asymmetric key, returns the corresponding Public key, otherwise error
	PublicKey() (Key, error)
}

// Key generation options for CryptoProvider
type KeyGenOpts interface {
	Algorithm() string
}

// HashOpts contains hash options for CryptoProvider
type HashOpts interface {
	Algorithm() string
}

// EncrypterOpts contains encrypting options
type EncrypterOpts interface {
}

// DecrypterOpts contains decrypting options
type DecrypterOpts interface {
}

// SignerOpts contain signing options
type SignerOpts interface {
	gocrypto.SignerOpts
}

type Signer interface {
	//Sign signs a message's hash
	Sign(k Key, digest []byte) ([]byte, error)
}

type Verifier interface {
	Verify(k Key, signature, digest []byte) (bool, error)
}

// Recover public key from signature
type Recoverer interface {
	Recover(hash, signature []byte) ([]byte, error)
}

type Encrypter interface {
	Encrypt(k Key, plaintext []byte) ([]byte, error)
}

type Decrypter interface {
	Decrypt(k Key, ciphertext []byte) ([]byte, error)
}

// bytehub+ crytograhic provider
type CryptoProvider interface {
	//KeyGen generates a new symmetric key
	KeyGen(opts KeyGenOpts) (Key, error)

	//KeyGen generates a new pair of asymmetric key
	KeyPairGen(opts KeyGenOpts) (Key, Key, error)

	//Hash hashes a message
	Hash(msg []byte, opts HashOpts) ([]byte, error)

	//GetHash returns the instance of hash function
	GetHash(opt HashOpts) (hash.Hash, error)

	Encrypter
	Decrypter
	Signer
	Verifier
}
