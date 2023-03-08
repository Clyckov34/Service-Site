package page_home

// Service виды услуг
type Service struct {
	Title    string `db:"title"`
	Url      string `db:"url"`
	FileName string `db:"file_name"`
}

type ServiceData struct {
	Data  []Service
	Error error
}

// SocialNetwork группы социальных сетей
type SocialNetwork struct {
	Url         string `db:"social_network.url"`
	FontAwesome string `db:"icon.teg"`
	Color       string `db:"icon.color"`
}

type SocialNetworkData struct {
	Data  []SocialNetwork
	Error error
}

// PortfolioData Портфолио
type PortfolioData struct {
	Data  []string
	Error error
}

