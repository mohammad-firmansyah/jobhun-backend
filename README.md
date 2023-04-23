
# Jobhun Academy 

Intern test for as backend developer at jobhun.id


## Installation

to install this project you need to install go version 1.19.2

```bash
  cd name-of-project
  go mod tidy
  go run main.go
```
    
## API Documentation

Postman : https://www.postman.com/science-technologist-7222748/workspace/mohammad-firmansyah/documentation/18626491-298b9c31-0891-49d3-b633-b63e638f8f57

#### Get all mahasiswa

```http
  GET /api/v1/mahasiswa
```

#### Get detail mahasiswa

```http
  GET /api/v1/mahasiswa/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `integer` | **Required**. Id of item to fetch |

#### Add new mahasiswa
```http
  POST /api/v1/mahasiswa
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nama`      | `string` | **Required**. name of the student |
| `usia`      | `integer` | **Required**. age of the student |
| `gender`      | `integer` | **Required**. gender of the student , 0 (man) 1 (women)|
| `tgl_registrasi`      | `string` | **Required**. registration date of the student|
| `hobi`      | `[]string` | **Required**. list of hoby of the student|
| `jurusan`      | `string` | **Required**. student major|

#### Update data mahasiswa
```http
  PUT /api/v1/mahasiswa
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nama`      | `string` | **Required**. name of the student |
| `usia`      | `integer` | **Required**. age of the student |
| `gender`      | `integer` | **Required**. gender of the student , 0 (man) 1 (women)|
| `tgl_registrasi`      | `string` | **Required**. registration date of the student|
| `hobi`      | `[]string` | **Required**. list of hoby of the student|
| `jurusan`      | `string` | **Required**. student major|

#### Delete data mahasiswa
```http
  DELETE /api/v1/mahasiswa/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `integer` | **Required**. Id of item to fetch |

#### Get all jurusan
```http
  GET /api/v1/jurusan
```

#### Add data jurusan
```http
  POST /api/v1/jurusan
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nama_jurusan`      | `string` | **Required**. name of major |

#### Get all jurusan
```http
  GET /api/v1/hobi
```

#### Add data hobi
```http
  POST /api/v1/hobi
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `nama_hobi`      | `string` | **Required**. name of hobby |

## Documentation

[API Documentation in Postman](https://www.postman.com/science-technologist-7222748/workspace/mohammad-firmansyah/documentation/18626491-298b9c31-0891-49d3-b633-b63e638f8f57)


## Authors

- [@mohammad-firmansyah](https://www.github.com/mohammad-firmansyah)

