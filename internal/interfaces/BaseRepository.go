package interfaces

type Repository[T any] interface {
	Create(entity T) error
	GetAll() ([]T, error)
	GetByID(id uint) (*T, error)
	Update(id uint, entity T) error
	Delete(id uint) error
}
