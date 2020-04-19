package seller_repository

type Factory interface {
	CreateRoRepository() Repository
	CreateRwRepository() Repository
}
