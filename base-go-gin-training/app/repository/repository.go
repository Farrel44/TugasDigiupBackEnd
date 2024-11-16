package repository

import "base-gin/storage"

var (
	accountRepo   *AccountRepository
	personRepo    *PersonRepository
	publisherRepo *PublisherRespository
	authorRepo    *AuthorRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
	publisherRepo = newPublisherRepo(db)
	authorRepo = newAuthorRepository(db)

}

func GetAccountRepo() *AccountRepository {
	return accountRepo
}

func GetPersonRepo() *PersonRepository {
	return personRepo
}

func GetPublisherRepo() *PublisherRespository {
	return publisherRepo
}

func GetAuthorRepo() *AuthorRepository {
	return authorRepo
}