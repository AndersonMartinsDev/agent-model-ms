package repository

import (
	"agent-model-ms/internal/domain/model"
	"agent-model-ms/internal/infrastructure/database"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
)

type AIAgentRepositoryImpl struct {
	*sql.DB
}

func NewAIAgentRepository() AIAgentRepositoryImpl {
	return AIAgentRepositoryImpl{
		DB: database.GetConnectionDatabase(),
	}
}

func (repo AIAgentRepositoryImpl) Createagent(agent model.AIAgent) error {
	database := repo.DB
	insert := `
	Insert INTO tb_agents(
		agent_name, 
		company_name, 
		company_description, 
		agent_behaviour, 
		company_url,
		model_id,
		user_uuid
		)
	VALUES($1, $2, $3, $4, $5, $6, $7)`

	statement, err := database.Prepare(insert)
	if err != nil {
		return err
	}

	_, erro := statement.Exec(agent.Name, agent.CompanyName, agent.CompanyDescription, agent.BehaviourIa, agent.CompanyUrl, agent.ModelId, uuid.MustParse(agent.UUID_user))

	defer statement.Close()
	return erro
}
func (repo AIAgentRepositoryImpl) UpdateAgent(agent model.AIAgent) error {
	database := repo.DB
	query := `
		UPDATE tb_agents
		SET
			agent_name = $1,
			model_id = $2,
			company_name = $3,
			company_description = $4,
			agent_behaviour = $5,
			company_url = $6,
			user_uuid = $7
		WHERE
			id = $8;`

	// Executa a query, passando os campos do struct como parâmetros
	_, err := database.Exec(query,
		agent.Name,
		agent.ModelId,
		agent.CompanyName,
		agent.CompanyDescription,
		agent.BehaviourIa,
		agent.CompanyUrl,
		agent.UUID_user,
		agent.Id, // O ID do registro que será atualizado
	)

	if err != nil {
		return fmt.Errorf("erro ao atualizar chatagent: %w", err)
	}
	return nil
}
func (repo AIAgentRepositoryImpl) GetAgent(agent_id uint) (model.AIAgent, error) {
	database := repo.DB

	query := fmt.Sprintf(`
		SELECT 
			b.agent_name,
			b.model_id,
			b.company_name,
			b.company_description,
			b.agent_behaviour,
			b.company_url 
		FROM 
			tb_agents b 
		WHERE 
			b.id = %d`, agent_id)

	row, erro := database.Query(query)
	if erro != nil {
		return model.AIAgent{}, erro
	}

	var agent model.AIAgent
	if row.Next() {
		if erro := row.Scan(
			&agent.Name,
			&agent.ModelId,
			&agent.CompanyName,
			&agent.CompanyDescription,
			&agent.BehaviourIa,
			&agent.CompanyUrl,
		); erro != nil {
			return model.AIAgent{}, erro
		}
	}

	if agent.Name == "" {
		return model.AIAgent{}, fmt.Errorf("agent with id=%d not exist", agent_id)
	}

	return agent, nil
}
func (repo AIAgentRepositoryImpl) List(user_uuid string) ([]model.AIAgent, error) {

	database := repo.DB
	query := `
		SELECT
			b.id,
			b.agent_name,
			b.model_id,
			b.company_name,
			b.company_description,
			b.agent_behaviour,
			b.company_url
		FROM
			tb_agents b
		WHERE 
			b.user_uuid = $1
			`

	rows, err := database.Query(query, user_uuid)
	if err != nil {
		log.Printf("Erro ao executar a query para listar agents do usuário %s: %v", user_uuid, err)
		return nil, fmt.Errorf("erro ao listar agents: %w", err)
	}
	defer rows.Close()

	var agents []model.AIAgent
	for rows.Next() {
		var agent model.AIAgent
		if err := rows.Scan(
			&agent.Id,
			&agent.Name,
			&agent.ModelId,
			&agent.CompanyName,
			&agent.CompanyDescription,
			&agent.BehaviourIa,
			&agent.CompanyUrl,
		); err != nil {
			return nil, fmt.Errorf("erro ao escanear agent: %w", err)
		}
		agents = append(agents, agent)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Erro durante a iteração das linhas para o usuário %s: %v", user_uuid, err)
		return nil, fmt.Errorf("erro ao iterar resultados: %w", err)
	}

	log.Printf("Listagem de %d agents para o usuário %s realizada com sucesso.", len(agents), user_uuid)
	return agents, nil

}
func (repo AIAgentRepositoryImpl) Delete(agent_id string) error {

	database := repo.DB
	query := `
		Delete
		FROM
			tb_agents b
		WHERE 
			b.id = $1
			`
	result, err := database.Exec(query, agent_id)
	if err != nil {
		log.Printf("Erro ao executar delete do usuário %s: %v", result, err)
		return nil
	}

	return nil

}
func (repository AIAgentRepositoryImpl) GetAgentAndPerfil(agent_id uint64) (model.AIAgent, error) {
	database := database.GetConnectionDatabase()

	query := fmt.Sprintf(`
		SELECT
			b.agent_name,
			b.company_name,
			b.company_description,
			b.agent_behaviour,
			b.company_url,
			mi.instructions
		FROM
			agents b
		JOIN
			tb_models AS mi ON b.model_id = mi.id
		WHERE
			b.id = '%d';
	`, agent_id)

	rows, erro := database.Query(query)
	if erro != nil {
		return model.AIAgent{}, erro
	}
	defer rows.Close()

	var agent model.AIAgent

	if rows.Next() {
		if erro := rows.Scan(
			&agent.Name,
			&agent.CompanyName,
			&agent.CompanyDescription,
			&agent.BehaviourIa,
			&agent.CompanyUrl,
			&agent.Instructions,
		); erro != nil {
			return model.AIAgent{}, erro
		}
	}
	return agent, nil
}
func (repository AIAgentRepositoryImpl) GetAgentFromPhone(phoneNumber string) (uint, error) {

	database := database.GetConnectionDatabase()

	query := fmt.Sprintf(`
		SELECT
			b.id
		FROM
			tb_agents b
		WHERE
			b.phone_number = '%s';
	`, phoneNumber)

	rows, erro := database.Query(query)
	if erro != nil {
		return 0, erro
	}
	defer rows.Close()

	var agentId uint

	if rows.Next() {
		if erro := rows.Scan(
			&agentId,
		); erro != nil {
			return 0, erro
		}
	}
	return agentId, nil
}
