package page_type_service

import "sync"

type Param struct {
	Id        int
	Title     string
	Url       string
	FileName  string
	KeyType   string
}

type ProfileDetail struct {
	Id        int    `db:"service_name.id"`
	Title     string `db:"service_name.title"`
	Url       string `db:"service_name.url"`
	FileName  string `db:"service_name.file_name"`
	KeyType   string `db:"service_name.key_type"`
}

type Profile struct {
	Id    int    `db:"service_name.id"`
	Title string `db:"service_name.title"`
}

type FileName struct {
	Comment []string
	Service []string
}

type WG struct {
	wp sync.WaitGroup
}
