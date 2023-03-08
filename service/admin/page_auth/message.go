package page_auth

import (
	"fmt"
)

//Message тело документа
func Message(code, dateTimeInterval string) string {
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
					<td colspan="2" class="Table-Title">Подтверждающий код</td>
				</tr>
				<tr>
					<td class="Table-Name-Category">Одноразовый код</td>
					<td>%v</td>
				</tr>
				<tr>
					<td class="Table-Name-Category">Срок действия кода</td>
					<td>%v</td>
				</tr>
			</table>
		</body>
		`, code, dateTimeInterval)
}
