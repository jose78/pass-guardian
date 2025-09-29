# Experiment K8s External Secrets

This project will try to documernt with examples how to implement and the CRD secret external easily. 


## Structure of Components

        ├── chart-test-secret-external
        │   ├── charts
        │   ├── Chart.yaml
        │   ├── templates
        │   │   ├── deployment.yaml
        │   │   ├── _helpers.tpl
        │   │   ├── hpa.yaml
        │   │   ├── httproute.yaml
        │   │   ├── ingress.yaml
        │   │   ├── NOTES.txt
        │   │   ├── serviceaccount.yaml
        │   │   ├── service.yaml
        │   │   └── tests
        │   │       └── test-connection.yaml
        │   └── values.yaml
        ├── kind-cluster-conf
        ├── README.md
        └── simple-project
            ├── Dockerfile
            ├── go.mod
            └── main.go

