package services

import (

	"github.com/harrycoa/apiUsuarios/entities"
	"github.com/harrycoa/apiUsuarios/repository"
)

type UsuarioService interface {
	FindAll() [] entities.Usuario
	FindOne(id string) (entities.Usuario, error)
	Save(entities.Usuario) entities.Usuario
	Update(id string, entities.Usuario) entities.Usuario
}

type usuarioService struct {
	repository repository.UsuarioRepository
}

func NewUsuarioService() UsuarioService{
	return &usuarioService{repository: repository.NewUsuarioRepository() }
}

func (s *usuarioService) FindAll() []entities.Usuario{
	return s.repository.FindAll()

}
func (s *usuarioService) FindOne(id string) (entities.Usuario, error){
	return s.repository.FindOne(id)
}
func (s *usuarioService) Save(usuario entities.Usuario) entities.Usuario{
	return s.repository.Save(usuario)
}
func (s *usuarioService) Update(id string, usuario entities.Usuario) entities.Usuario{
	return s.repository.Update(id, usuario)
}