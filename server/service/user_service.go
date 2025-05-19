package service

import (
	"errors"
	"golang-tutorial/contract"
	"golang-tutorial/dto"
	"golang-tutorial/entity"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepository contract.UserRepository
}

func implUserService(repo *contract.Repository) contract.UserService {
	return &UserService{
		UserRepository: repo.User,
	}
}

func (s *UserService) GetUser(id int) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetUser(id)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		StatusCode: http.StatusOK,
		Message:    "Berhasil mendapatkan data",
		Data: dto.UserData{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

func (s *UserService) Register(payload *dto.UserRequest) (*dto.UserResponse, error) {
	if !isValidEmail(payload.Email) {
		return nil, errors.New("email tidak valid")
	}

	if !isValidPassowrd(payload.Password) {
		return nil, errors.New("Password harus mengandung 1 uppercase, 1 angka, 1 simbol")
	}

	emailExists, err := s.UserRepository.CheckEmail(payload.Email)
	if err != nil {
		return nil, err
	}
	if emailExists {
		return nil, errors.New("email sudah terdaftar")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &entity.User{
		Email:    payload.Email,
		Password: string(hashPassword),
	}

	err = s.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponse{
		StatusCode: http.StatusCreated,
		Message:    "Akun berhasil dibuat",
		Data: dto.UserData{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}
	return response, nil
}

func isValidEmail(email string) bool {
	regex := `^[a-zA-Z0-9._%+-]+@unity\.com$`
	re := regexp.MustCompile(regex)
	return re.MatchString(email)
}
func isValidPassowrd(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSymbol := regexp.MustCompile(`[!@#\$%\^&\*\(\)_\+\-=\[\]\{\};:'"\\|,.<>\/?]`).MatchString(password)

	return hasUpper && hasDigit && hasSymbol
}

func (s *UserService) Login(payload *dto.UserRequest) (*dto.UserResponse, error) {
	user, err := s.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		return nil, errors.New("password salah")
	}

	response := &dto.UserResponse{
		StatusCode: http.StatusOK,
		Message:    "Login berhasil",
		Data: dto.UserData{
			ID:        user.ID,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		},
	}

	return response, nil
}

