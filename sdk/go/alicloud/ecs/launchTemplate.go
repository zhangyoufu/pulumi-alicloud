// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ecs

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides an ECS Launch Template resource.
//
// For information about Launch Template and how to use it, see [Launch Template](https://www.alibabacloud.com/help/doc-detail/73916.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-alicloud/sdk/v2/go/alicloud/ecs"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		opt0 := "system"
// 		images, err := ecs.GetImages(ctx, &ecs.GetImagesArgs{
// 			Owners: &opt0,
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		instances, err := ecs.GetInstances(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ecs.NewLaunchTemplate(ctx, "template", &ecs.LaunchTemplateArgs{
// 			DataDisks: ecs.LaunchTemplateDataDiskArray{
// 				&ecs.LaunchTemplateDataDiskArgs{
// 					Description: pulumi.String("test1"),
// 					Name:        pulumi.String("disk1"),
// 				},
// 				&ecs.LaunchTemplateDataDiskArgs{
// 					Description: pulumi.String("test2"),
// 					Name:        pulumi.String("disk2"),
// 				},
// 			},
// 			Description:             pulumi.String("test1"),
// 			HostName:                pulumi.String("tf-test-host"),
// 			ImageId:                 pulumi.String(images.Images[0].Id),
// 			InstanceChargeType:      pulumi.String("PrePaid"),
// 			InstanceName:            pulumi.String("tf-instance-name"),
// 			InstanceType:            pulumi.String(instances.Instances[0].InstanceType),
// 			InternetChargeType:      pulumi.String("PayByBandwidth"),
// 			InternetMaxBandwidthIn:  pulumi.Int(5),
// 			InternetMaxBandwidthOut: pulumi.Int(0),
// 			IoOptimized:             pulumi.String("none"),
// 			KeyPairName:             pulumi.String("test-key-pair"),
// 			NetworkInterfaces: &ecs.LaunchTemplateNetworkInterfacesArgs{
// 				Description:     pulumi.String("hello1"),
// 				Name:            pulumi.String("eth0"),
// 				PrimaryIp:       pulumi.String("10.0.0.2"),
// 				SecurityGroupId: pulumi.String("xxxx"),
// 				VswitchId:       pulumi.String("xxxxxxx"),
// 			},
// 			NetworkType:                 pulumi.String("vpc"),
// 			RamRoleName:                 pulumi.String("xxxxx"),
// 			ResourceGroupId:             pulumi.String("rg-zkdfjahg9zxncv0"),
// 			SecurityEnhancementStrategy: pulumi.String("Active"),
// 			SecurityGroupId:             pulumi.String("sg-zxcvj0lasdf102350asdf9a"),
// 			SpotPriceLimit:              pulumi.Float64(5),
// 			SpotStrategy:                pulumi.String("SpotWithPriceLimit"),
// 			SystemDiskCategory:          pulumi.String("cloud_ssd"),
// 			SystemDiskDescription:       pulumi.String("test disk"),
// 			SystemDiskName:              pulumi.String("hello"),
// 			SystemDiskSize:              pulumi.Int(40),
// 			Tags: pulumi.StringMap{
// 				"tag1": pulumi.String("hello"),
// 				"tag2": pulumi.String("world"),
// 			},
// 			Userdata:  pulumi.String("xxxxxxxxxxxxxx"),
// 			VpcId:     pulumi.String("vpc-asdfnbg0as8dfk1nb2"),
// 			VswitchId: pulumi.String("sw-ljkngaksdjfj0nnasdf"),
// 			ZoneId:    pulumi.String("beijing-a"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type LaunchTemplate struct {
	pulumi.CustomResourceState

	// Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
	AutoReleaseTime pulumi.StringPtrOutput `pulumi:"autoReleaseTime"`
	// The list of data disks created with instance.
	DataDisks LaunchTemplateDataDiskArrayOutput `pulumi:"dataDisks"`
	// The description of the data disk.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
	HostName pulumi.StringPtrOutput `pulumi:"hostName"`
	// Image ID.
	ImageId         pulumi.StringPtrOutput `pulumi:"imageId"`
	ImageOwnerAlias pulumi.StringPtrOutput `pulumi:"imageOwnerAlias"`
	// Billing methods. Optional values:
	// - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
	// - PostPaid: Pay-As-You-Go.
	InstanceChargeType pulumi.StringPtrOutput `pulumi:"instanceChargeType"`
	// The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	InstanceName pulumi.StringPtrOutput `pulumi:"instanceName"`
	// Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
	InstanceType pulumi.StringPtrOutput `pulumi:"instanceType"`
	// Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrOutput `pulumi:"internetChargeType"`
	// The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
	InternetMaxBandwidthIn pulumi.IntOutput `pulumi:"internetMaxBandwidthIn"`
	// Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
	InternetMaxBandwidthOut pulumi.IntPtrOutput `pulumi:"internetMaxBandwidthOut"`
	// Whether it is an I/O-optimized instance or not. Optional values:
	// - none
	// - optimized
	IoOptimized pulumi.StringPtrOutput `pulumi:"ioOptimized"`
	// The name of the key pair.
	// - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
	// - The password logon method for Linux instances is set to forbidden upon initialization.
	KeyPairName pulumi.StringPtrOutput `pulumi:"keyPairName"`
	// The name of the data disk.
	Name pulumi.StringOutput `pulumi:"name"`
	// The list of network interfaces created with instance.
	NetworkInterfaces LaunchTemplateNetworkInterfacesPtrOutput `pulumi:"networkInterfaces"`
	// Network type of the instance. Value options: `classic` | `vpc`.
	NetworkType pulumi.StringPtrOutput `pulumi:"networkType"`
	// The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
	RamRoleName     pulumi.StringPtrOutput `pulumi:"ramRoleName"`
	ResourceGroupId pulumi.StringPtrOutput `pulumi:"resourceGroupId"`
	// Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
	SecurityEnhancementStrategy pulumi.StringPtrOutput `pulumi:"securityEnhancementStrategy"`
	// The security group ID must be one in the same VPC.
	SecurityGroupId pulumi.StringPtrOutput `pulumi:"securityGroupId"`
	// -(Optional) 	Sets the maximum hourly instance price. Supports up to three decimal places.
	SpotPriceLimit pulumi.Float64PtrOutput `pulumi:"spotPriceLimit"`
	// The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
	// - NoSpot: Normal Pay-As-You-Go instance.
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	// - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
	SpotStrategy pulumi.StringPtrOutput `pulumi:"spotStrategy"`
	// The category of the system disk. System disk type. Optional values:
	// - cloud: Basic cloud disk.
	// - cloud_efficiency: Ultra cloud disk.
	// - cloud_ssd: SSD cloud Disks.
	// - ephemeral_ssd: local SSD Disks
	// - cloud_essd: ESSD cloud Disks.
	SystemDiskCategory pulumi.StringPtrOutput `pulumi:"systemDiskCategory"`
	// System disk description. It cannot begin with http:// or https://.
	SystemDiskDescription pulumi.StringPtrOutput `pulumi:"systemDiskDescription"`
	// System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	SystemDiskName pulumi.StringPtrOutput `pulumi:"systemDiskName"`
	// Size of the system disk, measured in GB. Value range: [20, 500].
	SystemDiskSize pulumi.IntPtrOutput `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapOutput `pulumi:"tags"`
	// User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
	Userdata pulumi.StringPtrOutput `pulumi:"userdata"`
	VpcId    pulumi.StringPtrOutput `pulumi:"vpcId"`
	// The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
	VswitchId pulumi.StringPtrOutput `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrOutput `pulumi:"zoneId"`
}

// NewLaunchTemplate registers a new resource with the given unique name, arguments, and options.
func NewLaunchTemplate(ctx *pulumi.Context,
	name string, args *LaunchTemplateArgs, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	if args == nil {
		args = &LaunchTemplateArgs{}
	}
	var resource LaunchTemplate
	err := ctx.RegisterResource("alicloud:ecs/launchTemplate:LaunchTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLaunchTemplate gets an existing LaunchTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLaunchTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LaunchTemplateState, opts ...pulumi.ResourceOption) (*LaunchTemplate, error) {
	var resource LaunchTemplate
	err := ctx.ReadResource("alicloud:ecs/launchTemplate:LaunchTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LaunchTemplate resources.
type launchTemplateState struct {
	// Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
	AutoReleaseTime *string `pulumi:"autoReleaseTime"`
	// The list of data disks created with instance.
	DataDisks []LaunchTemplateDataDisk `pulumi:"dataDisks"`
	// The description of the data disk.
	Description *string `pulumi:"description"`
	// Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
	HostName *string `pulumi:"hostName"`
	// Image ID.
	ImageId         *string `pulumi:"imageId"`
	ImageOwnerAlias *string `pulumi:"imageOwnerAlias"`
	// Billing methods. Optional values:
	// - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
	// - PostPaid: Pay-As-You-Go.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	InstanceName *string `pulumi:"instanceName"`
	// Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
	InstanceType *string `pulumi:"instanceType"`
	// Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
	InternetMaxBandwidthIn *int `pulumi:"internetMaxBandwidthIn"`
	// Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// Whether it is an I/O-optimized instance or not. Optional values:
	// - none
	// - optimized
	IoOptimized *string `pulumi:"ioOptimized"`
	// The name of the key pair.
	// - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
	// - The password logon method for Linux instances is set to forbidden upon initialization.
	KeyPairName *string `pulumi:"keyPairName"`
	// The name of the data disk.
	Name *string `pulumi:"name"`
	// The list of network interfaces created with instance.
	NetworkInterfaces *LaunchTemplateNetworkInterfaces `pulumi:"networkInterfaces"`
	// Network type of the instance. Value options: `classic` | `vpc`.
	NetworkType *string `pulumi:"networkType"`
	// The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
	RamRoleName     *string `pulumi:"ramRoleName"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
	SecurityEnhancementStrategy *string `pulumi:"securityEnhancementStrategy"`
	// The security group ID must be one in the same VPC.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// -(Optional) 	Sets the maximum hourly instance price. Supports up to three decimal places.
	SpotPriceLimit *float64 `pulumi:"spotPriceLimit"`
	// The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
	// - NoSpot: Normal Pay-As-You-Go instance.
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	// - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// The category of the system disk. System disk type. Optional values:
	// - cloud: Basic cloud disk.
	// - cloud_efficiency: Ultra cloud disk.
	// - cloud_ssd: SSD cloud Disks.
	// - ephemeral_ssd: local SSD Disks
	// - cloud_essd: ESSD cloud Disks.
	SystemDiskCategory *string `pulumi:"systemDiskCategory"`
	// System disk description. It cannot begin with http:// or https://.
	SystemDiskDescription *string `pulumi:"systemDiskDescription"`
	// System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	SystemDiskName *string `pulumi:"systemDiskName"`
	// Size of the system disk, measured in GB. Value range: [20, 500].
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
	Userdata *string `pulumi:"userdata"`
	VpcId    *string `pulumi:"vpcId"`
	// The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

type LaunchTemplateState struct {
	// Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
	AutoReleaseTime pulumi.StringPtrInput
	// The list of data disks created with instance.
	DataDisks LaunchTemplateDataDiskArrayInput
	// The description of the data disk.
	Description pulumi.StringPtrInput
	// Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
	HostName pulumi.StringPtrInput
	// Image ID.
	ImageId         pulumi.StringPtrInput
	ImageOwnerAlias pulumi.StringPtrInput
	// Billing methods. Optional values:
	// - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
	// - PostPaid: Pay-As-You-Go.
	InstanceChargeType pulumi.StringPtrInput
	// The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	InstanceName pulumi.StringPtrInput
	// Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
	InstanceType pulumi.StringPtrInput
	// Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
	InternetMaxBandwidthIn pulumi.IntPtrInput
	// Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// Whether it is an I/O-optimized instance or not. Optional values:
	// - none
	// - optimized
	IoOptimized pulumi.StringPtrInput
	// The name of the key pair.
	// - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
	// - The password logon method for Linux instances is set to forbidden upon initialization.
	KeyPairName pulumi.StringPtrInput
	// The name of the data disk.
	Name pulumi.StringPtrInput
	// The list of network interfaces created with instance.
	NetworkInterfaces LaunchTemplateNetworkInterfacesPtrInput
	// Network type of the instance. Value options: `classic` | `vpc`.
	NetworkType pulumi.StringPtrInput
	// The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
	RamRoleName     pulumi.StringPtrInput
	ResourceGroupId pulumi.StringPtrInput
	// Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
	SecurityEnhancementStrategy pulumi.StringPtrInput
	// The security group ID must be one in the same VPC.
	SecurityGroupId pulumi.StringPtrInput
	// -(Optional) 	Sets the maximum hourly instance price. Supports up to three decimal places.
	SpotPriceLimit pulumi.Float64PtrInput
	// The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
	// - NoSpot: Normal Pay-As-You-Go instance.
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	// - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
	SpotStrategy pulumi.StringPtrInput
	// The category of the system disk. System disk type. Optional values:
	// - cloud: Basic cloud disk.
	// - cloud_efficiency: Ultra cloud disk.
	// - cloud_ssd: SSD cloud Disks.
	// - ephemeral_ssd: local SSD Disks
	// - cloud_essd: ESSD cloud Disks.
	SystemDiskCategory pulumi.StringPtrInput
	// System disk description. It cannot begin with http:// or https://.
	SystemDiskDescription pulumi.StringPtrInput
	// System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	SystemDiskName pulumi.StringPtrInput
	// Size of the system disk, measured in GB. Value range: [20, 500].
	SystemDiskSize pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
	Userdata pulumi.StringPtrInput
	VpcId    pulumi.StringPtrInput
	// The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
	VswitchId pulumi.StringPtrInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (LaunchTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateState)(nil)).Elem()
}

type launchTemplateArgs struct {
	// Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
	AutoReleaseTime *string `pulumi:"autoReleaseTime"`
	// The list of data disks created with instance.
	DataDisks []LaunchTemplateDataDisk `pulumi:"dataDisks"`
	// The description of the data disk.
	Description *string `pulumi:"description"`
	// Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
	HostName *string `pulumi:"hostName"`
	// Image ID.
	ImageId         *string `pulumi:"imageId"`
	ImageOwnerAlias *string `pulumi:"imageOwnerAlias"`
	// Billing methods. Optional values:
	// - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
	// - PostPaid: Pay-As-You-Go.
	InstanceChargeType *string `pulumi:"instanceChargeType"`
	// The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	InstanceName *string `pulumi:"instanceName"`
	// Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
	InstanceType *string `pulumi:"instanceType"`
	// Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
	InternetChargeType *string `pulumi:"internetChargeType"`
	// The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
	InternetMaxBandwidthIn *int `pulumi:"internetMaxBandwidthIn"`
	// Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
	InternetMaxBandwidthOut *int `pulumi:"internetMaxBandwidthOut"`
	// Whether it is an I/O-optimized instance or not. Optional values:
	// - none
	// - optimized
	IoOptimized *string `pulumi:"ioOptimized"`
	// The name of the key pair.
	// - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
	// - The password logon method for Linux instances is set to forbidden upon initialization.
	KeyPairName *string `pulumi:"keyPairName"`
	// The name of the data disk.
	Name *string `pulumi:"name"`
	// The list of network interfaces created with instance.
	NetworkInterfaces *LaunchTemplateNetworkInterfaces `pulumi:"networkInterfaces"`
	// Network type of the instance. Value options: `classic` | `vpc`.
	NetworkType *string `pulumi:"networkType"`
	// The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
	RamRoleName     *string `pulumi:"ramRoleName"`
	ResourceGroupId *string `pulumi:"resourceGroupId"`
	// Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
	SecurityEnhancementStrategy *string `pulumi:"securityEnhancementStrategy"`
	// The security group ID must be one in the same VPC.
	SecurityGroupId *string `pulumi:"securityGroupId"`
	// -(Optional) 	Sets the maximum hourly instance price. Supports up to three decimal places.
	SpotPriceLimit *float64 `pulumi:"spotPriceLimit"`
	// The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
	// - NoSpot: Normal Pay-As-You-Go instance.
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	// - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
	SpotStrategy *string `pulumi:"spotStrategy"`
	// The category of the system disk. System disk type. Optional values:
	// - cloud: Basic cloud disk.
	// - cloud_efficiency: Ultra cloud disk.
	// - cloud_ssd: SSD cloud Disks.
	// - ephemeral_ssd: local SSD Disks
	// - cloud_essd: ESSD cloud Disks.
	SystemDiskCategory *string `pulumi:"systemDiskCategory"`
	// System disk description. It cannot begin with http:// or https://.
	SystemDiskDescription *string `pulumi:"systemDiskDescription"`
	// System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	SystemDiskName *string `pulumi:"systemDiskName"`
	// Size of the system disk, measured in GB. Value range: [20, 500].
	SystemDiskSize *int `pulumi:"systemDiskSize"`
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags map[string]interface{} `pulumi:"tags"`
	// User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
	Userdata *string `pulumi:"userdata"`
	VpcId    *string `pulumi:"vpcId"`
	// The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
	VswitchId *string `pulumi:"vswitchId"`
	// The zone ID of the instance.
	ZoneId *string `pulumi:"zoneId"`
}

// The set of arguments for constructing a LaunchTemplate resource.
type LaunchTemplateArgs struct {
	// Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
	AutoReleaseTime pulumi.StringPtrInput
	// The list of data disks created with instance.
	DataDisks LaunchTemplateDataDiskArrayInput
	// The description of the data disk.
	Description pulumi.StringPtrInput
	// Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
	HostName pulumi.StringPtrInput
	// Image ID.
	ImageId         pulumi.StringPtrInput
	ImageOwnerAlias pulumi.StringPtrInput
	// Billing methods. Optional values:
	// - PrePaid: Monthly, or annual subscription. Make sure that your registered credit card is invalid or you have insufficient balance in your PayPal account. Otherwise, InvalidPayMethod error may occur.
	// - PostPaid: Pay-As-You-Go.
	InstanceChargeType pulumi.StringPtrInput
	// The name of the instance. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	InstanceName pulumi.StringPtrInput
	// Instance type. For more information, call resourceAlicloudInstances to obtain the latest instance type list.
	InstanceType pulumi.StringPtrInput
	// Internet bandwidth billing method. Optional values: `PayByTraffic` | `PayByBandwidth`.
	InternetChargeType pulumi.StringPtrInput
	// The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
	InternetMaxBandwidthIn pulumi.IntPtrInput
	// Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
	InternetMaxBandwidthOut pulumi.IntPtrInput
	// Whether it is an I/O-optimized instance or not. Optional values:
	// - none
	// - optimized
	IoOptimized pulumi.StringPtrInput
	// The name of the key pair.
	// - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
	// - The password logon method for Linux instances is set to forbidden upon initialization.
	KeyPairName pulumi.StringPtrInput
	// The name of the data disk.
	Name pulumi.StringPtrInput
	// The list of network interfaces created with instance.
	NetworkInterfaces LaunchTemplateNetworkInterfacesPtrInput
	// Network type of the instance. Value options: `classic` | `vpc`.
	NetworkType pulumi.StringPtrInput
	// The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
	RamRoleName     pulumi.StringPtrInput
	ResourceGroupId pulumi.StringPtrInput
	// Whether or not to activate the security enhancement feature and install network security software free of charge. Optional values: Active | Deactive.
	SecurityEnhancementStrategy pulumi.StringPtrInput
	// The security group ID must be one in the same VPC.
	SecurityGroupId pulumi.StringPtrInput
	// -(Optional) 	Sets the maximum hourly instance price. Supports up to three decimal places.
	SpotPriceLimit pulumi.Float64PtrInput
	// The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Value range:
	// - NoSpot: Normal Pay-As-You-Go instance.
	// - SpotWithPriceLimit: Sets the maximum price for a spot instance.
	// - SpotAsPriceGo: The system automatically calculates the price. The maximum value is the Pay-As-You-Go price.
	SpotStrategy pulumi.StringPtrInput
	// The category of the system disk. System disk type. Optional values:
	// - cloud: Basic cloud disk.
	// - cloud_efficiency: Ultra cloud disk.
	// - cloud_ssd: SSD cloud Disks.
	// - ephemeral_ssd: local SSD Disks
	// - cloud_essd: ESSD cloud Disks.
	SystemDiskCategory pulumi.StringPtrInput
	// System disk description. It cannot begin with http:// or https://.
	SystemDiskDescription pulumi.StringPtrInput
	// System disk name. The name is a string of 2 to 128 characters. It must begin with an English or a Chinese character. It can contain A-Z, a-z, Chinese characters, numbers, periods (.), colons (:), underscores (_), and hyphens (-).
	SystemDiskName pulumi.StringPtrInput
	// Size of the system disk, measured in GB. Value range: [20, 500].
	SystemDiskSize pulumi.IntPtrInput
	// A mapping of tags to assign to the resource.
	// - Key: It can be up to 64 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It cannot be a null string.
	// - Value: It can be up to 128 characters in length. It cannot begin with "aliyun", "acs:", "http://", or "https://". It can be a null string.
	Tags pulumi.MapInput
	// User data of the instance, which is Base64-encoded. Size of the raw data cannot exceed 16 KB.
	Userdata pulumi.StringPtrInput
	VpcId    pulumi.StringPtrInput
	// The VSwitch ID for ENI. The instance must be in the same zone of the same VPC network as the ENI, but they may belong to different VSwitches.
	VswitchId pulumi.StringPtrInput
	// The zone ID of the instance.
	ZoneId pulumi.StringPtrInput
}

func (LaunchTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*launchTemplateArgs)(nil)).Elem()
}
