# Mini-Cloud

> 如果它看起来像 MINI，开起来像 MINI，那么它可能就是一台 MINI

> 微信小程序简单 HTTPS 接口，基于 gin 实现，使用 MySQ L持久化

## 准备工作

### 安装 MySQL(MariaDB)

```html
yum install mariadb-server -y
systemctl start mariadb
/usr/bin/mysql_secure_installation
systemctl enable mariadb
```

### Nginx

> 安装 Nginx 目的是为了使用 80 端口 可以支持多个域名

```html
yum install nginx -y
systemctl start nginx
systemctl enable nginx
```

```nginx
server
{
        listen  443;
        server_name  mini.your_domain.com;
        index index.html;
        root /data/workshop/src/Mini-Cloud/views/;

        //https support
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
        return 301 https://$host$request_uri; //http redirect to https

        location /api {
                proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr;
                proxy_pass http://127.0.0.1:8080/api;
        }
}
```

## 安装部署

```shell
glide install
```

## HTTPS 支持

> 使用免费 SSL 证书

```html
git clone https://github.com/certbot/certbot
cd certbot
./certbot-auto --help
./certbot-auto certonly --webroot --agree-tos -v -t --email yours@email.com -w /data/workshop/src/Mini-Cloud/views -d mini.your_domain.com
```

## 静态资源

> 使用 CDN 加速

> 吐槽下套路云: 几年没用的 OSS，一用就因欠费0.01元，不是有免费的5GB流量吗？

## 简单使用

- 启动、停止、重启、编译

> ./run.sh build {start|stop|restart|build}