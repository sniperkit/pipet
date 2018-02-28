// Copyright © 2018 Dhananjay Balan <mail@dbalan.in>
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
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package cmd

import (
	"fmt"
	"github.com/dbalan/pipet/pipetdata"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	tags = false
	full = false
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all snippets",
	Long:  `Lists all snippets, by default it only prints the uid and title`,
	Run: func(cmd *cobra.Command, args []string) {
		diskPath := viper.Get("document_dir").(string)
		dataStore, err := pipetdata.NewDataStore(expandHome(diskPath))
		errorGuard(err, "error accessing data store")

		sns, err := dataStore.List()
		errorGuard(err, "listing store failed")

		for _, snip := range sns {
			tagList := ""
			for _, t := range snip.Meta.Tags {
				tagList += t + " "
			}
			fmt.Printf("%s\t%s\t[%s]\n", snip.Meta.UID, snip.Meta.Title, tagList)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&tags, "tags", "t", false, "Enable list tags associated with the snippet.")
	listCmd.Flags().BoolVarP(&full, "full", "f", false, "Print the full snippet")
}