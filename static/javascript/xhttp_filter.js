const filterButton = document.getElementById('filterButton');
const resetButton = document.getElementById('filterButtonReset')
const xHttp = new XMLHttpRequest();

//Отчищает блок с данными по нажатию кнопки Reset
resetButton.onclick = () => {
    data = document.getElementById('block-data')
    removeTextPage(data)
}

//Запрос по нажатию на кнопку
filterButton.onclick = (e) => {
    e.preventDefault(); //Запрещает перезагрузку формы

    let data = {
        "status": formValue('status'),
        "category": formValue('category'),
        "family": formValue('family'),
        "task": formValue('task'),
        "dateStart": formValue('dateStart'),
        "dateEnd": formValue('dateEnd'),
        "address": formValue('address'),
        "manager": formValue('manager'),
        "phone": formValue('phone')
    };
    
    request(data);
};

//Запрос на сервер
function request(e) {
    const url = "/admin/cabinet/filter/task?Status=" + e.status +"&Category=" + e.category + "&Family=" + e.family + "&Phone=" + e.phone + "&Task=" + e.task +"&Address=" + e.address + "&Manager=" + e.manager + "&DateStart=" + e.dateStart + "&DateEnd=" + e.dateEnd;
    xHttp.open("GET", url);
    xHttp.responseType = "json";
    xHttp.send();
    xHttp.onload = dataTaskFilter;
}

//Обработка данных
function dataTaskFilter() {
    blockData = document.getElementById('block-data');
    removeTextPage(blockData);

    let data = xHttp.response;
    console.log(data)

    for (let i in data) {
        let rows = document.createElement("table")
        rows.innerHTML = '<tr><td colspan="2" class="Table-Title">Задача ' + data[i]["KeyType"] + '-' + data[i]["Id"] + '</td></tr>' +
        '<tr><td class="Table-Name-Category">Категория ремонта</td><td>' + data[i]["Category"] + '</td></tr>' +
            '<tr><td class="Table-Name-Category">ФИО Исполнитель</td><td>' + data[i]["Manager"]["String"] + '</td></tr>' +
        '<tr><td class="Table-Name-Category">ФИО Заказчика</td><td>' + data[i]["FirstName"] + '</td></tr>' +
        '<tr><td class="Table-Name-Category">Номер Заказчика</td><td>' + data[i]["Phone"] + '</td></tr>' +
        '<tr><td class="Table-Name-Category">Адрес</td><td>' + data[i]["Address"]["String"] + '</td></tr>' + 
        '<tr><td class="Table-Name-Category">Cтатус</td><td>' + data[i]["Status"] + ' (' + data[i]["StatusTranslate"] + ')</td></tr>' +
        '<tr><td class="Table-Name-Category">Дата создание задачи</td><td>' + data[i]["DateStart"] + '</td></tr>' +
        '<tr><td class="Table-Name-Category">Дата изменение статус задачи</td><td>' + data[i]["DateStatus"] + '</td></tr>' + 
        '<tr><td colspan="2" class="Table-Name-Category"><p><a href="/admin/cabinet/task/' + data[i]["Id"] +'" target="_blank">Подробно</a></p></td></tr>';
        blockData.appendChild(rows);
    }
}

//Выводит данные с определенных форм
function formValue(e) {
    return document.getElementById(e).value;
}

//Удаляет текст на странице
function removeTextPage(e) {
    if (e != "") {
        e.innerHTML = "";
    }
}