Demo and presentation by [Jhonas Mutton](https://github.com/JhonasMutton) and [Andreia Silva](https://github.com/andreiac-silva).

# Knative-demo
This repository contains projects used to demonstrate knative in "Knative and Cloud Events in Go" talk.


## Demo flow

We have an online store. To buy something, the applicating pass to the following steps:

1. (source) Purchase Service receives the initial payload - user input
2. (sequence) Payment Service receives the payload from Purchase, load customer information and approves payment
3. (parallel) Email Service sends an email with transaction with payment status
4. (parallel) Invoce Service will generate purchase invoice

These are ficticios functionalities. The most important here are how the services are communicating with each other and how we can maintain two or more versions of one service at the same time. 
To know more about Services individually, see their respective READMEs. 
