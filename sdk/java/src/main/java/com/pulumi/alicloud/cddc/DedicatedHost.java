// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cddc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cddc.DedicatedHostArgs;
import com.pulumi.alicloud.cddc.inputs.DedicatedHostState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ApsaraDB for MyBase Dedicated Host resource.
 * 
 * For information about ApsaraDB for MyBase Dedicated Host and how to use it, see [What is Dedicated Host](https://www.alibabacloud.com/help/doc-detail/210864.html).
 * 
 * &gt; **NOTE:** Available in v1.147.0+.
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
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.cloudconnect.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.cddc.CddcFunctions;
 * import com.pulumi.alicloud.adb.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.cddc.inputs.GetHostEcsLevelInfosArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.cddc.DedicatedHostGroup;
 * import com.pulumi.alicloud.cddc.DedicatedHostGroupArgs;
 * import com.pulumi.alicloud.cddc.DedicatedHost;
 * import com.pulumi.alicloud.cddc.DedicatedHostArgs;
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
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;default-NODELETING&#34;)
 *             .build());
 * 
 *         final var defaultZones = CddcFunctions.getZones();
 * 
 *         final var defaultHostEcsLevelInfos = CddcFunctions.getHostEcsLevelInfos(GetHostEcsLevelInfosArgs.builder()
 *             .dbType(&#34;mysql&#34;)
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .storageType(&#34;cloud_essd&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .build());
 * 
 *         var defaultDedicatedHostGroup = new DedicatedHostGroup(&#34;defaultDedicatedHostGroup&#34;, DedicatedHostGroupArgs.builder()        
 *             .engine(&#34;MySQL&#34;)
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .cpuAllocationRatio(101)
 *             .memAllocationRatio(50)
 *             .diskAllocationRatio(200)
 *             .allocationPolicy(&#34;Evenly&#34;)
 *             .hostReplacePolicy(&#34;Manual&#34;)
 *             .dedicatedHostGroupDesc(&#34;example_value&#34;)
 *             .build());
 * 
 *         var defaultDedicatedHost = new DedicatedHost(&#34;defaultDedicatedHost&#34;, DedicatedHostArgs.builder()        
 *             .hostName(&#34;example_value&#34;)
 *             .dedicatedHostGroupId(defaultDedicatedHostGroup.id())
 *             .hostClass(defaultHostEcsLevelInfos.applyValue(getHostEcsLevelInfosResult -&gt; getHostEcsLevelInfosResult.infos()[0].resClassCode()))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.ids()[0]))
 *             .vswitchId(defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .paymentType(&#34;Subscription&#34;)
 *             .tags(Map.ofEntries(
 *                 Map.entry(&#34;Created&#34;, &#34;TF&#34;),
 *                 Map.entry(&#34;For&#34;, &#34;CDDC_DEDICATED&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ApsaraDB for MyBase Dedicated Host can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cddc/dedicatedHost:DedicatedHost example &lt;dedicated_host_group_id&gt;:&lt;dedicated_host_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cddc/dedicatedHost:DedicatedHost")
public class DedicatedHost extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
     * 
     */
    @Export(name="allocationStatus", type=String.class, parameters={})
    private Output<String> allocationStatus;

    /**
     * @return Specifies whether instances can be created on the host. Valid values: `Allocatable` or `Suspended`. `Allocatable`: Instances can be created on the host. `Suspended`: Instances cannot be created on the host.
     * 
     */
    public Output<String> allocationStatus() {
        return this.allocationStatus;
    }
    /**
     * Specifies whether to enable the auto-renewal feature.
     * 
     */
    @Export(name="autoRenew", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> autoRenew;

    /**
     * @return Specifies whether to enable the auto-renewal feature.
     * 
     */
    public Output<Optional<Boolean>> autoRenew() {
        return Codegen.optional(this.autoRenew);
    }
    /**
     * The ID of the dedicated cluster.
     * 
     */
    @Export(name="dedicatedHostGroupId", type=String.class, parameters={})
    private Output<String> dedicatedHostGroupId;

    /**
     * @return The ID of the dedicated cluster.
     * 
     */
    public Output<String> dedicatedHostGroupId() {
        return this.dedicatedHostGroupId;
    }
    /**
     * The ID of the host.
     * 
     */
    @Export(name="dedicatedHostId", type=String.class, parameters={})
    private Output<String> dedicatedHostId;

    /**
     * @return The ID of the host.
     * 
     */
    public Output<String> dedicatedHostId() {
        return this.dedicatedHostId;
    }
    /**
     * The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
     * 
     */
    @Export(name="hostClass", type=String.class, parameters={})
    private Output<String> hostClass;

    /**
     * @return The instance type of the host. For more information about the supported instance types of hosts, see [Host specification details](https://www.alibabacloud.com/help/doc-detail/206343.htm).
     * 
     */
    public Output<String> hostClass() {
        return this.hostClass;
    }
    /**
     * The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    @Export(name="hostName", type=String.class, parameters={})
    private Output<String> hostName;

    /**
     * @return The name of the host. The name must be `1` to `64` characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
     * 
     */
    public Output<String> hostName() {
        return this.hostName;
    }
    /**
     * Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
     * 
     */
    @Export(name="imageCategory", type=String.class, parameters={})
    private Output</* @Nullable */ String> imageCategory;

    /**
     * @return Host Image Category. Valid values: `WindowsWithMssqlEntAlwaysonLicense`, `WindowsWithMssqlStdLicense`, `WindowsWithMssqlEntLicense`, `WindowsWithMssqlWebLicense`, `AliLinux`.
     * 
     */
    public Output<Optional<String>> imageCategory() {
        return Codegen.optional(this.imageCategory);
    }
    /**
     * Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
     * 
     */
    @Export(name="osPassword", type=String.class, parameters={})
    private Output</* @Nullable */ String> osPassword;

    /**
     * @return Host password. **NOTE:** The creation of a host password is supported only when the database type is `Tair-PMem`.
     * 
     */
    public Output<Optional<String>> osPassword() {
        return Codegen.optional(this.osPassword);
    }
    /**
     * The payment type of the resource. Valid values: `Subscription`.
     * 
     */
    @Export(name="paymentType", type=String.class, parameters={})
    private Output<String> paymentType;

    /**
     * @return The payment type of the resource. Valid values: `Subscription`.
     * 
     */
    public Output<String> paymentType() {
        return this.paymentType;
    }
    /**
     * The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
     * 
     */
    @Export(name="period", type=String.class, parameters={})
    private Output</* @Nullable */ String> period;

    /**
     * @return The unit of the subscription duration. Valid values: `Year`, `Month`, `Week`.
     * 
     */
    public Output<Optional<String>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return The state of the host. Valid values: `0:` The host is being created. `1`: The host is running. `2`: The host is faulty. `3`: The host is ready for deactivation. `4`: The host is being maintained. `5`: The host is deactivated. `6`: The host is restarting. `7`: The host is locked.
     * 
     */
    public Output<String> status() {
        return this.status;
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
     * The subscription duration of the host. Valid values:
     * * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
     * * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
     * * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
     * 
     */
    @Export(name="usedTime", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> usedTime;

    /**
     * @return The subscription duration of the host. Valid values:
     * * If the Period parameter is set to `Year`, the value of the UsedTime parameter ranges from `1` to `5`.
     * * If the Period parameter is set to `Month`, the value of the UsedTime parameter ranges from `1` to `9`.
     * * If the Period parameter is set to `Week`, the value of the UsedTime parameter ranges from `1`, `2` and `3`.
     * 
     */
    public Output<Optional<Integer>> usedTime() {
        return Codegen.optional(this.usedTime);
    }
    /**
     * The ID of the vSwitch to which the host is connected.
     * 
     */
    @Export(name="vswitchId", type=String.class, parameters={})
    private Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch to which the host is connected.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The ID of the zone.
     * 
     */
    @Export(name="zoneId", type=String.class, parameters={})
    private Output<String> zoneId;

    /**
     * @return The ID of the zone.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public DedicatedHost(String name) {
        this(name, DedicatedHostArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public DedicatedHost(String name, DedicatedHostArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public DedicatedHost(String name, DedicatedHostArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHost:DedicatedHost", name, args == null ? DedicatedHostArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private DedicatedHost(String name, Output<String> id, @Nullable DedicatedHostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cddc/dedicatedHost:DedicatedHost", name, state, makeResourceOptions(options, id));
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
    public static DedicatedHost get(String name, Output<String> id, @Nullable DedicatedHostState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new DedicatedHost(name, id, state, options);
    }
}
