package repository

//TODO wire this to controller and repo

import (
	"sync"
	"github.com/google/uuid"
	"github.com/milehighguy-dev/receipt-cruncher-go/pkg/model"
)

type ReceiptsRepository struct {
	mu       sync.RWMutex
	receipts map[uuid.UUID]model.Receipt
}

func NewReceiptsRepository() *ReceiptsRepository {
	return &ReceiptsRepository{
		receipts: make(map[uuid.UUID]model.Receipt),
	}
}

func (r *ReceiptsRepository) Save(receipt model.Receipt) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.receipts[receipt.ID] = receipt
}

func (r *ReceiptsRepository) Get(id uuid.UUID) (model.Receipt, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	receipt, exists := r.receipts[id]
	return receipt, exists
}
