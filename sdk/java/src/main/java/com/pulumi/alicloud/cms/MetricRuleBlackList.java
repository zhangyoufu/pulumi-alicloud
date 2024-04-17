// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cms;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cms.MetricRuleBlackListArgs;
import com.pulumi.alicloud.cms.inputs.MetricRuleBlackListState;
import com.pulumi.alicloud.cms.outputs.MetricRuleBlackListMetric;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Monitor Service Metric Rule Black List resource.
 * 
 * For information about Cloud Monitor Service Metric Rule Black List and how to use it, see [What is Metric Rule Black List](https://www.alibabacloud.com/help/en/cloudmonitor/latest/describemetricruleblacklist).
 * 
 * &gt; **NOTE:** Available since v1.194.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.ecs.EcsFunctions;
 * import com.pulumi.alicloud.ecs.inputs.GetInstanceTypesArgs;
 * import com.pulumi.alicloud.ecs.inputs.GetImagesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.Instance;
 * import com.pulumi.alicloud.ecs.InstanceArgs;
 * import com.pulumi.alicloud.cms.MetricRuleBlackList;
 * import com.pulumi.alicloud.cms.MetricRuleBlackListArgs;
 * import com.pulumi.alicloud.cms.inputs.MetricRuleBlackListMetricArgs;
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
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;Instance&#34;)
 *             .build());
 * 
 *         final var defaultGetInstanceTypes = EcsFunctions.getInstanceTypes(GetInstanceTypesArgs.builder()
 *             .availabilityZone(default_.zones()[0].id())
 *             .cpuCoreCount(1)
 *             .memorySize(2)
 *             .build());
 * 
 *         final var defaultGetImages = EcsFunctions.getImages(GetImagesArgs.builder()
 *             .nameRegex(&#34;^ubuntu_[0-9]+_[0-9]+_x64*&#34;)
 *             .owners(&#34;system&#34;)
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
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .name(name)
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .availabilityZone(default_.zones()[0].id())
 *             .instanceName(name)
 *             .imageId(defaultGetImages.applyValue(getImagesResult -&gt; getImagesResult.images()[0].id()))
 *             .instanceType(defaultGetInstanceTypes.applyValue(getInstanceTypesResult -&gt; getInstanceTypesResult.instanceTypes()[0].id()))
 *             .securityGroups(defaultSecurityGroup.id())
 *             .vswitchId(defaultSwitch.id())
 *             .build());
 * 
 *         var defaultMetricRuleBlackList = new MetricRuleBlackList(&#34;defaultMetricRuleBlackList&#34;, MetricRuleBlackListArgs.builder()        
 *             .instances(defaultInstance.id().applyValue(id -&gt; String.format(&#34;{{\&#34;instancceId\&#34;:\&#34;%s\&#34;}}&#34;, id)))
 *             .metrics(MetricRuleBlackListMetricArgs.builder()
 *                 .metricName(&#34;disk_utilization&#34;)
 *                 .build())
 *             .category(&#34;ecs&#34;)
 *             .enableEndTime(1799443209000)
 *             .namespace(&#34;acs_ecs_dashboard&#34;)
 *             .enableStartTime(1689243209000)
 *             .metricRuleBlackListName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Cloud Monitor Service Metric Rule Black List can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cms/metricRuleBlackList:MetricRuleBlackList example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cms/metricRuleBlackList:MetricRuleBlackList")
public class MetricRuleBlackList extends com.pulumi.resources.CustomResource {
    /**
     * Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     * 
     */
    @Export(name="category", refs={String.class}, tree="[0]")
    private Output<String> category;

    /**
     * @return Cloud service classification. For example, Redis includes kvstore_standard, kvstore_sharding, and kvstore_splitrw.
     * 
     */
    public Output<String> category() {
        return this.category;
    }
    /**
     * The timestamp for creating an alert blacklist policy.Unit: milliseconds.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The timestamp for creating an alert blacklist policy.Unit: milliseconds.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The effective time range of the alert blacklist policy.
     * 
     */
    @Export(name="effectiveTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> effectiveTime;

    /**
     * @return The effective time range of the alert blacklist policy.
     * 
     */
    public Output<Optional<String>> effectiveTime() {
        return Codegen.optional(this.effectiveTime);
    }
    /**
     * The start timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    @Export(name="enableEndTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enableEndTime;

    /**
     * @return The start timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    public Output<Optional<String>> enableEndTime() {
        return Codegen.optional(this.enableEndTime);
    }
    /**
     * The end timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    @Export(name="enableStartTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> enableStartTime;

    /**
     * @return The end timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    public Output<Optional<String>> enableStartTime() {
        return Codegen.optional(this.enableStartTime);
    }
    /**
     * The list of instances of cloud services specified in the alert blacklist policy.
     * 
     */
    @Export(name="instances", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> instances;

    /**
     * @return The list of instances of cloud services specified in the alert blacklist policy.
     * 
     */
    public Output<List<String>> instances() {
        return this.instances;
    }
    /**
     * The status of the alert blacklist policy. Value:-true: enabled.-false: disabled.
     * 
     */
    @Export(name="isEnable", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> isEnable;

    /**
     * @return The status of the alert blacklist policy. Value:-true: enabled.-false: disabled.
     * 
     */
    public Output<Boolean> isEnable() {
        return this.isEnable;
    }
    /**
     * The ID of the blacklist policy.
     * 
     */
    @Export(name="metricRuleBlackListId", refs={String.class}, tree="[0]")
    private Output<String> metricRuleBlackListId;

    /**
     * @return The ID of the blacklist policy.
     * 
     */
    public Output<String> metricRuleBlackListId() {
        return this.metricRuleBlackListId;
    }
    /**
     * The name of the alert blacklist policy.
     * 
     */
    @Export(name="metricRuleBlackListName", refs={String.class}, tree="[0]")
    private Output<String> metricRuleBlackListName;

    /**
     * @return The name of the alert blacklist policy.
     * 
     */
    public Output<String> metricRuleBlackListName() {
        return this.metricRuleBlackListName;
    }
    /**
     * Monitoring metrics in the instance. See `metrics` below.
     * 
     */
    @Export(name="metrics", refs={List.class,MetricRuleBlackListMetric.class}, tree="[0,1]")
    private Output</* @Nullable */ List<MetricRuleBlackListMetric>> metrics;

    /**
     * @return Monitoring metrics in the instance. See `metrics` below.
     * 
     */
    public Output<Optional<List<MetricRuleBlackListMetric>>> metrics() {
        return Codegen.optional(this.metrics);
    }
    /**
     * The data namespace of the cloud service.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output<String> namespace;

    /**
     * @return The data namespace of the cloud service.
     * 
     */
    public Output<String> namespace() {
        return this.namespace;
    }
    /**
     * The effective range of the alert blacklist policy. Value:-USER: The alert blacklist policy only takes effect in the current Alibaba cloud account.-GROUP: The alert blacklist policy takes effect in the specified application GROUP.
     * 
     */
    @Export(name="scopeType", refs={String.class}, tree="[0]")
    private Output<String> scopeType;

    /**
     * @return The effective range of the alert blacklist policy. Value:-USER: The alert blacklist policy only takes effect in the current Alibaba cloud account.-GROUP: The alert blacklist policy takes effect in the specified application GROUP.
     * 
     */
    public Output<String> scopeType() {
        return this.scopeType;
    }
    /**
     * Application Group ID list. The format is JSON Array.&gt; This parameter is displayed only when &#39;ScopeType&#39; is &#39;GROUP.
     * 
     */
    @Export(name="scopeValues", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> scopeValues;

    /**
     * @return Application Group ID list. The format is JSON Array.&gt; This parameter is displayed only when &#39;ScopeType&#39; is &#39;GROUP.
     * 
     */
    public Output<Optional<List<String>>> scopeValues() {
        return Codegen.optional(this.scopeValues);
    }
    /**
     * Modify the timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    @Export(name="updateTime", refs={String.class}, tree="[0]")
    private Output<String> updateTime;

    /**
     * @return Modify the timestamp of the alert blacklist policy.Unit: milliseconds.
     * 
     */
    public Output<String> updateTime() {
        return this.updateTime;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MetricRuleBlackList(String name) {
        this(name, MetricRuleBlackListArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MetricRuleBlackList(String name, MetricRuleBlackListArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MetricRuleBlackList(String name, MetricRuleBlackListArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/metricRuleBlackList:MetricRuleBlackList", name, args == null ? MetricRuleBlackListArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MetricRuleBlackList(String name, Output<String> id, @Nullable MetricRuleBlackListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cms/metricRuleBlackList:MetricRuleBlackList", name, state, makeResourceOptions(options, id));
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
    public static MetricRuleBlackList get(String name, Output<String> id, @Nullable MetricRuleBlackListState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MetricRuleBlackList(name, id, state, options);
    }
}
