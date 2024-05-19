# hryvnia-svc

***Note:*** the service structure has been generated using 
[this generator](https://gitlab.com/tokend/yo-generator-tokend-module).

## Description

The service allows users to check the current UAH exchange rate and subscribe to it.

## Documentation

To open documentation, run
```bash
  cd docs
  npm install
  npm start
```
and go to [swagger editor](http://localhost:8080/swagger-editor/).

To build documentation use 
```bash
npm run build
```

that will create open-api documentation in `web_deploy` folder.

## Running from docker compose
  
Make sure that docker and docker compose installed, then
```bash
docker compose up
```

If you make some changes in code and want to apply them, run
```bash
docker compose up --build
```

To stop it
```bash
docker compose down
```
