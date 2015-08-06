package service

type MailService interface {
	SendMail() error
}

type mailService struct {
}

func New() *mailService {
	return new(mailService)
}

func (m *mailService) SendMail() error {
	return nil
}
