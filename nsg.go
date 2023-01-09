package network

import (
	"context"
	"fmt"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/services/internal/iam"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-11-01/network"
	"github.com/Azure/go-autorest/autorest/to"
)

func (sg SecurityGroupClient) CreateNetworkSecurityGroup(
	name string,
	label string,
	location string) (management.OperationID, error)