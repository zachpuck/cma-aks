package k8sutil

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"

	"github.com/Azure/go-autorest/autorest/to"
)

func GetSecretList(namespace string, options v1.ListOptions) (result []corev1.Secret, err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	secrets, err := client.CoreV1().Secrets(namespace).List(options)

	result = secrets.Items
	return
}

func DeleteSecret(name string, namespace string) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	err = client.CoreV1().Secrets(namespace).Delete(name, &v1.DeleteOptions{})
	return
}

func GetSSHSecretList(namespace string) (result []corev1.Secret, err error) {
	listOption := v1.ListOptions{FieldSelector: "type=" + string(corev1.SecretTypeSSHAuth)}
	return GetSecretList(namespace, listOption)
}

func GetSecret(name string, namespace string) (secret corev1.Secret, err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	secretResult, err := client.CoreV1().Secrets(namespace).Get(name, v1.GetOptions{})
	if err != nil {
		return
	}
	secret = *secretResult
	return
}

func GetSSHSecret(name string, namespace string) (secret []byte, err error) {
	secretResult, err := GetSecret(name, namespace)
	if err != nil {
		return
	}
	if secretResult.Type != corev1.SecretTypeSSHAuth {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeSSHAuth, secretResult.Type)
		return
	}
	secret = secretResult.Data[corev1.SSHAuthPrivateKey]
	return
}

func DeleteSSHSecret(name string, namespace string) (err error) {
	return DeleteSecret(name, namespace)
}

func CreateSSHSecret(name string, namespace string, privateKey []byte) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, err = GenerateKubernetesConfig()
		if err != nil {
			return fmt.Errorf("error generating config: %v", err)
		}
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	labelMap := make(map[string]string)
	labelMap["cmaaks"] = "true"

	dataMap := make(map[string][]byte)
	dataMap[corev1.SSHAuthPrivateKey] = privateKey

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Labels: labelMap},
		Type:       corev1.SecretTypeSSHAuth,
		Data:       dataMap,
	}

	_, err = client.CoreV1().Secrets(namespace).Create(secret)
	return
}

func GetKubeconfigSecretList(namespace string) (result []corev1.Secret, err error) {
	listOption := v1.ListOptions{
		FieldSelector: "type=" + string(corev1.SecretTypeOpaque),
		LabelSelector: "kubeconfig=true",
	}
	return GetSecretList(namespace, listOption)
}

func GetKubeconfigSecret(name string, namespace string) (secret []byte, err error) {
	secretResult, err := GetSecret(name, namespace)
	if err != nil {
		return
	}
	if secretResult.Type != corev1.SecretTypeOpaque {
		err = fmt.Errorf("secret %s is not of type %s, but rather is of type %s", name, corev1.SecretTypeOpaque, secretResult.Type)
		return
	}
	if secretResult.Labels["kubeconfig"] != "true" {
		err = fmt.Errorf("secret %s does not have label kubeconfig=true", name)
		return
	}
	secret = secretResult.Data[corev1.ServiceAccountKubeconfigKey]
	return
}

func DeleteKubeconfigSecret(name string, namespace string) (err error) {
	return DeleteSecret(name, namespace)
}

func CreateKubeconfigSecret(name string, namespace string, kubeconfig []byte) (err error) {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return
	}

	labelMap := make(map[string]string)
	labelMap["cmaaks"] = "true"
	labelMap["kubeconfig"] = "true"

	dataMap := make(map[string][]byte)
	dataMap[corev1.ServiceAccountKubeconfigKey] = kubeconfig

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Labels: labelMap},
		Type:       corev1.SecretTypeOpaque,
		Data:       dataMap,
	}

	_, err = client.CoreV1().Secrets(namespace).Create(secret)
	return
}

