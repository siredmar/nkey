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

// Package cmd is the cobra command structure
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/siredmar/nkey/pkg/nkey"
)

var seed string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nkey",
	Short: "Convert NATS nkey to keypair",
	Long:  `Convert NATS nkey to keypair`,
	Run: func(cmd *cobra.Command, args []string) {
		if seed == "" {
			log.Fatal("No seed provided")
		}
		priv, pub, err := nkey.Convert(seed)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Private key: %s\n", priv)
		fmt.Printf("Public key: %s\n", pub)
	},
}

// Execute tells cobra what to do
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&seed, "seed", "s", "", "NKEY seed")
}
