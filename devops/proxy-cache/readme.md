```
docker compose -f ./proxy-cache/docker-compose.yaml up
```

get admin pass

```
docker exec -it nexus cat /nexus-data/admin.password
```

Nexus UI: http://localhost:8081

Squid Proxy: http://localhost:3128 (örneğin apt için proxy ayarı)
