// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.hbase;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.hbase.InstanceArgs;
import com.pulumi.alicloud.hbase.inputs.InstanceState;
import com.pulumi.alicloud.hbase.outputs.InstanceSlbConnAddr;
import com.pulumi.alicloud.hbase.outputs.InstanceUiProxyConnAddr;
import com.pulumi.alicloud.hbase.outputs.InstanceZkConnAddr;
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
 * Provides a HBase instance resource supports replica set instances only. The HBase provides stable, reliable, and automatic scalable database services.
 * It offers a full range of database solutions, such as disaster recovery, backup, recovery, monitoring, and alarms.
 * You can see detail product introduction [here](https://www.alibabacloud.com/help/en/apsaradb-for-hbase/latest/createcluster)
 * 
 * &gt; **NOTE:** Available since v1.67.0.
 * 
 * &gt; **NOTE:**  The following regions don&#39;t support create Classic network HBase instance.
 * [`cn-hangzhou`,`cn-shanghai`,`cn-qingdao`,`cn-beijing`,`cn-shenzhen`,`ap-southeast-1a`,.....]
 * The official website mark  more regions. or you can call [DescribeRegions](https://www.alibabacloud.com/help/en/apsaradb-for-hbase/latest/describeregions)
 * 
 * &gt; **NOTE:**  Create HBase instance or change instance type and storage would cost 15 minutes. Please make full preparation
 * 
 * ## Example Usage
 * 
 * ### Create a hbase instance
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.hbase.HbaseFunctions;
 * import com.pulumi.alicloud.hbase.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.hbase.Instance;
 * import com.pulumi.alicloud.hbase.InstanceArgs;
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
 *         final var name = config.get("name").orElse("tf-example");
 *         final var default = HbaseFunctions.getZones();
 * 
 *         final var defaultGetNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex("^default-NODELETING$")
 *             .build());
 * 
 *         final var defaultGetSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .zoneId(default_.zones()[1].id())
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()
 *             .name(name)
 *             .zoneId(default_.zones()[1].id())
 *             .vswitchId(defaultGetSwitches.applyValue(getSwitchesResult -> getSwitchesResult.ids()[0]))
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .engine("hbaseue")
 *             .engineVersion("2.0")
 *             .masterInstanceType("hbase.sn2.2xlarge")
 *             .coreInstanceType("hbase.sn2.2xlarge")
 *             .coreInstanceQuantity(2)
 *             .coreDiskType("cloud_efficiency")
 *             .coreDiskSize(400)
 *             .payType("PostPaid")
 *             .coldStorageSize(0)
 *             .deletionProtection("false")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * this is a example for class netType instance. you can find more detail with the examples/hbase dir.
 * 
 * ## Import
 * 
 * HBase can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:hbase/instance:Instance example hb-wz96815u13k659fvd
 * ```
 * 
 */
@ResourceType(type="alicloud:hbase/instance:Instance")
public class Instance extends com.pulumi.resources.CustomResource {
    /**
     * The account of the cluster web ui. Size [0-128].
     * 
     */
    @Export(name="account", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> account;

    /**
     * @return The account of the cluster web ui. Size [0-128].
     * 
     */
    public Output<Optional<String>> account() {
        return Codegen.optional(this.account);
    }
    /**
     * Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
     * 
     */
    @Export(name="autoRenew", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> autoRenew;

    /**
     * @return Valid values are `true`, `false`, system default to `false`, valid when pay_type = PrePaid.
     * 
     */
    public Output<Boolean> autoRenew() {
        return this.autoRenew;
    }
    /**
     * 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
     * 
     */
    @Export(name="coldStorageSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> coldStorageSize;

    /**
     * @return 0 or [800, 100000000], step:10-GB increments. 0 means is_cold_storage = false. [800, 100000000] means is_cold_storage = true.
     * 
     */
    public Output<Optional<Integer>> coldStorageSize() {
        return Codegen.optional(this.coldStorageSize);
    }
    /**
     * User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space, value range: [20, 64000].
     * - Cluster [400, 64000], step:40-GB increments.
     * - Single [20-500GB], step:1-GB increments.
     * 
     */
    @Export(name="coreDiskSize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> coreDiskSize;

    /**
     * @return User-defined HBase instance one core node&#39;s storage. Valid when engine=hbase/hbaseue. Bds engine no need core_disk_size, space.Unit: GB. Value range:
     * - Custom storage space, value range: [20, 64000].
     * - Cluster [400, 64000], step:40-GB increments.
     * - Single [20-500GB], step:1-GB increments.
     * 
     */
    public Output<Optional<Integer>> coreDiskSize() {
        return Codegen.optional(this.coreDiskSize);
    }
    /**
     * Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
     * 
     */
    @Export(name="coreDiskType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> coreDiskType;

    /**
     * @return Valid values are `cloud_ssd`, `cloud_essd_pl1`, `cloud_efficiency`, `local_hdd_pro`, `local_ssd_pro`，``, local_disk size is fixed. When engine=bds, no need to set disk type(or empty string).
     * 
     */
    public Output<Optional<String>> coreDiskType() {
        return Codegen.optional(this.coreDiskType);
    }
    /**
     * Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
     * 
     */
    @Export(name="coreInstanceQuantity", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> coreInstanceQuantity;

    /**
     * @return Default=2, [1-200]. If core_instance_quantity &gt; 1, this is cluster&#39;s instance. If core_instance_quantity = 1, this is a single instance.
     * 
     */
    public Output<Optional<Integer>> coreInstanceQuantity() {
        return Codegen.optional(this.coreInstanceQuantity);
    }
    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    @Export(name="coreInstanceType", refs={String.class}, tree="[0]")
    private Output<String> coreInstanceType;

    /**
     * @return Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    public Output<String> coreInstanceType() {
        return this.coreInstanceType;
    }
    /**
     * The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
     * 
     */
    @Export(name="deletionProtection", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> deletionProtection;

    /**
     * @return The switch of delete protection. True: delete protect, False: no delete protect. You must set false when you want to delete cluster.
     * 
     */
    public Output<Optional<Boolean>> deletionProtection() {
        return Codegen.optional(this.deletionProtection);
    }
    /**
     * 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
     * 
     */
    @Export(name="duration", refs={Integer.class}, tree="[0]")
    private Output<Integer> duration;

    /**
     * @return 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, valid when pay_type = PrePaid,  unit: month. 12, 24, 36 mean 1, 2, 3 years.
     * 
     */
    public Output<Integer> duration() {
        return this.duration;
    }
    /**
     * Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
     * 
     */
    @Export(name="engine", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> engine;

    /**
     * @return Valid values are &#34;hbase/hbaseue/bds&#34;. The following types are supported after v1.73.0: `hbaseue` and `bds`. Single hbase instance need to set engine=hbase, core_instance_quantity=1.
     * 
     */
    public Output<Optional<String>> engine() {
        return Codegen.optional(this.engine);
    }
    /**
     * HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://www.alibabacloud.com/help/en/data-lake-analytics/latest/createinstance).
     * 
     */
    @Export(name="engineVersion", refs={String.class}, tree="[0]")
    private Output<String> engineVersion;

    /**
     * @return HBase major version. hbase:1.1/2.0, hbaseue:2.0, bds:1.0, unsupport other engine temporarily. Value options can refer to the latest docs [CreateInstance](https://www.alibabacloud.com/help/en/data-lake-analytics/latest/createinstance).
     * 
     */
    public Output<String> engineVersion() {
        return this.engineVersion;
    }
    /**
     * The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
     * 
     */
    @Export(name="immediateDeleteFlag", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> immediateDeleteFlag;

    /**
     * @return The switch of delete immediate. True: delete immediate, False: delete delay. You will not found the cluster no matter set true or false.
     * 
     */
    public Output<Optional<Boolean>> immediateDeleteFlag() {
        return Codegen.optional(this.immediateDeleteFlag);
    }
    /**
     * The white ip list of the cluster.
     * 
     */
    @Export(name="ipWhite", refs={String.class}, tree="[0]")
    private Output<String> ipWhite;

    /**
     * @return The white ip list of the cluster.
     * 
     */
    public Output<String> ipWhite() {
        return this.ipWhite;
    }
    /**
     * The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
     * 
     */
    @Export(name="maintainEndTime", refs={String.class}, tree="[0]")
    private Output<String> maintainEndTime;

    /**
     * @return The end time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 04:00Z.
     * 
     */
    public Output<String> maintainEndTime() {
        return this.maintainEndTime;
    }
    /**
     * The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
     * 
     */
    @Export(name="maintainStartTime", refs={String.class}, tree="[0]")
    private Output<String> maintainStartTime;

    /**
     * @return The start time of the operation and maintenance time period of the instance, in the format of HH:mmZ (UTC time), for example 02:00Z.
     * 
     */
    public Output<String> maintainStartTime() {
        return this.maintainStartTime;
    }
    /**
     * Count nodes of the master node.
     * 
     */
    @Export(name="masterInstanceQuantity", refs={Integer.class}, tree="[0]")
    private Output<Integer> masterInstanceQuantity;

    /**
     * @return Count nodes of the master node.
     * 
     */
    public Output<Integer> masterInstanceQuantity() {
        return this.masterInstanceQuantity;
    }
    /**
     * Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    @Export(name="masterInstanceType", refs={String.class}, tree="[0]")
    private Output<String> masterInstanceType;

    /**
     * @return Instance specification. See [Instance specifications](https://help.aliyun.com/document_detail/53532.html), or you can call describeInstanceType api.
     * 
     */
    public Output<String> masterInstanceType() {
        return this.masterInstanceType;
    }
    /**
     * HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return HBase instance name. Length must be 2-128 characters long. Only Chinese characters, English letters, numbers, period (.), underline (_), or dash (-) are permitted.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The password of the cluster web ui account. Size [0-128].
     * 
     */
    @Export(name="password", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> password;

    /**
     * @return The password of the cluster web ui account. Size [0-128].
     * 
     */
    public Output<Optional<String>> password() {
        return Codegen.optional(this.password);
    }
    /**
     * Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
     * 
     */
    @Export(name="payType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> payType;

    /**
     * @return Valid values are `PrePaid`, `PostPaid`, System default to `PostPaid`. You can also convert PostPaid to PrePaid. And support convert PrePaid to PostPaid from 1.115.0+.
     * 
     */
    public Output<Optional<String>> payType() {
        return Codegen.optional(this.payType);
    }
    /**
     * The security group resource of the cluster.
     * 
     */
    @Export(name="securityGroups", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> securityGroups;

    /**
     * @return The security group resource of the cluster.
     * 
     */
    public Output<List<String>> securityGroups() {
        return this.securityGroups;
    }
    /**
     * The slb service addresses of the cluster. See `slb_conn_addrs` below.
     * 
     * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
     * 
     */
    @Export(name="slbConnAddrs", refs={List.class,InstanceSlbConnAddr.class}, tree="[0,1]")
    private Output<List<InstanceSlbConnAddr>> slbConnAddrs;

    /**
     * @return The slb service addresses of the cluster. See `slb_conn_addrs` below.
     * 
     * &gt; **NOTE:** Now only instance name can be change. The others(instance_type, disk_size, core_instance_quantity and so on) will be supported in the furture.
     * 
     */
    public Output<List<InstanceSlbConnAddr>> slbConnAddrs() {
        return this.slbConnAddrs;
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
     * The Web UI proxy addresses of the cluster. See `ui_proxy_conn_addrs` below.
     * 
     */
    @Export(name="uiProxyConnAddrs", refs={List.class,InstanceUiProxyConnAddr.class}, tree="[0,1]")
    private Output<List<InstanceUiProxyConnAddr>> uiProxyConnAddrs;

    /**
     * @return The Web UI proxy addresses of the cluster. See `ui_proxy_conn_addrs` below.
     * 
     */
    public Output<List<InstanceUiProxyConnAddr>> uiProxyConnAddrs() {
        return this.uiProxyConnAddrs;
    }
    /**
     * The id of the VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vpcId;

    /**
     * @return The id of the VPC.
     * 
     */
    public Output<Optional<String>> vpcId() {
        return Codegen.optional(this.vpcId);
    }
    /**
     * If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> vswitchId;

    /**
     * @return If vswitch_id is not empty, that mean net_type = vpc and has a same region. If vswitch_id is empty, net_type=classic. Intl site not support classic network.
     * 
     */
    public Output<Optional<String>> vswitchId() {
        return Codegen.optional(this.vswitchId);
    }
    /**
     * The zookeeper addresses of the cluster. See `zk_conn_addrs` below.
     * 
     */
    @Export(name="zkConnAddrs", refs={List.class,InstanceZkConnAddr.class}, tree="[0,1]")
    private Output<List<InstanceZkConnAddr>> zkConnAddrs;

    /**
     * @return The zookeeper addresses of the cluster. See `zk_conn_addrs` below.
     * 
     */
    public Output<List<InstanceZkConnAddr>> zkConnAddrs() {
        return this.zkConnAddrs;
    }
    /**
     * The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The Zone to launch the HBase instance. If vswitch_id is not empty, this zone_id can be &#34;&#34; or consistent.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Instance(java.lang.String name) {
        this(name, InstanceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Instance(java.lang.String name, InstanceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Instance(java.lang.String name, InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbase/instance:Instance", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private Instance(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:hbase/instance:Instance", name, state, makeResourceOptions(options, id), false);
    }

    private static InstanceArgs makeArgs(InstanceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? InstanceArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "password"
            ))
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
    public static Instance get(java.lang.String name, Output<java.lang.String> id, @Nullable InstanceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Instance(name, id, state, options);
    }
}
