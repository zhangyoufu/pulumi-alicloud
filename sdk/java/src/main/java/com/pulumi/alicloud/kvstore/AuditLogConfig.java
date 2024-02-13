// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.kvstore.AuditLogConfigArgs;
import com.pulumi.alicloud.kvstore.inputs.AuditLogConfigState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Redis And Memcache (KVStore) Audit Log Config resource.
 * 
 * &gt; **NOTE:** Available since v1.130.0.
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
 * import com.pulumi.alicloud.kvstore.KvstoreFunctions;
 * import com.pulumi.alicloud.kvstore.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.kvstore.Instance;
 * import com.pulumi.alicloud.kvstore.InstanceArgs;
 * import com.pulumi.alicloud.kvstore.AuditLogConfig;
 * import com.pulumi.alicloud.kvstore.AuditLogConfigArgs;
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
 *         final var defaultZones = KvstoreFunctions.getZones();
 * 
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups(GetResourceGroupsArgs.builder()
 *             .status(&#34;OK&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .dbInstanceName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .instanceClass(&#34;redis.master.large.default&#34;)
 *             .instanceType(&#34;Redis&#34;)
 *             .engineVersion(&#34;5.0&#34;)
 *             .securityIps(&#34;10.23.12.24&#34;)
 *             .config(Map.ofEntries(
 *                 Map.entry(&#34;appendonly&#34;, &#34;yes&#34;),
 *                 Map.entry(&#34;lazyfree-lazy-eviction&#34;, &#34;yes&#34;)
 *             ))
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;example&#34;)
 *             ))
 *             .build());
 * 
 *         var example = new AuditLogConfig(&#34;example&#34;, AuditLogConfigArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .dbAudit(true)
 *             .retention(1)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Redis And Memcache (KVStore) Audit Log Config can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:kvstore/auditLogConfig:AuditLogConfig example &lt;instance_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:kvstore/auditLogConfig:AuditLogConfig")
public class AuditLogConfig extends com.pulumi.resources.CustomResource {
    /**
     * Instance Creation Time.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return Instance Creation Time.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * Indicates Whether to Enable the Audit Log.  Valid value:
     * * true: Default Value, Open.
     * * false: Closed.
     * 
     * Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn&#39;t Support Separate Open.
     * 
     */
    @Export(name="dbAudit", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dbAudit;

    /**
     * @return Indicates Whether to Enable the Audit Log.  Valid value:
     * * true: Default Value, Open.
     * * false: Closed.
     * 
     * Note: When the Instance for the Cluster Architecture Or Read/Write Split Architecture, at the Same Time to Open Or Close the Data Node and the Proxy Node of the Audit Log Doesn&#39;t Support Separate Open.
     * 
     */
    public Output<Optional<Boolean>> dbAudit() {
        return Codegen.optional(this.dbAudit);
    }
    /**
     * Instance ID, Call the Describeinstances Get.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Instance ID, Call the Describeinstances Get.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Audit Log Retention Period Value: 1~365.
     * 
     * &gt; **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
     * 
     */
    @Export(name="retention", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> retention;

    /**
     * @return Audit Log Retention Period Value: 1~365.
     * 
     * &gt; **NOTE**: When the Instance dbaudit Value Is Set to True, This Parameter Entry into Force. The Parameter Setting of the Current Region of All an Apsaradb for Redis Instance for a Data Entry into Force.
     * 
     */
    public Output<Optional<Integer>> retention() {
        return Codegen.optional(this.retention);
    }
    /**
     * The status of the resource.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the resource.
     * 
     */
    public Output<String> status() {
        return this.status;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuditLogConfig(String name) {
        this(name, AuditLogConfigArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuditLogConfig(String name, AuditLogConfigArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuditLogConfig(String name, AuditLogConfigArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:kvstore/auditLogConfig:AuditLogConfig", name, args == null ? AuditLogConfigArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuditLogConfig(String name, Output<String> id, @Nullable AuditLogConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:kvstore/auditLogConfig:AuditLogConfig", name, state, makeResourceOptions(options, id));
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
    public static AuditLogConfig get(String name, Output<String> id, @Nullable AuditLogConfigState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuditLogConfig(name, id, state, options);
    }
}
