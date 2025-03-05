package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Name  string
    Email string
}

func main() {
    // Configuração da conexão com o MySQL
    dsn := "root:root@tcp(127.0.0.1:3306)/sakila?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Falha ao conectar ao banco de dados")
    }

    // Migração do schema
    db.AutoMigrate(&User{})

    // Configuração do Gin
    r := gin.Default()

    // Rota para criar um novo usuário
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }

        db.Create(&user)
        c.JSON(200, user)
    })

    // Rota para listar todos os usuários
    r.GET("/users", func(c *gin.Context) {
        var users []User
        db.Find(&users)
        c.JSON(200, users)
    })

    // Iniciar o servidor
    r.Run()
}