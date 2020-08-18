#/bin/bash
# *** ATTENTION *** This script is going to remove all of your knative components. It is important to follow the order to delete knative components.

echo "Removing sink bindings..."
kubectl delete -f ./flow/sink-binding-purchase.yaml -n demo

echo "Removing triggers..."
kubectl delete -f ./flow/trigger-process.yaml -n demo
kubectl delete -f ./flow/trigger-purchase.yaml -n demo

echo "Removing parallels..."
kubectl delete -f ./flow/parallel.yaml -n demo

echo "Removing sequences..."
kubectl delete -f ./flow/sequence.yaml -n demo

echo "Removing brokers..."
kubectl delete -f ./flow/broker-purchase.yaml -n demo
kubectl delete -f ./flow/broker-process.yaml -n demo
