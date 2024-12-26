# Project: Covid 19 NEO4J
# 📁 Collection: Extract, Transform, Load 


## End-point: Process CSV
### Method: POST
>```
>{{url}}/etl/csv
>```
### Body formdata

|Param|value|Type|
|---|---|---|
|vaccination|/D:/covid-19-neo4J/docs/data/vaccination-data.csv|file|
|vaccine|postman-cloud:///1efbfc64-69fc-43c0-a532-12152024084b|file|
|covid_global|postman-cloud:///1efbfc64-7ad6-4330-aa51-176f36fbd108|file|
|covid|postman-cloud:///1efbfc64-99d2-4540-b5d4-eefb0c5c784f|file|


### Response: 200
```json
{
    "message": "success",
    "warning": [
        "Warning: Unable to process the optional file 'covid_global': http: no such file"
    ]
}
```

### Response: 200
```json
{
    "message": "success",
    "warning": null
}
```

### Response: 422
```json
{
    "data": {},
    "errors": [
        "Error processing file 'covid': error unmarshalling CSV data: record on line 0; parse error on line 2, column 1: parsing time \"05/01/2020\" as \"2006-01-02T15:04:05Z07:00\": cannot parse \"05/01/2020\" as \"2006\""
    ],
    "message": "error processing some files",
    "warnings": null
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
# 📁 Collection: questions 


## End-point: Total Cases And Deaths By Country And Date
### Method: GET
>```
>{{url}}/questions/1?country_code=BR&date=10/01/2021
>```
### Query Params

|Param|value|
|---|---|
|country_code|BR|
|date|10/01/2021|


### Response: 200
```json
[
    {
        "country": "Brazil",
        "country_code": "BR",
        "created_at": "2024-12-26T04:33:43.5637373Z",
        "cumulative_cases": 8013708,
        "cumulative_deaths": 201460,
        "date_reported": "10/01/2021",
        "id": "C1PC9TFCBH3QF4",
        "new_cases": 313130,
        "new_deaths": 6049,
        "updated_at": "2024-12-26T04:33:43.5637373Z",
        "who_region": "AMRO"
    }
]
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Vaccinated People By Country And Date
### Method: GET
>```
>{{url}}/questions/2?country_code=BRA&date=17/01/2021
>```
### Query Params

|Param|value|
|---|---|
|country_code|BRA|
|date|17/01/2021|


### Response: 200
```json
{
    "data": {
        "c": {
            "Id": 14909,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:14909",
            "Labels": [
                "Country"
            ],
            "Props": {
                "code": "BRA",
                "created_at": "2024-12-26T12:10:51.5935835Z",
                "id": "C3974H06D3QE4Z",
                "name": "Brazil",
                "region": "AMRO",
                "updated_at": "2024-12-26T12:10:51.5935835Z"
            }
        },
        "d": {
            "Id": 14910,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:14910",
            "Labels": [
                "Date"
            ],
            "Props": {
                "date": "17/01/2021"
            }
        },
        "v": {
            "Id": 80755,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:80755",
            "Labels": [
                "Vaccination"
            ],
            "Props": {
                "country": "Brazil",
                "created_at": "2024-12-26T12:15:10.8025722Z",
                "data_source": "REPORTING",
                "date_updated": "29/12/2023",
                "first_vaccine_date": "17/01/2021",
                "id": "V53DA1YTDB1MAD",
                "iso3": "BRA",
                "persons_booster_add_dose": 1,
                "persons_booster_add_dose_per100": 52,
                "persons_last_dose": 1,
                "persons_last_dose_per100": 81,
                "persons_vaccinated_1plus_dose": 185000000,
                "persons_vaccinated_1plus_dose_per100": 87,
                "total_vaccinations": 5,
                "total_vaccinations_per100": 243,
                "updated_at": "2024-12-26T12:15:10.8025722Z",
                "who_region": "AMRO"
            }
        }
    },
    "message": "success"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Vaccines By Country And StartDate
### Method: GET
>```
>{{url}}/questions/3?country_code=GHA&start_date=03/02/2021
>```
### Query Params

|Param|value|
|---|---|
|country_code|GHA|
|start_date|03/02/2021|


### Response: 200
```json
{
    "data": {
        "c": {
            "Id": 15002,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:15002",
            "Labels": [
                "Country"
            ],
            "Props": {
                "code": "GHA",
                "created_at": "2024-12-26T12:10:51.8373515Z",
                "id": "CFZF1N8NGY7QXC",
                "name": "Ghana",
                "region": "AFRO",
                "updated_at": "2024-12-26T12:10:51.8373515Z"
            }
        },
        "d": {
            "Id": 14899,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:14899",
            "Labels": [
                "Date"
            ],
            "Props": {
                "date": "03/02/2021"
            }
        },
        "v": {
            "Id": 33298,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:33298",
            "Labels": [
                "Vaccine"
            ],
            "Props": {
                "authorization_date": "02/09/2021",
                "company": "Gamaleya Research Institute",
                "created_at": "2024-12-26T12:16:51.9086628Z",
                "data_source": "REPORTING",
                "end_date": "",
                "id": "VFGVSPK4KTQ8A7",
                "iso3": "GHA",
                "product": "Gam-Covid-Vac",
                "start_date": "03/02/2021",
                "updated_at": "2024-12-26T12:16:51.9086628Z",
                "vaccine": "Gamaleya - Gam-Covid-Vac"
            }
        }
    },
    "message": "success"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Vaccines By Country And StartDate Copy
### Method: GET
>```
>{{url}}/questions/4?country_code=GHA&start_date=03/02/2021
>```
### Query Params

|Param|value|
|---|---|
|country_code|GHA|
|start_date|03/02/2021|


### Response: 200
```json
{
    "data": {
        "c": {
            "Id": 15002,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:15002",
            "Labels": [
                "Country"
            ],
            "Props": {
                "code": "GHA",
                "created_at": "2024-12-26T12:10:51.8373515Z",
                "id": "CFZF1N8NGY7QXC",
                "name": "Ghana",
                "region": "AFRO",
                "updated_at": "2024-12-26T12:10:51.8373515Z"
            }
        },
        "d": {
            "Id": 14899,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:14899",
            "Labels": [
                "Date"
            ],
            "Props": {
                "date": "03/02/2021"
            }
        },
        "v": {
            "Id": 33298,
            "ElementId": "4:af72f23b-5f0d-47a2-9ee4-0cf6b872d50e:33298",
            "Labels": [
                "Vaccine"
            ],
            "Props": {
                "authorization_date": "02/09/2021",
                "company": "Gamaleya Research Institute",
                "created_at": "2024-12-26T12:16:51.9086628Z",
                "data_source": "REPORTING",
                "end_date": "",
                "id": "VFGVSPK4KTQ8A7",
                "iso3": "GHA",
                "product": "Gam-Covid-Vac",
                "start_date": "03/02/2021",
                "updated_at": "2024-12-26T12:16:51.9086628Z",
                "vaccine": "Gamaleya - Gam-Covid-Vac"
            }
        }
    },
    "message": "success"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Vaccines By Country And StartDate Copy 2
### Method: GET
>```
>{{url}}/questions/5?region=EMRO
>```
### Query Params

|Param|value|
|---|---|
|region|EMRO|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Ping
### Method: GET
>```
>{{url}}/ping
>```
### Response: 200
```json
{
    "ping": "pong"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
_________________________________________________
Powered By: [postman-to-markdown](https://github.com/bautistaj/postman-to-markdown/)
