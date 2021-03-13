/*
Copyright © 2021 Archit Bhonsle <abhonsle2000@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ArchitBhonsle/cp-bot/logic/fileops"
	"github.com/ArchitBhonsle/cp-bot/logic/scraping"
	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cp-bot [args]",
	Short: "A tool for simplifying getting started with a competitive programming contest.",
	Long: `This is a command line tool written in golang to streamline the process of
	participating in a competitive programming contest. The command can be used as follows:

	cp-bot https://codeforces.com/contest/1498
	cp-bot www.codechef.com/MARCH21B/
	cp-bot atcoder.jp/contests/arc114/
	`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		println(viper.GetString("directory"), viper.GetString("template"))
		contest, err := scraping.GetContest(args[0])
		cobra.CheckErr(err)

		send := make(chan *types.FetchedProblem, len(contest))
		for _, problem := range contest {
			go problem.Fetch(send)
		}

		for i := 1; i <= len(contest); i += 1 {
			fetchedProblem := <-send
			println(i, fileops.ProblemPath(fetchedProblem.ProblemInfo))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cp-bot.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cp-bot" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cp-bot")
		viper.SetConfigType("yaml")
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
