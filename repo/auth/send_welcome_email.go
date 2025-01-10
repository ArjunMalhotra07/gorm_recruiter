package repo

import (
	"context"
	"time"

	pb "github.com/ArjunMalhotra07/gorm_recruiter/proto"
)

func (r *AuthRepo) SendWelcomeEmail(to, subject, body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	req := &pb.SendEmailRequest{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	_, err := r.EmailClient.SendEmail(ctx, req)
	return err
}
