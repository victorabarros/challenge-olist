# Work at Olist

1. Deploy your project on a hosting service (we recommend [Heroku](https://heroku.com));
2. Apply for the position at our [career page](https://olist.gupy.io/) with:

- Link to the fork on Github (or bitbucket/gitlab);
- Link to the deployed project in a hosting service.

## Specification

### 1. Receive a CSV with authors and import to database

Given a CSV file with many authors (more than a million), you need to build a command to import the data into the database.

```csv
name
Luciano Ramalho
Osvaldo Santana Neto
...
```

Each author record in the database must have the following fields:

- id (self-generated)
- name

### 2. Expose authors' data in an endpoint

This endpoint needs to return a paginated list with the authors' data.
Optionally the authors can be searched by name.

### 3. CRUD (Create, Read, Update and Delete) of books

Each book record has the fields:

- id (self-generated)
- name
- edition
- publication_year
- authors (more than one author can write a book)

To create a book you need to send this payload (in json format) below:

```json
{
 "name": "string",
 "edition": "integer",
 "publication_year": "integer",
 "authors": "list of ids"
}
```

To retrieve a book (in easy mode) we can filter by 4 fields (or a composition of these four):

- name
- publication_year
- edition
- author

## Project Requirements

- [Effective Go](https://golang.org/doc/effective_go.html)
- Documentation:
  - Description;
  - Installing (setup) and testing instructions;
  - If you provide a [docker](https://www.docker.com/) solution for setup, ensure it works without docker too.
  - Brief description of the work environment used to run this project (Computer/operating system, text editor/IDE, libraries, etc).
- Provide API documentation (in English);
- Variables, code and strings must be all in English.

## Recommendations

- Write tests! Please make tests ... we appreciate tests <3... tests make the world better;
- Practice the [12 Factor-App](http://12factor.net) concepts;

## TODO

- [X] flag arg to load csv authors
- [X] Expose author in endpoint `/authors?offset=20&limit=10&name=Xpto` and `/author/<int:id>`
- [ ] db (mysql) https://dev.mysql.com/doc/mysql-getting-started/en/
- [ ] code server (try tdd)
- [ ] jmeter load test; >> 1000 tps
- [ ] code climate
- [ ] code tests 80%unit/20%integration
- [ ] commit postman json libary
- [ ] heroku
- [ ] docker-compose olist/api olist/postgress

### V2

- [ ] redis cache
- [ ] gke
