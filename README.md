# QuickAuth

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod&style=flat-square)](https://gitpod.io/#https://github.com/hello-jiangxiaoyu/QuickAuth)

![logo](docs/quick-auth.jpg)

oauth2, oidc auth

local develop cmd
```bash
docker run -p 5432:5432 -e POSTGRES_DB=quick_auth -e POSTGRES_PASSWORD=admin -e POSTGRES_USER=admin -d postgres
```


# 数据库规范
    1、表达是与否概念的字段，必须使用 is_xxx 的方式命名
    2、表名、字段名必须使用小写字母或数字，禁止出现数字开头，禁止两个下划线中间只出现数字
    3、model表名不使用复数名词
    4、主键索引名为 pk_字段名；唯一索引名为 uk_字段名；普通索引名则为 idx_字段名
    5、如果存储的字符串长度几乎相等，使用 CHARACTER 定长字符串类型
    6、表必备三字段：id, created_at, updated_at
    7、索引尽量不要用 TEXT 字段，合适的字符存储长度，不但节约数据库表空间、节约索引存储，更重要的是提升检索速度
    8、单表行数超过 500 万行或者单表容量超过 2GB，才推荐进行分库分表
    9、业务上具有唯一特性的字段，建议建成唯一索引
    10、需要 join 的字段，数据类型保持绝对一致；多表关联查询时，保证被关联的字段需要有索引
    11、数据修改或删除时，要先 select，避免出现误删除

