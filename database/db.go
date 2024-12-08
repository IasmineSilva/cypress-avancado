package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")

	// Validação básica
	if endereco == "" || usuario == "" || senha == "" || nomeBanco == "" || portaBanco == "" {
		log.Panic("Erro: Uma ou mais variáveis de ambiente do banco de dados estão ausentes")
	}

	stringDeConexao := "host=" + endereco + " user=" + usuario + " password=" + senha + " dbname=" + nomeBanco + " port=" + portaBanco + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados: ", err)
	}

	err = DB.AutoMigrate(&models.Aluno{})
	if err != nil {
		log.Panic("Erro ao fazer auto-migrate: ", err)
	}
}
