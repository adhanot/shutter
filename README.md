# Shutter

shutter is a RESTful API to control my home Shutters. I use it with alexa.

It use go-swagger (https://goswagger.io/), ENT Framework (https://entgo.io/) and JWT (https://github.com/dgrijalva/jwt-go).

shutter generate the client and the server. 

If you want modify the API go-swagger server, you need to modify the : 
```bash
/gen/api/swagger.yaml
```

If you want add ENT database schema, you can add file in : 
```bash
/gen/api/ent
```


## Installation

Use need to use docker-compose command :

```bash
docker-compose build --no-cache  && docker-compose up
```

## Usage

shutter launch API routes.
The routes are :
```bash
http://0.0.0.0:5000/register
http://0.0.0.0:5000/login
http://0.0.0.0:5000/salon/open
http://0.0.0.0:5000/salon/close
http://0.0.0.0:5000/salon/stop
http://0.0.0.0:5000/chambre/open
http://0.0.0.0:5000/chambre/close
http://0.0.0.0:5000/chambre/stop
http://0.0.0.0:5000/sdb/open
http://0.0.0.0:5000/sdb/close
http://0.0.0.0:5000/sdb/stop

```
After modify the swagger.yaml or add new database schema, you just need go in the "scripts" folder and launch :
```bash
sh generate.sh
```

You can access on the documentation here :

```bash
http://0.0.0.0:5000/v1/docs

