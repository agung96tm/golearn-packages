Golearn Packages [DB]
====================================
Learn any packages to support web development.


### How to run
```shell
docker-compose up -d --build
export DB_DSN=postgres://golearn_user:golearn_password@127.0.0.1:5432/golearn_db?sslmode=disable

make init
make dependencies

make migrate

make runapi  # api application, localhost:8001
make runweb  # web application, localhost:8000
```

### Contributors
Agung Yuliyanto: [LinkedIn](https://www.linkedin.com/in/agung96tm/) | [Github](https://github.com/agung96tm)