package usecase

import (
	"context"
	"fmt"
	"serch_encrypt/domain"
	"strings"

	"time"
)

type searchEncryptUseCase struct {
	customerRepository domain.SearchEncryptRepository
	contexTimeout      time.Duration
}

func NewSearchEncryptUseCase(contextTimeout time.Duration, customerRepository domain.SearchEncryptRepository) domain.SearchEncryptUseCase {
	return &searchEncryptUseCase{
		customerRepository: customerRepository,
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
	// make root
	root := domain.NewRoot(minId, maxId)
	// save invalid model encrypt
	var InvalidModelEncrypt []int
	// cari yang terdekat di kiri
	
	for {
		parent, child := domain.NearestSearch(nil, root.Root)
		if child == nil{
			break
		}
		if child != nil {
			if findEncrypt, err := r.SearchEncyrpt(ctx, parent, child); err != nil {
				return 0,0,err
			}else{
				if findEncrypt != nil{
					InvalidModelEncrypt = append(InvalidModelEncrypt,*findEncrypt )
				}
			}
		}
		domain.Stringify(root.Root, 0)
		fmt.Println("-> find and save", InvalidModelEncrypt)
	}


	return
}

func (r searchEncryptUseCase) SearchEncyrpt(ctx context.Context, parent *domain.Node, child *domain.Node) (*int, error) {
	// encrypt customer
	var modelEncrypt []domain.EncryptCustomer
	fmt.Println("-> child", child, child.Max-child.Min)
	if err := r.customerRepository.FetchByRange(ctx, &modelEncrypt, child.Min, child.Max); err != nil {
		fmt.Println("->modelEncrypt",modelEncrypt)
		if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
			//min dan max id jaraknya 1
			if (child.Max - child.Min) <= 1 {
				fmt.Println("-> masuk child")
				var encyrptCustomer domain.EncryptCustomer
				if err := r.customerRepository.FindById(ctx, &encyrptCustomer, child.Min); err != nil {
					if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
						//simpan invalid
						fmt.Println("->", child.Min,parent.Left,parent.Right)
						domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
						return &child.Min, nil
					} else {
						return nil, err
					}
				}

				if err := r.customerRepository.FindById(ctx, &encyrptCustomer, child.Max); err != nil {
					if strings.Contains(err.Error(), "SCP Encrypt Fuction") {
						//simpan invalid
						fmt.Println("->", child.Max,parent.Left,parent.Right)
						domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
						return &child.Max, nil
					} else {
						return nil, err
					}
				}

				// remove sesuai min max
				

			} else {
				// dibagi dua
				domain.DevidedEqually(child)
				Nearparent, NearChild := domain.NearestSearch(nil, child)
				if result, err := r.SearchEncyrpt(ctx, Nearparent, NearChild); err != nil {
					return nil, err
				} else {
					return result, nil
				}

			}
			//stringify
			// domain.Stringify(parent, 0)
		} else {
			return nil, err
		}
	} else {
		fmt.Println("-> min max", parent, child.Min, child.Max)
		removeSuccessStatus := domain.RemoveEqualByMinMax(parent, child, child.Min, child.Max)
		fmt.Println("-> remove status", removeSuccessStatus)
		if !removeSuccessStatus {
			// if parent.Max == child.Max && parent.Min == child.Min {
			// 	parent = nil
			// }
			if parent.Left == nil && parent.Right == nil{
				return nil, nil
			}	
		}
		Parent1, child1 := domain.NearestSearch(nil, parent)
		if result, err := r.SearchEncyrpt(ctx, Parent1, child1); err != nil {
			return nil, err
		} else {
			return result, nil
		}
	}
	return nil, nil
}
