let submitStatus = document.getElementById('status')
let manager = document.getElementById('manager')

submitStatus.onclick = () => {
    if (manager) {
        return true
    } else {
        alert("Добавьте исполнителя к задаче")
        return false
    }
}
