## Build w/xcaddy
```
xcaddy build 
    --with github.com/rsp2k/caddy-server-header
```

## Caddyfile
```
https://goiter.net {
    route {
        serverheader {
            header_value "GoiterBSD/1.0"
        }
    }
}
```
