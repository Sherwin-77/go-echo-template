package service

import (
	"context"
	"encoding/json"
	"github.com/sherwin-77/go-echo-template/pkg/query"
	"github.com/sherwin-77/go-echo-template/pkg/response"
	"net/http"
	"net/url"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/sherwin-77/go-echo-template/internal/entity"
	"github.com/sherwin-77/go-echo-template/internal/http/dto"
	"github.com/sherwin-77/go-echo-template/internal/repository"
	"github.com/sherwin-77/go-echo-template/pkg/caches"
	"github.com/sherwin-77/go-echo-template/pkg/tokens"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService interface {
	GetUsers(ctx context.Context, queryParams url.Values) ([]entity.User, *response.Meta, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, request dto.UserRequest) (*entity.User, error)
	UpdateUser(ctx context.Context, request dto.UpdateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	Login(ctx context.Context, request dto.LoginRequest) (string, error)
	Register(ctx context.Context, request dto.UserRequest) (*entity.User, bool, error)
	ChangeRole(ctx context.Context, request dto.ChangeRoleRequest) error
}

type userService struct {
	tokenService   tokens.TokenService
	userRepository repository.UserRepository
	roleRepository repository.RoleRepository
	userBuilder    query.Builder
	cache          caches.Cache
}

func NewUserService(
	tokenService tokens.TokenService,
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	userBuilder query.Builder,
	cache caches.Cache,
) UserService {
	return &userService{tokenService, userRepository, roleRepository, userBuilder, cache}
}

func (s *userService) GetUsers(ctx context.Context, queryParams url.Values) ([]entity.User, *response.Meta, error) {
	var users []entity.User
	var err error
	var meta *response.Meta
	db := s.userRepository.SingleTransaction()

	db, meta = s.userBuilder.ApplyBuilder(db, queryParams, &entity.User{})
	users, err = s.userRepository.GetUsers(ctx, db)
	if err != nil {
		return nil, nil, err
	}

	return users, meta, nil
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	userKey := "users:" + id
	user := &entity.User{}
	cachedData := s.cache.Get(userKey)
	if cachedData != "" {
		if err := json.Unmarshal([]byte(cachedData), user); err != nil {
			return nil, err
		}
	} else {
		var err error
		db := s.userRepository.SingleTransaction()
		user, err = s.userRepository.GetUserByID(ctx, db, id)
		if err != nil {
			return nil, err
		}

		data, _ := json.Marshal(user)

		if err := s.cache.Set(userKey, string(data), 5*time.Minute); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (s *userService) CreateUser(ctx context.Context, request dto.UserRequest) (*entity.User, error) {
	db := s.userRepository.SingleTransaction()
	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
	}

	if err := s.userRepository.CreateUser(ctx, db, user); err != nil {
		return nil, err
	}
	if err := s.cache.Del("users:all"); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, request dto.UpdateUserRequest) (*entity.User, error) {
	db := s.userRepository.SingleTransaction()
	user, err := s.userRepository.GetUserByID(ctx, db, request.ID)
	if err != nil {
		return nil, err
	}

	if request.Email != "" {
		user.Email = request.Email
	}
	if request.Username != "" {
		user.Username = request.Username
	}
	if request.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, err
		}

		user.Password = string(hashedPassword)
	}

	if err := s.userRepository.UpdateUser(ctx, db, user); err != nil {
		return nil, err
	}

	if err := s.cache.Del("users:" + user.ID.String()); err != nil {
		return nil, err
	}

	if err := s.cache.Del("users:all"); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	db := s.userRepository.SingleTransaction()
	user, err := s.userRepository.GetUserByID(ctx, db, id)
	if err != nil {
		return err
	}

	if err := s.userRepository.DeleteUser(ctx, db, user); err != nil {
		return err
	}

	if err := s.cache.Del("users:" + id); err != nil {
		return err
	}

	if err := s.cache.Del("users:all"); err != nil {
		return err
	}

	return nil
}

func (s *userService) ChangeRole(ctx context.Context, request dto.ChangeRoleRequest) error {
	return s.userRepository.WithTransaction(func(tx *gorm.DB) error {
		user, err := s.userRepository.GetUserByID(ctx, tx, request.UserID)
		if err != nil {
			return err
		}

		var addItems []*entity.Role
		var removeItems []*entity.Role

		for _, item := range request.Items {
			role, err := s.roleRepository.GetRoleByID(ctx, tx, item.ID)
			if err != nil {
				return err
			}

			if item.Action == "add" {
				addItems = append(addItems, role)
			} else if item.Action == "remove" {
				removeItems = append(removeItems, role)
			} else {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid action")
			}
		}

		if len(addItems) > 0 {
			if err := s.userRepository.AddRoles(ctx, tx, user, addItems); err != nil {
				return err
			}
		}

		if len(removeItems) > 0 {
			if err := s.userRepository.RemoveRoles(ctx, tx, user, removeItems); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *userService) Login(ctx context.Context, request dto.LoginRequest) (string, error) {
	db := s.userRepository.SingleTransaction()
	user, err := s.userRepository.GetUserByEmail(ctx, db, request.Email)
	var userPassword string
	if err != nil {
		userPassword = "$2a$10$pRe6SEQi6edG0bEYzAaMF.S1oszSANbZORukCi7j3QFku5jC1frFW"
	} else {
		userPassword = user.Password
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(request.Password)); err != nil || user == nil {
		return "", echo.NewHTTPError(http.StatusUnauthorized, "Invalid email or password")
	}
	expiredTime := time.Now().Add(24 * time.Hour)
	token, err := s.tokenService.GenerateAccessToken(tokens.JWTCustomClaims{
		ID:       user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredTime),
		},
	})

	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Register(ctx context.Context, request dto.UserRequest) (*entity.User, bool, error) {
	user := &entity.User{
		Username: request.Username,
		Email:    request.Email,
	}
	var isFirstUser bool
	if err := s.userRepository.WithTransaction(func(tx *gorm.DB) error {
		users, err := s.userRepository.GetUsersFiltered(ctx, tx, 1, 0, "id", "email != ?", request.Email)
		if err != nil {
			return err
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
		if err := s.userRepository.CreateUser(ctx, tx, user); err != nil {
			return err
		}
		if len(users) == 0 {
			isFirstUser = true
			roles, err := s.roleRepository.GetRolesFiltered(ctx, tx, 1, 0, "id", "auth_level >= 3")
			if err != nil {
				return err
			}
			var role *entity.Role
			if len(roles) == 0 {
				role = &entity.Role{
					Name:      "Admin",
					AuthLevel: 3,
				}
				if err := s.roleRepository.CreateRole(ctx, tx, role); err != nil {
					return err
				}
			} else {
				role = &roles[0]
			}

			if err := s.userRepository.AddRoles(ctx, tx, user, []*entity.Role{role}); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return nil, isFirstUser, err
	}

	return user, isFirstUser, nil
}
