package page_setting

type Param struct {
	Id       int    `db:"id"`
	Email    string `db:"email"`
	FullName string `db:"first_name"`
	Login    string `db:"login"`
}

type CheckLogin struct {
	Login string `db:"login"`
}

type CheckPassword struct {
	Password string `db:"password"`
}
