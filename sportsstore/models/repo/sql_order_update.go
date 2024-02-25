package repo

import "sportsstore/models"

func (repo *SqlRepository) SetOrderShipped(o *models.Order) {
	result, err := repo.Commands.UpdateOrder.ExecContext(repo.Context, o.Shipped, o.ID)
	if err != nil {
		repo.Logger.Panicf("Cannot exec UpdateOrder command: %v", err.Error())
		return
	}
	rows, err := result.RowsAffected()
	if err != nil {
		repo.Logger.Panicf("Cannot get updated ID: %v", err.Error())
		return
	}
	if rows != 1 {
		repo.Logger.Panicf("Got unexpected rows affected: %v", err.Error())
	}
}
