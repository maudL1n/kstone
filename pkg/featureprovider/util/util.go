/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2023 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

package util

import (
	"strconv"
	"strings"
	kstoneapiv1 "tkestack.io/kstone/pkg/apis/kstone/v1alpha1"
)

func IsFeatureGateEnabled(annotations map[string]string, feature string) bool {
	if gates, found := annotations[kstoneapiv1.KStoneFeatureAnno]; found && gates != "" {
		features := strings.Split(gates, ",")
		for _, f := range features {
			ff := strings.Split(f, "=")
			if len(ff) != 2 {
				continue
			}

			g, _ := strconv.ParseBool(ff[1])
			if ff[0] == feature && g {
				return true
			}
		}
	}
	return false
}
