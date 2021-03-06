// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package servicemanagement_test

import (
	"context"

	servicemanagement "cloud.google.com/go/servicemanagement/apiv1"
	"google.golang.org/api/iterator"
	servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
)

func ExampleNewServiceManagerClient() {
	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleServiceManagerClient_ListServices() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.ListServicesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServices(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleServiceManagerClient_GetService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.GetServiceRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_CreateService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.CreateServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_DeleteService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.DeleteServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	err = op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleServiceManagerClient_UndeleteService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.UndeleteServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.UndeleteService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_ListServiceConfigs() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.ListServiceConfigsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServiceConfigs(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleServiceManagerClient_GetServiceConfig() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.GetServiceConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetServiceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_CreateServiceConfig() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.CreateServiceConfigRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateServiceConfig(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_SubmitConfigSource() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.SubmitConfigSourceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.SubmitConfigSource(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_ListServiceRollouts() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.ListServiceRolloutsRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListServiceRollouts(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleServiceManagerClient_GetServiceRollout() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.GetServiceRolloutRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetServiceRollout(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_CreateServiceRollout() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.CreateServiceRolloutRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.CreateServiceRollout(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_GenerateConfigReport() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.GenerateConfigReportRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GenerateConfigReport(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_EnableService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.EnableServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.EnableService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleServiceManagerClient_DisableService() {
	// import servicemanagementpb "google.golang.org/genproto/googleapis/api/servicemanagement/v1"

	ctx := context.Background()
	c, err := servicemanagement.NewServiceManagerClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &servicemanagementpb.DisableServiceRequest{
		// TODO: Fill request struct fields.
	}
	op, err := c.DisableService(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}

	resp, err := op.Wait(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}
