// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package redshiftdata

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/redshiftdataapiservice"
	"github.com/hashicorp/aws-sdk-go-base/v2/awsv1shim/v2/tfawserr"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func FindStatementByID(ctx context.Context, conn *redshiftdataapiservice.RedshiftDataAPIService, id string) (*redshiftdataapiservice.DescribeStatementOutput, error) {
	input := &redshiftdataapiservice.DescribeStatementInput{
		Id: aws.String(id),
	}

	output, err := conn.DescribeStatementWithContext(ctx, input)

	if tfawserr.ErrCodeEquals(err, redshiftdataapiservice.ErrCodeResourceNotFoundException) {
		return nil, &retry.NotFoundError{
			LastError:   err,
			LastRequest: input,
		}
	}

	if err != nil {
		return nil, err
	}

	if output == nil {
		return nil, tfresource.NewEmptyResultError(input)
	}

	return output, nil
}
