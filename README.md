# golaburo

## run locally
```sh
go run main.go
```

## Deployment
```sh
git push dokku main
```

### migration guide

https://github.com/golang-migrate/migrate/blob/master/database/postgres/TUTORIAL.md


```sh
# migrate create -ext sql -dir db/migrations -seq create_users_table

migrate -database ${POSTGRESQL_URL} -path db/migrations up
# migrate -database ${POSTGRESQL_URL} -path db/migrations down
```

### Database access

For now it's defined as global

https://techinscribed.com/different-approaches-to-pass-database-connection-into-controllers-in-golang/
https://www.alexedwards.net/blog/organising-database-access


## dokku

```
# update dependencies
sudo apt update
sudo apt upgrade

# add swapfile
sudo swapon --show
sudo fallocate -l 1G /swapfile
sudo chmod 600 /swapfile
sudo mkswap /swapfile
sudo swapon /swapfile
sudo cp /etc/fstab /etc/fstab.bak
echo '/swapfile none swap sw 0 0' | sudo tee -a /etc/fstab

# install dokku
wget https://raw.githubusercontent.com/dokku/dokku/v0.23.0/bootstrap.sh
sudo DOKKU_TAG=v0.23.0 bash bootstrap.sh

# create dokku app
dokku apps:create laburo
dokku proxy:ports laburo
dokku proxy:ports-set laburo http:80:5000
dokku proxy:ports-set laburo http:80:43594
dokku proxy:enable laburo

dokku config:set laburo GOVERSION=go1.15.6

dokku buildpacks:add laburo https://github.com/heroku/heroku-buildpack-go.git
dokku buildpacks:add laburo https://github.com/heroku/heroku-buildpack-nodejs.git
```

### debug container
```
docker ps
docker exec -it e628307fa2fa bash
```
