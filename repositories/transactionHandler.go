package repositories

import (
	"context"
	"database/sql"
)

func BeginTransaction(repoMgr *RepositoryManager) error {
	ctx := context.Background()
	transaction, err := repoMgr.CategoryRepository.dbHandler.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	repoMgr.CategoryRepository.transaction = transaction

	return nil
}

func RollbackTransaction(repoManager *RepositoryManager) error {
	transaction := repoManager.CategoryRepository.transaction

	repoManager.CategoryRepository.transaction = nil

	return transaction.Rollback()
}

func CommitTransaction(repoManager *RepositoryManager) error {
	transaction := repoManager.CategoryRepository.transaction

	repoManager.CategoryRepository.transaction = nil

	return transaction.Commit()
}
