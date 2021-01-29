/*

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

package base

import (
	datav1alpha1 "github.com/fluid-cloudnative/fluid/api/v1alpha1"
	cruntime "github.com/fluid-cloudnative/fluid/pkg/runtime"
)

// Load the data
func (t *TemplateEngine) LoadData(ctx cruntime.ReconcileRequestContext, targetDataload datav1alpha1.DataLoad) (string, string, error) {
	return t.Implement.LoadData(ctx, targetDataload)
}

// Check if the runtime is ready
func (t *TemplateEngine) Ready(ctx cruntime.ReconcileRequestContext) bool{
	return t.Implement.Ready(ctx)
}
