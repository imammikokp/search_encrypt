/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"search_encrypt/domain"
	"search_encrypt/pkg/database"

	searchEncryptHandler "search_encrypt/internal/search_encrypt/handler"
	searchEncryptRepo "search_encrypt/internal/search_encrypt/repository"
	searchEncryptUseCase "search_encrypt/internal/search_encrypt/usecase"

	searchHistoryEncryptRepo "search_encrypt/internal/search_history_encrypt/repository"
	searchHistoryEncryptUseCase "search_encrypt/internal/search_history_encrypt/usecase"

	invalidEncryptionRepo "search_encrypt/internal/invalid_encryption/repository"
	invalidHistoryEncryptionRepo "search_encrypt/internal/invalid_history_encryption/repository"

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
	invalidEncryptDB.Conn().AutoMigrate(&domain.InvalidHistoryEncryption{})


	invalidEncryptionRepository := invalidEncryptionRepo.NewInvalidEncryptReposiotry(invalidEncryptDB.Conn())
	invalidHistoryEncryptionRepository := invalidHistoryEncryptionRepo.NewInvalidHistoryEncryptReposiotry(invalidEncryptDB.Conn())

	searchEncryptRepository := searchEncryptRepo.NewSearchEncryptRepository(customerDomainDb.Conn())
	searchEncryptUseCase := searchEncryptUseCase.NewSearchEncryptUseCase(time.Duration(3)*time.Minute, searchEncryptRepository, invalidEncryptionRepository)
	
	searchHistoryEncryptRepository := searchHistoryEncryptRepo.NewSearchHistoryEncryptRepository(customerDomainDb.Conn())
	searchHistoryEncryptUseCase := searchHistoryEncryptUseCase.NewSearchHistoryEncryptUseCase(time.Duration(3)*time.Minute, searchHistoryEncryptRepository,invalidHistoryEncryptionRepository)

	searchHandler := searchEncryptHandler.NewSearchEncryptCmdHandler(searchEncryptUseCase,searchHistoryEncryptUseCase ,rootCmd)
	searchEncryptHandler.Initialize(rootCmd, searchHandler)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Execute()
}
