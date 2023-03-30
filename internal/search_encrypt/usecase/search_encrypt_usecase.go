package usecase

import "serch_encrypt/domain"

type searchEncryptUseCase struct {
	customerRepository domain.CustomerRepository
}

func NewSearchEncryptUseCase( customerRepository domain.CustomerRepository) domain.CustomerUseCase{
	return &searchEncryptUseCase{
		customerRepository: customerRepository,
	}
}


func (r searchEncryptUseCase) checkLength()(*int,error){
	
	return nil, nil
}