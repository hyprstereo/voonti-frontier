package cmd

type Command = cobra.Command
var (
	rootCmd *Command
	cmdPattern string `*.cmd.voo`
)