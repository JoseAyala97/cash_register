package repositories

import "gorm.io/gorm"

// GenericRepository es una implementación genérica de la interfaz Repository
type GenericRepository[T any] struct {
	db *gorm.DB
}

// NewGenericRepository es un constructor para el repositorio genérico
func NewGenericRepository[T any](db *gorm.DB) *GenericRepository[T] {
	return &GenericRepository[T]{db: db}
}

// Crear una nueva entidad
func (r *GenericRepository[T]) Create(entity T) error {
	return r.db.Create(&entity).Error
}

// Obtener todas las entidades
func (r *GenericRepository[T]) GetAll() ([]T, error) {
	var entities []T
	if err := r.db.Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

// Obtener una entidad por ID
func (r *GenericRepository[T]) GetByID(id uint) (*T, error) {
	var entity T
	if err := r.db.First(&entity, id).Error; err != nil {
		return nil, err
	}
	return &entity, nil
}

// Actualizar una entidad
func (r *GenericRepository[T]) Update(id uint, updatedEntity T) error {
	return r.db.Model(new(T)).Where("id = ?", id).Updates(updatedEntity).Error
}

// Eliminar una entidad
func (r *GenericRepository[T]) Delete(id uint) error {
	return r.db.Delete(new(T), id).Error
}
