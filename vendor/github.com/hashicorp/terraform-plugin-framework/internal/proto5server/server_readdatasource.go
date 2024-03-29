// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package proto5server

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/internal/fromproto5"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwserver"
	"github.com/hashicorp/terraform-plugin-framework/internal/logging"
	"github.com/hashicorp/terraform-plugin-framework/internal/toproto5"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
)

// ReadDataSource satisfies the tfprotov5.ProviderServer interface.
func (s *Server) ReadDataSource(ctx context.Context, proto5Req *tfprotov5.ReadDataSourceRequest) (*tfprotov5.ReadDataSourceResponse, error) {
	ctx = s.registerContext(ctx)
	ctx = logging.InitContext(ctx)

	fwResp := &fwserver.ReadDataSourceResponse{}

	dataSource, diags := s.FrameworkServer.DataSource(ctx, proto5Req.TypeName)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto5.ReadDataSourceResponse(ctx, fwResp), nil
	}

	dataSourceSchema, diags := s.FrameworkServer.DataSourceSchema(ctx, proto5Req.TypeName)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto5.ReadDataSourceResponse(ctx, fwResp), nil
	}

	providerMetaSchema, diags := s.FrameworkServer.ProviderMetaSchema(ctx)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto5.ReadDataSourceResponse(ctx, fwResp), nil
	}

	fwReq, diags := fromproto5.ReadDataSourceRequest(ctx, proto5Req, dataSource, dataSourceSchema, providerMetaSchema)

	fwResp.Diagnostics.Append(diags...)

	if fwResp.Diagnostics.HasError() {
		return toproto5.ReadDataSourceResponse(ctx, fwResp), nil
	}

	s.FrameworkServer.ReadDataSource(ctx, fwReq, fwResp)

	return toproto5.ReadDataSourceResponse(ctx, fwResp), nil
}
