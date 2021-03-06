/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/*
Package loadbalancer contains external loadbalancer related constants and configuration.

Having a direct control on loadbalancer config is a specific necessity for kinder, because
in kinder all the actions for setting up a working cluster can happen
at different time while in kind everything - from create to a working K8s cluster -
happens within an atomic operation.

Nevertheless, kinder interally relies on loadbalancer config implemented in kind (temporary
from a fork of kind an internal package)e, but this will be hopefully addressed).
*/
package loadbalancer
