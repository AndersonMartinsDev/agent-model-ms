CREATE TABLE tb_models (
    id SERIAL PRIMARY KEY,
    perfil_name VARCHAR(100),
    instructions TEXT
);

CREATE TABLE IF NOT EXISTS tb_agents (
  id SERIAL PRIMARY KEY,
  agent_name VARCHAR(255) NOT NULL,
  company_name VARCHAR(255),
  company_description TEXT,
  agent_behaviour TEXT,
  company_url VARCHAR(255),
  phone_number VARCHAR(20) UNIQUE,
  model_id INTEGER REFERENCES tb_models(id),
  user_uuid UUID
);