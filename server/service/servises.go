package service

type Service interface {
	CheckingTheIntegrityOfAmazonInformation(email string, password string) error
	CheckingTheIntegrityOfRakutenInformation(userId string, password string) error
	CheckingContextContents() error
}
