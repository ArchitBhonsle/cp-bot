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
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ArchitBhonsle/cp-bot/logic/fileops"
	"github.com/ArchitBhonsle/cp-bot/logic/scraping"
	"github.com/ArchitBhonsle/cp-bot/logic/types"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cp-bot [args] [flags]",
	Short: "A tool for simplifying getting started with a competitive programming contest.",
	Long: `This is a command line tool written in golang to streamline the process of
participating in a competitive programming contest. The command can be used as follows:

cp-bot https://codeforces.com/contest/1498
cp-bot atcoder.jp/contests/arc114/
	`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		contest, err := scraping.GetContest(args[0])
		if err != nil {
			return err
		}

		send := make(chan *types.FetchedProblem, len(contest))
		for _, problem := range contest {
			go problem.Fetch(send)
		}

		for i := 1; i <= len(contest); i += 1 {
			fetchedProblem := <-send
			fileops.CreateFiles(fetchedProblem)
		}

		return nil
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	home, err := homedir.Dir()
	cobra.CheckErr(err)

	// Which config file to use
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", fmt.Sprintf("config file (default \"%v\")", path.Join(home, ".cp-bot")))

	// Which directory to use for competitive programming
	rootCmd.PersistentFlags().StringP("directory", "d", path.Join(home, "cp"), "the directory to use for competitive programming")
	viper.BindPFlag("directory", rootCmd.PersistentFlags().Lookup("directory"))

	// Which template to use
	rootCmd.PersistentFlags().StringP("template", "t", path.Join(home, "cp", "template.cpp"), "the template to use")
	viper.BindPFlag("template", rootCmd.PersistentFlags().Lookup("template"))

	// Whether to print various statements
	// TODO actually use this
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.SetConfigType("yaml")
		viper.AddConfigPath(path.Join(home, ".config", "cp-bot"))
		viper.SetConfigName("config.yaml")
	}

	viper.SetEnvPrefix("CPB")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			cobra.CheckErr(err)
		}
	}
}
