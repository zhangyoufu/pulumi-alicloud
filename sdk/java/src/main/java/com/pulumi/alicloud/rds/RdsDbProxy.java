// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.RdsDbProxyArgs;
import com.pulumi.alicloud.rds.inputs.RdsDbProxyState;
import com.pulumi.alicloud.rds.outputs.RdsDbProxyReadOnlyInstanceWeight;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Information about RDS database exclusive agent and its usage, see [Dedicated proxy (read/write splitting).](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/dedicated-proxy).
 * &gt; **NOTE:** Available in 1.193.0+.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.ReadOnlyInstance;
 * import com.pulumi.alicloud.rds.ReadOnlyInstanceArgs;
 * import com.pulumi.alicloud.rds.RdsDbProxy;
 * import com.pulumi.alicloud.rds.RdsDbProxyArgs;
 * import com.pulumi.alicloud.rds.inputs.RdsDbProxyReadOnlyInstanceWeightArgs;
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
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;Rds&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;dbInstancevpc&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;5.7&#34;)
 *             .instanceType(&#34;rds.mysql.c1.large&#34;)
 *             .instanceStorage(&#34;20&#34;)
 *             .instanceChargeType(&#34;Postpaid&#34;)
 *             .instanceName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .dbInstanceStorageType(&#34;local_ssd&#34;)
 *             .build());
 * 
 *         var defaultReadOnlyInstance = new ReadOnlyInstance(&#34;defaultReadOnlyInstance&#34;, ReadOnlyInstanceArgs.builder()        
 *             .masterDbInstanceId(defaultInstance.id())
 *             .zoneId(defaultInstance.zoneId())
 *             .engineVersion(defaultInstance.engineVersion())
 *             .instanceType(&#34;rds.mysql.s3.large&#34;)
 *             .instanceStorage(&#34;20&#34;)
 *             .instanceName(String.format(&#34;%sro&#34;, name))
 *             .vswitchId(defaultSwitch.id())
 *             .build());
 * 
 *         var defaultRdsDbProxy = new RdsDbProxy(&#34;defaultRdsDbProxy&#34;, RdsDbProxyArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .instanceNetworkType(&#34;VPC&#34;)
 *             .vpcId(defaultInstance.vpcId())
 *             .vswitchId(defaultInstance.vswitchId())
 *             .dbProxyInstanceNum(2)
 *             .dbProxyConnectionPrefix(&#34;ttest001&#34;)
 *             .dbProxyConnectStringPort(3306)
 *             .dbProxyEndpointReadWriteMode(&#34;ReadWrite&#34;)
 *             .readOnlyInstanceMaxDelayTime(90)
 *             .dbProxyFeatures(&#34;TransactionReadSqlRouteOptimizeStatus:1;ConnectionPersist:1;ReadWriteSpliting:1&#34;)
 *             .readOnlyInstanceDistributionType(&#34;Custom&#34;)
 *             .readOnlyInstanceWeights(            
 *                 RdsDbProxyReadOnlyInstanceWeightArgs.builder()
 *                     .instanceId(defaultInstance.id())
 *                     .weight(&#34;100&#34;)
 *                     .build(),
 *                 RdsDbProxyReadOnlyInstanceWeightArgs.builder()
 *                     .instanceId(defaultReadOnlyInstance.id())
 *                     .weight(&#34;500&#34;)
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * &gt; **NOTE:** Resource `alicloud.rds.RdsDbProxy` should be created after `alicloud.rds.ReadOnlyInstance`, so the `depends_on` statement is necessary.
 * ## Block read_only_instance_weight
 * 
 * The read_only_instance_weight mapping supports the following:
 * 
 * * `instance_id` - (Required) The Id of the instance and its read-only instances that can run database.
 * * `weight` - (Required) Weight of instances that can run the database and their read-only instances. Read weights increase in increments of 100, and the maximum read weight is 10000.
 * 
 * ## Import
 * 
 * RDS database proxy feature can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rds/rdsDbProxy:RdsDbProxy example abc12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/rdsDbProxy:RdsDbProxy")
public class RdsDbProxy extends com.pulumi.resources.CustomResource {
    /**
     * The port number that is associated with the proxy endpoint.
     * 
     */
    @Export(name="dbProxyConnectStringPort", type=Integer.class, parameters={})
    private Output<Integer> dbProxyConnectStringPort;

    /**
     * @return The port number that is associated with the proxy endpoint.
     * 
     */
    public Output<Integer> dbProxyConnectStringPort() {
        return this.dbProxyConnectStringPort;
    }
    /**
     * The new prefix of the proxy endpoint. Enter a prefix.
     * 
     */
    @Export(name="dbProxyConnectionPrefix", type=String.class, parameters={})
    private Output<String> dbProxyConnectionPrefix;

    /**
     * @return The new prefix of the proxy endpoint. Enter a prefix.
     * 
     */
    public Output<String> dbProxyConnectionPrefix() {
        return this.dbProxyConnectionPrefix;
    }
    /**
     * Connection instance string.
     * 
     */
    @Export(name="dbProxyConnectionString", type=String.class, parameters={})
    private Output<String> dbProxyConnectionString;

    /**
     * @return Connection instance string.
     * 
     */
    public Output<String> dbProxyConnectionString() {
        return this.dbProxyConnectionString;
    }
    /**
     * Remarks of agent terminal.
     * 
     */
    @Export(name="dbProxyEndpointAliases", type=String.class, parameters={})
    private Output<String> dbProxyEndpointAliases;

    /**
     * @return Remarks of agent terminal.
     * 
     */
    public Output<String> dbProxyEndpointAliases() {
        return this.dbProxyEndpointAliases;
    }
    /**
     * Proxy connection address ID.
     * 
     */
    @Export(name="dbProxyEndpointId", type=String.class, parameters={})
    private Output<String> dbProxyEndpointId;

    /**
     * @return Proxy connection address ID.
     * 
     */
    public Output<String> dbProxyEndpointId() {
        return this.dbProxyEndpointId;
    }
    /**
     * The read and write attributes of the proxy terminal. Valid values:
     * - ReadWrite: The proxy terminal connects to the primary instance and can receive both read and write requests.
     * - ReadOnly: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
     * 
     */
    @Export(name="dbProxyEndpointReadWriteMode", type=String.class, parameters={})
    private Output<String> dbProxyEndpointReadWriteMode;

    /**
     * @return The read and write attributes of the proxy terminal. Valid values:
     * - ReadWrite: The proxy terminal connects to the primary instance and can receive both read and write requests.
     * - ReadOnly: The proxy terminal does not connect to the primary instance and can receive only read requests. This is the default value.
     * 
     */
    public Output<String> dbProxyEndpointReadWriteMode() {
        return this.dbProxyEndpointReadWriteMode;
    }
    /**
     * The features that you want to enable for the proxy endpoint. If you specify more than one feature, separate the features with semicolons (;). Format: Feature 1:Status;Feature 2:Status;.... Do not add a semicolon (;) at the end of the last value. Valid feature values:
     * - ReadWriteSpliting: read/write splitting.
     * - ConnectionPersist: connection pooling.
     * - TransactionReadSqlRouteOptimizeStatus: transaction splitting.
     *   Valid status values:
     * - 1: enabled.
     * - 0: disabled.
     * 
     */
    @Export(name="dbProxyFeatures", type=String.class, parameters={})
    private Output<String> dbProxyFeatures;

    /**
     * @return The features that you want to enable for the proxy endpoint. If you specify more than one feature, separate the features with semicolons (;). Format: Feature 1:Status;Feature 2:Status;.... Do not add a semicolon (;) at the end of the last value. Valid feature values:
     * - ReadWriteSpliting: read/write splitting.
     * - ConnectionPersist: connection pooling.
     * - TransactionReadSqlRouteOptimizeStatus: transaction splitting.
     *   Valid status values:
     * - 1: enabled.
     * - 0: disabled.
     * 
     */
    public Output<String> dbProxyFeatures() {
        return this.dbProxyFeatures;
    }
    /**
     * The number of proxy instances that are enabled. Valid values: 1 to 60.
     * 
     */
    @Export(name="dbProxyInstanceNum", type=Integer.class, parameters={})
    private Output<Integer> dbProxyInstanceNum;

    /**
     * @return The number of proxy instances that are enabled. Valid values: 1 to 60.
     * 
     */
    public Output<Integer> dbProxyInstanceNum() {
        return this.dbProxyInstanceNum;
    }
    /**
     * The SSL configuration setting that you want to apply on the instance. Valid values:
     * - Close: disables SSL encryption.
     * - Open: enables SSL encryption or modifies the endpoint that requires SSL encryption.
     * - Update: updates the validity period of the SSL certificate.
     * 
     */
    @Export(name="dbProxySslEnabled", type=String.class, parameters={})
    private Output<String> dbProxySslEnabled;

    /**
     * @return The SSL configuration setting that you want to apply on the instance. Valid values:
     * - Close: disables SSL encryption.
     * - Open: enables SSL encryption or modifies the endpoint that requires SSL encryption.
     * - Update: updates the validity period of the SSL certificate.
     * 
     */
    public Output<String> dbProxySslEnabled() {
        return this.dbProxySslEnabled;
    }
    /**
     * The point in time at which you want to apply the new database proxy settings. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
     * 
     */
    @Export(name="effectiveSpecificTime", type=String.class, parameters={})
    private Output<String> effectiveSpecificTime;

    /**
     * @return The point in time at which you want to apply the new database proxy settings. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
     * 
     */
    public Output<String> effectiveSpecificTime() {
        return this.effectiveSpecificTime;
    }
    /**
     * When modifying the number of proxy instances,The time when you want to apply the new database proxy settings.Valid values:
     * - Immediate: ApsaraDB RDS immediately applies the new settings.
     * - MaintainTime: ApsaraDB RDS applies the new settings during the maintenance window that you specified. For more information, see Modify the maintenance window.
     * - SpecificTime: ApsaraDB RDS applies the new settings at a specified point in time.
     * 
     */
    @Export(name="effectiveTime", type=String.class, parameters={})
    private Output<String> effectiveTime;

    /**
     * @return When modifying the number of proxy instances,The time when you want to apply the new database proxy settings.Valid values:
     * - Immediate: ApsaraDB RDS immediately applies the new settings.
     * - MaintainTime: ApsaraDB RDS applies the new settings during the maintenance window that you specified. For more information, see Modify the maintenance window.
     * - SpecificTime: ApsaraDB RDS applies the new settings at a specified point in time.
     * 
     */
    public Output<String> effectiveTime() {
        return this.effectiveTime;
    }
    /**
     * The Id of instance that can run database.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The Id of instance that can run database.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The network type of the instance. Set the value to VPC.
     * 
     */
    @Export(name="instanceNetworkType", type=String.class, parameters={})
    private Output<String> instanceNetworkType;

    /**
     * @return The network type of the instance. Set the value to VPC.
     * 
     */
    public Output<String> instanceNetworkType() {
        return this.instanceNetworkType;
    }
    /**
     * Network type of proxy connection address.
     * 
     */
    @Export(name="netType", type=String.class, parameters={})
    private Output<String> netType;

    /**
     * @return Network type of proxy connection address.
     * 
     */
    public Output<String> netType() {
        return this.netType;
    }
    /**
     * The policy that is used to allocate read weights. Valid values:
     * - Standard: ApsaraDB RDS automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
     * - Custom: You must manually allocate read weights to the instance and its read-only instances.
     * 
     */
    @Export(name="readOnlyInstanceDistributionType", type=String.class, parameters={})
    private Output<String> readOnlyInstanceDistributionType;

    /**
     * @return The policy that is used to allocate read weights. Valid values:
     * - Standard: ApsaraDB RDS automatically allocates read weights to the instance and its read-only instances based on the specifications of the instances.
     * - Custom: You must manually allocate read weights to the instance and its read-only instances.
     * 
     */
    public Output<String> readOnlyInstanceDistributionType() {
        return this.readOnlyInstanceDistributionType;
    }
    /**
     * The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, ApsaraDB RDS no longer forwards read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained. Unit: seconds. Valid values: 0 to 3600.
     * 
     */
    @Export(name="readOnlyInstanceMaxDelayTime", type=Integer.class, parameters={})
    private Output<Integer> readOnlyInstanceMaxDelayTime;

    /**
     * @return The maximum latency threshold that is allowed for read/write splitting. If the latency on a read-only instance exceeds the threshold that you specified, ApsaraDB RDS no longer forwards read requests to the read-only instance. If you do not specify this parameter, the default value of this parameter is retained. Unit: seconds. Valid values: 0 to 3600.
     * 
     */
    public Output<Integer> readOnlyInstanceMaxDelayTime() {
        return this.readOnlyInstanceMaxDelayTime;
    }
    /**
     * A list of the read weights of the instance and its read-only instances.  It contains two sub-fields(instance_id and weight). Read weights increase in increments of 100, and the maximum read weight is 10000.
     * 
     */
    @Export(name="readOnlyInstanceWeights", type=List.class, parameters={RdsDbProxyReadOnlyInstanceWeight.class})
    private Output<List<RdsDbProxyReadOnlyInstanceWeight>> readOnlyInstanceWeights;

    /**
     * @return A list of the read weights of the instance and its read-only instances.  It contains two sub-fields(instance_id and weight). Read weights increase in increments of 100, and the maximum read weight is 10000.
     * 
     */
    public Output<List<RdsDbProxyReadOnlyInstanceWeight>> readOnlyInstanceWeights() {
        return this.readOnlyInstanceWeights;
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The time when the certificate expires.
     * 
     */
    @Export(name="sslExpiredTime", type=String.class, parameters={})
    private Output<String> sslExpiredTime;

    /**
     * @return The time when the certificate expires.
     * 
     */
    public Output<String> sslExpiredTime() {
        return this.sslExpiredTime;
    }
    /**
     * The point in time at which you want to upgrade the database proxy version of the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
     * 
     */
    @Export(name="switchTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> switchTime;

    /**
     * @return The point in time at which you want to upgrade the database proxy version of the instance. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
     * 
     */
    public Output<Optional<String>> switchTime() {
        return Codegen.optional(this.switchTime);
    }
    /**
     * The time when you want to upgrade the database proxy version of the instance. Valid values:
     * - MaintainTime: ApsaraDB RDS performs the upgrade during the maintenance window that you specified. This is the default value. For more information, see Modify the maintenance window.
     * - Immediate: ApsaraDB RDS immediately performs the upgrade.
     * - SpecificTime: ApsaraDB RDS performs the upgrade at a specified point in time.
     * 
     */
    @Export(name="upgradeTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> upgradeTime;

    /**
     * @return The time when you want to upgrade the database proxy version of the instance. Valid values:
     * - MaintainTime: ApsaraDB RDS performs the upgrade during the maintenance window that you specified. This is the default value. For more information, see Modify the maintenance window.
     * - Immediate: ApsaraDB RDS immediately performs the upgrade.
     * - SpecificTime: ApsaraDB RDS performs the upgrade at a specified point in time.
     * 
     */
    public Output<Optional<String>> upgradeTime() {
        return Codegen.optional(this.upgradeTime);
    }
    /**
     * The ID of the virtual private cloud (VPC) to which the instance belongs.
     * 
     */
    @Export(name="vpcId", type=String.class, parameters={})
    private Output<String> vpcId;

    /**
     * @return The ID of the virtual private cloud (VPC) to which the instance belongs.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The ID of the vSwitch that is associated with the specified VPC.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch that is associated with the specified VPC.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdsDbProxy(String name) {
        this(name, RdsDbProxyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdsDbProxy(String name, RdsDbProxyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdsDbProxy(String name, RdsDbProxyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsDbProxy:RdsDbProxy", name, args == null ? RdsDbProxyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdsDbProxy(String name, Output<String> id, @Nullable RdsDbProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsDbProxy:RdsDbProxy", name, state, makeResourceOptions(options, id));
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
    public static RdsDbProxy get(String name, Output<String> id, @Nullable RdsDbProxyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdsDbProxy(name, id, state, options);
    }
}
