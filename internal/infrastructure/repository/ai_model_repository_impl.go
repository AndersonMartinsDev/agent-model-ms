package repository

import (
	"agent-model-ms/internal/domain/model"
	"agent-model-ms/internal/infrastructure/database"
	"database/sql"
)

type AIModelRepositoryImpl struct {
	*sql.DB
	Table string
}

func NewAIModelRepository() AIModelRepositoryImpl {
	return AIModelRepositoryImpl{}
}

func (repo AIModelRepositoryImpl) CreateModel(model model.AIModel) error {
	database := database.GetConnectionDatabase()
	insert := "Insert Into tb_models(perfil_name, instructions) values($1, $2)"

	statement, err := database.Prepare(insert)
	if err != nil {
		return err
	}

	_, erro := statement.Exec(model.PerFilName, model.Instructions)

	defer statement.Close()
	return erro

}
func (repo AIModelRepositoryImpl) GetModel(model_id string) (model.AIModel, error) {
	database := database.GetConnectionDatabase()
	insert := `
		SELECT 
			id,
			perfil_name
		FROM 
			tb_models md
		WHERE 
			md.id = $1
	`

	row, err := database.Query(insert, model_id)
	if err != nil {
		return model.AIModel{}, err
	}

	var modelAi model.AIModel
	if row.Next() {
		if erro := row.Scan(
			&modelAi.ID,
			&modelAi.PerFilName,
		); erro != nil {
			return model.AIModel{}, erro
		}
	}

	defer row.Close()
	return modelAi, nil

}
func (repo AIModelRepositoryImpl) GetModelList() ([]model.AIModel, error) {
	database := database.GetConnectionDatabase()
	insert := `
		SELECT 
			id,
			perfil_name
		FROM 
			tb_models
	`

	row, err := database.Query(insert)
	if err != nil {
		return []model.AIModel{}, err
	}

	var modelsAi []model.AIModel

	for row.Next() {
		var modelAi model.AIModel
		if erro := row.Scan(
			&modelAi.ID,
			&modelAi.PerFilName,
		); erro != nil {
			return []model.AIModel{}, erro
		}
		modelsAi = append(modelsAi, modelAi)
	}

	defer row.Close()
	return modelsAi, nil

}
