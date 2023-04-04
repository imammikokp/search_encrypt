package usecase

import (
	"context"
	"serch_encrypt/domain"
	"time"
)

type searchEncryptUseCase struct {
	customerRepository domain.CustomerRepository
	contexTimeout time.Duration
}

func NewSearchEncryptUseCase( contextTimeout time.Duration,customerRepository domain.CustomerRepository) domain.CustomerUseCase{
	return &searchEncryptUseCase{
		customerRepository: customerRepository,
		contexTimeout: contextTimeout,
	}
}


func (r searchEncryptUseCase) CheckLength()(int64,error){
	ctxB := context.Background()
	ctx,cancel := context.WithTimeout(ctxB,r.contexTimeout)
	defer cancel()
	count, err := r.customerRepository.GetCountAll(ctx)
	if err != nil{
		return 0 , err
	}
	return count, nil
}