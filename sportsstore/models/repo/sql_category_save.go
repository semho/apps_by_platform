package repo

import "sportsstore/models"

func (repo *SqlRepository) SaveCategory(c *models.Category) {
	if c.ID == 0 {
		result, err := repo.Commands.SaveCategory.ExecContext(repo.Context, c.CategoryName)
		if err != nil {
			repo.Logger.Panicf("Cannot exec SaveCategory command: %v", err.Error())
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			repo.Logger.Panicf("Cannot get inserted ID: %v", err.Error())
			return
		}
		c.ID = int(id)
		return
	} else {
		result, err := repo.Commands.UpdateCategory.ExecContext(repo.Context, c.CategoryName, c.ID)
		if err != nil {
			repo.Logger.Panicf("Cannot exec UpdateCategory command: %v", err.Error())
			return
		}
		affected, err := result.RowsAffected()
		if err != nil {
			repo.Logger.Panicf("Cannot get rows affected: %v", err)
			return
		}
		if affected != 1 {
			repo.Logger.Panicf("Got unexpected rows affected: %v", err)
		}
	}
}
