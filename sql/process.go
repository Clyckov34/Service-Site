package sql

// Install
const (
	//CreateServiceName Виды Услуг
	CreateServiceName = `CREATE TABLE service_name(
		id INT(2) AUTO_INCREMENT PRIMARY KEY,
		key_type VARCHAR(5) NOT NULL,
		title VARCHAR(256) NOT NULL,
		url VARCHAR(256) NOT NULL,
		file_name VARCHAR(256) NOT NULL
	);`

	//CreateService Список Услуг
	CreateService = `CREATE TABLE service(
		id INT(10) AUTO_INCREMENT PRIMARY KEY,
		id_name INT(10) NOT NULL,
		title VARCHAR(256) NOT NULL,
		price INT(10) NOT NULL,
		sale INT(10) NULL,
		file_name VARCHAR(256) NOT NULL,
		text VARCHAR(1000) NOT NULL,
	
		FOREIGN KEY (id_name) REFERENCES service_name (id) ON DELETE CASCADE
	);`

	//CreateSocialNetwork Социальные сети
	CreateSocialNetwork = `CREATE TABLE social_network(
		id INT(2) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		id_icon INT(2) NOT NULL,
		url VARCHAR(256) NOT NULL,

		FOREIGN KEY (id_icon) REFERENCES icon (id) ON DELETE CASCADE
	);`

	//CreateAuth Менеджеры
	CreateAuth = `CREATE TABLE auth (
		id INT(2) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		login VARCHAR(256) NOT NULL,
		password VARCHAR(256) NOT NULL,
		first_name VARCHAR(100) NOT NULL,
		email VARCHAR(100) NOT NULL,
		code VARCHAR(256) NULL,
		date DATETIME,
		admin BOOLEAN
	);`

	//CreateStatus Статусы
	CreateStatus = `CREATE TABLE status (
		id INT(1) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		name VARCHAR(256) NOT NULL,
		translate VARCHAR(256) NOT NULL
	);`

	//CreateTask Задачи
	CreateTask = `CREATE TABLE task (
		id INT(6) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		id_status INT(1) NOT NULL,
		id_type INT(1) NULL,
		id_service INT(10) NULL,
		id_auth INT(2) NULL,
		first_name VARCHAR(256) NOT NULL,
		phone VARCHAR(12) NOT NULL,
		email VARCHAR(100) NULL,
		address VARCHAR(256) NULL,
		date_start DATETIME,
		date_status DATETIME,
		price INT(10),

		FOREIGN KEY (id_status) REFERENCES status (id) ON DELETE CASCADE,
		FOREIGN KEY (id_type) REFERENCES service_name (id) ON DELETE CASCADE,
		FOREIGN KEY (id_service) REFERENCES service (id) ON DELETE SET NULL,
		FOREIGN KEY (id_auth) REFERENCES auth (id) ON DELETE SET NULL
	);`

	//CreateComment Комментарии
	CreateComment = `CREATE TABLE comment (
		id INT(10) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		id_task INT(6) NOT NULL,
		id_auth INT(2) NULL,
		text TEXT(1000) NOT NULL,
		file_name VARCHAR(256) NULL,
		date DATETIME,

		FOREIGN KEY (id_task) REFERENCES task (id) ON DELETE CASCADE,
		FOREIGN KEY (id_auth) REFERENCES auth (id) ON DELETE SET NULL
	);`

	//CreateIcon Иконки
	CreateIcon = `CREATE TABLE icon (
		id INT(2) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		name VARCHAR(50) NOT NULL,
		teg VARCHAR(256) NOT NULL,
		color VARCHAR(50)
	);`

	//CreateHistoryTask история с задачей
	CreateHistoryTask = `CREATE TABLE history_task (
		id INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,
		id_auth INT(2) NULL,
		id_task INT(6) NOT NULL,
		title VARCHAR(500) NOT NULL,
		date DATETIME,
		ip VARCHAR(15) NOT NULL,
		
		FOREIGN KEY (id_auth) REFERENCES auth (id) ON DELETE SET NULL,
		FOREIGN KEY (id_task) REFERENCES task (id) ON DELETE CASCADE
	);`
)
