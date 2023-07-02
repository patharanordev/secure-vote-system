# **PostgreSQL**

## **Encrypt Password**

`gen_salt` Generates a new random salt string for use in crypt(). The salt string also tells crypt() which algorithm to use.

The type parameter specifies the hashing algorithm. The accepted types are: des, xdes, md5 and bf.

```golang
// ...
var lastInsertId int
err := p.db.QueryRow(
    "INSERT INTO user_info(username, password) VALUES( $1, crypt($2, gen_salt('bf')) ) returning uid;",
    usr, pwd,
).Scan(&lastInsertId)

return lastInsertId, err
```

## **Contribute**

### **Data initialization**

Normally `docker-compose.dev.yml` file, make initial table and setup extension which required in this service, please refer to  `./common/init.sql`. The `*.sql` file will be executed in `/docker-entrypoint-initdb.d` of the `postgres` container image :

```yml
# ...

    db:
        container_name: db
        image: postgres
        # ...
        volumes:
        - ./common/init.sql:/docker-entrypoint-initdb.d/init.sql

# ...
```

![init-db](../assets/init-db.png)


In case database didn't create (or initial) correctly, please removing `pgdata` directory then re-run command below again :

```sh
sh run.dev.sh
```
