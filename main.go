/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"serch_encrypt/domain"
	"serch_encrypt/pkg/database"

	searchEncryptHandler "serch_encrypt/internal/search_encrypt/handler"
	searchEncryptRepo "serch_encrypt/internal/search_encrypt/repository"
	searchEncryptUseCase "serch_encrypt/internal/search_encrypt/usecase"

	invalidEncryptionRepo "serch_encrypt/internal/invalid_encryption/repository"

	"time"

	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "search_encrypt",
		Short: "search invalid in encrypt ",
		Long:  `search invalid in encrypt`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}


	viper.SetConfigName("app.json")
	viper.SetConfigType("json")
	viper.AddConfigPath("./conf/")
	err := viper.ReadInConfig()
    if err != nil {
        fmt.Printf("Error reading config file: %s\n", err)
        return
    }
	databaseCustomersConf:=viper.GetStringMapString("database_customers")
	customerDomainDb, err := database.New(database.ConfigFromEnvironment(databaseCustomersConf))
	if err != nil {
		panic(err.Error())
	}

	databaseInvalidEncryptConf:=viper.GetStringMapString("database_invalid_encryption_customers")
	invalidEncryptDB, err := database.New(database.ConfigFromEnvironment(databaseInvalidEncryptConf))
	if err != nil {
		panic(err.Error())
	}

	invalidEncryptDB.Conn().AutoMigrate(&domain.InvalidEncryption{})
	invalidEncryptionRepository := invalidEncryptionRepo.NewInvalidEncryptReposiotry(invalidEncryptDB.Conn())

	searchRepository := searchEncryptRepo.NewSearchEncryptRepository(customerDomainDb.Conn())
	searchUseCase := searchEncryptUseCase.NewSearchEncryptUseCase(time.Duration(3)*time.Minute, searchRepository, invalidEncryptionRepository)
	searchHandler := searchEncryptHandler.NewSearchEncryptCmdHandler(searchUseCase, rootCmd)
	searchEncryptHandler.Initialize(rootCmd, searchHandler)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Execute()
}
