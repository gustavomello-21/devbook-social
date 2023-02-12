package controllers

import (
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário criado"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Retornando todos os usuários"))
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário selecionado"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário atualizado"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Usuário Deletado"))
}
