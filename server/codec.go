package server

import "github.com/matthieutran/leafre-login/pkg/crypto"

func generateCodecs(version int, ivRecv, ivSend [4]byte) (encrypter, decrypter func(d []byte) []byte) {
	// Create codecs
	c := crypto.NewCodec(ivRecv, ivSend, version)

	// Create encrypter
	encrypter = func(d []byte) (res []byte) {
		res, _ = c.Encrypt(d, true, true)
		return
	}

	// Create decrypter
	decrypter = func(d []byte) (res []byte) {
		res, _ = c.Decrypt(d, true, true)
		return
	}

	return
}