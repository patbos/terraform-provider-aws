// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/listpages/main.go -ListOps=DescribeDirectConnectGateways,DescribeDirectConnectGatewayAssociations,DescribeDirectConnectGatewayAssociationProposals
//go:generate go run ../../generate/tags/main.go -ListTags -ListTagsOp=DescribeTags -ListTagsInIDElem=ResourceArns -ListTagsInIDNeedSlice=yes -ListTagsOutTagsElem=ResourceTags[0].Tags -ServiceTagsSlice -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package directconnect
