package repository

import "base-gin/storage"

var (
	accountRepo   *AccountRepository
	personRepo    *PersonRepository
	publisherRepo *PublisherRespository
	authorRepo    *AuthorRepository
	bookRepo      *BookRepository
	borrowRepo    *BorrowRepository
)

func SetupRepositories() {
	db := storage.GetDB()
	accountRepo = newAccountRepository(db)
	personRepo = newPersonRepository(db)
	publisherRepo = newPublisherRepo(db)
	authorRepo = newAuthorRepository(db)
	bookRepo = NewRepository(db)
	borrowRepo = NewBorrowRepository(db)
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

func GetBookRepo() *BookRepository {
	return bookRepo
}

func GetBorrowRepo() *BorrowRepository {
	return borrowRepo
}
