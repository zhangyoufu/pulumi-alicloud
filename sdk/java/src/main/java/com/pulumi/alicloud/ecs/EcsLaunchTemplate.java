// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsLaunchTemplateArgs;
import com.pulumi.alicloud.ecs.inputs.EcsLaunchTemplateState;
import com.pulumi.alicloud.ecs.outputs.EcsLaunchTemplateDataDisk;
import com.pulumi.alicloud.ecs.outputs.EcsLaunchTemplateNetworkInterfaces;
import com.pulumi.alicloud.ecs.outputs.EcsLaunchTemplateSystemDisk;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Launch Template resource.
 * 
 * For information about ECS Launch Template and how to use it, see [What is Launch Template](https://www.alibabacloud.com/help/en/doc-detail/74686.htm).
 * 
 * &gt; **NOTE:** Available since v1.120.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.EcsLaunchTemplate;
 * import com.pulumi.alicloud.ecs.EcsLaunchTemplateArgs;
 * import com.pulumi.alicloud.ecs.inputs.EcsLaunchTemplateSystemDiskArgs;
 * import com.pulumi.alicloud.ecs.inputs.EcsLaunchTemplateNetworkInterfacesArgs;
 * import com.pulumi.alicloud.ecs.inputs.EcsLaunchTemplateDataDiskArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_[0-9]+_[0-9]+_x64*&#34;)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(&#34;terraform-example&#34;)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .name(&#34;terraform-example&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultEcsLaunchTemplate = new EcsLaunchTemplate(&#34;defaultEcsLaunchTemplate&#34;, EcsLaunchTemplateArgs.builder()        
 *             .launchTemplateName(&#34;terraform-example&#34;)
 *             .description(&#34;terraform-example&#34;)
 *             .imageId(defaultGetImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .hostName(&#34;terraform-example&#34;)
 *             .instanceChargeType(&#34;PrePaid&#34;)
 *             .instanceName(&#34;terraform-example&#34;)
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .internetChargeType(&#34;PayByBandwidth&#34;)
 *             .internetMaxBandwidthIn(&#34;5&#34;)
 *             .internetMaxBandwidthOut(&#34;5&#34;)
 *             .ioOptimized(&#34;optimized&#34;)
 *             .keyPairName(&#34;key_pair_name&#34;)
 *             .ramRoleName(&#34;ram_role_name&#34;)
 *             .networkType(&#34;vpc&#34;)
 *             .securityEnhancementStrategy(&#34;Active&#34;)
 *             .spotPriceLimit(&#34;5&#34;)
 *             .spotStrategy(&#34;SpotWithPriceLimit&#34;)
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .systemDisk(EcsLaunchTemplateSystemDiskArgs.builder()
 *                 .category(&#34;cloud_ssd&#34;)
 *                 .description(&#34;Test For Terraform&#34;)
 *                 .name(&#34;terraform-example&#34;)
 *                 .size(&#34;40&#34;)
 *                 .deleteWithInstance(&#34;false&#34;)
 *                 .build())
 *             .userData(&#34;xxxxxxx&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(default_.zones()[0].id())
 *             .templateTags(Map.ofEntries(
 *                 Map.entry(&#34;Create&#34;, &#34;Terraform&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;example&#34;)
 *             ))
 *             .networkInterfaces(EcsLaunchTemplateNetworkInterfacesArgs.builder()
 *                 .name(&#34;eth0&#34;)
 *                 .description(&#34;hello1&#34;)
 *                 .primaryIp(&#34;10.0.0.2&#34;)
 *                 .securityGroupId(defaultSecurityGroup.id())
 *                 .vswitchId(defaultSwitch.id())
 *                 .build())
 *             .dataDisks(            
 *                 EcsLaunchTemplateDataDiskArgs.builder()
 *                     .name(&#34;disk1&#34;)
 *                     .description(&#34;description&#34;)
 *                     .deleteWithInstance(&#34;true&#34;)
 *                     .category(&#34;cloud&#34;)
 *                     .encrypted(&#34;false&#34;)
 *                     .performanceLevel(&#34;PL0&#34;)
 *                     .size(&#34;20&#34;)
 *                     .build(),
 *                 EcsLaunchTemplateDataDiskArgs.builder()
 *                     .name(&#34;disk2&#34;)
 *                     .description(&#34;description2&#34;)
 *                     .deleteWithInstance(&#34;true&#34;)
 *                     .category(&#34;cloud&#34;)
 *                     .encrypted(&#34;false&#34;)
 *                     .performanceLevel(&#34;PL0&#34;)
 *                     .size(&#34;20&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ECS Launch Template can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ecs/ecsLaunchTemplate:EcsLaunchTemplate example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsLaunchTemplate:EcsLaunchTemplate")
public class EcsLaunchTemplate extends com.pulumi.resources.CustomResource {
    /**
     * Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
     * 
     */
    @Export(name="autoReleaseTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> autoReleaseTime;

    /**
     * @return Instance auto release time. The time is presented using the ISO8601 standard and in UTC time. The format is  YYYY-MM-DDTHH:MM:SSZ.
     * 
     */
    public Output<Optional<String>> autoReleaseTime() {
        return Codegen.optional(this.autoReleaseTime);
    }
    /**
     * The list of data disks created with instance. See `data_disks` below.
     * 
     */
    @Export(name="dataDisks", refs={List.class,EcsLaunchTemplateDataDisk.class}, tree="[0,1]")
    private Output</* @Nullable */ List<EcsLaunchTemplateDataDisk>> dataDisks;

    /**
     * @return The list of data disks created with instance. See `data_disks` below.
     * 
     */
    public Output<Optional<List<EcsLaunchTemplateDataDisk>>> dataDisks() {
        return Codegen.optional(this.dataDisks);
    }
    /**
     * The Deployment Set Id.
     * 
     */
    @Export(name="deploymentSetId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deploymentSetId;

    /**
     * @return The Deployment Set Id.
     * 
     */
    public Output<Optional<String>> deploymentSetId() {
        return Codegen.optional(this.deploymentSetId);
    }
    /**
     * Description of instance launch template version 1. It can be [2, 256] characters in length. It cannot start with &#34;http://&#34; or &#34;https://&#34;. The default value is null.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return Description of instance launch template version 1. It can be [2, 256] characters in length. It cannot start with &#34;http://&#34; or &#34;https://&#34;. The default value is null.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether to enable the instance operating system configuration.
     * 
     */
    @Export(name="enableVmOsConfig", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableVmOsConfig;

    /**
     * @return Whether to enable the instance operating system configuration.
     * 
     */
    public Output<Optional<Boolean>> enableVmOsConfig() {
        return Codegen.optional(this.enableVmOsConfig);
    }
    /**
     * Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
     * 
     */
    @Export(name="hostName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> hostName;

    /**
     * @return Instance host name.It cannot start or end with a period (.) or a hyphen (-) and it cannot have two or more consecutive periods (.) or hyphens (-).For Windows: The host name can be [2, 15] characters in length. It can contain A-Z, a-z, numbers, periods (.), and hyphens (-). It cannot only contain numbers. For other operating systems: The host name can be [2, 64] characters in length. It can be segments separated by periods (.). It can contain A-Z, a-z, numbers, and hyphens (-).
     * 
     */
    public Output<Optional<String>> hostName() {
        return Codegen.optional(this.hostName);
    }
    /**
     * The Image ID.
     * 
     */
    @Export(name="imageId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageId;

    /**
     * @return The Image ID.
     * 
     */
    public Output<Optional<String>> imageId() {
        return Codegen.optional(this.imageId);
    }
    /**
     * Mirror source. Valid values: `system`, `self`, `others`, `marketplace`, `&#34;&#34;`. Default to: `&#34;&#34;`.
     * 
     */
    @Export(name="imageOwnerAlias", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> imageOwnerAlias;

    /**
     * @return Mirror source. Valid values: `system`, `self`, `others`, `marketplace`, `&#34;&#34;`. Default to: `&#34;&#34;`.
     * 
     */
    public Output<Optional<String>> imageOwnerAlias() {
        return Codegen.optional(this.imageOwnerAlias);
    }
    /**
     * Billing methods. Valid values: `PostPaid`, `PrePaid`.
     * 
     */
    @Export(name="instanceChargeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceChargeType;

    /**
     * @return Billing methods. Valid values: `PostPaid`, `PrePaid`.
     * 
     */
    public Output<Optional<String>> instanceChargeType() {
        return Codegen.optional(this.instanceChargeType);
    }
    /**
     * The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * Instance type. For more information, call resource_alicloud_instances to obtain the latest instance type list.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceType;

    /**
     * @return Instance type. For more information, call resource_alicloud_instances to obtain the latest instance type list.
     * 
     */
    public Output<Optional<String>> instanceType() {
        return Codegen.optional(this.instanceType);
    }
    /**
     * Internet bandwidth billing method. Valid values: `PayByTraffic`, `PayByBandwidth`.
     * 
     */
    @Export(name="internetChargeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> internetChargeType;

    /**
     * @return Internet bandwidth billing method. Valid values: `PayByTraffic`, `PayByBandwidth`.
     * 
     */
    public Output<Optional<String>> internetChargeType() {
        return Codegen.optional(this.internetChargeType);
    }
    /**
     * The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
     * 
     */
    @Export(name="internetMaxBandwidthIn", refs={Integer.class}, tree="[0]")
    private Output<Integer> internetMaxBandwidthIn;

    /**
     * @return The maximum inbound bandwidth from the Internet network, measured in Mbit/s. Value range: [1, 200].
     * 
     */
    public Output<Integer> internetMaxBandwidthIn() {
        return this.internetMaxBandwidthIn;
    }
    /**
     * Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
     * 
     */
    @Export(name="internetMaxBandwidthOut", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> internetMaxBandwidthOut;

    /**
     * @return Maximum outbound bandwidth from the Internet, its unit of measurement is Mbit/s. Value range: [0, 100].
     * 
     */
    public Output<Optional<Integer>> internetMaxBandwidthOut() {
        return Codegen.optional(this.internetMaxBandwidthOut);
    }
    /**
     * Whether it is an I/O-optimized instance or not. Valid values: `none`, `optimized`.
     * 
     */
    @Export(name="ioOptimized", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ioOptimized;

    /**
     * @return Whether it is an I/O-optimized instance or not. Valid values: `none`, `optimized`.
     * 
     */
    public Output<Optional<String>> ioOptimized() {
        return Codegen.optional(this.ioOptimized);
    }
    /**
     * The name of the key pair.
     * - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
     * - The password logon method for Linux instances is set to forbidden upon initialization.
     * 
     */
    @Export(name="keyPairName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> keyPairName;

    /**
     * @return The name of the key pair.
     * - Ignore this parameter for Windows instances. It is null by default. Even if you enter this parameter, only the  Password content is used.
     * - The password logon method for Linux instances is set to forbidden upon initialization.
     * 
     */
    public Output<Optional<String>> keyPairName() {
        return Codegen.optional(this.keyPairName);
    }
    /**
     * The name of Launch Template.
     * 
     */
    @Export(name="launchTemplateName", refs={String.class}, tree="[0]")
    private Output<String> launchTemplateName;

    /**
     * @return The name of Launch Template.
     * 
     */
    public Output<String> launchTemplateName() {
        return this.launchTemplateName;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `launch_template_name` instead.
     * 
     * @deprecated
     * Field &#39;name&#39; has been deprecated from provider version 1.120.0. New field &#39;launch_template_name&#39; instead.
     * 
     */
    @Deprecated /* Field 'name' has been deprecated from provider version 1.120.0. New field 'launch_template_name' instead. */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `launch_template_name` instead.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The list of network interfaces created with instance. See `network_interfaces` below.
     * 
     */
    @Export(name="networkInterfaces", refs={EcsLaunchTemplateNetworkInterfaces.class}, tree="[0]")
    private Output</* @Nullable */ EcsLaunchTemplateNetworkInterfaces> networkInterfaces;

    /**
     * @return The list of network interfaces created with instance. See `network_interfaces` below.
     * 
     */
    public Output<Optional<EcsLaunchTemplateNetworkInterfaces>> networkInterfaces() {
        return Codegen.optional(this.networkInterfaces);
    }
    /**
     * Network type of the instance. Valid values: `classic`, `vpc`.
     * 
     */
    @Export(name="networkType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> networkType;

    /**
     * @return Network type of the instance. Valid values: `classic`, `vpc`.
     * 
     */
    public Output<Optional<String>> networkType() {
        return Codegen.optional(this.networkType);
    }
    /**
     * Whether to use the password preset by the mirror.
     * 
     */
    @Export(name="passwordInherit", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> passwordInherit;

    /**
     * @return Whether to use the password preset by the mirror.
     * 
     */
    public Output<Optional<Boolean>> passwordInherit() {
        return Codegen.optional(this.passwordInherit);
    }
    /**
     * The subscription period of the instance. Unit: months. This parameter takes effect and is required only when InstanceChargeType is set to PrePaid. If the DedicatedHostId parameter is specified, the value of the Period parameter must be within the subscription period of the dedicated host.
     * - When the PeriodUnit parameter is set to `Week`, the valid values of the Period parameter are `1`, `2`, `3`, and `4`.
     * - When the PeriodUnit parameter is set to `Month`, the valid values of the Period parameter are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, and `60`.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The subscription period of the instance. Unit: months. This parameter takes effect and is required only when InstanceChargeType is set to PrePaid. If the DedicatedHostId parameter is specified, the value of the Period parameter must be within the subscription period of the dedicated host.
     * - When the PeriodUnit parameter is set to `Week`, the valid values of the Period parameter are `1`, `2`, `3`, and `4`.
     * - When the PeriodUnit parameter is set to `Month`, the valid values of the Period parameter are `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, and `60`.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The private IP address of the instance.
     * 
     */
    @Export(name="privateIpAddress", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> privateIpAddress;

    /**
     * @return The private IP address of the instance.
     * 
     */
    public Output<Optional<String>> privateIpAddress() {
        return Codegen.optional(this.privateIpAddress);
    }
    /**
     * The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
     * 
     */
    @Export(name="ramRoleName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ramRoleName;

    /**
     * @return The RAM role name of the instance. You can use the RAM API ListRoles to query instance RAM role names.
     * 
     */
    public Output<Optional<String>> ramRoleName() {
        return Codegen.optional(this.ramRoleName);
    }
    /**
     * The ID of the resource group to which to assign the instance, Elastic Block Storage (EBS) device, and ENI.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The ID of the resource group to which to assign the instance, Elastic Block Storage (EBS) device, and ENI.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * Whether or not to activate the security enhancement feature and install network security software free of charge. Valid values: `Active`, `Deactive`.
     * 
     */
    @Export(name="securityEnhancementStrategy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityEnhancementStrategy;

    /**
     * @return Whether or not to activate the security enhancement feature and install network security software free of charge. Valid values: `Active`, `Deactive`.
     * 
     */
    public Output<Optional<String>> securityEnhancementStrategy() {
        return Codegen.optional(this.securityEnhancementStrategy);
    }
    /**
     * The security group ID.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> securityGroupId;

    /**
     * @return The security group ID.
     * 
     */
    public Output<Optional<String>> securityGroupId() {
        return Codegen.optional(this.securityGroupId);
    }
    /**
     * The ID of security group N to which to assign the instance.
     * 
     */
    @Export(name="securityGroupIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> securityGroupIds;

    /**
     * @return The ID of security group N to which to assign the instance.
     * 
     */
    public Output<Optional<List<String>>> securityGroupIds() {
        return Codegen.optional(this.securityGroupIds);
    }
    /**
     * The protection period of the preemptible instance. Unit: hours. Valid values: `0`, `1`, `2`, `3`, `4`, `5`, and `6`. Default to: `1`.
     * 
     */
    @Export(name="spotDuration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> spotDuration;

    /**
     * @return The protection period of the preemptible instance. Unit: hours. Valid values: `0`, `1`, `2`, `3`, `4`, `5`, and `6`. Default to: `1`.
     * 
     */
    public Output<Optional<String>> spotDuration() {
        return Codegen.optional(this.spotDuration);
    }
    /**
     * Sets the maximum hourly instance price. Supports up to three decimal places.
     * 
     */
    @Export(name="spotPriceLimit", refs={Double.class}, tree="[0]")
    private Output</* @Nullable */ Double> spotPriceLimit;

    /**
     * @return Sets the maximum hourly instance price. Supports up to three decimal places.
     * 
     */
    public Output<Optional<Double>> spotPriceLimit() {
        return Codegen.optional(this.spotPriceLimit);
    }
    /**
     * The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Valid values: `NoSpot`, `SpotAsPriceGo`, `SpotWithPriceLimit`.
     * 
     */
    @Export(name="spotStrategy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> spotStrategy;

    /**
     * @return The spot strategy for a Pay-As-You-Go instance. This parameter is valid and required only when InstanceChargeType is set to PostPaid. Valid values: `NoSpot`, `SpotAsPriceGo`, `SpotWithPriceLimit`.
     * 
     */
    public Output<Optional<String>> spotStrategy() {
        return Codegen.optional(this.spotStrategy);
    }
    /**
     * The System Disk. See `system_disk` below.
     * 
     */
    @Export(name="systemDisk", refs={EcsLaunchTemplateSystemDisk.class}, tree="[0]")
    private Output<EcsLaunchTemplateSystemDisk> systemDisk;

    /**
     * @return The System Disk. See `system_disk` below.
     * 
     */
    public Output<EcsLaunchTemplateSystemDisk> systemDisk() {
        return this.systemDisk;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     * @deprecated
     * Field &#39;system_disk_category&#39; has been deprecated from provider version 1.120.0. New field &#39;system_disk&#39; instead.
     * 
     */
    @Deprecated /* Field 'system_disk_category' has been deprecated from provider version 1.120.0. New field 'system_disk' instead. */
    @Export(name="systemDiskCategory", refs={String.class}, tree="[0]")
    private Output<String> systemDiskCategory;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     */
    public Output<String> systemDiskCategory() {
        return this.systemDiskCategory;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     * @deprecated
     * Field &#39;system_disk_description&#39; has been deprecated from provider version 1.120.0. New field &#39;system_disk&#39; instead.
     * 
     */
    @Deprecated /* Field 'system_disk_description' has been deprecated from provider version 1.120.0. New field 'system_disk' instead. */
    @Export(name="systemDiskDescription", refs={String.class}, tree="[0]")
    private Output<String> systemDiskDescription;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     */
    public Output<String> systemDiskDescription() {
        return this.systemDiskDescription;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     * @deprecated
     * Field &#39;system_disk_name&#39; has been deprecated from provider version 1.120.0. New field &#39;system_disk&#39; instead.
     * 
     */
    @Deprecated /* Field 'system_disk_name' has been deprecated from provider version 1.120.0. New field 'system_disk' instead. */
    @Export(name="systemDiskName", refs={String.class}, tree="[0]")
    private Output<String> systemDiskName;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     */
    public Output<String> systemDiskName() {
        return this.systemDiskName;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     * @deprecated
     * Field &#39;system_disk_size&#39; has been deprecated from provider version 1.120.0. New field &#39;system_disk&#39; instead.
     * 
     */
    @Deprecated /* Field 'system_disk_size' has been deprecated from provider version 1.120.0. New field 'system_disk' instead. */
    @Export(name="systemDiskSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> systemDiskSize;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `system_disk` instead.
     * 
     */
    public Output<Integer> systemDiskSize() {
        return this.systemDiskSize;
    }
    /**
     * A mapping of tags to assign to instance, block storage, and elastic network.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to instance, block storage, and elastic network.
     * - Key: It can be up to 64 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It cannot be a null string.
     * - Value: It can be up to 128 characters in length. It cannot begin with &#34;aliyun&#34;, &#34;acs:&#34;, &#34;http://&#34;, or &#34;https://&#34;. It can be a null string.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The template resource group id.
     * 
     */
    @Export(name="templateResourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> templateResourceGroupId;

    /**
     * @return The template resource group id.
     * 
     */
    public Output<Optional<String>> templateResourceGroupId() {
        return Codegen.optional(this.templateResourceGroupId);
    }
    /**
     * A mapping of tags to assign to the launch template.
     * 
     */
    @Export(name="templateTags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> templateTags;

    /**
     * @return A mapping of tags to assign to the launch template.
     * 
     */
    public Output<Optional<Map<String,Object>>> templateTags() {
        return Codegen.optional(this.templateTags);
    }
    /**
     * The User Data.
     * 
     */
    @Export(name="userData", refs={String.class}, tree="[0]")
    private Output<String> userData;

    /**
     * @return The User Data.
     * 
     */
    public Output<String> userData() {
        return this.userData;
    }
    /**
     * It has been deprecated from version 1.120.0, and use field `user_data` instead.
     * 
     * @deprecated
     * Field &#39;userdata&#39; has been deprecated from provider version 1.120.0. New field &#39;user_data&#39; instead.
     * 
     */
    @Deprecated /* Field 'userdata' has been deprecated from provider version 1.120.0. New field 'user_data' instead. */
    @Export(name="userdata", refs={String.class}, tree="[0]")
    private Output<String> userdata;

    /**
     * @return It has been deprecated from version 1.120.0, and use field `user_data` instead.
     * 
     */
    public Output<String> userdata() {
        return this.userdata;
    }
    /**
     * The description of the launch template version. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    @Export(name="versionDescription", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> versionDescription;

    /**
     * @return The description of the launch template version. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
     * 
     */
    public Output<Optional<String>> versionDescription() {
        return Codegen.optional(this.versionDescription);
    }
    /**
     * The ID of the VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * When creating a VPC-Connected instance, you must specify its VSwitch ID.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return When creating a VPC-Connected instance, you must specify its VSwitch ID.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The zone ID of the instance.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The zone ID of the instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsLaunchTemplate(String name) {
        this(name, EcsLaunchTemplateArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsLaunchTemplate(String name, @Nullable EcsLaunchTemplateArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsLaunchTemplate(String name, @Nullable EcsLaunchTemplateArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsLaunchTemplate:EcsLaunchTemplate", name, args == null ? EcsLaunchTemplateArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EcsLaunchTemplate(String name, Output<String> id, @Nullable EcsLaunchTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsLaunchTemplate:EcsLaunchTemplate", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static EcsLaunchTemplate get(String name, Output<String> id, @Nullable EcsLaunchTemplateState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsLaunchTemplate(name, id, state, options);
    }
}
