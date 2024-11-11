<html>
<div>
    <center>
        <h1>HOME - Услуги для дома</h1>
    </center>
</div>
<div>
    <h2>Описание:</h2>    
    <p>Сайт для обслуживание от покупок до заказ услуг. Принимает, и обработывает заявки от клиентов. Личный кабинет:</p>
    <img src="https://s1.hostingkartinok.com/uploads/images/2023/02/1072ddf2307385cd14adbf055a6d37f7.png">
</div>
<div>
    <h2>Docker-compose images</h2>
    <ul>
        <li>nginx:alpine (Файл конфигурации: <code>volumes/etc/nginx/conf.d/nginx_repair.conf</code>)</li>
        <li>mysql:8.0</li>
        <li>phpmyadmin:5.2</li>
        <li>golang:1.20</li>
    </ul>
</div>
<div>
    <h2>Функционал:</h2>
    <ol>
        <li>Планировщик задач.</li>
            <ul>
                <li>Изменять статусы</li>
                <ul>
                    <li>To Do (К выполнению)</li>
                    <li>In Progress (В процессе)</li>
                    <li>Pause (Пауза)</li>
                    <li>Denied (Отклонено)</li>
                    <li>Done (Готово)</li>
                </ul>
                <li>Оставлять заметки в задачях + фото</li>
                <li>Назначать отвественного менеджера</li>
                <li>История изменений</li>
                <li>Назначать стоимость объекта</li>
                <li>Редактировать, Удалять задачи</li>
            </ul>
        <li>Фильтр по (Задачам):</li>
            <ul>
                <li>Статусом</li>
                <li>Категориям</li>
                <li>Исполнителям (Отвественные за выполнения задач)</li>
                <li>ФИО заказчика</li>
                <li>Номер телефона заказчика</li>
                <li>Адрес заказчика</li>
                <li>Номер задачи</li>
                <li>Дата "ОТ" (По умолчанию: Первого числа текущего месяца)</li>
                <li>Дата "ДО" (По умолчанию: Текущая дата)</li>
            </ul>
        <li>Статистика</li>
            <ul>
                <li>Общая статистика (С графиком)</li>
            </ul>
        <li>Управление менеджерами</li>
            <ul>
                <li>Добавить, Редактировать, Удалить</li>
                <li>Выдавать права:</li>
                <table style="border: 1px solid">
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid">№</td>
                        <td style="border: 1px solid">Наименование</td>
                        <td style="border: 1px solid">Admin</td>
                        <td style="border: 1px solid">Manager</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">1</td>
                        <td style="border: 1px solid; text-align: center">Планировщик задач</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">2</td>
                        <td style="border: 1px solid; text-align: center">Фильтр по (Задачам)</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">3</td>
                        <td style="border: 1px solid; text-align: center">Статистика</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">4</td>
                        <td style="border: 1px solid; text-align: center">Настройка личного кабинета</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">5</td>
                        <td style="border: 1px solid; text-align: center">Управление менеджерами</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">-</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">6</td>
                        <td style="border: 1px solid; text-align: center">Категории услуг</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">-</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">7</td>
                        <td style="border: 1px solid; text-align: center">Список услуг</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">-</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">8</td>
                        <td style="border: 1px solid; text-align: center">Портфолио</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">-</td>
                    </tr>
                    <tr style="border: 1px solid">
                        <td style="border: 1px solid; text-align: center">9</td>
                        <td style="border: 1px solid; text-align: center">Соц. Сети</td>
                        <td style="border: 1px solid; text-align: center">+</td>
                        <td style="border: 1px solid; text-align: center">-</td>
                    </tr>
                </table>
            </ul>
        <li>Категории услуг (Добавить, Редактировать, Удалить)</li>
        <li>Список услуг (Добавить, Редактировать, Удалить)</li>
        <li>Портфолио (Добавить, Удалить)</li>
        <li>Соц. Сети (Добавить, Редактировать, Удалить)</li>
        <li>Настройка личного кабинета:</li>
            <ul>
                <li>Пароль (Редактировать)</li>
                <li>Логин (Редактировать)</li>
                <li>Email (Редактировать)</li>
                <li>ФИО (Редактировать)</li>
            </ul>
    </ol>
</div>
<div>
    <h2>Первоначальная настройка сайта для запуска</h2>
    <ol>
        <li>Скачать проект: > <code>git clone https://github.com/Clyckov34/Repair.git</code></li>
        <li>Виртуальное окружение</li>
        <img src="https://s1.hostingkartinok.com/uploads/images/2023/02/cc801e8301ab3412bfbeb02ad865a601.png">
            <ol>
                <li>Открыть и отредактировать файл <i>.env</i></li>
                <li>Отредактировать переменные MYSQL:</li>
                    <ul>
                        <li><code>MYSQL_CREATE_USER</code> - Логин. Создает нового пользователя от базы данных MySQL</li>
                        <li><code>MYSQL_CREATE_PASSWORD</code> - Пароль. Для нового пользователя <a href="https://1password.com/ru/password-generator/" target="_blank">Генератор пароля</a></li>
                    </ul>
                <li>Отредактировать переменные Почтового клиента Яндекс.Почта:</li>
                    <ul>
                        <li><code>MAIL_LOGIN</code> - Логин почты</li>
                        <li><code>MAIL_PASSWORD</code> - <a href="https://passport.yandex.ru/profile/access/apppasswords/create?retpath=https://mail.yandex.ru&scope=mail&uid=122011994" target="_blank">Пароль от внешнего приложения</a></li>
                        <li><code>MAIL_HOST</code> - Хост почты SMTP. (По умолчанию: smtp.yandex.ru)</li>
                        <li><code>MAIL_PORT</code> - Порт. (По умолчанию: 465)</li>
                        <a href="https://yandex.ru/support/mail/mail-clients/others.html" target="_blank">Документация Яндекс.Почта</a>, можно использовать алтернативные почтовые сервера например от Google, Mail.RU
                    </ul>
            </ol>
    </ol>
</div>
<div>
    <h2>Build and Run APP Docker-compose (OC Linux) VPS - Server</h2>
    <ol>
        <li>> <code>sudo docker-compose build</code> - Сборка контейнеров</li>
        <li>> <code>sudo docker-compose up -d</code> - Запуск контейнеров</li>
    </ol>
</div>
<div>
    <h2>Одноразовая настройка личного кабинета</h2>
    <ol>
        <li>Создание личного кабинета <i>https://myserver/install</i></li>
        <li>Личный кабинет <i>https://myserver/admin</i></li>
    </ol>
</div>