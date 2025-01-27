package service

//TODO finish this

import (
	"github.com/google/uuid"
	"github.com/milehighguy-dev/receipt-cruncher-go/pkg/repository"
	"github.com/milehighguy-dev/receipt-cruncher-go/pkg/model"
)

var receiptsRepository repository.ReceiptsRepository

func SaveReceipt(receipt model.Receipt) (model.ProcessResponse, error) {
	err := receiptsRepository.Save(receipt)
	if err != nil {
		return ProcessResponse{}, err
	}
	return ProcessResponse{ID: receipt.ID}, nil
}

func GetReceipt(id uuid.UUID) (model.Receipt, error) {
	return receiptsRepository.GetByID(id)
}
func DeleteReceipt(id uuid.UUID) error {
	return receiptsRepository.Delete(id)
}

func UpdateReceipt(receipt model.Receipt) error {
	return receiptsRepository.Update(receipt)
}