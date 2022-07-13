# SSTU document processing REST API

#### GET documents/download/:filename - скачать документ
#### GET documents/ - получить все документы
```json
{
    "error": false,
    "message": "documents found",
    "data": {
        "data": [
            {
                "id": "62cec54551fcd8a452809e90",
                "name": "Пушкин Александр Сергеевич",
                "job": "программист старшийю",
                "education": "СГТУ им Гагарина Ю.А. 2131212",
                "passport": "123123 123123",
                "address": "г Саратов ул Пушкина дом Калатушкина 1231",
                "snils": "1231231",
                "email": "pushkinas@mail.ru",
                "filepath": "document157988100",
                "street": "Пушкина",
                "house": "10",
                "flat": "42",
                "city": "Саратов",
                "index": "51501",
                "phone": "8800553535",
                "course": "sss"
            }
        ]
    }
}
```
#### POST documents/create - создать новый документ
```json
{
    "name": "Пушкин Александр Сергеевич",
    "job": "программист старшийю",
    "education": "СГТУ им Гагарина Ю.А. 2131212",
    "passport": "123123 123123",
    "address": "г Саратов ул Пушкина дом Калатушкина 1231",
    "snils": "1231231",
    "email": "pushkinas@mail.ru",
    "street": "Пушкина",
    "house": "10",
    "flat": "42",
    "city": "Саратов",
    "index": "51501",
    "phone": "8800553535",
    "doctype": "aptech_course",
    "course": "sss"
}
```


#### GET clients/ - получить всех клиентов
```json
{
    "error": false,
    "message": "clients found",
    "data": {
        "data": [
            {
                "_id": {
                    "name": "Пушкин Александр Сергеевич",
                    "passport": "123123 123123"
                },
                "documentsCount": 1
            }
        ]
    }
}
```
#### POST courses/create/ - создать новый курс
```json
{
    "name": "sss",
    "cost": 12,
    "hours": 13,
    "courseType": "ddd"
}
```
#### GET courses/ - получить все курсы
```json
{
    "error": false,
    "message": "courses found",
    "data": {
        "data": [
            {
                "name": "«Информационные технологии» по профилю направления подготовки 09.03.02 «Информационные системы и технологии»",
                "hours": 622,
                "cost": 106000,
                "courseType": "Дополнительные профессиональные программы профессиональной переподготовки"
            },
            {
                "name": "sss",
                "hours": 13,
                "cost": 12,
                "courseType": "ddd"
            }
        ]
    }
}
```
