# Kubernetes keystone authentication webhook

- [Kubernetes keystone authentication webhook](#kubernetes-keystone-authentication-webhook)
    - [Prerequisites](#prerequisites)
    - [Limitations](#limitations)
    - [Build instructions](#build-instructions)
        - [Docker image](#docker-image)
        - [Basic executable](#basic-executable)
    - [Local testing](#local-testing)
    - [Kubernetes setup](#kubernetes-setup)
    - [License](#license)

## Prerequisites
1. Docker.
2. Openstack Keystone serving over http or https.
3. The properly configured permissions for 'identity:validate_token'
   keystone operation.
4. The properly signed SSL certificate and the private key placed in the 'build' dir.


## Limitations
1. Only v3 keystone authentication method is supported.
2. User's groups are not listed in the response.


## Build instructions
### Docker image
```bash
$ make docker

# Start the webhook from the newly built image
$ docker create -it --rm -p 2000:2000 -e OS_AUTH_URL="http://<address>:<port>/v3" -e OS_USERNAME="admin" -e OS_PASSWORD="password" -e OS_DOMAIN_ID="default" --name  keystone-auth-hook keystone-auth-hook
```

### Basic executable
```bash
# Build for Linux
$ make

# Specify GOOS env variable to build for other platforms,
# e.g. to build for OSX
$ GOOS=darwin make

# Prepare environemnt
$ export OS_AUTH_URL="http://<address>:<port>/v3"
$ export OS_USERNAME="admin"
$ export OS_PASSWORD="password"
$ export OS_DOMAIN_ID="default"

# Run auth hook
$ cd build
$ ./keystone-auth-hook --certfile webhook.pem --keyfile webhook.key
```


## Local testing
```bash
# Issue a token (assume that access to keystone is configured in a proper way)
$ openstack token issue --os-username <username> --os-tenant-name <username>

# Try to authenticate
$ curl -D - --cacert ca-bundle.pem -X POST https://<address>:<port>/authenticate -d '{"apiVersion": "authentication.k8s.io/v1beta1", "kind": "TokenReview", "spec": {"token": "<token>"}}' -H "Content-Type: application/json"
```


## Kubernetes setup
1. Copy docs/auth-webhook-config.yaml to /etc/kubernetes/manifests/
2. Copy your SSL certificate, private key and CA authority bundle to /etc/kubernetes/pki/webhook/
3. Add the following option to /etc/kubernetes/manifests/kube-apiserver.yaml file:
```bash
 - --authentication-token-webhook-config-file=/etc/kubernetes/manifests/auth-webhook-config.yaml
```
4. Restart kubelet:
```bash
$ systemctl restart kubelet
```
5. Try simple command using the token generated before:
```bash
$ kubectl -s="https://<k8s-master-address>:6443" --insecure-skip-tls-verify=true get pods --token <token> --namespace kube-system
```

For more detailed information regarding kubernetes setup please refer to
[the official documentation](https://kubernetes.io/docs/admin/authentication/#webhook-token-authentication).

## License

[Apache 2.0](LICENSE)
