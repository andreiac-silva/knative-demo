#/bin/bash
# *** ATTENTION *** Before executing that, ensure if there are sufficient nodes aware on Kubernetes

echo "Creating brokers..."
kubectl apply -f ./flow/broker-purchase.yaml -n demo
kubectl apply -f ./flow/broker-process.yaml -n demo
sleep 10

echo "Creating sequences..."
kubectl apply -f ./flow/sequence.yaml -n demo
sleep 10

echo "Creating parallels..."
kubectl apply -f ./flow/parallel.yaml -n demo
sleep 10

echo "Creating triggers..."
kubectl apply -f ./flow/trigger-process.yaml -n demo
kubectl apply -f ./flow/trigger-purchase.yaml -n demo
sleep 10

echo "Creating sink bindings..."
kubectl apply -f ./flow/sink-binding-purchase.yaml -n demo
