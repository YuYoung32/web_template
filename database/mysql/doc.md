## mysql模块

封装了一个全局的mysql连接池，提供自动迁移。

### 作用

### 提供ORM对象关系映射

建议在每个（同类型）模型单独放在一个文件里，并定义`autoMigrate`方法，在connection.go中调用。

### 提供数据库连接

使用`GetDBInstance()`