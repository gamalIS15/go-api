package simple

type SimpleRepository struct {
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleService(repository *SimpleRepository) *SimpleService {
	return &SimpleService{SimpleRepository: repository}
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}
