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
    xHttp.open("GET", "price-list/search?p=" + input, true);
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

    for (let i in data){
        if (data[i]["Sale"]["Int64"] != 0) {
            res = 'от <del>' + data[i]["Sale"]["Int64"] + '</del> <span class="Sale">' + data[i]["Price"] + ' руб</span>'
        } else {
            res = 'от ' + data[i]["Price"] + ' руб'
        }

        let rows = document.createElement("a")
        rows.setAttribute("href", '/store/' + data[i]["Url"] + '/detail?id=' + data[i]["Id"])
        rows.setAttribute("title", 'Перейти в раздел ' + data[i]["Title"])
        rows.className = "Ofer"

        rows.innerHTML = '<img src="/img/service/' + data[i]["FileName"] + '" alt="' + data[i]["Title"] + '">' +
        '<h3>' + data[i]["Title"] + '</h3>' +
        '<div class="Price">' + res + '</div>';
        namesBlock.appendChild(rows);
    }
};