#!/usr/bin/env bash

# Copyright 2019 The Skaffold Authors
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

echo "Please enter new config version:"
read NEW_VERSION

go run ./hack/new_config_version/version.go ${NEW_VERSION}

goimports -w ./pkg/skaffold/schema
make generate-schemas
git --no-pager diff --minimal
make test

echo
echo "---------------------------------------"
echo
echo "Files generated for $NEW_VERSION."
echo "All tests should have passed. For the docs change, commit the results and rerun 'make test'."
echo "Please double check manually the generated files as well: the upgrade functionality, and all the examples:"
echo
git status -s
echo
echo "---------------------------------------"
