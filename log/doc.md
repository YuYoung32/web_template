## Log模块

封装一个全局的、根据配置文件的日志等级设定的、预设格式的Logrus对象。

### 作用

#### 用于记录程序运行过程中的日志信息

使用`GetLogger()`或者`GetLoggerWithSkip(int)`获取一个使用WithField方法 设定 日志Field有调用栈 的Logrus对象。

#### 提供中间件，供Gin框架使用

从`gin.ctx`中获取信息并输出到日志中。

