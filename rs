apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  containers:
  - name: nginx
    image: nginx:1.14.2
    ports:
    - containerPort: 80


apiVersion: v1
kind: Pod
metadata:
  name: nx
spec:
  containers:
  - name: nx
    image: nginx:1.14.2
    ports:
    - containerPort: 80

------label both the pods diffrenetly----dev---test

apiVersion: apps/v1
kind: ReplicaSet
metadata:
  name: myrs
spec:
  replicas: 3
  selector:
    matchExpressions:
      - key: team
        operator: In
        values:
           - dev
           - test
  template:
    metadata:
      name: dev
      labels:
        team: dev

    spec:
      containers:
      - name: php-redis
        image: nginx


     




