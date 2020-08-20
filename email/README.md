# Email
Email Service is a dummy service that receives a cloud event and simulate sending email.

## Execution

- ### Local

#### First of all, you must download Wire
O Wire is responsible for dependency injection. Download it:
```
go get -u github.com/google/wire/cmd/wire
```
After the previous step, execute:
```
wire
```

---

#### Run

```
go run ./
```
- ### Pulling the project image
```
docker pull andreiacsilva/email:latest
```

## How can I directly send a cloud event to Email Service?
If you are running the application locally, you can simply send a POST request such as bellow.

```
curl -X POST \
  http://localhost:8080/ \
  -H 'content-type: application/json' \
  -H 'ce-id: 610b6dd4-c85d-417b-b58f-3771e532' \
  -H 'ce-source: http://store.com/buy' \
  -H 'ce-specversion: 1.0' \
  -H 'ce-type: com.store.buy' \
  -H 'ce-status: approved' \
  -d '{
   "customer":{
      "cpf":"123.222.333-X",
      "name":"Robert Neville",
      "email":"robert.neville@company.com"
   },
   "items":[
      {
         "id":2000,
         "description":"Covid-19 Vaccine",
         "quantity":70,
         "price":10.00
      },
      {
         "id":2001,
         "description":"Pale Ale Beer",
         "quantity":70,
         "price":15.00
      }
   ]
}'
```
