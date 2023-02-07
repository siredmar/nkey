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

package nkey

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	assert := assert.New(t)
	priv, pub, err := Convert("SAAOW7V3L5YRJGBYY3ZTXB5WHYU4W3HFQ44MBYIBWAECRMD6RMZJTS426E")
	assert.Nil(err)
	assert.Equal(priv, "PDVX5O27OEKJQOGG6M5YPNR6FHFWZZMHHDAOCANQBAULA7ULGKM4XXDL2O2N6GNXC7SYZDSCXEWR5NN4SDECHPU77HZX6K7JCJJV5EQFFIZQ")
	assert.Equal(pub, "ADOGXU5U34M3OF7FRSHEFOJND223ZEGIEO7J76PTP4V6SESTL2JALU4U")
}
