package project

var SampleProjectJson = []byte(`{
  "name": "sample"
}
`)

var SampleDeploymentJson = []byte(`{
  "apiVersion": "v1",
  "kind": "DeploymentConfig",
  "metadata": {
    "name": "sample",
    "labels": {
      "app": "sample",
      "template": "sample"
    }
  },
  "spec": {
    "triggers": [
      {
        "type": "ConfigChange"
      }
    ],
    "replicas": 0,
    "selector": {
      "app": "sample",
      "deploymentconfig": "sample"
    },
    "template": {
      "metadata": {
        "labels": {
          "app": "sample",
          "deploymentconfig": "sample"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "sample",
            "image": "sample:1.0",
            "imagePullPolicy": "Always",
            "env": [
              {
                "name": "ENV_VAR",
                "value": "$ENV_VAL"
              },
              {
                "name": "KUBERNETES_NAMESPACE",
                "valueFrom": {
                  "fieldRef": {
                    "apiVersion": "v1",
                    "fieldPath": "metadata.namespace"
                  }
                }
              }
            ],
            "ports": [
              {
                "containerPort": 8080,
                "name": "8080-http"
              }
            ],
            "resources": {
              "limits": {
                "cpu": "1",
                "memory": "1Gi"
              },
              "requests": {
                "cpu": "100m",
                "memory": "100Mi"
              }
            },
            "livenessProbe": {
              "httpGet": {
                "path": "/health",
                "scheme": "HTTP",
                "port": 8080
              },
              "initialDelaySeconds": 30,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 6
            },
            "readinessProbe": {
              "httpGet": {
                "path": "/info",
                "scheme": "HTTP",
                "port": 8080
              },
              "initialDelaySeconds": 30,
              "timeoutSeconds": 15,
              "periodSeconds": 10,
              "successThreshold": 1,
              "failureThreshold": 3
            }
          }
        ],
        "dnsPolicy": "ClusterFirst",
        "imagePullSecrets": [
          {
            "name": "sample-secret"
          }
        ]
      }
    }
  }
}
`)

var SampleRouteJson = []byte(`{
  "apiVersion": "v1",
  "kind": "Route",
  "metadata": {
    "name": "sample-route",
    "labels": {
      "app": "sample"
    }
  },
  "spec": {
    "port": {
      "targetPort": "8080-http"
    },
    "tls": {
      "insecureEdgeTerminationPolicy": "Redirect",
      "termination": "edge"
    },
    "to": {
      "kind": "Service",
      "name": "sample-service",
      "weight": 100
    }
  }
}
`)

var SampleSecretJson = []byte(`{
  "apiVersion": "v1",
  "kind": "Secret",
  "metadata": {
    "name": "sample-secret"
  },
  "data": {
    "SAMPLE_ENV": "YWI=",
    "SAMPLE_ENV_OTHER": "YWM="
  }
}
`)

var SampleServiceJson = []byte(`{
  "apiVersion": "v1",
  "kind": "Service",
  "metadata": {
    "name": "sample-service",
    "labels": {
      "app": "sample"
    }
  },
  "spec": {
    "ports": [
      {
        "name": "8080-http",
        "port": 8080,
        "protocol": "TCP",
        "targetPort": 8080
      }
    ],
    "selector": {
      "app": "sample",
      "deploymentconfig": "sample"
    },
    "type": "ClusterIP",
    "sessionAffinity": "None"
  }
}
`)
