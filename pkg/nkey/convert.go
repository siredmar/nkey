/*
Copyright © 2023 Armin Schlegel <armin.schlegel@gmx.de>
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package nkey contains nkey specific metods
package nkey

import (
	"github.com/nats-io/nkeys"
)

// Convert can convert a given nkey seed to public/private keypair
func Convert(seed string) (string, string, error) {
	kp, err := nkeys.FromSeed([]byte(seed))
	if err != nil {
		return "", "", err
	}
	pub, err := kp.PublicKey()
	if err != nil {
		return "", "", err
	}

	priv, err := kp.PrivateKey()
	if err != nil {
		return "", "", err
	}

	return string(priv), pub, nil
}
