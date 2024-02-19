Golearn Packages [DB - Gorm]
====================================
Learn any packages to support web development.

### How to run

#### Migration Tools
```shell
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz
mv migrate.linux-amd64 $GOPATH/bin/migrate

# sqlite support
go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

#### Run application
```shell
export DB_DSN=sqlite3://db.sqlite

# -- MIGRATION --
# make makemigrations name=create_users_table
make migrate

# -- RUN APPLICATIONS --
make runapi  # api application
make runweb  # web application
```

### Additional Info
* Support DB Transaction
* Support CRUD API and Web

### Contributors
Agung Yuliyanto: [LinkedIn](https://www.linkedin.com/in/agung96tm/) | [Github](https://github.com/agung96tm)