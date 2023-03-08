package page_store

import (
	"fmt"
)

//message состовлет сообщения и оправляет отправителю | раздел Форма заказа
func message(firstName, phone, email, street, categoryName, serviceName string) string {
	return fmt.Sprintf(`
	<head>
		<style>
			
			h2 {
				text-align: center;
			}

			table {
				width: 100%%;
				margin-bottom: 20px;
				background: #e2e2e2;
				padding: 10px;
			}

			table tr, table td {
				border: 1px solid #2e2e2e;
				padding: 5px;
				width: 50%%;
			}
	
			.Table-Title, .Table-Title-2 td {
				text-align: center;
				background: #f7b537;
				font-weight: bold;
			}
	
			.Table-Body-2-Price {
				text-align: center;
			}
	
			.Table-Name-Category, .Table-Name-Price {
				background: #e2e2e2;
			}

		</style>
	</head>
	<body>

		<table>
			<tr>
				<td colspan="2" class="Table-Title">Заявка на оформления услуги</td>
			</tr>
			<tr>
				<td class="Table-Name-Category">Фамилия Имя Отчество</td>
				<td>%v</td>
			</tr>
			<tr>
				<td class="Table-Name-Price">Телефон</td>
				<td><a href="tel:%v">%v</a></td>
			</tr>
			<tr>
				<td class="Table-Name-Price">Почта</td>
				<td><a href="mailto:%v">%v</a></td>
			</tr>
			<tr>
				<td class="Table-Name-Price">Адрес</td>
				<td>%v</td>
			</tr>
			<tr>
				<td class="Table-Name-Price">Категория</td>
				<td>%v</td>
			</tr>
			<tr>
				<td class="Table-Name-Price">Товар (Услуга)</td>
				<td>%v</td>
			</tr>
		</table>

	</body>
	`, firstName, phone, phone, email, email, street, categoryName, serviceName)
}
