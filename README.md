# Work at Olist

1. Deploy your project on a hosting service (we recommend [Heroku](https://heroku.com));

## Specification

### 3. CRUD of books

- Update -> put and patch
- Delete

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
