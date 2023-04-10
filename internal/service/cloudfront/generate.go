// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate go run ../../generate/tags/main.go -ListTags -ListTagsInIDElem=Resource -ListTagsOutTagsElem=Tags.Items -ServiceTagsSlice "-TagInCustomVal=&cloudfront.Tags{Items: Tags(updatedTags.IgnoreAWS())}" -TagInIDElem=Resource "-UntagInCustomVal=&cloudfront.TagKeys{Items: aws.StringSlice(removedTags.IgnoreAWS().Keys())}" -UpdateTags
// ONLY generate directives and package declaration! Do not add anything else to this file.

package cloudfront
