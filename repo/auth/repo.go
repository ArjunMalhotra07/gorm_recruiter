package repo

import (
	"log"

	"github.com/ArjunMalhotra07/gorm_recruiter/models"
	pb "github.com/ArjunMalhotra07/gorm_recruiter/proto"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type Authentication interface {
	CreateUserID() (string, error)
	CreateEncryptedPassword(userPassword, passwordHash string) (string, error)
	CreateUser(user *models.User) error
	CreateJwtToken(userID string, isEmployer bool) (string, error)
	LoginUser(email, password string) (*models.User, error)
	SendWelcomeEmail(to, subject, body string) error
}

type AuthRepo struct {
	Driver      *gorm.DB
	EmailClient pb.EmailServiceClient
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to email service: %v", err)
	}
	emailClient := pb.NewEmailServiceClient(conn)
	return &AuthRepo{Driver: db, EmailClient: emailClient}
}
