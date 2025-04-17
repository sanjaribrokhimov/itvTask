package models

import (
	"time"

	"gorm.io/gorm"
)

// Movie represents a movie in the database
type Movie struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"size:255;not null"`
	Director  string         `json:"director" gorm:"size:255;not null"`
	Year      int            `json:"year" gorm:"not null"`
	Plot      string         `json:"plot" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// MovieRepository defines the interface for movie operations
type MovieRepository interface {
	Create(movie *Movie) error
	GetAll() ([]Movie, error)
	GetByID(id uint) (*Movie, error)
	Update(movie *Movie) error
	Delete(id uint) error
	SearchByTitle(title string) ([]Movie, error)
}

type movieRepository struct {
	db *gorm.DB
}

func NewMovieRepository(db *gorm.DB) MovieRepository {
	return &movieRepository{db: db}
}

func (r *movieRepository) Create(movie *Movie) error {
	return r.db.Create(movie).Error
}

func (r *movieRepository) GetAll() ([]Movie, error) {
	var movies []Movie
	err := r.db.Find(&movies).Error
	return movies, err
}

func (r *movieRepository) GetByID(id uint) (*Movie, error) {
	var movie Movie
	err := r.db.First(&movie, id).Error
	return &movie, err
}

func (r *movieRepository) Update(movie *Movie) error {
	return r.db.Save(movie).Error
}

func (r *movieRepository) Delete(id uint) error {
	return r.db.Delete(&Movie{}, id).Error
}

func (r *movieRepository) SearchByTitle(title string) ([]Movie, error) {
	var movies []Movie
	err := r.db.Where("title LIKE ?", "%"+title+"%").Find(&movies).Error
	return movies, err
}
