package configuration

import (
	"agent-model-ms/internal/infrastructure/commons/logger"
	"agent-model-ms/internal/infrastructure/database"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

var (
	Porta         = 8080
	NodeClientURL = ""
)

func LoadEnv() {
	// para debug local
	if erro := godotenv.Load(); erro != nil {
		panic("Error ao carregar as variáveis de ambiente!")
	}
	slog.Info("Variáveis de ambiente carregadas com sucesso!")
}

func LoadLogger() {
	custom_log := slog.New(logger.NewHandler(nil))
	slog.SetDefault(custom_log)
	slog.Info("Logger Carregado com sucesso!")
}

func LoadDatabase() {
	db_name, _ := os.LookupEnv("DATABASE_NAME")
	db_host := os.Getenv("DATABASE_HOST")
	db_port := os.Getenv("DATABASE_PORT")
	db_user := GetSecret("DATABASE_USER")
	db_password := GetSecret("DATABASE_PASSWORD")

	database.SetDatabaseEnv(db_name, db_host, db_port, db_user, db_password)
	if erro := database.GetConnectionDatabase().Ping(); erro != nil {
		panic(fmt.Sprintf("Não foi possivel conectar-se com o banco de dados: %s", erro))
	}

	slog.Info("Conexão com Banco de dados Estabelecida")

	database.InitalStrucuture()
}
