package repository

import (
	"errors"
	"github.com/harrycoa/apiUsuarios/entities"
	"strconv"
)

type UsuarioRepository interface {
	FindAll() [] entities.Usuario
	FindOne(id string) (entities.Usuario, error)
	Save(entities.Usuario) entities.Usuario
	Update(id string, entities.Usuario) entities.Usuario
}

type usuarioRepository struct {
	usuarios []entities.Usuario
}

func NewUsuarioRepository() UsuarioRepository{
	usuario1 := entities.Usuario {
		Id:	1,
		Nombre: "carlos",
		Apellido: "sanc",
		Edad: "40",
		Estado: true,
	}

	usuario2 := entities.Usuario {
		Id:	2,
		Nombre: "carlos2",
		Apellido: "sanc2",
		Edad: "40",
		Estado: true,
	}

	usuario3 := entities.Usuario {
		Id:	1,
		Nombre: "carlos3",
		Apellido: "sanc3",
		Edad: "40",
		Estado: true,
	}

	usuarios := []entities.Usuario{}
	usuarios = append(usuarios, usuario1)
	usuarios = append(usuarios, usuario2)
	usuarios = append(usuarios, usuario3)

	return &usuarioRepository{usuarios: usuarios}
}

func (repo *usuarioRepository) FindAll() [] entities.Usuario{
	return repo.usuarios

}
func (repo *usuarioRepository) FindOne(id string) (entities.Usuario , error) {
	idEntero, _ := strconv.Atoi(id)
	for _, usuario := range repo.usuarios {
		if idEntero == usuario.Id {
			return usuario, nil
		}
	}
	return entities.Usuario{}, errors.New("No se encontro usuario")
}
func (repo *usuarioRepository) Save(usuario entities.Usuario) entities.Usuario{
	repo.usuarios = append(repo.usuarios, usuario)
	return usuario
}
func (repo *usuarioRepository) Update(id string, usuario entities.Usuario) entities.Usuario{
	idEntero, _ := strconv.Atoi(id)
	usuariosFinal := []entities.Usuario
	for _, usuarioOriginal := range repo.usuarios {
		if idEntero == usuario.Id {
			usuariosFinal = append(usuariosFinal, usuario)
		} else {
			usuariosFinal = append(usuariosFinal, usuarioOriginal)
		}
	}
	repo.usuarios = usuariosFinal
	return usuario
}