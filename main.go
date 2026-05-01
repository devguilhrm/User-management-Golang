package main

import (
	"fmt"
	"log"
	"user-management/db"
	"user-management/models"
	"user-management/repository"
	"user-management/services"
	"user-management/utils"
)

func main() {

	database := db.ConectDB("db/users.db")
	err := db.CreateTable(database)
	if err != nil {
		log.Fatal(err)
	}

	defer database.Close()

	for {
		fmt.Println("------- Cadastro de Usuários -------")
		fmt.Println("(1) Criar Usuário")
		fmt.Println("(2) Ver Usuários Cadastrados")
		fmt.Println("(3) Atualizar Usuário")
		fmt.Println("(4) Deletar Usuário")
		fmt.Println("(5) Sair")

		var input int

		fmt.Print("Escolha a opção: ")
		fmt.Scanln(&input)

		switch input {
		case 1:
			var (
				nome  string = utils.ReadInputStr("Digite o Nome: ")
				email string = utils.ReadInputStr("Digite o Email: ")
				idade int    = utils.ReadInputInt("Digite a idade: ")
			)

			user := *services.NewUser(nome, email, idade)
			err := repository.InsertUser(database, &user)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Usuário criado e cadastrado com sucesso!")

		case 2:
			users, err := repository.ReadUser(database)
			if err != nil {
				log.Fatal(err)
			}
			// Conta os Usuários dentro do Slice users
			var contUsers int = 0
			for _, user := range users {
				contUsers++
				fmt.Printf("Usuário: (%s) | Email: (%s) | ID: (%s) | Idade: (%d) \n", user.Name, user.Email, user.ID, user.Age)
			}
			fmt.Println("O total de:", contUsers, "Cadastrados")

		case 3:
			var userInputName string = utils.ReadInputStr("Digite o nome completo do usuário para a busca: ")

			rows, err := database.Query("SELECT id, name FROM users")
			if err != nil {
				log.Fatal(err)
			}

			defer rows.Close()

			var foundID, foundName string
			userFound := false
			for rows.Next() {
				var id, name string
				err := rows.Scan(&id, &name)
				if err != nil {
					log.Fatal(err)
				}
				if userInputName == name {
					foundID = id
					foundName = name
					userFound = true
					// Fechar o processo manualmente para evitar o "database is locked"
					rows.Close()
					break
				}
			}

			if !userFound {
				fmt.Println("Nenhum usuário encontrado com esse nome!")
				break
			}

			fmt.Printf("\n👤 Usuário encontrado!\nNome atual: %s (ID: %s)\n", foundName, foundID)

			newName := utils.ReadInputStr("Novo nome (deixe em branco para não alterar): ")

			newEmail := utils.ReadInputStr("Novo e-mail (deixe em branco para não alterar): ")

			var newAge int = utils.ReadInputInt("Nova idade (deixe com o valor 0 para não alterar): ")

			currentUser := models.User{ID: foundID, Name: foundName}
			currentDataRow := database.QueryRow("SELECT email, age FROM users WHERE id = ?", foundID)
			err = currentDataRow.Scan(&currentUser.Email, &currentUser.Age)
			if err != nil {
				log.Fatal(err)
			}

			if newName != "" {
				currentUser.Name = newName
			}
			if newEmail != "" {
				currentUser.Email = newEmail
			}
			if newAge != 0 {
				currentUser.Age = newAge
			}

			// Salva atualizações no BD
			err = repository.UpdateUser(database, &currentUser)
			if err != nil {
				fmt.Println("Erro ao atualizar o usuário:", err)
			} else {
				fmt.Println("Usuário atualizado com sucesso!")
			}

		case 4:
			var nameToDelete string = utils.ReadInputStr("Digite o nome para Deletar: ")

			rows, err := database.Query("SELECT name FROM users")
			if err != nil {
				log.Fatal(err)
			}
			defer rows.Close()

			userFound := false
			for rows.Next() {
				var name string
				err := rows.Scan(&name)
				if err != nil {
					log.Fatal(err)
				}
				if nameToDelete == name {
					userFound = true
					rows.Close()
					break
				}
			}
			if !userFound {
				fmt.Printf("Usuário %s Não encontrado! \n", nameToDelete)
				break
			}

			fmt.Println("Usuário encontrado!")
			var deleteQuestion string = utils.ReadInputStr_oneF("Deletar o Usuário %s (s) (n): \n", nameToDelete)

			if deleteQuestion == "s" {
				err := repository.DeleteUser(database, nameToDelete)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Usuário Deletado com sucesso!")
			} else if deleteQuestion == "n" {
				fmt.Println("O usuário não foi deletado!")
			} else {
				fmt.Println("Opção inválida!")
			}

		case 5:
			fmt.Println("Você saiu do programa!")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