func CreateAutoScaleSecret(name string, namespace string, data map[string][]byte) error {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return fmt.Errorf("error getting kubeconfig: %v", err)
	}
	labelMap := make(map[string]string)
	labelMap["cmaaks"] = "true"
	labelMap["clusterautoscaling"] = "true"

	secret := &corev1.Secret{
		ObjectMeta: v1.ObjectMeta{Name: name, Labels: labelMap},
		Type:       corev1.SecretTypeOpaque,
		Data:       data,
	}
	_, err = client.CoreV1().Secrets(namespace).Create(secret)
	if err != nil {
		return fmt.Errorf("error creating cluster autoscale secret: %v", err)
	}
	return nil
}

func CreateAutoScaleDeployment(agentPool string, min int32, max int32) error {
	if DefaultConfig == nil {
		DefaultConfig, _ = GenerateKubernetesConfig()
	}
	client, err := kubernetes.NewForConfig(DefaultConfig)
	if err != nil {
		return fmt.Errorf("error getting kubeconfig: %v", err)
	}

	// ServiceAccount
	serviceAccount := &corev1.ServiceAccount{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"k8s-addon": "cluster-autoscaler.addons.k8s.io", "k8s-app": "cluster-autoscaler"},
		},
	}
	_, err = client.CoreV1().ServiceAccounts("kube-system").Create(serviceAccount)
	if err != nil {
		if err.Error() != `serviceaccounts "cluster-autoscaler" already exists` {
			return fmt.Errorf("error creating cluster autoscale deployment: %v", err)
		}
	}

	// ClusterRole
	clusterRole := &rbacv1.ClusterRole{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"k8s-addon": "cluster-autoscaler.addons.k8s.io", "k8s-app": "cluster-autoscaler"},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"events", "endpoints"},
				Verbs:     []string{"create", "patch"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"pods/eviction"},
				Verbs:     []string{"create"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"pods/status"},
				Verbs:     []string{"update"},
			},
			{
				APIGroups:     []string{""},
				Resources:     []string{"endpoints"},
				ResourceNames: []string{"cluster-autoscaler"},
				Verbs:         []string{"get", "update"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"nodes"},
				Verbs:     []string{"watch", "list", "get", "update"},
			},
			{
				APIGroups: []string{""},
				Resources: []string{"pods", "services", "replicationcontrollers", "persistentvolumeclaims", "persistentvolumes"},
				Verbs:     []string{"watch", "list", "get"},
			},
			{
				APIGroups: []string{"extensions"},
				Resources: []string{"replicasets", "daemonsets"},
				Verbs:     []string{"watch", "list", "get"},
			},
			{
				APIGroups: []string{"policy"},
				Resources: []string{"poddisruptionbudgets"},
				Verbs:     []string{"watch", "list"},
			},
			{
				APIGroups: []string{"apps"},
				Resources: []string{"statefulsets"},
				Verbs:     []string{"watch", "list", "get"},
			},
			{
				APIGroups: []string{"storage.k8s.io"},
				Resources: []string{"storageclasses"},
				Verbs:     []string{"get", "list", "watch"},
			},
		},
	}
	_, err = client.Rbac().ClusterRoles().Create(clusterRole)
	if err != nil {
		if err.Error() != `clusterroles.rbac.authorization.k8s.io "cluster-autoscaler" already exists` {
			return fmt.Errorf("err creating cluster role for autoscaling: %v", err)
		}
	}

	// Role
	role := &rbacv1.Role{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"k8s-addon": "cluster-autoscaler.addons.k8s.io", "k8s-app": "cluster-autoscaler"},
		},
		Rules: []rbacv1.PolicyRule{
			{
				APIGroups: []string{""},
				Resources: []string{"configmaps"},
				Verbs:     []string{"create"},
			},
			{
				APIGroups:     []string{""},
				Resources:     []string{"configmaps"},
				ResourceNames: []string{"cluster-autoscaler-status"},
				Verbs:         []string{"delete", "get", "update"},
			},
		},
	}
	_, err = client.RbacV1().Roles("kube-system").Create(role)
	if err != nil {
		if err.Error() != `roles.rbac.authorization.k8s.io "cluster-autoscaler" already exists` {
			return fmt.Errorf("error creating cluster autoscale role: %v", err)
		}
	}
	// ClusterRoleBinding
	clusterRoleBinding := &rbacv1.ClusterRoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"k8s-addon": "cluster-autoscaler.addons.k8s.io", "k8s-app": "cluster-autoscaler"},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "ClusterRole",
			Name:     "cluster-autoscaler",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "cluster-autoscaler",
				Namespace: "kube-system",
			},
		},
	}
	_, err = client.RbacV1().ClusterRoleBindings().Create(clusterRoleBinding)
	if err != nil {
		if err.Error() != `clusterrolebindings.rbac.authorization.k8s.io "cluster-autoscaler" already exists` {
			return fmt.Errorf("error creating cluster autoscale cluster role binding: %v", err)
		}
	}
	// RoleBinding
	roleBinding := &rbacv1.RoleBinding{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"k8s-addon": "cluster-autoscaler.addons.k8s.io", "k8s-app": "cluster-autoscaler"},
		},
		RoleRef: rbacv1.RoleRef{
			APIGroup: "rbac.authorization.k8s.io",
			Kind:     "Role",
			Name:     "cluster-autoscaler",
		},
		Subjects: []rbacv1.Subject{
			{
				Kind:      "ServiceAccount",
				Name:      "cluster-autoscaler",
				Namespace: "kube-system",
			},
		},
	}
	_, err = client.RbacV1().RoleBindings("kube-system").Create(roleBinding)
	if err != nil {
		if err.Error() != `rolebindings.rbac.authorization.k8s.io "cluster-autoscaler" already exists` {
			return fmt.Errorf("error creating cluster autoscale role binding: %v", err)
		}
	}
	// Deployment:
	deployment := &appsv1.Deployment{
		ObjectMeta: v1.ObjectMeta{
			Name:   "cluster-autoscaler",
			Labels: map[string]string{"app": "cluster-autoscaler"},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: to.Int32Ptr(1),
			Selector: &v1.LabelSelector{
				MatchLabels: map[string]string{"app": "cluster-autoscaler"},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Labels: map[string]string{"app": "cluster-autoscaler"},
				},
				Spec: corev1.PodSpec{
					ServiceAccountName: "cluster-autoscaler",
					RestartPolicy:      "Always",
					Containers: []corev1.Container{
						{
							Image:           "gcr.io/google-containers/cluster-autoscaler:v1.3.3",
							ImagePullPolicy: "Always",
							Name:            "cluster-autoscaler",
							// Resources: corev1.ResourceRequirements{
							// 	Limits: corev1.ResourceList{

							// 		"cpu": resource.Quantity{
							// 			Format: "100m",
							// 		},
							// 	},
							// },
							Command: []string{
								"./cluster-autoscaler",
								"--v=3",
								"--logtostderr=true",
								"--cloud-provider=azure",
								"--skip-nodes-with-local-storage=false",
								"--nodes=" + fmt.Sprint(min) + ":" + fmt.Sprint(max) + ":" + agentPool,
							},
							Env: []corev1.EnvVar{
								{
									Name: "ARM_SUBSCRIPTION_ID",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "SubscriptionID",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "ARM_RESOURCE_GROUP",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "ResourceGroup",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "ARM_TENANT_ID",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "TenantID",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "ARM_CLIENT_ID",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "ClientID",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "ARM_CLIENT_SECRET",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "ClientSecret",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "ARM_VM_TYPE",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "VMType",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "AZURE_CLUSTER_NAME",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "ClusterName",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
								{
									Name: "AZURE_NODE_RESOURCE_GROUP",
									ValueFrom: &corev1.EnvVarSource{
										SecretKeyRef: &corev1.SecretKeySelector{
											Key: "NodeResourceGroup",
											LocalObjectReference: corev1.LocalObjectReference{
												Name: "cluster-autoscaler-azure",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	_, err = client.AppsV1().Deployments("kube-system").Create(deployment)
	if err != nil {
		if err.Error() != `deployments.apps "cluster-autoscaler" already exists` {
			return fmt.Errorf("error creating cluster autoscale deployment: %v", err)
		}
	}
	return nil
}
