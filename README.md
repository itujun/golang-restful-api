# BELAJAR MEMBUAT CRUD RESTFUL-API

## LIBRARY TAMBAHAN

- Sql Driver : (https://github.com/go-sql-driver/mysql/)[https://github.com/go-sql-driver/mysql/]
  Installation

```bash
go get -u github.com/go-sql-driver/mysql
```

- Http Router : (https://github.com/julienschmidt/httprouter)[https://github.com/julienschmidt/httprouter]
  Installation

```bash
go get github.com/julienschmidt/httprouter
```

- Validation : (https://github.com/go-playground/validator)[https://github.com/go-playground/validator]
  Installation

```bash
go get github.com/go-playground/validator/v10
```

- Unit Test / Testify (https://github.com/stretchr/testify)[https://github.com/stretchr/testify]
  Installation

```bash
go get github.com/stretchr/testify
```

## MEMBUAT DATABASE

query

```bash
create table category
(
  id  integer PRIMARY KEY auto_increment,
  name VARCHAR(200) not NULL
)engine = InnoDB
```
