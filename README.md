# Mini-Cloud

> 如果它看起来像MINI，开起来像MINI，那么它可能就是一台MINI

> 微信小程序简单HTTPS接口，基于gin实现，使用MySQL持久化

## 准备工作

### 安装 MySQL(MariaDB)

> yum install mariadb-server -y
> systemctl start mariadb
> /usr/bin/mysql_secure_installation
> systemctl enable mariadb

### Nginx

> 安装 Nginx 目的是为了使用 80 端口 可以支持多个域名

> yum install nginx -y
> systemctl start nginx
> systemctl enable nginx

```nginx
server
{
        listen  443;
        server_name  mini.your_domain.com;
        index index.html;
        root /data/workshop/src/Mini-Cloud/views/;

		//https支持
        ssl on;
        ssl_certificate /etc/letsencrypt/live/mini.your_domain.com/fullchain.pem;
        ssl_certificate_key /etc/letsencrypt/live/mini.your_domain.com/privkey.pem;

        location /api {
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass http://127.0.0.1:8080/api;
        }

        access_log  /var/log/nginx/access.log  main;
}
server
{
        listen 80;
        server_name mini.your_domain.com;
        index index.html;
        root /data/workshop/src/Mini-Cloud/views/;
        return 301 https://$host$request_uri; //http重定向至https

        location /api {
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass http://127.0.0.1:8080/api;
        }
}
```

## 安装部署

#### gin安装

> go get -u github.com/gin-gonic/gin
> go get github.com/kardianos/govendor
> govendor init
> govendor fetch github.com/gin-gonic/gin@v1.3

#### 依赖包

> go get -u github.com/go-ini/ini
> go get -u github.com/Unknwon/com
> go get -u github.com/jinzhu/gorm
> go get -u github.com/go-sql-driver/mysql

## HTTPS支持

> 使用免费SSL证书

```HTML
git clone https://github.com/certbot/certbot
cd certbot
./certbot-auto --help
./certbot-auto certonly --webroot --agree-tos -v -t --email yours@email.com -w /data/workshop/src/Mini-Cloud/views -d mini.your_domain.com
```

## 图片分离

> 使用CDN

## 简单使用

- 启动、停止、重启、编译

> ./run.sh build {start|stop|restart|build}