# Regal Riches

Regal Riche是一款社交游戏的后端服务。提供后端处理接口

## 接口设计

| 接口 | 说明 | 参数 | 返回值 |
| - | - | - | - |
| /registry | 注册 | username, password | 200: 注册成功, 400: 用户名已存在 |
| /login | 登录 | username, password | 200: 登录成功, 400: 用户名或密码错误 |