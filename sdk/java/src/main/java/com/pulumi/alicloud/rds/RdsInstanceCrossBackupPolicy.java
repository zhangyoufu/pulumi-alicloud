// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.RdsInstanceCrossBackupPolicyArgs;
import com.pulumi.alicloud.rds.inputs.RdsInstanceCrossBackupPolicyState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Integer;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides an RDS instance emote disaster recovery strategy policy resource and used to configure instance emote disaster recovery strategy policy.
 * 
 * For information about RDS cross region backup settings and how to use them, see [What is cross region backup](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/modify-cross-region-backup-settings).
 * 
 * &gt; **NOTE:** Available in 1.195.0+.
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
 * import com.pulumi.alicloud.rds.RdsFunctions;
 * import com.pulumi.alicloud.rds.inputs.GetCrossRegionsArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.RdsInstanceCrossBackupPolicy;
 * import com.pulumi.alicloud.rds.RdsInstanceCrossBackupPolicyArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf-testAccRdsCrossRegionBackupPolicy&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(creation)
 *             .build());
 * 
 *         final var regions = RdsFunctions.getCrossRegions();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/24&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[5].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var instance = new Instance(&#34;instance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .dbInstanceStorageType(&#34;local_ssd&#34;)
 *             .instanceType(&#34;rds.mysql.c1.large&#34;)
 *             .instanceStorage(&#34;10&#34;)
 *             .vswitchId(defaultSwitch.id())
 *             .instanceName(name)
 *             .build());
 * 
 *         var policy = new RdsInstanceCrossBackupPolicy(&#34;policy&#34;, RdsInstanceCrossBackupPolicyArgs.builder()        
 *             .instanceId(instance.id())
 *             .crossBackupRegion(regions.applyValue(getCrossRegionsResult -&gt; getCrossRegionsResult.ids()[0]))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * RDS remote disaster recovery policies can be imported using id or instance id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy example &#34;rm-12345678&#34;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy")
public class RdsInstanceCrossBackupPolicy extends com.pulumi.resources.CustomResource {
    /**
     * The status of the overall cross-region backup switch on the instance. Valid values:
     * - Disabled
     * - Enable
     * 
     */
    @Export(name="backupEnabled", type=String.class, parameters={})
    private Output<String> backupEnabled;

    /**
     * @return The status of the overall cross-region backup switch on the instance. Valid values:
     * - Disabled
     * - Enable
     * 
     */
    public Output<String> backupEnabled() {
        return this.backupEnabled;
    }
    /**
     * The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="backupEnabledTime", type=String.class, parameters={})
    private Output<String> backupEnabledTime;

    /**
     * @return The time when cross-region backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> backupEnabledTime() {
        return this.backupEnabledTime;
    }
    /**
     * The ID of the destination region where the cross-region backup files of the instance are stored.
     * 
     */
    @Export(name="crossBackupRegion", type=String.class, parameters={})
    private Output<String> crossBackupRegion;

    /**
     * @return The ID of the destination region where the cross-region backup files of the instance are stored.
     * 
     */
    public Output<String> crossBackupRegion() {
        return this.crossBackupRegion;
    }
    /**
     * The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
     * 
     */
    @Export(name="crossBackupType", type=String.class, parameters={})
    private Output<String> crossBackupType;

    /**
     * @return The policy that is used to save cross-region backups of the instance. Default value: 1. The default value 1 indicates that all cross-region backups are saved.
     * 
     */
    public Output<String> crossBackupType() {
        return this.crossBackupType;
    }
    /**
     * The state of the instance. For more information, see Instance status.
     * 
     */
    @Export(name="dbInstanceStatus", type=String.class, parameters={})
    private Output<String> dbInstanceStatus;

    /**
     * @return The state of the instance. For more information, see Instance status.
     * 
     */
    public Output<String> dbInstanceStatus() {
        return this.dbInstanceStatus;
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", type=String.class, parameters={})
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The lock status of the instance. Valid values:
     * - Unlock: The instance is not locked.
     * - ManualLock: The instance is manually locked.
     * - LockByExpiration: The instance is locked upon expiration.
     * - LockByRestoration: The instance is automatically locked before a rollback.
     * - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
     * 
     */
    @Export(name="lockMode", type=String.class, parameters={})
    private Output<String> lockMode;

    /**
     * @return The lock status of the instance. Valid values:
     * - Unlock: The instance is not locked.
     * - ManualLock: The instance is manually locked.
     * - LockByExpiration: The instance is locked upon expiration.
     * - LockByRestoration: The instance is automatically locked before a rollback.
     * - LockByDiskQuota: The instance is automatically locked because its storage space is exhausted. In this situation, the instance is inaccessible.
     * 
     */
    public Output<String> lockMode() {
        return this.lockMode;
    }
    /**
     * The status of the cross-region log backup feature on the instance. Valid values:
     * - Enable: Enables the feature.
     * - Disabled: Disables the feature.
     * 
     */
    @Export(name="logBackupEnabled", type=String.class, parameters={})
    private Output<String> logBackupEnabled;

    /**
     * @return The status of the cross-region log backup feature on the instance. Valid values:
     * - Enable: Enables the feature.
     * - Disabled: Disables the feature.
     * 
     */
    public Output<String> logBackupEnabled() {
        return this.logBackupEnabled;
    }
    /**
     * The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    @Export(name="logBackupEnabledTime", type=String.class, parameters={})
    private Output<String> logBackupEnabledTime;

    /**
     * @return The time when cross-region log backup was enabled on the instance. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
     * 
     */
    public Output<String> logBackupEnabledTime() {
        return this.logBackupEnabledTime;
    }
    /**
     * The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
     * 
     */
    @Export(name="retentType", type=String.class, parameters={})
    private Output<String> retentType;

    /**
     * @return The policy that is used to retain cross-region backups of the instance. Default value: 1. The default value 1 indicate that cross-region backups are retained based on the specified retention period.
     * 
     */
    public Output<String> retentType() {
        return this.retentType;
    }
    /**
     * The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
     * 
     */
    @Export(name="retention", type=Integer.class, parameters={})
    private Output<Integer> retention;

    /**
     * @return The number of days for which the cross-region backup files of the instance are retained. Valid values: 7 to 1825. Default value: 7.
     * 
     */
    public Output<Integer> retention() {
        return this.retention;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RdsInstanceCrossBackupPolicy(String name) {
        this(name, RdsInstanceCrossBackupPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RdsInstanceCrossBackupPolicy(String name, RdsInstanceCrossBackupPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RdsInstanceCrossBackupPolicy(String name, RdsInstanceCrossBackupPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy", name, args == null ? RdsInstanceCrossBackupPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RdsInstanceCrossBackupPolicy(String name, Output<String> id, @Nullable RdsInstanceCrossBackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/rdsInstanceCrossBackupPolicy:RdsInstanceCrossBackupPolicy", name, state, makeResourceOptions(options, id));
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
    public static RdsInstanceCrossBackupPolicy get(String name, Output<String> id, @Nullable RdsInstanceCrossBackupPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RdsInstanceCrossBackupPolicy(name, id, state, options);
    }
}
