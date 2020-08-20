Demo and presentation by [Jhonas Mutton](https://github.com/JhonasMutton) and [Andreia Silva](https://github.com/andreiac-silva).

# Knative-demo
This repository contains projects used to demonstrate knative in "Knative and Cloud Events in Go" talk.


## Demo flow

We have an online store. To buy something, the applicating pass to the following steps:

1. (source) Purchase Service receives the initial payload - user input
2. (sequence) Payment Service receives the payload from Purchase, load customer information and approves payment
3. (parallel) Email Service sends an email with transaction with payment status
4. (parallel) Invoce Service will surprisingly generate purchase invoice :smile:

These are ficticios functionalities. The most important here are how the services are communicating with each other and how we can maintain two or more versions of one service at the same time. 
To know more about Services individually, see their respective READMEs. 

## Installation

All installation files are in [services-install](https://github.com/andreiac-silva/knative-demo/tree/master/services-install) directory. You'll see two sub-directories:

:file_folder: helm
:file_folder: knative


## Prerequisites

You need a cluster with Kubernetes and Knative installed.

### Services
Access [helm](https://github.com/andreiac-silva/knative-demo/tree/master/services-install/helm) directory. You can use [helm](https://helm.sh/docs/intro/install/) to install demo services:

``` helm install purchase purchase/ -n demo ```
``` helm install payment payment/ -n demo ```
``` helm install email email/ -n demo ```
``` helm install invoice invoice/ -n demo ```

Other option is just apply kservice.yaml of each module by Kubernetes (remember to replace values on each kservice.yaml):

```kubectl apply -f purchase/templates/kservice.yaml ```
```kubectl apply -f payment/templates/kservice.yaml ```
```kubectl apply -f email/templates/kservice.yaml ```
```kubectl apply -f invoice/templates/kservice.yaml ```

### Knative Components
Access [knative](https://github.com/andreiac-silva/knative-demo/tree/master/services-install/knative) directory and execute create-flows.sh.
