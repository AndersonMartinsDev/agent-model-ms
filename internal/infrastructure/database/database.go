package database

import (
	"database/sql"
	"fmt"
	"log"
	"os" // Adicione para lidar com sinais de interrupção

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lib/pq"
)

var PostgreSQL *sql.DB

func conectar() {
	psqlconn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)

	db, erro := sql.Open("postgres", psqlconn)
	if erro != nil {
		panic(fmt.Sprintf("Não foi possível abrir a conexão com o banco de dados: %v", erro))
	}

	if erro = db.Ping(); erro != nil {
		db.Close() // Fecha a conexão que acabou de ser aberta se o ping falhar
		panic(fmt.Sprintf("Não foi possível conectar ao banco de dados: %v", erro))
	}

	db.SetMaxOpenConns(25) // Número máximo de conexões abertas
	db.SetMaxIdleConns(10) // Número máximo de conexões ociosas
	PostgreSQL = db
	log.Println("Conexão principal com o banco de dados estabelecida com sucesso.")
}

// conectarAdminDB conecta ao PostgreSQL usando um banco de dados padrão (como 'postgres')
func conectarAdminDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		DB_HOST, DB_PORT, DB_USER, DB_PASSWORD) // Conectando a 'postgres' ou 'template1'
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com o DB admin: %w", err)
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("erro ao pingar DB admin: %w", err)
	}
	return db, nil
}

// GetConnectionDatabase retorna a conexão principal com o banco de dados.
func GetConnectionDatabase() *sql.DB {
	if PostgreSQL == nil {
		conectar()
	}
	return PostgreSQL
}

// CloseConnectionDatabase fecha a conexão principal com o banco de dados.
func CloseConnectionDatabase() {
	if PostgreSQL != nil {
		err := PostgreSQL.Close()
		if err != nil {
			log.Printf("Erro ao fechar a conexão do banco de dados: %v", err)
		} else {
			log.Println("Conexão principal do banco de dados fechada com sucesso.")
		}
	}
}

// InitalStrucuture configura a estrutura inicial do banco de dados,
func InitalStrucuture() {
	log.Println("Iniciando a verificação e criação da estrutura do banco de dados...")
	bancoAdmin, erro := conectarAdminDB() // Renomeado para maior clareza
	if erro != nil {
		panic(fmt.Sprintf("Problema ao conectar ao banco de dados admin: %v", erro))
	}
	defer bancoAdmin.Close() // Garante que a conexão admin seja fechada

	// Tenta criar o banco de dados.
	_, err := bancoAdmin.Exec(fmt.Sprintf("CREATE DATABASE %s", DB_NAME))
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "42P04" {
			log.Printf("Banco de dados '%s' já existe, prosseguindo.", DB_NAME)
		} else {
			panic(fmt.Sprintf("Erro ao criar banco de dados '%s': %v", DB_NAME, err))
		}
	} else {
		log.Printf("Banco de dados '%s' criado com sucesso.", DB_NAME)
	}

	// Conecta-se ao banco de dados recém-criado (ou existente) para aplicar as migrações.
	conectar() // Isso vai inicializar a variável global PostgreSQL

	// 4. Configura o driver do migrate com a conexão principal
	// Use PostgreSQL diretamente aqui, pois 'conectar()' já o inicializou.
	config := &postgres.Config{}
	driver, err := postgres.WithInstance(PostgreSQL, config)
	if err != nil {
		log.Fatalf("Erro ao configurar driver de migração: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/api/migrations", // Caminho para suas migrações
		DB_NAME,                     // Nome do banco de dados
		driver,
	)

	if err != nil {
		log.Fatalf("Erro ao criar instância de migração: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Erro ao aplicar migrações: %v", err)
	}

	log.Println("Migrações aplicadas com sucesso!")
}

// SetupDatabase deve ser chamada no início da sua aplicação (e.g., na main).
func SetupDatabase() {
	InitalStrucuture()
	// Configurar um listener para fechar a conexão no encerramento da aplicação
	// Isso é útil para lidar com Ctrl+C ou outros sinais de interrupção
	go func() {
		sigChan := make(chan os.Signal, 1)
		// signal.Notify(sigChan, os.Interrupt, os.Kill) // Adicione "os/signal" import
		<-sigChan
		CloseConnectionDatabase()
		os.Exit(0)
	}()
}
