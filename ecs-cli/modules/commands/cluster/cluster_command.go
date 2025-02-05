// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// package clusterCommand defines all the commands for cluster workflows
package clusterCommand

import (
	ecscli "github.com/aws/amazon-ecs-cli/ecs-cli/modules"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/cli/cluster"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/commands/flags"
	"github.com/aws/amazon-ecs-cli/ecs-cli/modules/commands/usage"
	"github.com/urfave/cli"
)

func UpCommand() cli.Command {
	return cli.Command{
		Name:         "up",
		Usage:        usage.ClusterUp,
		Before:       ecscli.BeforeApp,
		Action:       cluster.ClusterUp,
		Flags:        flags.AppendFlags(clusterUpFlags(), flags.OptionalConfigFlags(), flags.OptionalLaunchTypeFlag(), flags.DebugFlag()),
		OnUsageError: flags.UsageErrorFactory("up"),
	}
}

func DownCommand() cli.Command {
	return cli.Command{
		Name:         "down",
		Usage:        usage.ClusterDown,
		Action:       cluster.ClusterDown,
		Flags:        flags.AppendFlags(clusterDownFlags(), flags.OptionalConfigFlags()),
		OnUsageError: flags.UsageErrorFactory("down"),
	}
}

func ScaleCommand() cli.Command {
	return cli.Command{
		Name:         "scale",
		Usage:        usage.ClusterScale,
		Action:       cluster.ClusterScale,
		Flags:        flags.AppendFlags(clusterScaleFlags(), flags.OptionalConfigFlags()),
		OnUsageError: flags.UsageErrorFactory("scale"),
	}
}

func PsCommand() cli.Command {
	return cli.Command{
		Name:         "ps",
		Usage:        usage.ClusterPs,
		Action:       cluster.ClusterPS,
		Flags:        flags.AppendFlags(flags.OptionalConfigFlags(), flags.OptionalDesiredStatusFlag()),
		OnUsageError: flags.UsageErrorFactory("ps"),
	}
}

func clusterUpFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  flags.CapabilityIAMFlag,
			Usage: "Acknowledges that this command may create IAM resources. Required if --instance-role is not specified. NOTE: Not applicable for launch type FARGATE or when creating an empty cluster.",
		},
		cli.BoolFlag{
			Name:  flags.EmptyFlag + ",e",
			Usage: "[Optional] Specifies that an ECS cluster will be created with no resources.",
		},
		cli.StringFlag{
			Name:  flags.InstanceRoleFlag,
			Usage: "[Optional] Specifies a custom IAM Role for instances in your cluster. Required if --capability-iam is not specified. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.StringFlag{
			Name:  flags.KeypairNameFlag,
			Usage: "[Optional] Specifies the name of an existing Amazon EC2 key pair to enable SSH access to the EC2 instances in your cluster. Recommended for EC2 launch type. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.StringFlag{
			Name:  flags.InstanceTypeFlag,
			Usage: "[Optional] Specifies the EC2 instance type for your container instances. If you specify the A1 instance family, the ECS optimized arm64 AMI will be used, otherwise the x86 AMI will be used. Defaults to t2.micro. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.StringFlag{
			Name:  flags.SpotPriceFlag,
			Usage: "[Optional] If filled and greater than 0, EC2 Spot instances will be requested.",
		},
		cli.StringFlag{
			Name:  flags.ImageIdFlag,
			Usage: "[Optional] Specify the AMI ID for your container instances. Defaults to amazon-ecs-optimized AMI. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.BoolFlag{
			Name:  flags.NoAutoAssignPublicIPAddressFlag,
			Usage: "[Optional] Do not assign public IP addresses to new instances in this VPC. Unless this option is specified, new instances in this VPC receive an automatically assigned public IP address. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.StringFlag{
			Name:  flags.AsgMaxSizeFlag,
			Usage: "[Optional] Specifies the number of instances to launch and register to the cluster. Defaults to 1. NOTE: Not applicable for launch type FARGATE.",
		},
		cli.StringFlag{
			Name:  flags.VpcAzFlag,
			Usage: "[Optional] Specifies a comma-separated list of 2 VPC Availability Zones in which to create subnets (these zones must have the available status). This option is recommended if you do not specify a VPC ID with the --vpc option. WARNING: Leaving this option blank can result in failure to launch container instances if an unavailable zone is chosen at random.",
		},
		cli.StringFlag{
			Name:  flags.SecurityGroupFlag,
			Usage: "[Optional] Specifies a comma-separated list of existing security groups to associate with your container instances. If you do not specify a security group here, then a new one is created.",
		},
		cli.StringFlag{
			Name:  flags.SourceCidrFlag,
			Usage: "[Optional] Specifies a CIDR/IP range for the security group to use for container instances in your cluster. This parameter is ignored if an existing security group is specified with the --security-group option. Defaults to 0.0.0.0/0.",
		},
		cli.StringFlag{
			Name:  flags.EcsPortFlag,
			Usage: "[Optional] Specifies a port to open on the security group to use for container instances in your cluster. This parameter is ignored if an existing security group is specified with the --security-group option. Defaults to port 80.",
		},
		cli.StringFlag{
			Name:  flags.SubnetIdsFlag,
			Usage: "[Optional] Specifies a comma-separated list of existing VPC Subnet IDs in which to launch your container instances. This option is required if you specify a VPC with the --vpc option.",
		},
		cli.StringFlag{
			Name:  flags.VpcIdFlag,
			Usage: "[Optional] Specifies the ID of an existing VPC in which to launch your container instances. If you specify a VPC ID, you must specify a list of existing subnets in that VPC with the --subnets option. If you do not specify a VPC ID, a new VPC is created with two subnets.",
		},
		cli.StringSliceFlag{
			Name:  flags.UserDataFlag,
			Usage: "[Optional] Specifies additional User Data for your EC2 instances. Files can be shell scripts or cloud-init directives and are packaged into a MIME Multipart Archive along with ECS CLI provided User Data which directs instances to join your cluster.",
			Value: &cli.StringSlice{},
		},
		cli.BoolFlag{
			Name:  flags.ForceFlag + ", f",
			Usage: "[Optional] Forces the recreation of any existing resources that match your current configuration. This option is useful for cleaning up stale resources from previous failed attempts.",
		},
		cli.StringFlag{
			Name:  flags.ResourceTagsFlag,
			Usage: "[Optional] Specify tags which will be added to AWS Resources created for your cluster. Specify in the format 'key1=value1,key2=value2,key3=value3'",
		},
	}
}

func clusterDownFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  flags.ForceFlag + ", f",
			Usage: "[Optional] Acknowledges that this command permanently deletes resources.",
		},
	}
}

func clusterScaleFlags() []cli.Flag {
	return []cli.Flag{
		cli.BoolFlag{
			Name:  flags.CapabilityIAMFlag,
			Usage: "Acknowledges that this command may create IAM resources.",
		},
		cli.StringFlag{
			Name:  flags.AsgMaxSizeFlag,
			Usage: "Specifies the number of instances to maintain in your cluster.",
		},
	}
}
