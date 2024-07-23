// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package tailsamplingprocessor // import "github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor"

import (
	"go.opentelemetry.io/collector/component"
	"go.uber.org/zap"

	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/tailsamplingprocessor/internal/sampling"
)

func getNewAndPolicy(settings component.TelemetrySettings, config *AndCfg) (sampling.PolicyEvaluator, error) {
	settings.Logger.Debug("HELLO FROM AND_HELPER")
	subPolicyEvaluators := make([]sampling.PolicyEvaluator, len(config.SubPolicyCfg))
	for i := range config.SubPolicyCfg {
		policyCfg := &config.SubPolicyCfg[i]
		settings.Logger.Debug("SUBPOLICY NAME:", zap.Any("subpolicy_name", policyCfg.Name))
		policy, err := getAndSubPolicyEvaluator(settings, policyCfg)
		if err != nil {
			return nil, err
		}
		subPolicyEvaluators[i] = policy
	}
	return sampling.NewAnd(settings.Logger, subPolicyEvaluators), nil
}

// Return instance of and sub-policy
func getAndSubPolicyEvaluator(settings component.TelemetrySettings, cfg *AndSubPolicyCfg) (sampling.PolicyEvaluator, error) {
	return getSharedPolicyEvaluator(settings, &cfg.sharedPolicyCfg)
}
