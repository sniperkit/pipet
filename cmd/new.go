// Copyright © 2018 Dhananjay Balan
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
//    may be used to endorse or promote products derived from this software
//    without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
// FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
// DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
// SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
// CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
// OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"

	"github.com/dbalan/pipet/pipetdata"
	"github.com/satori/go.uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	snippetTags = []string{}
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: `Creates a new snippet and opens $EDITOR to edit content`,
	Run: func(cmd *cobra.Command, args []string) {
		title := cmd.Flag("title").Value.String()
		diskPath := viper.Get("document_dir").(string)

		if title == "" {
			title = uuid.Must(uuid.NewV4()).String()
		}

		if len(snippetTags) == 0 {
			snippetTags = append(snippetTags, "untagged")
		}

		dataStore, err := pipetdata.NewDataStore(expandHome(diskPath))
		errorGuard(err, "error accessing data store")

		fn, err := dataStore.New(title, snippetTags...)
		errorGuard(err, "creating snippet failed")

		fmt.Println("created snippet: ", fn)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	newCmd.PersistentFlags().String("title", "", "title for the snippet, if unset a random uuid is used.")
	newCmd.PersistentFlags().StringArray("tags", snippetTags, "tags for snippet, if unset, a single tag `untagged` is set.")

}
