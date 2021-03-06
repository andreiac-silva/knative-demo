# Payment
Payment Service is a dummy service that receives a cloud event, process and propagate it thought the flow.
This Service will approve purchase (dummy) and will load customer's name and e-mail into current event.

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
To simulate various versions running for the same application, there are 3 versions available on docker hub.
So, you can pull some one of this tags v1, v2 or latest.
```
docker pull andreiacsilva/payment:<v1,v2 or latest>
```

## How can I directly send a cloud event to Payment Service?
If you are running the application locally, you can simply send a POST request such as one of bellow.

Binary
```
curl -X POST \
  http://localhost:8080/ \
  -H 'content-type: application/json' \
  -H 'ce-id: 610b6dd4-c85d-417b-b58f-3771e532' \
  -H 'ce-source: http://store.com/buy' \
  -H 'ce-specversion: 1.0' \
  -H 'ce-type: com.store.buy' \
  -d '{
   "customer":{
      "cpf":"123.222.333-X"
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

Structured
```
curl -v "http://localhost:8080" \
-X POST \
-H "Content-Type: application/cloudevents+json" \
-d '{"data": {"customer":{ "cpf":"123.222.333-X" }, "items":[ { "id":2000, "description":"Covid-19 Vaccine", "quantity":70, "price":10.00 }, { "id":2001, "description":"Pale Ale Beer", "quantity":70, "price":15.00 }]},"datacontenttype": "application/json","id": "e587cbc20dcd4e4b80eaf9ebbb2c5efc","source": "http://sensedia.com/purchase/#demo/purchase-6f7b55865d-42jck","specversion": "1.0","time": "2020-08-18T11:27:00.000350563Z","type": "dev.knative.eventing.purchase"}'
```
