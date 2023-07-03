package handler

import (
	"fmt"
	"search_encrypt/domain"

	"search_encrypt/pkg"

	"github.com/spf13/cobra"
)

type searchEnctyptCmdHandler struct {
	SearchEncryptUseCase domain.SearchEncryptUseCase
	SearchHistoryEncryptUseCase  domain.SearchHistoryEncryptUseCase
}

func NewSearchEncryptCmdHandler(SearchEncryptUseCase domain.SearchEncryptUseCase,SearchHistoryEncryptUseCase domain.SearchHistoryEncryptUseCase ,rootCmd *cobra.Command) domain.SearchEncryptCmdHandler {
	return &searchEnctyptCmdHandler{
		SearchEncryptUseCase: SearchEncryptUseCase,
		SearchHistoryEncryptUseCase: SearchHistoryEncryptUseCase,
	}
}

func (r searchEnctyptCmdHandler) CountAll() *cobra.Command {
	var countAll = &cobra.Command{
		Use:     "countAll",
		Short:   "countAll",
		Aliases: []string{"CAll"},
		Run: func(cmd *cobra.Command, args []string) {
			LengthValue, err := r.SearchEncryptUseCase.CheckLength()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("the amount of datas of database are ", LengthValue)
		},
	}

	return countAll
}

func (r searchEnctyptCmdHandler) FindInvalidEncryptByRange() *cobra.Command {
	var countAll = &cobra.Command{
		Use:     "findInvalid",
		Short:   "find Invalid Encrypt By Range",
		Aliases: []string{"min Id","max Id"},
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			r.SearchEncryptUseCase.FindAndSaveInvalidEncryptByRange(pkg.StringToInt(args[0]),pkg.StringToInt(args[1]))

		},
	}

	return countAll
}

func (r searchEnctyptCmdHandler) CountAllHistory() *cobra.Command {
	var countAll = &cobra.Command{
		Use:     "countAllHistory",
		Short:   "countAllHistory",
		Aliases: []string{"CAll"},
		Run: func(cmd *cobra.Command, args []string) {
			LengthValue, err := r.SearchHistoryEncryptUseCase.CheckLength()
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println("the amount of datas of database are ", LengthValue)
		},
	}

	return countAll
}

func (r searchEnctyptCmdHandler) FindInvalidHistoryEncryptByRange() *cobra.Command {
	var countAll = &cobra.Command{
		Use:     "findInvalidHistory",
		Short:   "find Invalid Encrypt By Range",
		Aliases: []string{"min Id","max Id"},
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {

			r.SearchHistoryEncryptUseCase.FindAndSaveInvalidEncryptByRange(pkg.StringToInt(args[0]),pkg.StringToInt(args[1]))

		},
	}

	return countAll
}