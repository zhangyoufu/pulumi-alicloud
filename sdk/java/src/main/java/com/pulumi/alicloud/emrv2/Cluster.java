// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.emrv2;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.emrv2.ClusterArgs;
import com.pulumi.alicloud.emrv2.inputs.ClusterState;
import com.pulumi.alicloud.emrv2.outputs.ClusterApplicationConfig;
import com.pulumi.alicloud.emrv2.outputs.ClusterBootstrapScript;
import com.pulumi.alicloud.emrv2.outputs.ClusterNodeAttribute;
import com.pulumi.alicloud.emrv2.outputs.ClusterNodeGroup;
import com.pulumi.alicloud.emrv2.outputs.ClusterSubscriptionConfig;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a EMR cluster resource. This resource is based on EMR&#39;s new version OpenAPI.
 * 
 * For information about EMR New and how to use it, see [Add a domain](https://www.alibabacloud.com/help/doc-detail/28068.htm).
 * 
 * &gt; **NOTE:** Available since v1.199.0.
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
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.EcsKeyPair;
 * import com.pulumi.alicloud.ecs.EcsKeyPairArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ram.Role;
 * import com.pulumi.alicloud.ram.RoleArgs;
 * import com.pulumi.alicloud.emrv2.Cluster;
 * import com.pulumi.alicloud.emrv2.ClusterArgs;
 * import com.pulumi.alicloud.emrv2.inputs.ClusterApplicationConfigArgs;
 * import com.pulumi.alicloud.emrv2.inputs.ClusterNodeAttributeArgs;
 * import com.pulumi.alicloud.emrv2.inputs.ClusterNodeGroupArgs;
 * import com.pulumi.alicloud.emrv2.inputs.ClusterNodeGroupSystemDiskArgs;
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
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableInstanceType(&#34;ecs.g7.xlarge&#34;)
 *             .build());
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .cidrBlock(&#34;172.16.0.0/21&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .vswitchName(name)
 *             .build());
 * 
 *         var defaultEcsKeyPair = new EcsKeyPair(&#34;defaultEcsKeyPair&#34;, EcsKeyPairArgs.builder()        
 *             .keyPairName(name)
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultRole = new Role(&#34;defaultRole&#34;, RoleArgs.builder()        
 *             .document(&#34;&#34;&#34;
 *     {
 *         &#34;Statement&#34;: [
 *         {
 *             &#34;Action&#34;: &#34;sts:AssumeRole&#34;,
 *             &#34;Effect&#34;: &#34;Allow&#34;,
 *             &#34;Principal&#34;: {
 *             &#34;Service&#34;: [
 *                 &#34;emr.aliyuncs.com&#34;,
 *                 &#34;ecs.aliyuncs.com&#34;
 *             ]
 *             }
 *         }
 *         ],
 *         &#34;Version&#34;: &#34;1&#34;
 *     }
 *             &#34;&#34;&#34;)
 *             .description(&#34;this is a role example.&#34;)
 *             .force(true)
 *             .build());
 * 
 *         var defaultCluster = new Cluster(&#34;defaultCluster&#34;, ClusterArgs.builder()        
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .clusterType(&#34;DATALAKE&#34;)
 *             .releaseVersion(&#34;EMR-5.10.0&#34;)
 *             .clusterName(name)
 *             .deployMode(&#34;NORMAL&#34;)
 *             .securityMode(&#34;NORMAL&#34;)
 *             .applications(            
 *                 &#34;HADOOP-COMMON&#34;,
 *                 &#34;HDFS&#34;,
 *                 &#34;YARN&#34;,
 *                 &#34;HIVE&#34;,
 *                 &#34;SPARK3&#34;,
 *                 &#34;TEZ&#34;)
 *             .applicationConfigs(            
 *                 ClusterApplicationConfigArgs.builder()
 *                     .applicationName(&#34;HIVE&#34;)
 *                     .configFileName(&#34;hivemetastore-site.xml&#34;)
 *                     .configItemKey(&#34;hive.metastore.type&#34;)
 *                     .configItemValue(&#34;DLF&#34;)
 *                     .configScope(&#34;CLUSTER&#34;)
 *                     .build(),
 *                 ClusterApplicationConfigArgs.builder()
 *                     .applicationName(&#34;SPARK3&#34;)
 *                     .configFileName(&#34;hive-site.xml&#34;)
 *                     .configItemKey(&#34;hive.metastore.type&#34;)
 *                     .configItemValue(&#34;DLF&#34;)
 *                     .configScope(&#34;CLUSTER&#34;)
 *                     .build())
 *             .nodeAttributes(ClusterNodeAttributeArgs.builder()
 *                 .ramRole(defaultRole.name())
 *                 .securityGroupId(defaultSecurityGroup.id())
 *                 .vpcId(defaultNetwork.id())
 *                 .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *                 .keyPairName(defaultEcsKeyPair.id())
 *                 .build())
 *             .tags(Map.of(&#34;created&#34;, &#34;tf&#34;))
 *             .nodeGroups(            
 *                 ClusterNodeGroupArgs.builder()
 *                     .nodeGroupType(&#34;MASTER&#34;)
 *                     .nodeGroupName(&#34;emr-master&#34;)
 *                     .paymentType(&#34;PayAsYouGo&#34;)
 *                     .vswitchIds(defaultSwitch.id())
 *                     .withPublicIp(false)
 *                     .instanceTypes(&#34;ecs.g7.xlarge&#34;)
 *                     .nodeCount(1)
 *                     .systemDisk(ClusterNodeGroupSystemDiskArgs.builder()
 *                         .category(&#34;cloud_essd&#34;)
 *                         .size(80)
 *                         .count(1)
 *                         .build())
 *                     .dataDisks(ClusterNodeGroupDataDiskArgs.builder()
 *                         .category(&#34;cloud_essd&#34;)
 *                         .size(80)
 *                         .count(3)
 *                         .build())
 *                     .build(),
 *                 ClusterNodeGroupArgs.builder()
 *                     .nodeGroupType(&#34;CORE&#34;)
 *                     .nodeGroupName(&#34;emr-core&#34;)
 *                     .paymentType(&#34;PayAsYouGo&#34;)
 *                     .vswitchIds(defaultSwitch.id())
 *                     .withPublicIp(false)
 *                     .instanceTypes(&#34;ecs.g7.xlarge&#34;)
 *                     .nodeCount(3)
 *                     .systemDisk(ClusterNodeGroupSystemDiskArgs.builder()
 *                         .category(&#34;cloud_essd&#34;)
 *                         .size(80)
 *                         .count(1)
 *                         .build())
 *                     .dataDisks(ClusterNodeGroupDataDiskArgs.builder()
 *                         .category(&#34;cloud_essd&#34;)
 *                         .size(80)
 *                         .count(3)
 *                         .build())
 *                     .build())
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.ids()[0]))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Aliclioud E-MapReduce cluster can be imported using the id e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:emrv2/cluster:Cluster default &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:emrv2/cluster:Cluster")
public class Cluster extends com.pulumi.resources.CustomResource {
    /**
     * The application configurations of EMR cluster. See `application_configs` below.
     * 
     */
    @Export(name="applicationConfigs", type=List.class, parameters={ClusterApplicationConfig.class})
    private Output</* @Nullable */ List<ClusterApplicationConfig>> applicationConfigs;

    /**
     * @return The application configurations of EMR cluster. See `application_configs` below.
     * 
     */
    public Output<Optional<List<ClusterApplicationConfig>>> applicationConfigs() {
        return Codegen.optional(this.applicationConfigs);
    }
    /**
     * The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
     * 
     */
    @Export(name="applications", type=List.class, parameters={String.class})
    private Output<List<String>> applications;

    /**
     * @return The applications of EMR cluster to be installed, e.g. HADOOP-COMMON, HDFS, YARN, HIVE, SPARK2, SPARK3, ZOOKEEPER etc. You can find all valid applications in emr web console.
     * 
     */
    public Output<List<String>> applications() {
        return this.applications;
    }
    /**
     * The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster. See `bootstrap_scripts` below.
     * 
     */
    @Export(name="bootstrapScripts", type=List.class, parameters={ClusterBootstrapScript.class})
    private Output</* @Nullable */ List<ClusterBootstrapScript>> bootstrapScripts;

    /**
     * @return The bootstrap scripts to be effected when creating emr-cluster or resize emr-cluster. See `bootstrap_scripts` below.
     * 
     */
    public Output<Optional<List<ClusterBootstrapScript>>> bootstrapScripts() {
        return Codegen.optional(this.bootstrapScripts);
    }
    /**
     * The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, &#34;-&#34;, &#34;_&#34;.
     * 
     */
    @Export(name="clusterName", type=String.class, parameters={})
    private Output<String> clusterName;

    /**
     * @return The name of emr cluster. The name length must be less than 64. Supported characters: chinese character, english character, number, &#34;-&#34;, &#34;_&#34;.
     * 
     */
    public Output<String> clusterName() {
        return this.clusterName;
    }
    /**
     * EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
     * 
     */
    @Export(name="clusterType", type=String.class, parameters={})
    private Output<String> clusterType;

    /**
     * @return EMR Cluster Type, e.g. DATALAKE, OLAP, DATAFLOW, DATASERVING, CUSTOM etc. You can find all valid EMR cluster type in emr web console.
     * 
     */
    public Output<String> clusterType() {
        return this.clusterType;
    }
    /**
     * The deploy mode of EMR cluster. Supported value: NORMAL or HA.
     * 
     */
    @Export(name="deployMode", type=String.class, parameters={})
    private Output<String> deployMode;

    /**
     * @return The deploy mode of EMR cluster. Supported value: NORMAL or HA.
     * 
     */
    public Output<String> deployMode() {
        return this.deployMode;
    }
    /**
     * The node attributes of ecs instances which the emr-cluster belongs. See `node_attributes` below.
     * 
     */
    @Export(name="nodeAttributes", type=List.class, parameters={ClusterNodeAttribute.class})
    private Output<List<ClusterNodeAttribute>> nodeAttributes;

    /**
     * @return The node attributes of ecs instances which the emr-cluster belongs. See `node_attributes` below.
     * 
     */
    public Output<List<ClusterNodeAttribute>> nodeAttributes() {
        return this.nodeAttributes;
    }
    /**
     * Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example). See `node_groups` below.
     * 
     */
    @Export(name="nodeGroups", type=List.class, parameters={ClusterNodeGroup.class})
    private Output<List<ClusterNodeGroup>> nodeGroups;

    /**
     * @return Groups of node, You can specify MASTER as a group, CORE as a group (just like the above example). See `node_groups` below.
     * 
     */
    public Output<List<ClusterNodeGroup>> nodeGroups() {
        return this.nodeGroups;
    }
    /**
     * Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return Payment Type for this cluster. Supported value: PayAsYouGo or Subscription.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
     * 
     */
    @Export(name="releaseVersion", type=String.class, parameters={})
    private Output<String> releaseVersion;

    /**
     * @return EMR Version, e.g. EMR-5.10.0. You can find the all valid EMR Version in emr web console.
     * 
     */
    public Output<String> releaseVersion() {
        return this.releaseVersion;
    }
    /**
     * The Id of resource group which the emr-cluster belongs.
     * 
     */
    @Export(name="resourceGroupId", type=String.class, parameters={})
    private Output<String> resourceGroupId;

    /**
     * @return The Id of resource group which the emr-cluster belongs.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
     * 
     */
    @Export(name="securityMode", type=String.class, parameters={})
    private Output<String> securityMode;

    /**
     * @return The security mode of EMR cluster. Supported value: NORMAL or KERBEROS.
     * 
     */
    public Output<String> securityMode() {
        return this.securityMode;
    }
    /**
     * The detail configuration of subscription payment type. See `subscription_config` below.
     * 
     */
    @Export(name="subscriptionConfig", type=ClusterSubscriptionConfig.class, parameters={})
    private Output</* @Nullable */ ClusterSubscriptionConfig> subscriptionConfig;

    /**
     * @return The detail configuration of subscription payment type. See `subscription_config` below.
     * 
     */
    public Output<Optional<ClusterSubscriptionConfig>> subscriptionConfig() {
        return Codegen.optional(this.subscriptionConfig);
    }
    /**
     * A mapping of tags to assign to the resource.
     * 
     */
    @Export(name="tags", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> tags;

    /**
     * @return A mapping of tags to assign to the resource.
     * 
     */
    public Output<Optional<Map<String,Object>>> tags() {
        return Codegen.optional(this.tags);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Cluster(String name) {
        this(name, ClusterArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Cluster(String name, ClusterArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Cluster(String name, ClusterArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:emrv2/cluster:Cluster", name, args == null ? ClusterArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Cluster(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:emrv2/cluster:Cluster", name, state, makeResourceOptions(options, id));
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
    public static Cluster get(String name, Output<String> id, @Nullable ClusterState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Cluster(name, id, state, options);
    }
}
