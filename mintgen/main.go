package main

import (
	"fmt"

	"github.com/eris-ltd/mint-client/Godeps/_workspace/src/github.com/spf13/cobra"
)

var (
	DirFlag string
	//AddrsFlag  string
	CsvPathFlag string
	PubkeyFlag  string
)

func main() {

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "check mintgen version",
		Run: func(cmd *cobra.Command, args []string) {
			// major: alpha
			// minor: drop pub, csvs with empty fields, no config
			// revision:
			fmt.Println("0.3.0")
		},
	}

	var randomCmd = &cobra.Command{
		Use:   "random",
		Short: "mintgen random [flags] <N> <chain_id>",
		Long:  "Create <N> keys and a genesis.json with corresponding validators and chain_id <name>",
		Run:   cliRandom,
	}

	var knownCmd = &cobra.Command{
		Use:   "known",
		Short: "mintgen known <chain_id> [flags] ",
		Long:  "Create a genesis.json with --pub <pub_1> <pub_2> <pub_N> or --csv <path_to_file>, or pass stdin",
		Run:   cliKnown,
	}

	randomCmd.Flags().StringVarP(&DirFlag, "dir", "d", "", "Directory to save genesis and priv_validators in. Default is ~/.eris/data/<chain_id>")

	knownCmd.Flags().StringVarP(&PubkeyFlag, "pub", "", "", "pubkeys to include when generating genesis.json. flag is req'd")
	knownCmd.Flags().StringVarP(&DirFlag, "dir", "d", "", "Directory to save genesis.json in. Default is ~/.eris/data/<chain_id>")
	knownCmd.Flags().StringVarP(&CsvPathFlag, "csv", "", "", "Path to .csv with the following params: (pubkey, starting balance, name, permissions, setbit")

	var rootCmd = &cobra.Command{
		Use:   "mintgen",
		Short: "a tool for generating tendermint genesis files",
		Long:  "a tool for generating tendermint genesis files",
	}
	rootCmd.AddCommand(randomCmd, knownCmd, versionCmd)
	rootCmd.Execute()
}
