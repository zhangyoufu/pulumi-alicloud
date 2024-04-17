// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.mongodb;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.mongodb.AuditPolicyArgs;
import com.pulumi.alicloud.mongodb.inputs.AuditPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a MongoDB Audit Policy resource.
 * 
 * For information about MongoDB Audit Policy and how to use it, see [What is Audit Policy](https://www.alibabacloud.com/help/doc-detail/131941.html).
 * 
 * &gt; **NOTE:** Available since v1.148.0.
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
 * import com.pulumi.alicloud.mongodb.MongodbFunctions;
 * import com.pulumi.alicloud.mongodb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.mongodb.Instance;
 * import com.pulumi.alicloud.mongodb.InstanceArgs;
 * import com.pulumi.alicloud.mongodb.AuditPolicy;
 * import com.pulumi.alicloud.mongodb.AuditPolicyArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         final var default = MongodbFunctions.getZones();
 * 
 *         final var index = default_.zones().length() - 1;
 * 
 *         final var zoneId = default_.zones()[index].id();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;172.17.3.0/24&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(zoneId)
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engineVersion(&#34;4.2&#34;)
 *             .dbInstanceClass(&#34;dds.mongo.mid&#34;)
 *             .dbInstanceStorage(10)
 *             .vswitchId(defaultSwitch.id())
 *             .securityIpLists(            
 *                 &#34;10.168.1.12&#34;,
 *                 &#34;100.69.7.112&#34;)
 *             .name(name)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;example&#34;)
 *             ))
 *             .build());
 * 
 *         var defaultAuditPolicy = new AuditPolicy(&#34;defaultAuditPolicy&#34;, AuditPolicyArgs.builder()        
 *             .dbInstanceId(defaultInstance.id())
 *             .auditStatus(&#34;disabled&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * MongoDB Audit Policy can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:mongodb/auditPolicy:AuditPolicy example &lt;db_instance_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:mongodb/auditPolicy:AuditPolicy")
public class AuditPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The status of the audit log. Valid values: `disabled`, `enable`.
     * 
     */
    @Export(name="auditStatus", refs={String.class}, tree="[0]")
    private Output<String> auditStatus;

    /**
     * @return The status of the audit log. Valid values: `disabled`, `enable`.
     * 
     */
    public Output<String> auditStatus() {
        return this.auditStatus;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="dbInstanceId", refs={String.class}, tree="[0]")
    private Output<String> dbInstanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }
    /**
     * The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
     * 
     */
    @Export(name="storagePeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> storagePeriod;

    /**
     * @return The retention period of audit logs. Valid values: `1` to `30`. Default value: `30`.
     * 
     */
    public Output<Optional<Integer>> storagePeriod() {
        return Codegen.optional(this.storagePeriod);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuditPolicy(String name) {
        this(name, AuditPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuditPolicy(String name, AuditPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuditPolicy(String name, AuditPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/auditPolicy:AuditPolicy", name, args == null ? AuditPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuditPolicy(String name, Output<String> id, @Nullable AuditPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:mongodb/auditPolicy:AuditPolicy", name, state, makeResourceOptions(options, id));
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
    public static AuditPolicy get(String name, Output<String> id, @Nullable AuditPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuditPolicy(name, id, state, options);
    }
}
