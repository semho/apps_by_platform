package repo

import "sportsstore/models"

func (repo *SqlRepository) SaveProduct(p *models.Product) {
	if p.ID == 0 {
		result, err := repo.Commands.SaveProduct.ExecContext(repo.Context, p.Name, p.Description, p.Category.ID,
			p.Price)
		if err != nil {
			repo.Logger.Panicf("Cannot exec SaveProduct command: %v", err.Error())
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			repo.Logger.Panicf("Cannot get inserted ID: %v", err.Error())
			return
		}
		p.ID = int(id)
		return
	} else {
		result, err := repo.Commands.UpdateProduct.ExecContext(repo.Context, p.Name, p.Description, p.Category.ID,
			p.Price, p.ID)
		if err != nil {
			repo.Logger.Panicf("Cannot exec Update command: %v", err.Error())
			return
		}
		affected, err := result.RowsAffected()
		if err != nil {
			repo.Logger.Panicf("Cannot get rows affected: %v", err)
			return
		}
		if affected != 1 {
			repo.Logger.Panicf("Cannot unexpected row affected: %v", affected)
		}
	}
}
