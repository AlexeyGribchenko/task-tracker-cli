[![CI Pipeline](https://github.com/AlexeyGribchenko/pos-counter/actions/workflows/ci.yml/badge.svg)](https://github.com/AlexeyGribchenko/pos-counter/actions/workflows/ci.yml)
[![Coverage Status](https://coveralls.io/repos/github/AlexeyGribchenko/task-tracker-cli/badge.svg?branch=main)](https://coveralls.io/github/AlexeyGribchenko/task-tracker-cli?branch=main)
[![LICENSE](https://img.shields.io/github/license/AlexeyGribchenko/task-tracker-cli)](https://github.com/AlexeyGribchenko/task-tracker-cli/blob/main/LICENSE)

# Task tracker
Тут будет краткое описание проекта
## Оглавнение
1. Функционал
2. Конфигурирование 
3. Запуск / Установка

## Функционал
### 1. Получить полный список задач
Команда:
```bash
task list
    -s "Column name" # Сортировка по колонке
    -f "filter value" # Фильтрация по значению
    # филтрация пока не реализована
```
> **s** - опциональный параметр для сортировки списка задач по названию колонки

>**f** - опциональный параметр для фильтрации списка задач по значению

Пример:
```bash
# добавить пример вывода
```
#### 1.1 Сортировка списка задач
Пример:
```bash
task list -s "status"
```
Пример:
```bash
# добавить пример вывода
```
#### 1.2 Фильтрация списка задач список задач
Команда:
```bash
task list -f "value" # доработать логиту фильтрации
```
Пример:
```bash
# добавить пример вывода
```
### 2. Добавить новую задачу
Команда:
```bash
task add "Task name"
    -d "Description for the task" # Добавить описание
```
> **d** - опциональный парамерт для добавления описания к задаче

Пример:
```bash
# добавить пример вывода
```

### 3. Обновить статус задачи
Допустимые статусы задач:
1. `created`
2. `in_progress` # будет заменено на active|pending или что-то еще
3. `completed`
4. `cancelled`

Команда:
```bash
task status "completed"
```
Пример:
```bash
task status "completed"

# добавить пример вывода
```
### 4. Удаление задачи
Команда:
```bash
task rm task_id
```
Пример:
```bash
task rm 10

# добавить пример вывода
```
## Запуск / Установка
### Сборка из исходного кода
Клонируйте репозиторий
```bash
git clone https://github.com/AlexeyGribchenko/task-tracker-cli
```
Если у вас есть `make`, выполните команду
```bash
make install
```
Пока не реализовано.

На этом этапе должна будет пройти сборка приложения (для чего потребуется установленный язык программирования go).

Затем произойдет установка файлов в соответствующие папки в Linux.

Вариант установки на windows также будет проработан и описан позднее. 