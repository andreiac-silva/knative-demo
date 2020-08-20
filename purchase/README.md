# Purchase
Purchase Service is a dummy service that receives a rest call and start an event in knative services.

## Execution

### Local

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
### Pulling the project image
```
docker pull andreiacsilva/purchase:latest
```

### Environment Variables
The service must receive some environment variables, by default those env vars are in .env file.
Example:
```
SERVER_PORT=8080

POD_NAMESPACE=demo 
POD_NAME=purchase
```
Obs: Some env vars are set by Knative in deployment process.

### API
Purchase service receives a **POST** in path "**/purchase/api/v1/buy**" with this body:

```
{
   "customer":{
      "cpf":"111.222.333-X"
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
}
```
