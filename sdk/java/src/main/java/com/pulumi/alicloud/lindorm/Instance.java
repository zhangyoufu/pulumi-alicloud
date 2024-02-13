// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.lindorm;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.lindorm.InstanceArgs;
import com.pulumi.alicloud.lindorm.inputs.InstanceState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Lindorm Instance resource.
 * 
 * For information about Lindorm Instance and how to use it, see [What is Instance](https://www.alibabacloud.com/help/en/lindorm/latest/product-introduction-overview).
 * 
 * &gt; **NOTE:** Available since v1.132.0.
 * 
 * &gt; **NOTE:**  The Lindorm Instance does not support updating the specifications of multiple different engines, or the number of nodes at the same time.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.lindorm.Instance;
 * import com.pulumi.alicloud.lindorm.InstanceArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-example&#34;);
 *         final var region = &#34;cn-hangzhou&#34;;
 * 
 *         final var zoneId = &#34;cn-hangzhou-h&#34;;
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;^default-NODELETING$&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .diskCategory(&#34;cloud_efficiency&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .zoneId(zoneId)
 *             .vswitchId(defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .instanceName(name)
 *             .tableEngineSpecification(&#34;lindorm.g.4xlarge&#34;)
 *             .tableEngineNodeCount(&#34;2&#34;)
 *             .instanceStorage(&#34;1920&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Lindorm Instance can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:lindorm/instance:Instance example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:lindorm/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The multi-availability zone instance, coordinating the virtual switch ID of the availability zone, the switch must be located under the availability zone corresponding to the ArbiterZoneId. This parameter is required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="arbiterVswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> arbiterVswitchId;

    /**
     * @return The multi-availability zone instance, coordinating the virtual switch ID of the availability zone, the switch must be located under the availability zone corresponding to the ArbiterZoneId. This parameter is required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> arbiterVswitchId() {
        return Codegen.optional(this.arbiterVswitchId);
    }
    /**
     * The multiple Availability Zone Instance, the availability zone ID of the coordinating availability zone. required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="arbiterZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> arbiterZoneId;

    /**
     * @return The multiple Availability Zone Instance, the availability zone ID of the coordinating availability zone. required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> arbiterZoneId() {
        return Codegen.optional(this.arbiterZoneId);
    }
    /**
     * The deployment architecture. If you do not fill in this parameter, the default is 1.0. to create multiple availability instances, fill in 2.0. if you need to create multiple availability instances, this parameter is required. Valid values: `1.0` to `2.0`.
     * 
     */
    @Export(name="archVersion", refs={String.class}, tree="[0]")
    private Output<String> archVersion;

    /**
     * @return The deployment architecture. If you do not fill in this parameter, the default is 1.0. to create multiple availability instances, fill in 2.0. if you need to create multiple availability instances, this parameter is required. Valid values: `1.0` to `2.0`.
     * 
     */
    public Output<String> archVersion() {
        return this.archVersion;
    }
    /**
     * The cold storage capacity of the instance. Unit: GB. Valid values: [800, 1000000].
     * 
     */
    @Export(name="coldStorage", refs={Integer.class}, tree="[0]")
    private Output<Integer> coldStorage;

    /**
     * @return The cold storage capacity of the instance. Unit: GB. Valid values: [800, 1000000].
     * 
     */
    public Output<Integer> coldStorage() {
        return this.coldStorage;
    }
    /**
     * The multiple availability zone instances, CORE single node capacity. required if you want to create multiple availability zone instances. Valid values: `400` to `64000`.
     * 
     */
    @Export(name="coreSingleStorage", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> coreSingleStorage;

    /**
     * @return The multiple availability zone instances, CORE single node capacity. required if you want to create multiple availability zone instances. Valid values: `400` to `64000`.
     * 
     */
    public Output<Optional<Integer>> coreSingleStorage() {
        return Codegen.optional(this.coreSingleStorage);
    }
    /**
     * The core spec. When `disk_category` is `local_ssd_pro` or `local_hdd_pro`, this filed is valid.
     * - When `disk_category` is `local_ssd_pro`, the valid values is `lindorm.i2.xlarge`, `lindorm.i2.2xlarge`, `lindorm.i2.4xlarge`, `lindorm.i2.8xlarge`.
     * - When `disk_category` is `local_hdd_pro`, the valid values is `lindorm.d2c.6xlarge`, `lindorm.d2c.12xlarge`, `lindorm.d2c.24xlarge`, `lindorm.d2s.5xlarge`, `lindorm.d2s.10xlarge`, `lindorm.d1.2xlarge`, `lindorm.d1.4xlarge`, `lindorm.d1.6xlarge`.
     * 
     */
    @Export(name="coreSpec", refs={String.class}, tree="[0]")
    private Output<String> coreSpec;

    /**
     * @return The core spec. When `disk_category` is `local_ssd_pro` or `local_hdd_pro`, this filed is valid.
     * - When `disk_category` is `local_ssd_pro`, the valid values is `lindorm.i2.xlarge`, `lindorm.i2.2xlarge`, `lindorm.i2.4xlarge`, `lindorm.i2.8xlarge`.
     * - When `disk_category` is `local_hdd_pro`, the valid values is `lindorm.d2c.6xlarge`, `lindorm.d2c.12xlarge`, `lindorm.d2c.24xlarge`, `lindorm.d2s.5xlarge`, `lindorm.d2s.10xlarge`, `lindorm.d1.2xlarge`, `lindorm.d1.4xlarge`, `lindorm.d1.6xlarge`.
     * 
     */
    public Output<String> coreSpec() {
        return this.coreSpec;
    }
    /**
     * The deletion protection of instance.
     * 
     */
    @Export(name="deletionProection", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> deletionProection;

    /**
     * @return The deletion protection of instance.
     * 
     */
    public Output<Boolean> deletionProection() {
        return this.deletionProection;
    }
    /**
     * The disk type of instance. Valid values: `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_pl0`, `capacity_cloud_storage`, `local_ssd_pro`, `local_hdd_pro`. **NOTE:** From version 1.207.0, `disk_category` can be set to `cloud_essd_pl0`.
     * 
     */
    @Export(name="diskCategory", refs={String.class}, tree="[0]")
    private Output<String> diskCategory;

    /**
     * @return The disk type of instance. Valid values: `cloud_efficiency`, `cloud_ssd`, `cloud_essd`, `cloud_essd_pl0`, `capacity_cloud_storage`, `local_ssd_pro`, `local_hdd_pro`. **NOTE:** From version 1.207.0, `disk_category` can be set to `cloud_essd_pl0`.
     * 
     */
    public Output<String> diskCategory() {
        return this.diskCategory;
    }
    /**
     * The duration of paid. Valid when the `payment_type` is `Subscription`.  When `pricing_cycle` set to `Month`, the valid value id `1` to `9`.  When `pricing_cycle` set to `Year`, the valid value id `1` to `3`.
     * 
     */
    @Export(name="duration", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> duration;

    /**
     * @return The duration of paid. Valid when the `payment_type` is `Subscription`.  When `pricing_cycle` set to `Month`, the valid value id `1` to `9`.  When `pricing_cycle` set to `Year`, the valid value id `1` to `3`.
     * 
     */
    public Output<Optional<String>> duration() {
        return Codegen.optional(this.duration);
    }
    /**
     * (Available since v1.163.0) Whether to enable file engine.
     * 
     */
    @Export(name="enabledFileEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledFileEngine;

    /**
     * @return (Available since v1.163.0) Whether to enable file engine.
     * 
     */
    public Output<Boolean> enabledFileEngine() {
        return this.enabledFileEngine;
    }
    /**
     * (Available since v1.163.0) Whether to enable lts engine.
     * 
     */
    @Export(name="enabledLtsEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledLtsEngine;

    /**
     * @return (Available since v1.163.0) Whether to enable lts engine.
     * 
     */
    public Output<Boolean> enabledLtsEngine() {
        return this.enabledLtsEngine;
    }
    /**
     * (Available since v1.163.0) Whether to enable search engine.
     * 
     */
    @Export(name="enabledSearchEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledSearchEngine;

    /**
     * @return (Available since v1.163.0) Whether to enable search engine.
     * 
     */
    public Output<Boolean> enabledSearchEngine() {
        return this.enabledSearchEngine;
    }
    /**
     * (Available since v1.211.0) Whether to enable streaming engine.
     * 
     */
    @Export(name="enabledStreamEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledStreamEngine;

    /**
     * @return (Available since v1.211.0) Whether to enable streaming engine.
     * 
     */
    public Output<Boolean> enabledStreamEngine() {
        return this.enabledStreamEngine;
    }
    /**
     * (Available since v1.163.0) Whether to enable table engine.
     * 
     */
    @Export(name="enabledTableEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledTableEngine;

    /**
     * @return (Available since v1.163.0) Whether to enable table engine.
     * 
     */
    public Output<Boolean> enabledTableEngine() {
        return this.enabledTableEngine;
    }
    /**
     * (Available since v1.163.0) Whether to enable time serires engine.
     * 
     */
    @Export(name="enabledTimeSeriresEngine", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabledTimeSeriresEngine;

    /**
     * @return (Available since v1.163.0) Whether to enable time serires engine.
     * 
     */
    public Output<Boolean> enabledTimeSeriresEngine() {
        return this.enabledTimeSeriresEngine;
    }
    /**
     * The count of file engine.
     * 
     */
    @Export(name="fileEngineNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> fileEngineNodeCount;

    /**
     * @return The count of file engine.
     * 
     */
    public Output<Integer> fileEngineNodeCount() {
        return this.fileEngineNodeCount;
    }
    /**
     * The specification of file engine. Valid values: `lindorm.c.xlarge`.
     * 
     */
    @Export(name="fileEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> fileEngineSpecification;

    /**
     * @return The specification of file engine. Valid values: `lindorm.c.xlarge`.
     * 
     */
    public Output<String> fileEngineSpecification() {
        return this.fileEngineSpecification;
    }
    /**
     * The name of the instance.
     * 
     */
    @Export(name="instanceName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> instanceName;

    /**
     * @return The name of the instance.
     * 
     */
    public Output<Optional<String>> instanceName() {
        return Codegen.optional(this.instanceName);
    }
    /**
     * The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
     * 
     */
    @Export(name="instanceStorage", refs={String.class}, tree="[0]")
    private Output<String> instanceStorage;

    /**
     * @return The storage capacity of the instance. Unit: GB. For example, the value 50 indicates 50 GB.
     * 
     */
    public Output<String> instanceStorage() {
        return this.instanceStorage;
    }
    /**
     * The ip white list of instance.
     * 
     */
    @Export(name="ipWhiteLists", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ipWhiteLists;

    /**
     * @return The ip white list of instance.
     * 
     */
    public Output<Optional<List<String>>> ipWhiteLists() {
        return Codegen.optional(this.ipWhiteLists);
    }
    /**
     * The multi-available zone instance, log node disk type. required if you need to create multiple availability zone instances. Valid values: `cloud_efficiency`, `cloud_ssd`.
     * 
     */
    @Export(name="logDiskCategory", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logDiskCategory;

    /**
     * @return The multi-available zone instance, log node disk type. required if you need to create multiple availability zone instances. Valid values: `cloud_efficiency`, `cloud_ssd`.
     * 
     */
    public Output<Optional<String>> logDiskCategory() {
        return Codegen.optional(this.logDiskCategory);
    }
    /**
     * The multiple Availability Zone Instance, number of log nodes. this parameter is required if you want to create multiple availability zone instances. Valid values: `4` to `400`.
     * 
     */
    @Export(name="logNum", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> logNum;

    /**
     * @return The multiple Availability Zone Instance, number of log nodes. this parameter is required if you want to create multiple availability zone instances. Valid values: `4` to `400`.
     * 
     */
    public Output<Optional<Integer>> logNum() {
        return Codegen.optional(this.logNum);
    }
    /**
     * The multi-availability instance, log single-node disk capacity. This parameter is required if you want to create multiple availability zone instances. Valid values: `400` to `64000`.
     * 
     */
    @Export(name="logSingleStorage", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> logSingleStorage;

    /**
     * @return The multi-availability instance, log single-node disk capacity. This parameter is required if you want to create multiple availability zone instances. Valid values: `400` to `64000`.
     * 
     */
    public Output<Optional<Integer>> logSingleStorage() {
        return Codegen.optional(this.logSingleStorage);
    }
    /**
     * The multiple availability zone instances, log node specification. required if you need to create multiple availability zone instances. Valid values: `lindorm.sn1.large`, `lindorm.sn1.2xlarge`.
     * 
     */
    @Export(name="logSpec", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logSpec;

    /**
     * @return The multiple availability zone instances, log node specification. required if you need to create multiple availability zone instances. Valid values: `lindorm.sn1.large`, `lindorm.sn1.2xlarge`.
     * 
     */
    public Output<Optional<String>> logSpec() {
        return Codegen.optional(this.logSpec);
    }
    /**
     * The count of lindorm tunnel service.
     * 
     */
    @Export(name="ltsNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> ltsNodeCount;

    /**
     * @return The count of lindorm tunnel service.
     * 
     */
    public Output<Integer> ltsNodeCount() {
        return this.ltsNodeCount;
    }
    /**
     * The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
     * 
     */
    @Export(name="ltsNodeSpecification", refs={String.class}, tree="[0]")
    private Output<String> ltsNodeSpecification;

    /**
     * @return The specification of lindorm tunnel service. Valid values: `lindorm.g.2xlarge`, `lindorm.g.xlarge`.
     * 
     */
    public Output<String> ltsNodeSpecification() {
        return this.ltsNodeSpecification;
    }
    /**
     * The multi-zone combinations. Availability zone combinations are supported on the sale page. required if you need to create multiple availability zone instances. Valid values: `ap-southeast-5abc-aliyun`, `cn-hangzhou-ehi-aliyun`, `cn-beijing-acd-aliyun`, `ap-southeast-1-abc-aliyun`, `cn-zhangjiakou-abc-aliyun`, `cn-shanghai-efg-aliyun`, `cn-shanghai-abd-aliyun`, `cn-hangzhou-bef-aliyun`, `cn-hangzhou-bce-aliyun`, `cn-beijing-fgh-aliyun`, `cn-shenzhen-abc-aliyun`.
     * 
     */
    @Export(name="multiZoneCombination", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> multiZoneCombination;

    /**
     * @return The multi-zone combinations. Availability zone combinations are supported on the sale page. required if you need to create multiple availability zone instances. Valid values: `ap-southeast-5abc-aliyun`, `cn-hangzhou-ehi-aliyun`, `cn-beijing-acd-aliyun`, `ap-southeast-1-abc-aliyun`, `cn-zhangjiakou-abc-aliyun`, `cn-shanghai-efg-aliyun`, `cn-shanghai-abd-aliyun`, `cn-hangzhou-bef-aliyun`, `cn-hangzhou-bce-aliyun`, `cn-beijing-fgh-aliyun`, `cn-shenzhen-abc-aliyun`.
     * 
     */
    public Output<Optional<String>> multiZoneCombination() {
        return Codegen.optional(this.multiZoneCombination);
    }
    /**
     * The billing method. Valid values: `PayAsYouGo` and `Subscription`.
     * 
     */
    @Export(name="paymentType", refs={String.class}, tree="[0]")
    private Output<String> paymentType;

    /**
     * @return The billing method. Valid values: `PayAsYouGo` and `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The pricing cycle. Valid when the `payment_type` is `Subscription`. Valid values: `Month` and `Year`.
     * 
     */
    @Export(name="pricingCycle", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> pricingCycle;

    /**
     * @return The pricing cycle. Valid when the `payment_type` is `Subscription`. Valid values: `Month` and `Year`.
     * 
     */
    public Output<Optional<String>> pricingCycle() {
        return Codegen.optional(this.pricingCycle);
    }
    /**
     * Multi-available zone instances, the virtual switch ID of the primary available zone, must be under the available zone corresponding to the PrimaryZoneId. required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="primaryVswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> primaryVswitchId;

    /**
     * @return Multi-available zone instances, the virtual switch ID of the primary available zone, must be under the available zone corresponding to the PrimaryZoneId. required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> primaryVswitchId() {
        return Codegen.optional(this.primaryVswitchId);
    }
    /**
     * Multi-availability zone instance with the availability zone ID of the main availability zone. required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="primaryZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> primaryZoneId;

    /**
     * @return Multi-availability zone instance with the availability zone ID of the main availability zone. required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> primaryZoneId() {
        return Codegen.optional(this.primaryZoneId);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The count of search engine.
     * 
     */
    @Export(name="searchEngineNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> searchEngineNodeCount;

    /**
     * @return The count of search engine.
     * 
     */
    public Output<Integer> searchEngineNodeCount() {
        return this.searchEngineNodeCount;
    }
    /**
     * The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     * 
     */
    @Export(name="searchEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> searchEngineSpecification;

    /**
     * @return The specification of search engine. Valid values: `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.g.xlarge`.
     * 
     */
    public Output<String> searchEngineSpecification() {
        return this.searchEngineSpecification;
    }
    /**
     * (Available since v1.196.0) The instance type.
     * 
     */
    @Export(name="serviceType", refs={String.class}, tree="[0]")
    private Output<String> serviceType;

    /**
     * @return (Available since v1.196.0) The instance type.
     * 
     */
    public Output<String> serviceType() {
        return this.serviceType;
    }
    /**
     * The multiple availability zone instances, the virtual switch ID of the ready availability zone must be under the availability zone corresponding to the StandbyZoneId. required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="standbyVswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> standbyVswitchId;

    /**
     * @return The multiple availability zone instances, the virtual switch ID of the ready availability zone must be under the availability zone corresponding to the StandbyZoneId. required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> standbyVswitchId() {
        return Codegen.optional(this.standbyVswitchId);
    }
    /**
     * The multiple availability zone instances with availability zone IDs for the prepared availability zones. required if you need to create multiple availability zone instances.
     * 
     */
    @Export(name="standbyZoneId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> standbyZoneId;

    /**
     * @return The multiple availability zone instances with availability zone IDs for the prepared availability zones. required if you need to create multiple availability zone instances.
     * 
     */
    public Output<Optional<String>> standbyZoneId() {
        return Codegen.optional(this.standbyZoneId);
    }
    /**
     * The status of Instance.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of Instance.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The number of LindormStream nodes in the instance.
     * 
     */
    @Export(name="streamEngineNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> streamEngineNodeCount;

    /**
     * @return The number of LindormStream nodes in the instance.
     * 
     */
    public Output<Integer> streamEngineNodeCount() {
        return this.streamEngineNodeCount;
    }
    /**
     * The specification of the LindormStream nodes in the instance. Valid values: `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`.
     * 
     */
    @Export(name="streamEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> streamEngineSpecification;

    /**
     * @return The specification of the LindormStream nodes in the instance. Valid values: `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`.
     * 
     */
    public Output<String> streamEngineSpecification() {
        return this.streamEngineSpecification;
    }
    /**
     * The count of table engine.
     * 
     */
    @Export(name="tableEngineNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> tableEngineNodeCount;

    /**
     * @return The count of table engine.
     * 
     */
    public Output<Integer> tableEngineNodeCount() {
        return this.tableEngineNodeCount;
    }
    /**
     * The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`.
     * 
     */
    @Export(name="tableEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> tableEngineSpecification;

    /**
     * @return The specification of  table engine. Valid values: `lindorm.c.2xlarge`, `lindorm.c.4xlarge`, `lindorm.c.8xlarge`, `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`.
     * 
     */
    public Output<String> tableEngineSpecification() {
        return this.tableEngineSpecification;
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }
    /**
     * The count of time series engine.
     * 
     */
    @Export(name="timeSeriesEngineNodeCount", refs={Integer.class}, tree="[0]")
    private Output<Integer> timeSeriesEngineNodeCount;

    /**
     * @return The count of time series engine.
     * 
     */
    public Output<Integer> timeSeriesEngineNodeCount() {
        return this.timeSeriesEngineNodeCount;
    }
    /**
     * The specification of time series engine. Valid values: `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.r.8xlarge`.
     * 
     */
    @Export(name="timeSeriesEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> timeSeriesEngineSpecification;

    /**
     * @return The specification of time series engine. Valid values: `lindorm.g.xlarge`, `lindorm.g.2xlarge`, `lindorm.g.4xlarge`, `lindorm.g.8xlarge`, `lindorm.r.8xlarge`.
     * 
     */
    public Output<String> timeSeriesEngineSpecification() {
        return this.timeSeriesEngineSpecification;
    }
    /**
     * Field `time_serires_engine_specification` has been deprecated from provider version 1.182.0. New field `time_series_engine_specification` instead.
     * 
     * @deprecated
     * Field `time_serires_engine_specification` has been deprecated from provider version 1.182.0. New field `time_series_engine_specification` instead.
     * 
     */
    @Deprecated /* Field `time_serires_engine_specification` has been deprecated from provider version 1.182.0. New field `time_series_engine_specification` instead. */
    @Export(name="timeSeriresEngineSpecification", refs={String.class}, tree="[0]")
    private Output<String> timeSeriresEngineSpecification;

    /**
     * @return Field `time_serires_engine_specification` has been deprecated from provider version 1.182.0. New field `time_series_engine_specification` instead.
     * 
     */
    public Output<String> timeSeriresEngineSpecification() {
        return this.timeSeriresEngineSpecification;
    }
    /**
     * The VPC ID of the instance.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The VPC ID of the instance.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The vswitch id.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
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
    public Instance(String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:lindorm/instance:Instance", name, args == null ? InstanceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Instance(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:lindorm/instance:Instance", name, state, makeResourceOptions(options, id));
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
    public static Instance get(String name, Output<String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
