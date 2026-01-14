/*
 * @Author: lwnmengjing<lwnmengjing@qq.com>
 * @Date: 2022/10/8 15:57:10
 * @Last Modified by: lwnmengjing<lwnmengjing@qq.com>
 * @Last Modified time: 2022/10/8 15:57:10
 */

package cdk8s

import (
	"fmt"
	"github.com/WangDe7/cd-template/pkg/config"
	"github.com/WangDe7/cd-template/stage"
)

// Generate the k8s yaml file
func Generate(configPath, stageStr, image string, servicePath []string) {
	config.NewConfig(&configPath)
	switch stageStr {
	case "dev", "test", "local", "alpha", "beta", "staging":
		config.Cfg.Hpa.Enabled = false
		if config.Cfg.Resources == nil {
			config.Cfg.Resources = make(map[string]config.Resource)
			config.Cfg.Resources["limits"] = config.Resource{
				CPU:    "1",
				Memory: "2Gi",
			}
		}
		config.Cfg.Resources["requests"] = config.Resource{
			CPU:    "100m",
			Memory: "128Mi",
		}
		config.Cfg.Replicas = config.Cfg.TestReplicas
	}
	config.Cfg.Image.Path = image
	fmt.Printf("************************* %s *************************\n", "configmap")
	fmt.Println(config.Cfg.ConfigmapResource)
	stage.Synth(stageStr, servicePath...)
}
