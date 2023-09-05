# 轻巧的Golang代码生成器

**功能支持：**

- 生成实体类【包含json】

- 支持Mysql

- 默认格式：
    
    使用方式：`./gen -dns="root:pwd@tcp(127.0.0.1:3306)/dbName?charset=utf8"`
	
	```go
	type User struct {
	  Id        int       `db:"id" json:"id,omitempty"`
	  Username  string    `db:"username" json:"username,omitempty"`
	  Password  string    `db:"password" json:"password,omitempty"`
      Salt      string    `db:"salt" json:"salt,omitempty"`
      Email     string    `db:"email" json:"email,omitempty"`
      CreatedAt time.Time `db:"created_at" json:"createdAt,omitempty"`
    }
    ```
    
- 支持xorm框架

  使用方式：`./gen -dns="root:pwd@tcp(127.0.0.1:3306)/dbName?charset=utf8" -orm=xorm `

  生成格式：

  ```go
  type User struct {
  	Id        int       `xorm:"int(11) notnull pk autoincr" json:"id,omitempty"`
  	Username  string    `xorm:"varchar(50) notnull" json:"username,omitempty"`
  	Password  string    `xorm:"varchar(255) notnull" json:"password,omitempty"`
  	Salt      string    `xorm:"varchar(20) notnull" json:"salt,omitempty"`
  	Email     string    `xorm:"varchar(100)" json:"email,omitempty"`
  	CreatedAt time.Time `xorm:"datetime" json:"createdAt,omitempty"`
  }
  ```

  

**生成原理：**

使用golang 原生的`template`进行渲染，参考文档：

https://cloud.tencent.com/developer/article/1683688

`SHOW TABLES`查询库中所有表

`db.Query("SHOW COLUMNS FROM " + tableName)`查询表字段信息