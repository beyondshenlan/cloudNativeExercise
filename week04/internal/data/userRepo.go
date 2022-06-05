package data

type UserRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

type DB struct {
	url string
}

func NewDB(str string) *DB {
	return &DB{url: str}

}
