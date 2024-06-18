# 根据已有的数据库表生成结构体
```
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
go get github.com/volatiletech/null/v8

sqlboiler psql -c config.toml -p order_dao

config.toml:
output   = "my_models"
wipe     = true
no-tests = true
add-enum-types = false
no-auto-timestamps = true
no-back-referencing = true
no-context = true
no-driver-templates = true
no-hooks = true
no-rows-affected = true

[psql]
  dbname = "dbname"
  host   = "localhost"
  port   = 5432
  user   = "dbusername"
  pass   = "dbpassword"
  schema = "myschema"
  sslmode = "disable"  
  blacklist = ["migrations", "other"]

[mysql]
  dbname  = "dbname"
  host    = "localhost"
  port    = 3306
  user    = "dbusername"
  pass    = "dbpassword"
  sslmode = "false"
  tinyint_as_int = true

[mssql]
  dbname  = "dbname"
  host    = "localhost"
  port    = 1433
  user    = "dbusername"
  pass    = "dbpassword"
  sslmode = "disable"
  schema  = "notdbo"
```
