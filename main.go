/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"serch_encrypt/domain"

	searchEncryptHandler "serch_encrypt/internal/search_encrypt/handler"
	searchEncryptRepo "serch_encrypt/internal/search_encrypt/repository"
	searchEncryptUseCase "serch_encrypt/internal/search_encrypt/usecase"

	invalidEncryptionRepo "serch_encrypt/internal/invalid_encryption/repository"

	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "search_encrypt",
		Short: "search invalid in encrypt ",
		Long: `search invalid in encrypt`,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
	}

	customerDomainDsn := "host= user= password= dbname= port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	customerDomainDb, err := gorm.Open(postgres.Open(customerDomainDsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}

	invalidEncryptDsn := "host=localhost user=postgres password= dbname=invalid_encryption port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	invalidEncryptDB, err := gorm.Open(postgres.Open(invalidEncryptDsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	invalidEncryptDB.AutoMigrate(&domain.InvalidEncryption{})

	invalidEncryptionRepository:=invalidEncryptionRepo.NewInvalidEncryptReposiotry(invalidEncryptDB)



	searchRepository:=searchEncryptRepo.NewSearchEncryptRepository(customerDomainDb)
	searchUseCase:=searchEncryptUseCase.NewSearchEncryptUseCase(time.Duration(3)*time.Minute, searchRepository,invalidEncryptionRepository)
	searchHandler:=searchEncryptHandler.NewSearchEncryptCmdHandler(searchUseCase,rootCmd)
	searchEncryptHandler.Initialize(rootCmd,searchHandler)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Execute()
}
