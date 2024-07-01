package user

type storage interface {
	Create(u User)
}

type Service struct {
	s storage
}

func NewService(s storage) *Service {
	return &Service{s: s}
}

func (s *Service) SignUp(email, password string) {

	// complex logic of gathering user data
	user := New(email, password)

	// user notifications: emails, SMS, etc.
	s.s.Create(user)

	// sending internal events to share user creation

	// etc
}
