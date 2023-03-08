// Получить кнопку:
let mybutton = document.getElementById("myBtn");

// Когда пользователь прокручивает вниз 20px от верхней части документа, покажите кнопку
window.onscroll = () => {
    let position = 100;

    if (document.body.scrollTop > position || document.documentElement.scrollTop > position) {
        mybutton.style.display = "block";
    } else {
        mybutton.style.display = "none";
    }
}

mybutton.onclick = () => {
    document.body.scrollTop = 0; // Для Safari
    document.documentElement.scrollTop = 0; // Для Chrome, Firefox, IE и Opera
}
