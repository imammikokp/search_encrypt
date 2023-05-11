package handler

import (
	"serch_encrypt/domain"

	"github.com/spf13/cobra"
)

func Initialize(cli *cobra.Command, handler domain.SearchEncryptCmdHandler) {
	cli.AddCommand(handler.CountAll())
	cli.AddCommand(handler.FindInvalidEncryptByRange())
}
