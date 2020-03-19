/*
 * Copyright (c) 2019 ubirch GmbH.
 *
 * ```
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 * ```
 */

package ubirch

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/google/uuid"
	"github.com/paypal/go.crypto/keystore"
)

// TestCreateKeyStore tests, if a new keystore can be created
func TestCreateKeystore(t *testing.T) {
	asserter := assert.New(t)
	//create new crypto context and check, if the kystore is correct TODO not sure if this test is valid
	var context = &CryptoContext{
		Keystore: &keystore.Keystore{},
		Names:    map[string]uuid.UUID{},
	}
	asserter.IsTypef(&keystore.Keystore{}, context.Keystore, "Keystore creation failed")
}

// TODO saveProtocolContext why is this function in the main
// TODO loadProtocolContext, why is this function in the main
// TODO: Answer, the load and store functions are outside, to keep the protocol outside the keystore

func TestLoadKeystore(t *testing.T) {
	asserter := assert.New(t)
	//	requirer := require.New(t)
	//Set up test objects and parameters
	var context = &CryptoContext{
		Keystore: &keystore.Keystore{},
		Names:    map[string]uuid.UUID{},
	}
	p := Protocol{
		Crypto:     context,
		Signatures: map[uuid.UUID][]byte{},
	}
	asserter.NoErrorf(loadProtocolContext(&p, "../test.json"), "Failed loading")
	id := uuid.MustParse(defaultUUID)
	asserter.Nilf(context.GenerateKey(defaultName, id), "Failed to generate Key")
}

func TestSetKey(t *testing.T) {
	asserter := assert.New(t)
	//Set up test objects and parameters
	var context = &CryptoContext{
		Keystore: &keystore.Keystore{},
		Names:    map[string]uuid.UUID{},
	}
	id := uuid.MustParse(defaultUUID)
	privBytesCorrect, err := hex.DecodeString(defaultPriv)
	if err != nil {
		panic(err)
	}
	privBytesTooLong := append(privBytesCorrect, 0xFF)
	privBytesTooShort := privBytesCorrect[1:]

	//Test valid key length
	asserter.Nilf(context.SetKey(defaultName, id, privBytesCorrect), "SetKey() failed with error: %v", err)
	// test to short key
	asserter.Errorf(context.SetKey(defaultName, id, privBytesTooShort), "SetKey() accepts too short keys.")
	// test too long key
	asserter.Errorf(context.SetKey(defaultName, id, privBytesTooLong), "SetKey() accepts too long keys")
}

func TestSetPublicKey(t *testing.T) {
	//Set up test objects and parameters
	var context = &CryptoContext{
		Keystore: &keystore.Keystore{},
		Names:    map[string]uuid.UUID{},
	}
	id := uuid.MustParse(defaultUUID)
	pubBytesCorrect, err := hex.DecodeString(defaultPub)
	if err != nil {
		panic(err)
	}
	pubBytesTooLong := append(pubBytesCorrect, 0xFF)
	pubBytesTooShort := pubBytesCorrect[1:]

	//Test valid key length
	err = context.SetPublicKey(defaultName, id, pubBytesCorrect)
	if err != nil {
		t.Errorf("SetPublicKey() failed with error: %v", err)
	}
	err = context.SetPublicKey(defaultName, id, pubBytesTooShort)
	if err == nil {
		t.Errorf("SetPublicKey() accepts too short keys.")
	}
	err = context.SetPublicKey(defaultName, id, pubBytesTooLong)
	if err == nil {
		t.Errorf("SetPublicKey() accepts too long keys")
	}
}
