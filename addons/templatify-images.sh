#!/usr/bin/env bash

# Copyright 2022 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -euo pipefail

dqreplacement='"{{ Image \1 }}"'
sqreplacement=$'\'{{ Image "\1" }}\''
replacement='{{ Image "\1" }}'

for registry in docker.io quay.io public.ecr.aws; do
  sed -i "s/\(\"$registry.*\"\)/$dqreplacement/g" $@
  sed -i "s/\('$registry.*'\)/$sqreplacement/g" $@

  # the space makes it so that this does not match the
  # output of the previous two sed invocations
  sed -i "s/ \($registry.*\)/ $replacement/g" $@
done
