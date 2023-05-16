package usecase

import (
	"context"
	"errors"
	"fmt"
	"serch_encrypt/domain"
	"strings"

	"time"

	"gorm.io/gorm"
)

type searchEncryptUseCase struct {
	customerRepository domain.SearchEncryptRepository
	InvalidEncryptionRepository domain.InvalidEncryptionRepository
	contexTimeout      time.Duration
}

func NewSearchEncryptUseCase(contextTimeout time.Duration, customerRepository domain.SearchEncryptRepository,InvalidEncryptionRepository domain.InvalidEncryptionRepository) domain.SearchEncryptUseCase {
	return &searchEncryptUseCase{
		customerRepository: customerRepository,
		InvalidEncryptionRepository:InvalidEncryptionRepository,
		contexTimeout:      contextTimeout,
	}
}

func (r searchEncryptUseCase) CheckLength() (int64, error) {
	ctxB := context.Background()
	ctx, cancel := context.WithTimeout(ctxB, r.contexTimeout)
	defer cancel()
	count, err := r.customerRepository.GetCountAll(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r searchEncryptUseCase) FindAndSaveInvalidEncryptByRange(minId int, maxId int) (validAmount int, invalidAmount int, funcError error) {
	ctxB := context.Background()
	ctx, cancel := context.WithTimeout(ctxB, r.contexTimeout)
	defer cancel()
	root := domain.NewRoot(minId, maxId)
	var InvalidEncrypts []int
	for {

		if root.Root == nil {
			break
		}

		parent, child := domain.NearestSearch(root.Root, root.Root)
		if child == nil {
			break
		}
		if child != nil {
			if findEncrypt, finish, err := r.searchEncyrpt(ctx, parent, child, InvalidEncrypts); err != nil {
				return 0, 0, err
			} else {
				if finish {
					break
				}
				if findEncrypt != nil {
					InvalidEncrypts = append(InvalidEncrypts, *findEncrypt...)
				}
			}
		}

		var invalidEncryptModels []domain.InvalidEncryption
		for _, v := range InvalidEncrypts {
			invalidEncryptModels = append(invalidEncryptModels, domain.InvalidEncryption{CustomerID: v})
		}

		if err:=r.InvalidEncryptionRepository.Inserts(ctx,invalidEncryptModels);err!=nil{
			return 0,0,err
		}

		domain.Stringify(root.Root, 0)

	}
	
	// fmt.Println("->result", InvalidModelEncrypt)
	return
}

func (r searchEncryptUseCase) searchEncyrpt(ctx context.Context, parent *domain.Node, child *domain.Node, inEncryptResults []int) (*[]int, bool, error) {
	// encrypt customer
	var modelEncrypt []domain.EncryptCustomer

	if err := r.customerRepository.FetchByRange(ctx, &modelEncrypt, inEncryptResults, child.Min, child.Max); err != nil {

		if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
			//min dan max id jaraknya 1
			if (child.Max - child.Min) <= 1 {
				var inEncryptResult []int
				var encyrptCustomer domain.EncryptCustomer
				if err := r.customerRepository.FindById(ctx, &encyrptCustomer, child.Min); err != nil {
					if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
						//simpan invalid
						fmt.Println("->child test",parent,child)
						domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
						inEncryptResult = append(inEncryptResult, child.Min)
					} else if errors.Is(err, gorm.ErrRecordNotFound) {
						

					} else {
						return nil, false, err
					}
				}

				if child.Min != child.Max{
					if err := r.customerRepository.FindById(ctx, &encyrptCustomer, child.Max); err != nil {
						if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
							//simpan invalid
							fmt.Println("->child test",parent,child)
							domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
							inEncryptResult = append(inEncryptResult, child.Max)
						} else if errors.Is(err, gorm.ErrRecordNotFound) {
						
	
						} else {
							return nil, false, err
						}
					}
				}

				if len(inEncryptResult)!=0{
					return &inEncryptResult,false,nil
				}

			} else {
				domain.DevidedEqually(child)
				Nearparent, NearChild := domain.NearestSearch(nil, child)
				if result, _, err := r.searchEncyrpt(ctx, Nearparent, NearChild, inEncryptResults); err != nil {
					return nil, false, err
				} else {
					return result, false, nil
				}

			}
		} else {
			return nil, false, err
		}
	} else {
		domain.Stringify(parent, 0)
		removeSuccessStatus := domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
		fmt.Println("->remove",removeSuccessStatus)
		domain.Stringify(parent, 0)
		if !removeSuccessStatus {
			if parent.Max == child.Max && parent.Min == child.Min {
				return nil, true, nil
			}
		}
	}
	return nil, false, nil
}
