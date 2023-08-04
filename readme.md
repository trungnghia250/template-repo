# Sample repo golang + sql database + orm

### 1. References
1. For ORM (Object-relational mapping) you can consider using that library: [Link](https://gorm.io/)
2. Project Layout: [Link](https://github.com/golang-standards/project-layout)
3. Style guideline: [Link](https://github.com/uber-go/guide/blob/master/style.md)

### 2. Concepts
Rule of Clean Architecture by Uncle Bob:

- Independent of Frameworks. The architecture does not depend on the existence of some library of feature laden software. This allows you to use such frameworks as tools, rather than having to cram your system into their limited constraints.

- Testable. The business rules can be tested without the UI, Database, Web Server, or any other external element.

- Independent of UI. The UI can change easily, without changing the rest of the system. A Web UI could be replaced with a console UI, for example, without changing the business rules.

- Independent of Database. You can swap out Oracle or SQL Server, for Mongo, BigTable, CouchDB, or something else. Your business rules are not bound to the database.

- Independent of any external agency. In fact your business rules simply don’t know anything at all about the outside world.

More at https://8thlight.com/blog/uncle-bob/2012/08/13/the-clean-architecture.html

This project has 4 Domain layer:

- Models Layer

- Repository Layer

- Usecase Layer

- Delivery Layer

*The diagram:*
![img.png](img.png)
