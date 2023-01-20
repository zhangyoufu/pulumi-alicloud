// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.kvstore;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.kvstore.BackupPolicyArgs;
import com.pulumi.alicloud.kvstore.inputs.BackupPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.kvstore.Instance;
 * import com.pulumi.alicloud.kvstore.InstanceArgs;
 * import com.pulumi.alicloud.kvstore.BackupPolicy;
 * import com.pulumi.alicloud.kvstore.BackupPolicyArgs;
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
 *         final var creation = config.get(&#34;creation&#34;).orElse(&#34;KVStore&#34;);
 *         final var multiAz = config.get(&#34;multiAz&#34;).orElse(&#34;false&#34;);
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;kvstorebackuppolicyvpc&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .instanceClass(&#34;Memcache&#34;)
 *             .instanceName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .privateIp(&#34;172.16.0.10&#34;)
 *             .securityIps(&#34;10.0.0.1&#34;)
 *             .instanceType(&#34;memcache.master.small.default&#34;)
 *             .engineVersion(&#34;2.8&#34;)
 *             .build());
 * 
 *         var defaultBackupPolicy = new BackupPolicy(&#34;defaultBackupPolicy&#34;, BackupPolicyArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .backupPeriods(            
 *                 &#34;Tuesday&#34;,
 *                 &#34;Wednesday&#34;)
 *             .backupTime(&#34;10:00Z-11:00Z&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * KVStore backup policy can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:kvstore/backupPolicy:BackupPolicy example r-abc12345678
 * ```
 * 
 */
@ResourceType(type="alicloud:kvstore/backupPolicy:BackupPolicy")
public class BackupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     * 
     */
    @Export(name="backupPeriods", type=List.class, parameters={String.class})
    private Output<List<String>> backupPeriods;

    /**
     * @return Backup Cycle. Allowed values: Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday
     * 
     */
    public Output<List<String>> backupPeriods() {
        return this.backupPeriods;
    }
    /**
     * Backup time, in the format of HH:mmZ- HH:mm Z
     * 
     */
    @Export(name="backupTime", type=String.class, parameters={})
    private Output</* @Nullable */ String> backupTime;

    /**
     * @return Backup time, in the format of HH:mmZ- HH:mm Z
     * 
     */
    public Output<Optional<String>> backupTime() {
        return Codegen.optional(this.backupTime);
    }
    /**
     * The id of ApsaraDB for Redis or Memcache intance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The id of ApsaraDB for Redis or Memcache intance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public BackupPolicy(String name) {
        this(name, BackupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public BackupPolicy(String name, BackupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public BackupPolicy(String name, BackupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:kvstore/backupPolicy:BackupPolicy", name, args == null ? BackupPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private BackupPolicy(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:kvstore/backupPolicy:BackupPolicy", name, state, makeResourceOptions(options, id));
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
    public static BackupPolicy get(String name, Output<String> id, @Nullable BackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new BackupPolicy(name, id, state, options);
    }
}
