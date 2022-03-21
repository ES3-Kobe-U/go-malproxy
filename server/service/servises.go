package service

import "context"

type Service interface {
	CheckingTheIntegrityOfAmazonInformation(ctx context.Context, email string, password string) (context.Context, error)
	CheckingTheIntegrityOfRakutenInformation(ctx context.Context, userId string, password string) (context.Context, error)
	CheckingContextContents(ctx context.Context) error
	CheckingTheIntegrityOfAmazonCaptcha(ctx context.Context, password string, guess string) (context.Context, error)
}
