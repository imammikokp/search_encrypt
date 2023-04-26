/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"serch_encrypt/domain"
	"serch_encrypt/internal/search_encrypt/handler"
	"serch_encrypt/internal/search_encrypt/repository"
	"serch_encrypt/internal/search_encrypt/usecase"
	"time"

	"github.com/spf13/cobra"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

	customerDomainDsn := "host=35.219.112.128 user=developer password=kreditmu30 dbname=customer_domain port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	customerDomainDb, err := gorm.Open(postgres.Open(customerDomainDsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	invalidEncryptDsn := "host=localhost user=postgres password=admin dbname=invalid_encryption port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	invalidEncryptDB, err := gorm.Open(postgres.Open(invalidEncryptDsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	invalidEncryptDB.AutoMigrate(&domain.InvalidEncryption{})

	searchRepository:=repository.NewSearchEncryptRepository(customerDomainDb)
	searchUseCase:=usecase.NewSearchEncryptUseCase(time.Duration(3)*time.Minute, searchRepository)
	searchHandler:=handler.NewSearchEncryptCmdHandler(searchUseCase,rootCmd)
	handler.Initialize(rootCmd,searchHandler)
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Execute()
}
