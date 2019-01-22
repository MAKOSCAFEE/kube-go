# kube-go

Collection of golang and kubernetes learnings.

- [go-pod](go-pod.yaml) - This demostrate how to create a pod separately.
- [node-services](node-services.yaml) - Demonstrate how to create services
- [node-controller](node-controller.yaml) - Demonstrate how to create `ReplicationController` with three replicas.
- [internal-service](sample-internal-service.yaml) - Demonstate how to create internal service. Very useful for backend service that you want to be available to othe containers running in your cluster but not open to the world(`no external ip`)
- [NodePort-service](custom-load-balancing.yaml) - This type of service allows service to be exposed through the host/node on a specific port. This way you can use IP address to access service on the assigne node port.
- [Multiple-ports](multiple-ports.yaml) - This shows how you can create multiple ports and expose them at the same time.
