// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ecs;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ecs.EcsInstanceSetArgs;
import com.pulumi.alicloud.ecs.inputs.EcsInstanceSetState;
import com.pulumi.alicloud.ecs.outputs.EcsInstanceSetDataDisk;
import com.pulumi.alicloud.ecs.outputs.EcsInstanceSetExcludeInstanceFilter;
import com.pulumi.alicloud.ecs.outputs.EcsInstanceSetNetworkInterface;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECS Instance Set resource.
 * 
 * For information about ECS Instance Set and how to use it, see [What is Instance Set](https://www.alibabacloud.com/help/en/elastic-compute-service/latest/runinstances).
 * 
 * &gt; **NOTE:** Available in v1.173.0+.
 * 
 * &gt; **NOTE:** This resource is used to batch create a group of instance resources with the same configuration. However, this resource is not recommended. `alicloud.ecs.Instance` is preferred.
 * 
 * &gt; **NOTE:** In the instances managed by this resource, names are automatically generated based on `instance_name` and `unique_suffix`.
 * 
 * &gt; **NOTE:** Only `tags` support batch modification.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.adb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecp.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.cloudconnect.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.EcsInstanceSet;
 * import com.pulumi.alicloud.ecs.EcsInstanceSetArgs;
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
 *         final var config = ctx.config();
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-testaccecsset&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableDiskCategory(&#34;cloud_efficiency&#34;)
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         final var defaultImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_[0-9]+_[0-9]+_x64*&#34;)
 *             .mostRecent(true)
 *             .owners(&#34;system&#34;)
 *             .build());
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;default-NODELETING&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .build());
 * 
 *         var beijingK = new EcsInstanceSet(&#34;beijingK&#34;, EcsInstanceSetArgs.builder()        
 *             .amount(100)
 *             .imageId(defaultImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(defaultInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .instanceName(name)
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .systemDiskPerformanceLevel(&#34;PL0&#34;)
 *             .systemDiskCategory(&#34;cloud_essd&#34;)
 *             .systemDiskSize(200)
 *             .vswitchId(defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .securityGroupIds(defaultSecurityGroup.stream().map(element -&gt; element.id()).collect(toList()))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="alicloud:ecs/ecsInstanceSet:EcsInstanceSet")
public class EcsInstanceSet extends com.pulumi.resources.CustomResource {
    /**
     * The number of instances that you want to create. Valid values: `1` to `100`.
     * 
     */
    @Export(name="amount", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> amount;

    /**
     * @return The number of instances that you want to create. Valid values: `1` to `100`.
     * 
     */
    public Output<Optional<Integer>> amount() {
        return Codegen.optional(this.amount);
    }
    /**
     * The automatic release time of the `PostPaid` instance.
     * 
     */
    @Export(name="autoReleaseTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> autoReleaseTime;

    /**
     * @return The automatic release time of the `PostPaid` instance.
     * 
     */
    public Output<Optional<String>> autoReleaseTime() {
        return Codegen.optional(this.autoReleaseTime);
    }
    /**
     * Whether to enable auto-renewal for the instance. This parameter is valid only when the `instance_charge_type` is set to `PrePaid`.
     * 
     */
    @Export(name="autoRenew", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Whether to enable auto-renewal for the instance. This parameter is valid only when the `instance_charge_type` is set to `PrePaid`.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * Auto renewal period of an instance, in the unit of month. It is valid when `instance_charge_type` is `PrePaid`.
     * - When `period_unit` is `Month`, Valid values: `1`, `2`, `3`, `6`, `12`.
     * - When `period_unit` is `Week`, Valid values: `1`, `2`, `3`.
     * 
     */
    @Export(name="autoRenewPeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> autoRenewPeriod;

    /**
     * @return Auto renewal period of an instance, in the unit of month. It is valid when `instance_charge_type` is `PrePaid`.
     * - When `period_unit` is `Month`, Valid values: `1`, `2`, `3`, `6`, `12`.
     * - When `period_unit` is `Week`, Valid values: `1`, `2`, `3`.
     * 
     */
    public Output<Optional<Integer>> autoRenewPeriod() {
        return Codegen.optional(this.autoRenewPeriod);
    }
    /**
     * Indicate how to check instance ready to use.
     * - `false`: Default value. Means that the instances are ready when their DescribeInstances status is Running, at which time guestOS(Ecs os) may not be ready yet.
     * - `true`: Checking instance ready with Ecs assistant, which means guestOs boots successfully. Premise is that the specified image `image_id` has built-in Ecs assistant. Most of the public images have assistant installed already.
     * 
     */
    @Export(name="bootCheckOsWithAssistant", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> bootCheckOsWithAssistant;

    /**
     * @return Indicate how to check instance ready to use.
     * - `false`: Default value. Means that the instances are ready when their DescribeInstances status is Running, at which time guestOS(Ecs os) may not be ready yet.
     * - `true`: Checking instance ready with Ecs assistant, which means guestOs boots successfully. Premise is that the specified image `image_id` has built-in Ecs assistant. Most of the public images have assistant installed already.
     * 
     */
    public Output<Optional<Boolean>> bootCheckOsWithAssistant() {
        return Codegen.optional(this.bootCheckOsWithAssistant);
    }
    /**
     * The list of data disks created with instance. See the following `Block data_disks`.
     * 
     */
    @Export(name="dataDisks", type=List.class, parameters={EcsInstanceSetDataDisk.class})
    private Output</* @Nullable */ List<EcsInstanceSetDataDisk>> dataDisks;

    /**
     * @return The list of data disks created with instance. See the following `Block data_disks`.
     * 
     */
    public Output<Optional<List<EcsInstanceSetDataDisk>>> dataDisks() {
        return Codegen.optional(this.dataDisks);
    }
    /**
     * The ID of the dedicated host on which to create the instance. If the `dedicated_host_id` is specified, the `spot_strategy` and `spot_price_limit`  are ignored. This is because preemptible instances cannot be created on dedicated hosts.
     * 
     */
    @Export(name="dedicatedHostId", type=String.class, parameters={})
    private Output</* @Nullable */ String> dedicatedHostId;

    /**
     * @return The ID of the dedicated host on which to create the instance. If the `dedicated_host_id` is specified, the `spot_strategy` and `spot_price_limit`  are ignored. This is because preemptible instances cannot be created on dedicated hosts.
     * 
     */
    public Output<Optional<String>> dedicatedHostId() {
        return Codegen.optional(this.dedicatedHostId);
    }
    /**
     * Whether to enable release protection for the instance.
     * 
     */
    @Export(name="deletionProtection", type=Boolean.class, parameters={})
    private Output<Boolean> deletionProtection;

    /**
     * @return Whether to enable release protection for the instance.
     * 
     */
    public Output<Boolean> deletionProtection() {
        return this.deletionProtection;
    }
    /**
     * The ID of the deployment set to which to deploy the instance.
     * 
     */
    @Export(name="deploymentSetId", type=String.class, parameters={})
    private Output</* @Nullable */ String> deploymentSetId;

    /**
     * @return The ID of the deployment set to which to deploy the instance.
     * 
     */
    public Output<Optional<String>> deploymentSetId() {
        return Codegen.optional(this.deploymentSetId);
    }
    /**
     * The description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with `http://` or `https://`.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the instance, This description can have a string of 2 to 256 characters, It cannot begin with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The instances that need to be excluded from the Instance Set. See the following `Block exclude_instance_filter`.
     * 
     */
    @Export(name="excludeInstanceFilter", type=EcsInstanceSetExcludeInstanceFilter.class, parameters={})
    private Output</* @Nullable */ EcsInstanceSetExcludeInstanceFilter> excludeInstanceFilter;

    /**
     * @return The instances that need to be excluded from the Instance Set. See the following `Block exclude_instance_filter`.
     * 
     */
    public Output<Optional<EcsInstanceSetExcludeInstanceFilter>> excludeInstanceFilter() {
        return Codegen.optional(this.excludeInstanceFilter);
    }
    /**
     * The hostname of instance.
     * 
     */
    @Export(name="hostName", type=String.class, parameters={})
    private Output<String> hostName;

    /**
     * @return The hostname of instance.
     * 
     */
    public Output<String> hostName() {
        return this.hostName;
    }
    /**
     * The ID of the Elastic High Performance Computing (E-HPC) cluster to which to assign the instance.
     * 
     */
    @Export(name="hpcClusterId", type=String.class, parameters={})
    private Output</* @Nullable */ String> hpcClusterId;

    /**
     * @return The ID of the Elastic High Performance Computing (E-HPC) cluster to which to assign the instance.
     * 
     */
    public Output<Optional<String>> hpcClusterId() {
        return Codegen.optional(this.hpcClusterId);
    }
    /**
     * The Image to use for the instance.
     * 
     */
    @Export(name="imageId", type=String.class, parameters={})
    private Output<String> imageId;

    /**
     * @return The Image to use for the instance.
     * 
     */
    public Output<String> imageId() {
        return this.imageId;
    }
    /**
     * The billing method of the instance. Valid values: `PrePaid`, `PostPaid`.
     * 
     */
    @Export(name="instanceChargeType", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceChargeType;

    /**
     * @return The billing method of the instance. Valid values: `PrePaid`, `PostPaid`.
     * 
     */
    public Output<Optional<String>> instanceChargeType() {
        return Codegen.optional(this.instanceChargeType);
    }
    /**
     * A list of ECS Instance ID.
     * 
     */
    @Export(name="instanceIds", type=List.class, parameters={String.class})
    private Output<List<String>> instanceIds;

    /**
     * @return A list of ECS Instance ID.
     * 
     */
    public Output<List<String>> instanceIds() {
        return this.instanceIds;
    }
    /**
     * The name of the ECS. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen, and must not begin with `http://` or `https://`.
     * 
     */
    @Export(name="instanceName", type=String.class, parameters={})
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The name of the ECS. This instance_name can have a string of 2 to 128 characters, must contain only alphanumeric characters or hyphens, such as &#34;-&#34;,&#34;.&#34;,&#34;_&#34;, and must not begin with a hyphen, and must not begin with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * The type of instance to start.
     * 
     */
    @Export(name="instanceType", type=String.class, parameters={})
    private Output<String> instanceType;

    /**
     * @return The type of instance to start.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The Internet charge type of the instance. Valid values are `PayByBandwidth`, `PayByTraffic`.
     * 
     */
    @Export(name="internetChargeType", type=String.class, parameters={})
    private Output<String> internetChargeType;

    /**
     * @return The Internet charge type of the instance. Valid values are `PayByBandwidth`, `PayByTraffic`.
     * 
     */
    public Output<String> internetChargeType() {
        return this.internetChargeType;
    }
    /**
     * The Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value values: `1` to `100`.
     * 
     */
    @Export(name="internetMaxBandwidthOut", type=Integer.class, parameters={})
    private Output<Integer> internetMaxBandwidthOut;

    /**
     * @return The Maximum outgoing bandwidth to the public network, measured in Mbps (Mega bit per second). Value values: `1` to `100`.
     * 
     */
    public Output<Integer> internetMaxBandwidthOut() {
        return this.internetMaxBandwidthOut;
    }
    /**
     * The name of key pair that can login ECS instance successfully without password.
     * 
     */
    @Export(name="keyPairName", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyPairName;

    /**
     * @return The name of key pair that can login ECS instance successfully without password.
     * 
     */
    public Output<Optional<String>> keyPairName() {
        return Codegen.optional(this.keyPairName);
    }
    /**
     * The ID of the launch template.
     * 
     */
    @Export(name="launchTemplateId", type=String.class, parameters={})
    private Output</* @Nullable */ String> launchTemplateId;

    /**
     * @return The ID of the launch template.
     * 
     */
    public Output<Optional<String>> launchTemplateId() {
        return Codegen.optional(this.launchTemplateId);
    }
    /**
     * The name of the launch template. To use a launch template to create an instance, you must use the `launch_template_id` or `launch_template_name` parameter to specify the launch template.
     * 
     */
    @Export(name="launchTemplateName", type=String.class, parameters={})
    private Output</* @Nullable */ String> launchTemplateName;

    /**
     * @return The name of the launch template. To use a launch template to create an instance, you must use the `launch_template_id` or `launch_template_name` parameter to specify the launch template.
     * 
     */
    public Output<Optional<String>> launchTemplateName() {
        return Codegen.optional(this.launchTemplateName);
    }
    /**
     * The version of the launch template.
     * 
     */
    @Export(name="launchTemplateVersion", type=String.class, parameters={})
    private Output</* @Nullable */ String> launchTemplateVersion;

    /**
     * @return The version of the launch template.
     * 
     */
    public Output<Optional<String>> launchTemplateVersion() {
        return Codegen.optional(this.launchTemplateVersion);
    }
    /**
     * A list of NetworkInterface. See the following `Block network_interfaces`.
     * 
     */
    @Export(name="networkInterfaces", type=List.class, parameters={EcsInstanceSetNetworkInterface.class})
    private Output</* @Nullable */ List<EcsInstanceSetNetworkInterface>> networkInterfaces;

    /**
     * @return A list of NetworkInterface. See the following `Block network_interfaces`.
     * 
     */
    public Output<Optional<List<EcsInstanceSetNetworkInterface>>> networkInterfaces() {
        return Codegen.optional(this.networkInterfaces);
    }
    /**
     * The password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols.
     * 
     */
    @Export(name="password", type=String.class, parameters={})
    private Output</* @Nullable */ String> password;

    /**
     * @return The password to an instance is a string of 8 to 30 characters. It must contain uppercase/lowercase letters and numerals, but cannot contain special symbols.
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Whether to use the password preset in the image.
     * 
     */
    @Export(name="passwordInherit", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> passwordInherit;

    /**
     * @return Whether to use the password preset in the image.
     * 
     */
    public Output<Optional<Boolean>> passwordInherit() {
        return Codegen.optional(this.passwordInherit);
    }
    /**
     * The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`.
     * - When `period_unit` is `Month`, Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     * - When `period_unit` is `Week`, Valid values: `1`, `2`, `3`.
     * 
     */
    @Export(name="period", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The duration that you will buy the resource, in month. It is valid when `instance_charge_type` is `PrePaid`.
     * - When `period_unit` is `Month`, Valid values: `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `12`, `24`, `36`, `48`, `60`.
     * - When `period_unit` is `Week`, Valid values: `1`, `2`, `3`.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The duration unit that you will buy the resource. It is valid when `instance_charge_type` is &#39;PrePaid&#39;. Valid value: `Week`, `Month`.
     * 
     */
    @Export(name="periodUnit", type=String.class, parameters={})
    private Output</* @Nullable */ String> periodUnit;

    /**
     * @return The duration unit that you will buy the resource. It is valid when `instance_charge_type` is &#39;PrePaid&#39;. Valid value: `Week`, `Month`.
     * 
     */
    public Output<Optional<String>> periodUnit() {
        return Codegen.optional(this.periodUnit);
    }
    /**
     * The Instance RAM role name.
     * 
     */
    @Export(name="ramRoleName", type=String.class, parameters={})
    private Output</* @Nullable */ String> ramRoleName;

    /**
     * @return The Instance RAM role name.
     * 
     */
    public Output<Optional<String>> ramRoleName() {
        return Codegen.optional(this.ramRoleName);
    }
    /**
     * The ID of resource group which the instance belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The ID of resource group which the instance belongs.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The security enhancement strategy.
     * - `Active`: Enable security enhancement strategy, it only works on system images.
     * - `Deactive`: Disable security enhancement strategy, it works on all images.
     * 
     */
    @Export(name="securityEnhancementStrategy", type=String.class, parameters={})
    private Output</* @Nullable */ String> securityEnhancementStrategy;

    /**
     * @return The security enhancement strategy.
     * - `Active`: Enable security enhancement strategy, it only works on system images.
     * - `Deactive`: Disable security enhancement strategy, it works on all images.
     * 
     */
    public Output<Optional<String>> securityEnhancementStrategy() {
        return Codegen.optional(this.securityEnhancementStrategy);
    }
    /**
     * A list of security group ids to associate with.
     * 
     */
    @Export(name="securityGroupIds", type=List.class, parameters={String.class})
    private Output<List<String>> securityGroupIds;

    /**
     * @return A list of security group ids to associate with.
     * 
     */
    public Output<List<String>> securityGroupIds() {
        return this.securityGroupIds;
    }
    /**
     * The hourly price threshold of a instance, and it takes effect only when parameter &#39;spot_strategy&#39; is &#39;SpotWithPriceLimit&#39;. Three decimals is allowed at most.
     * 
     */
    @Export(name="spotPriceLimit", type=Double.class, parameters={})
    private Output<Double> spotPriceLimit;

    /**
     * @return The hourly price threshold of a instance, and it takes effect only when parameter &#39;spot_strategy&#39; is &#39;SpotWithPriceLimit&#39;. Three decimals is allowed at most.
     * 
     */
    public Output<Double> spotPriceLimit() {
        return this.spotPriceLimit;
    }
    /**
     * The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter `instance_charge_type` is &#39;PostPaid&#39;.
     * - `NoSpot`: A regular Pay-As-You-Go instance.
     * - `SpotWithPriceLimit`: A price threshold for a spot instance.
     * - `SpotAsPriceGo`: A price that is based on the highest Pay-As-You-Go instance
     * 
     */
    @Export(name="spotStrategy", type=String.class, parameters={})
    private Output<String> spotStrategy;

    /**
     * @return The spot strategy of a Pay-As-You-Go instance, and it takes effect only when parameter `instance_charge_type` is &#39;PostPaid&#39;.
     * - `NoSpot`: A regular Pay-As-You-Go instance.
     * - `SpotWithPriceLimit`: A price threshold for a spot instance.
     * - `SpotAsPriceGo`: A price that is based on the highest Pay-As-You-Go instance
     * 
     */
    public Output<String> spotStrategy() {
        return this.spotStrategy;
    }
    /**
     * The ID of the automatic snapshot policy applied to the system disk.
     * 
     */
    @Export(name="systemDiskAutoSnapshotPolicyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskAutoSnapshotPolicyId;

    /**
     * @return The ID of the automatic snapshot policy applied to the system disk.
     * 
     */
    public Output<Optional<String>> systemDiskAutoSnapshotPolicyId() {
        return Codegen.optional(this.systemDiskAutoSnapshotPolicyId);
    }
    /**
     * The category of the system disk. Valid values are `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud`.
     * 
     */
    @Export(name="systemDiskCategory", type=String.class, parameters={})
    private Output<String> systemDiskCategory;

    /**
     * @return The category of the system disk. Valid values are `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud`.
     * 
     */
    public Output<String> systemDiskCategory() {
        return this.systemDiskCategory;
    }
    /**
     * The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="systemDiskDescription", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskDescription;

    /**
     * @return The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> systemDiskDescription() {
        return Codegen.optional(this.systemDiskDescription);
    }
    /**
     * The name of the system disk.
     * 
     */
    @Export(name="systemDiskName", type=String.class, parameters={})
    private Output</* @Nullable */ String> systemDiskName;

    /**
     * @return The name of the system disk.
     * 
     */
    public Output<Optional<String>> systemDiskName() {
        return Codegen.optional(this.systemDiskName);
    }
    /**
     * The performance level of the ESSD used as the system disk. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     * 
     */
    @Export(name="systemDiskPerformanceLevel", type=String.class, parameters={})
    private Output<String> systemDiskPerformanceLevel;

    /**
     * @return The performance level of the ESSD used as the system disk. Valid values: `PL0`, `PL1`, `PL2`, `PL3`.
     * 
     */
    public Output<String> systemDiskPerformanceLevel() {
        return this.systemDiskPerformanceLevel;
    }
    /**
     * The size of the system disk, measured in GiB. Value range:  values: `20` to `500`.
     * 
     */
    @Export(name="systemDiskSize", type=Integer.class, parameters={})
    private Output<Integer> systemDiskSize;

    /**
     * @return The size of the system disk, measured in GiB. Value range:  values: `20` to `500`.
     * 
     */
    public Output<Integer> systemDiskSize() {
        return this.systemDiskSize;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, String.class})
    private Output</* @Nullable */ Map<String,String>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,String>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * Whether to automatically append incremental suffixes to the hostname specified by the HostName parameter and to the instance name specified by the InstanceName parameter when you batch create instances. The incremental suffixes can range from `001` to `999`.
     * 
     */
    @Export(name="uniqueSuffix", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> uniqueSuffix;

    /**
     * @return Whether to automatically append incremental suffixes to the hostname specified by the HostName parameter and to the instance name specified by the InstanceName parameter when you batch create instances. The incremental suffixes can range from `001` to `999`.
     * 
     */
    public Output<Optional<Boolean>> uniqueSuffix() {
        return Codegen.optional(this.uniqueSuffix);
    }
    /**
     * The virtual switch ID to launch in VPC.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return The virtual switch ID to launch in VPC.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The ID of the zone in which to create the instance.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    /**
     * @return The ID of the zone in which to create the instance.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public EcsInstanceSet(String name) {
        this(name, EcsInstanceSetArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public EcsInstanceSet(String name, EcsInstanceSetArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public EcsInstanceSet(String name, EcsInstanceSetArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsInstanceSet:EcsInstanceSet", name, args == null ? EcsInstanceSetArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private EcsInstanceSet(String name, Output<String> id, @Nullable EcsInstanceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ecs/ecsInstanceSet:EcsInstanceSet", name, state, makeResourceOptions(options, id));
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
    public static EcsInstanceSet get(String name, Output<String> id, @Nullable EcsInstanceSetState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new EcsInstanceSet(name, id, state, options);
    }
}
