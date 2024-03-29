# Golang Web App Example
[![Go Report Card](https://goreportcard.com/badge/github.com/dustyrat/go-webapp?style=flat-square)](https://goreportcard.com/report/github.com/dustyrat/go-webapp)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/dustyrat/go-webapp)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/dustyrat/go-webapp)](https://pkg.go.dev/github.com/dustyrat/go-webapp)
[![Release](https://img.shields.io/github/release/dustyrat/go-webapp.svg?style=flat-square)](https://github.com/dustyrat/go-webapp/releases/latest)


# Notes:
* Ingress Controller install: `helm upgrade --install ingress-nginx ingress-nginx --repo https://kubernetes.github.io/ingress-nginx --namespace ingress-nginx --create-namespace`
* Local Docker registry: `docker run -d -p 5000:5000 --restart=always --name registry registry:latest`
* Add DNS to hosts file:
  Entry: `127.0.0.1 example.internal`
  * Windows Location: `c:\windows\system32\drivers\etc\hosts`
  * Linux: `/etc/hosts`
* Kubernetes Dashboard setup: [kubernetes-dashboard](https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/)

