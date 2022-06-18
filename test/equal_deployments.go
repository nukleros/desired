package test

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func StructuredUnorderedDesired() map[interface{}]interface{} {
	file, err := ioutil.ReadFile("test/yaml/1-unordered-array-1.yaml")
	if err != nil {
		panic(err)
	}

	data := make(map[interface{}]interface{})

	if err := yaml.Unmarshal(file, &data); err != nil {
		panic(err)
	}

	return data
}

func StructuredUnorderedActual() map[interface{}]interface{} {
	file, err := ioutil.ReadFile("test/yaml/1-unordered-array-2.yaml")
	if err != nil {
		panic(err)
	}

	data := make(map[interface{}]interface{})

	if err := yaml.Unmarshal(file, &data); err != nil {
		panic(err)
	}

	return data
}

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

func NilDeployment() map[string]interface{} {
	return nil
}
