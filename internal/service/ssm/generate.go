// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/tags/main.go -ListTags -ListTagsInIDElem=ResourceId -ListTagsOutTagsElem=TagList -ServiceTagsSlice -TagOp=AddTagsToResource -TagInIDElem=ResourceId -TagResTypeElem=ResourceType -UntagOp=RemoveTagsFromResource -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package ssm
