package page_social_network

type ListNetworkAll struct {
	Id    int    `db:"social_network.id"`
	Icon  string `db:"icon.teg"`
	Url   string `db:"social_network.url"`
	Title string `db:"icon.name"`
}

type ListNetworkAllChan struct {
	Data []ListNetworkAll
	Err  error
}

type Param struct {
	Id     int
	IdIcon int
	Url    string
}

type ListNetworkId struct {
	Id  int    `db:"id"`
	Url string `db:"url"`
}

type ListNetworkIdChan struct {
	Data ListNetworkId
	Err  error
}

type Icon struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type IconChan struct {
	Data []Icon
	Err  error
}
