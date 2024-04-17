// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.rds;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.rds.DbNodeArgs;
import com.pulumi.alicloud.rds.inputs.DbNodeState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provide RDS cluster instance to increase node resources, see [What is RDS DB Node](https://www.alibabacloud.com/help/en/apsaradb-for-rds/latest/api-rds-2014-08-15-createdbnodes).
 * 
 * &gt; **NOTE:** Available since v1.202.0.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.rds.RdsFunctions;
 * import com.pulumi.alicloud.rds.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.rds.inputs.GetInstanceClassesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.rds.Instance;
 * import com.pulumi.alicloud.rds.InstanceArgs;
 * import com.pulumi.alicloud.rds.DbNode;
 * import com.pulumi.alicloud.rds.DbNodeArgs;
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
 *         final var default = RdsFunctions.getZones(GetZonesArgs.builder()
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .instanceChargeType(&#34;PostPaid&#34;)
 *             .category(&#34;cluster&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceClasses = RdsFunctions.getInstanceClasses(GetInstanceClassesArgs.builder()
 *             .zoneId(default_.ids()[0])
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .category(&#34;cluster&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .instanceChargeType(&#34;PostPaid&#34;)
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
 *             .zoneId(default_.ids()[0])
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .name(name)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .engineVersion(&#34;8.0&#34;)
 *             .instanceType(defaultGetInstanceClasses.applyValue(getInstanceClassesResult -&gt; getInstanceClassesResult.instanceClasses()[0].instanceClass()))
 *             .instanceStorage(defaultGetInstanceClasses.applyValue(getInstanceClassesResult -&gt; getInstanceClassesResult.instanceClasses()[0].storageRange().min()))
 *             .instanceChargeType(&#34;Postpaid&#34;)
 *             .instanceName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .monitoringPeriod(&#34;60&#34;)
 *             .dbInstanceStorageType(&#34;cloud_essd&#34;)
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .zoneId(default_.ids()[0])
 *             .zoneIdSlaveA(default_.ids()[0])
 *             .build());
 * 
 *         var defaultDbNode = new DbNode(&#34;defaultDbNode&#34;, DbNodeArgs.builder()        
 *             .dbInstanceId(defaultInstance.id())
 *             .classCode(defaultInstance.instanceType())
 *             .zoneId(defaultSwitch.zoneId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * RDS MySQL database cluster node agent function can be imported using id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:rds/dbNode:DbNode example &lt;db_instance_id&gt;:&lt;node_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:rds/dbNode:DbNode")
public class DbNode extends com.pulumi.resources.CustomResource {
    /**
     * The specification information of the node.
     * 
     */
    @Export(name="classCode", refs={String.class}, tree="[0]")
    private Output<String> classCode;

    /**
     * @return The specification information of the node.
     * 
     */
    public Output<String> classCode() {
        return this.classCode;
    }
    /**
     * The Id of instance that can run database.
     * 
     */
    @Export(name="dbInstanceId", refs={String.class}, tree="[0]")
    private Output<String> dbInstanceId;

    /**
     * @return The Id of instance that can run database.
     * 
     */
    public Output<String> dbInstanceId() {
        return this.dbInstanceId;
    }
    /**
     * The ID of the node.
     * 
     */
    @Export(name="nodeId", refs={String.class}, tree="[0]")
    private Output<String> nodeId;

    /**
     * @return The ID of the node.
     * 
     */
    public Output<String> nodeId() {
        return this.nodeId;
    }
    /**
     * The region ID of the node.
     * 
     */
    @Export(name="nodeRegionId", refs={String.class}, tree="[0]")
    private Output<String> nodeRegionId;

    /**
     * @return The region ID of the node.
     * 
     */
    public Output<String> nodeRegionId() {
        return this.nodeRegionId;
    }
    /**
     * The role of node.
     * 
     */
    @Export(name="nodeRole", refs={String.class}, tree="[0]")
    private Output<String> nodeRole;

    /**
     * @return The role of node.
     * 
     */
    public Output<String> nodeRole() {
        return this.nodeRole;
    }
    /**
     * The zone ID of the node.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The zone ID of the node.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DbNode(String name) {
        this(name, DbNodeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DbNode(String name, DbNodeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DbNode(String name, DbNodeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/dbNode:DbNode", name, args == null ? DbNodeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DbNode(String name, Output<String> id, @Nullable DbNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:rds/dbNode:DbNode", name, state, makeResourceOptions(options, id));
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
    public static DbNode get(String name, Output<String> id, @Nullable DbNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DbNode(name, id, state, options);
    }
}
