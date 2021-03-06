apiVersion: v1
kind: Secret
metadata:
  name: remote-env-certificate
  namespace: kyma-installer
type: Opaque
data:
  remote_env_ca: "__REMOTE_ENV_CA__"
  remote_env_ca_key: "__REMOTE_ENV_CA_KEY__"
---
apiVersion: v1
kind: Secret
metadata:
  name: remote-env-certificate-verrides
  namespace: kyma-installer
  labels:
    installer: overrides
type: Opaque
data:
  global.remoteEnvCa: "__REMOTE_ENV_CA__"
  global.remoteEnvCaKey: "__REMOTE_ENV_CA_KEY__"
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cluster-certificate
  namespace: kyma-installer
data:
  tls_cert: "__TLS_CERT__"
  tls_key: "__TLS_KEY__"
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: cluster-certificate-overrides
  namespace: kyma-installer
  labels:
    installer: overrides
data:
  global.tlsCrt: "__TLS_CERT__"
  global.tlsKey: "__TLS_KEY__"
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: installation-config
  namespace: kyma-installer
data:
  is_local_installation: "__IS_LOCAL_INSTALLATION__"
  external_public_ip: "__EXTERNAL_PUBLIC_IP__"
  domain: "__DOMAIN__"
  remote_env_ip: "__REMOTE_ENV_IP__"
  k8s_apiserver_url: "__K8S_APISERVER_URL__"
  k8s_apiserver_ca: "__K8S_APISERVER_CA__"
  admin_group: "__ADMIN_GROUP__"
  enable_etcd_backup_operator: "__ENABLE_ETCD_BACKUP_OPERATOR__"
  etcd_backup_abs_container_name: "__ETCD_BACKUP_ABS_CONTAINER_NAME__"
  slack_api_url: "__SLACK_API_URL_VALUE__"
  slack_channel: "__SLACK_CHANNEL_VALUE__"
  victor_ops_routing_key: "__VICTOR_OPS_ROUTING_KEY_VALUE__"
  victor_ops_api_key: "__VICTOR_OPS_API_KEY_VALUE__"
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: installation-config-overrides
  namespace: kyma-installer
  labels:
    installer: overrides
data:
  global.isLocalEnv: "__IS_LOCAL_INSTALLATION__"
  ingressgateway.service.externalPublicIp: "__EXTERNAL_PUBLIC_IP__"
  global.domainName: "__DOMAIN__"
  nginx-ingress.controller.service.loadBalancerIP: "__REMOTE_ENV_IP__"
  configurations-generator.kubeConfig.url: "__K8S_APISERVER_URL__"
  configurations-generator.kubeConfig.ca: "__K8S_APISERVER_CA__"
  cluster-users.users.adminGroup: "__ADMIN_GROUP__"
  etcd-operator.backupOperator.enabled: "__ENABLE_ETCD_BACKUP_OPERATOR__"
  global.etcdBackupABS.containerName: "__ETCD_BACKUP_ABS_CONTAINER_NAME__"
  global.alertTools.credentials.slack.apiurl: "__SLACK_API_URL_VALUE__"
  global.alertTools.credentials.slack.channel: "__SLACK_CHANNEL_VALUE__"
  global.alertTools.credentials.victorOps.routingkey: "__VICTOR_OPS_ROUTING_KEY_VALUE__"
  global.alertTools.credentials.victorOps.apikey: "__VICTOR_OPS_API_KEY_VALUE__"
