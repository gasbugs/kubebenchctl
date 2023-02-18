package diagnostics

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"

	batchv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type CheckConfig struct {
	Text     string `json:"text"`
	Audit    string `json:"audit"`
	Type     string `json:"type"`
	ID       string `json:"id"`
	Level    string `json:"level"`
	Message  string `json:"message"`
	Metadata struct {
		Remediation string `json:"remediation"`
		CisVersion  string `json:"cis_version"`
	} `json:"metadata"`
}

func RunKubeBench(ctx context.Context, clientset *kubernetes.Clientset, node string) ([]CheckConfig, error) {
	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "kube-bench-",
			Namespace:    "default",
		},
		Spec: batchv1.JobSpec{
			Template: corev1.PodTemplateSpec{
				Spec: corev1.PodSpec{
					RestartPolicy: "Never",
					NodeSelector: map[string]string{
						"kubernetes.io/hostname": node,
					},
					Containers: []corev1.Container{
						{
							Name:            "kube-bench",
							Image:           "aquasec/kube-bench:latest",
							ImagePullPolicy: "Always",
							Command: []string{
								"/usr/local/bin/kube-bench",
								"--json",
							},
						},
					},
				},
			},
		},
	}

	job, err := clientset.BatchV1().Jobs("default").Create(ctx, job, metav1.CreateOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to create kube-bench job: %v", err)
	}

	if err = waitUntilJobCompleted(ctx, clientset, job); err != nil {
		return nil, fmt.Errorf("kube-bench job failed: %v", err)
	}

	podName, err := getPodName(ctx, clientset, job)
	if err != nil {
		return nil, fmt.Errorf("failed to get kube-bench pod name: %v", err)
	}

	pod, err := clientset.CoreV1().Pods("default").Get(ctx, podName, metav1.GetOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get kube-bench pod: %v", err)
	}

	if pod.Status.Phase != corev1.PodSucceeded {
		return nil, fmt.Errorf("kube-bench pod failed: %v", pod.Status.Message)
	}

	stdout := pod.Status.ContainerStatuses[0].State.Terminated.Stdout

	var results []CheckConfig
	if err := json.NewDecoder(strings.NewReader(stdout)).Decode(&results); err != nil {
		return nil, fmt.Errorf("failed to decode kube-bench results: %v", err)
	}

	return results, nil
}

