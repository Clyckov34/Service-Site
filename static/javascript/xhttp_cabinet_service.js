const input = document.getElementById("Service-search");
const namesBlock = document.getElementById("Block-service-search");
const keyboard = document.getElementById("form-search");

const xHttp = new XMLHttpRequest();

// Отключаем Enter
input.onkeydown = (e) => {
    if (e.keyCode === 13) return false;
};


input.oninput = () => { 
    RequestJson(input.value) 
};


function RequestJson(input){
    xHttp.open("GET", "/admin/cabinet/list-service/search?p=" + input, true);
    xHttp.responseType = "json";
    xHttp.send();
    
    xHttp.onload = DataServiceJson;
};


function DataServiceJson(){
    let data = xHttp.response;

    // Если данные существуют на HTML страницы, то заменяем на пустые строки
    if (namesBlock != ""){
        namesBlock.innerHTML = "";
    }

    console.log(data);

    for (let i in data){
        if (data[i]["Sale"]["Int64"] != 0) {
            res = '<del>' + data[i]["Sale"]["Int64"] + '</del> <span class="Sale">' + data[i]["Price"] + '</span>'
        } else {
            res = data[i]["Price"]
        }

        let rows = document.createElement("table")
        rows.innerHTML = '<tr><td colspan="2" class="Table-Title">' + data[i]["Title"] + '</td></tr>'+
        '<tr><td class="Table-Name-Category">Категория</td><td>' + data[i]["TitleTypeRepair"] + '</td></tr>'+
        '<tr><td class="Table-Name-Price">Цена</td><td>' + res + '</td></tr>' +
        '<tr><td colspan="2" class="Document"><a href="/admin/cabinet/list-service/detail/' + data[i]["Id"] + '">Подробно</a></td></tr>';
        namesBlock.appendChild(rows);
    }
};