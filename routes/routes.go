package routes

import (
	"github.com/bdeolive/gin-api-rest/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.BuscarAlunos)
	r.GET("/alunos/:id", controllers.BuscarAlunoId)
	r.GET("/alunos/cpf/:cpf", controllers.BuscarAlunoCPF)
	r.GET("/:nome", controllers.Saudacao)
	r.POST("/alunos", controllers.InserirAluno)
	r.PATCH("/alunos/:id", controllers.EditarAluno)
	r.DELETE("/alunos/:id", controllers.DeletarAluno)

	r.Run(":5000")
}
