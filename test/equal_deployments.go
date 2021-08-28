package test

func EqualDeploymentDesired() map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app.kubernetes.io/name": "test",
			},
			"namespace": "equal",
			"name":      "equal",
		},
		"spec": map[string]interface{}{
			"replicas": 1,
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app.kubernetes.io/name": "test",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app.kubernetes.io/name": "test",
					},
					"name": "test",
				},
				"spec": map[string]interface{}{
					"containers": []interface{}{
						map[string]interface{}{
							"name":  "test",
							"image": "test",
						},
					},
				},
			},
		},
	}
}

func EqualDeploymentActual() map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app.kubernetes.io/name": "test",
			},
			"namespace": "equal",
			"name":      "equal",
		},
		"spec": map[string]interface{}{
			"replicas": 1,
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app.kubernetes.io/name": "test",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app.kubernetes.io/name": "test",
					},
					"name": "test",
				},
				"spec": map[string]interface{}{
					"serviceAccountName": "test",
					"containers": []interface{}{
						map[string]interface{}{
							"name":  "test",
							"image": "test",
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									"memory": "26m",
									"cpu":    "50",
								},
								"limits": map[string]interface{}{
									"memory": "25m",
									"cpu":    "50",
								},
							},
							"securityContext": map[string]interface{}{
								"runAsNonRoot": true,
								"runAsUser":    65532,
								"runAsGroup":   65532,
							},
						},
					},
				},
			},
		},
	}
}

func InequalDeployment() map[string]interface{} {
	return map[string]interface{}{
		"apiVersion": "apps/v1",
		"kind":       "Deployment",
		"metadata": map[string]interface{}{
			"labels": map[string]interface{}{
				"app.kubernetes.io/name": "testdata",
			},
			"namespace": "inequal",
			"name":      "equal",
		},
		"spec": map[string]interface{}{
			"replicas": 5,
			"selector": map[string]interface{}{
				"matchLabels": map[string]interface{}{
					"app.kubernetes.io/name": "test",
				},
			},
			"template": map[string]interface{}{
				"metadata": map[string]interface{}{
					"labels": map[string]interface{}{
						"app.kubernetes.io/name": "test",
					},
					"name": "test",
				},
				"spec": map[string]interface{}{
					"serviceAccountName": "test",
					"containers": []interface{}{
						map[string]interface{}{
							"name":  "test",
							"image": "test",
							"resources": map[string]interface{}{
								"requests": map[string]interface{}{
									"memory": "64MB",
									"cpu":    "51m",
								},
								"limits": map[string]interface{}{
									"memory": "25m",
									"cpu":    "50",
								},
							},
							"securityContext": map[string]interface{}{
								"runAsNonRoot": true,
								"runAsUser":    65532,
								"runAsGroup":   65532,
							},
						},
					},
				},
			},
		},
	}
}
