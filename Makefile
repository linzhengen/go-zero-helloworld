dev:
	skaffold dev -v debug --kube-context=minikube

deploy-etcd:
	helm repo add bitnami https://charts.bitnami.com/bitnami --force-update
	helm repo update bitnami
	helm upgrade --install etcd bitnami/etcd \
	--kube-context minikube \
	--create-namespace \
	--namespace etcd-dns \
	--set auth.rbac.create=false \
	--set livenessProbe.enabled=false \
	--set readinessProbe.enabled=false \
	--set startupProbe.enabled=false
