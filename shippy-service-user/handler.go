package main

import (
	"context"
	"github.com/asim/go-micro/v3/errors"
	"github.com/asim/go-micro/v3/logger"
	pb "github.com/dbielecki97/shippy/shippy-service-user/proto/user"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authable interface {
	Decode(token string) (*CustomClaims, error)
	Encode(user *pb.User) (string, error)
}

type handler struct {
	repository   Repository
	tokenService authable
}

func (s *handler) Get(ctx context.Context, req *pb.User, res *pb.Response) error {
	result, err := s.repository.Get(ctx, req.Id)
	if err != nil {
		return err
	}

	user := UnmarshalUser(result)
	res.User = user

	return nil
}

func (s *handler) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) error {
	results, err := s.repository.GetAll(ctx)
	if err != nil {
		return err
	}

	users := UnmarshalUserCollection(results)
	res.Users = users

	return nil
}

func (s *handler) Auth(ctx context.Context, req *pb.User, res *pb.Token) error {
	user, err := s.repository.GetByEmail(ctx, req.Email)
	if err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return status.Errorf(codes.Unauthenticated, "invalid password")

	}
	token, err := s.tokenService.Encode(UnmarshalUser(user))
	if err != nil {
		return err
	}

	res.Token = token
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.User, res *pb.Response) error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	req.Password = string(hashedPass)
	req.Id = uuid.NewV4().String()
	if err := s.repository.Create(ctx, MarshalUser(req)); err != nil {
		return err
	}

	// Strip the password back out, so's we're not returning it
	req.Password = ""
	res.User = req
	return nil
}

func (s *handler) ValidateToken(ctx context.Context, req *pb.Token, res *pb.Token) error {
	logger.Info("ValidateToken")
	claims, err := s.tokenService.Decode(req.Token)
	if err != nil {
		logger.Error(err)
		return err
	}

	if claims.User.Id == "" {
		return errors.Unauthorized("shippy.service.user", "invalid user")
	}

	res.Valid = true
	return nil
}
