//go:build !generate
// +build !generate

package route53

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

// Custom Route 53 service lister functions using the same format as generated code.

//nolint:unused // This function is called from a sweeper.
func listTrafficPolicyInstancesPages(conn *route53.Route53, input *route53.ListTrafficPolicyInstancesInput, fn func(*route53.ListTrafficPolicyInstancesOutput, bool) bool) error {
	return listTrafficPolicyInstancesPagesWithContext(context.Background(), conn, input, fn)
}

//nolint:unused // This function is called from a sweeper.
func listTrafficPolicyInstancesPagesWithContext(ctx context.Context, conn *route53.Route53, input *route53.ListTrafficPolicyInstancesInput, fn func(*route53.ListTrafficPolicyInstancesOutput, bool) bool) error {
	for {
		output, err := conn.ListTrafficPolicyInstancesWithContext(ctx, input)
		if err != nil {
			return err
		}

		lastPage := !aws.BoolValue(output.IsTruncated)
		if !fn(output, lastPage) || lastPage {
			break
		}

		input.HostedZoneIdMarker = output.HostedZoneIdMarker
		input.TrafficPolicyInstanceNameMarker = output.TrafficPolicyInstanceNameMarker
		input.TrafficPolicyInstanceTypeMarker = output.TrafficPolicyInstanceTypeMarker
	}
	return nil
}
