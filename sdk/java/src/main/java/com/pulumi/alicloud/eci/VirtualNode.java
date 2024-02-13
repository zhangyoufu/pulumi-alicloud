// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.eci;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.eci.VirtualNodeArgs;
import com.pulumi.alicloud.eci.inputs.VirtualNodeState;
import com.pulumi.alicloud.eci.outputs.VirtualNodeTaint;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a ECI Virtual Node resource.
 * 
 * For information about ECI Virtual Node and how to use it, see [What is Virtual Node](https://www.alibabacloud.com/help/en/doc-detail/89129.html).
 * 
 * &gt; **NOTE:** Available since v1.145.0.
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
 * import com.pulumi.alicloud.eci.EciFunctions;
 * import com.pulumi.alicloud.eci.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.ecs.EipAddress;
 * import com.pulumi.alicloud.ecs.EipAddressArgs;
 * import com.pulumi.alicloud.resourcemanager.ResourcemanagerFunctions;
 * import com.pulumi.alicloud.resourcemanager.inputs.GetResourceGroupsArgs;
 * import com.pulumi.alicloud.eci.VirtualNode;
 * import com.pulumi.alicloud.eci.VirtualNodeArgs;
 * import com.pulumi.alicloud.eci.inputs.VirtualNodeTaintArgs;
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
 *         final var defaultZones = EciFunctions.getZones();
 * 
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .vpcName(name)
 *             .cidrBlock(&#34;10.0.0.0/8&#34;)
 *             .build());
 * 
 *         var defaultSwitch = new Switch(&#34;defaultSwitch&#34;, SwitchArgs.builder()        
 *             .vswitchName(name)
 *             .cidrBlock(&#34;10.1.0.0/16&#34;)
 *             .vpcId(defaultNetwork.id())
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].zoneIds()[0]))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetwork.id())
 *             .build());
 * 
 *         var defaultEipAddress = new EipAddress(&#34;defaultEipAddress&#34;, EipAddressArgs.builder()        
 *             .isp(&#34;BGP&#34;)
 *             .addressName(name)
 *             .netmode(&#34;public&#34;)
 *             .bandwidth(&#34;1&#34;)
 *             .securityProtectionTypes(&#34;AntiDDoS_Enhanced&#34;)
 *             .paymentType(&#34;PayAsYouGo&#34;)
 *             .build());
 * 
 *         final var defaultResourceGroups = ResourcemanagerFunctions.getResourceGroups();
 * 
 *         var defaultVirtualNode = new VirtualNode(&#34;defaultVirtualNode&#34;, VirtualNodeArgs.builder()        
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .virtualNodeName(name)
 *             .vswitchId(defaultSwitch.id())
 *             .enablePublicNetwork(false)
 *             .eipInstanceId(defaultEipAddress.id())
 *             .resourceGroupId(defaultResourceGroups.applyValue(getResourceGroupsResult -&gt; getResourceGroupsResult.groups()[0].id()))
 *             .kubeConfig(&#34;kube_config&#34;)
 *             .tags(Map.of(&#34;Created&#34;, &#34;TF&#34;))
 *             .taints(VirtualNodeTaintArgs.builder()
 *                 .effect(&#34;NoSchedule&#34;)
 *                 .key(&#34;TF&#34;)
 *                 .value(&#34;example&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * ECI Virtual Node can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:eci/virtualNode:VirtualNode example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:eci/virtualNode:VirtualNode")
public class VirtualNode extends com.pulumi.resources.CustomResource {
    /**
     * The Id of eip.
     * 
     */
    @Export(name="eipInstanceId", refs={String.class}, tree="[0]")
    private Output<String> eipInstanceId;

    /**
     * @return The Id of eip.
     * 
     */
    public Output<String> eipInstanceId() {
        return this.eipInstanceId;
    }
    /**
     * Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
     * 
     */
    @Export(name="enablePublicNetwork", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enablePublicNetwork;

    /**
     * @return Whether to enable public network. **NOTE:** If `eip_instance_id` is not configured and `enable_public_network` is true, the system will create an elastic public network IP.
     * 
     */
    public Output<Optional<Boolean>> enablePublicNetwork() {
        return Codegen.optional(this.enablePublicNetwork);
    }
    /**
     * The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
     * 
     */
    @Export(name="kubeConfig", refs={String.class}, tree="[0]")
    private Output<String> kubeConfig;

    /**
     * @return The kube config for the k8s cluster. It needs to be connected after Base64 encoding.
     * 
     */
    public Output<String> kubeConfig() {
        return this.kubeConfig;
    }
    /**
     * The resource group ID.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> resourceGroupId;

    /**
     * @return The resource group ID.
     * 
     */
    public Output<Optional<String>> resourceGroupId() {
        return Codegen.optional(this.resourceGroupId);
    }
    /**
     * The security group ID.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return The security group ID.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }
    /**
     * The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The Status of the virtual node. Valid values: `Cleaned`, `Failed`, `Pending`, `Ready`.
     * 
     */
    public Output<String> status() {
        return this.status;
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
     * The taint. See `taints` below.
     * 
     */
    @Export(name="taints", refs={List.class,VirtualNodeTaint.class}, tree="[0,1]")
    private Output</* @Nullable */ List<VirtualNodeTaint>> taints;

    /**
     * @return The taint. See `taints` below.
     * 
     */
    public Output<Optional<List<VirtualNodeTaint>>> taints() {
        return Codegen.optional(this.taints);
    }
    /**
     * The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
     * 
     */
    @Export(name="virtualNodeName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> virtualNodeName;

    /**
     * @return The name of the virtual node. The length of the name is limited to `2` to `128` characters. It can contain uppercase and lowercase letters, Chinese characters, numbers, half-width colon (:), underscores (_), or hyphens (-), and must start with letters.
     * 
     */
    public Output<Optional<String>> virtualNodeName() {
        return Codegen.optional(this.virtualNodeName);
    }
    /**
     * The vswitch id.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The vswitch id.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }
    /**
     * The Zone.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The Zone.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VirtualNode(String name) {
        this(name, VirtualNodeArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VirtualNode(String name, VirtualNodeArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VirtualNode(String name, VirtualNodeArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/virtualNode:VirtualNode", name, args == null ? VirtualNodeArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VirtualNode(String name, Output<String> id, @Nullable VirtualNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:eci/virtualNode:VirtualNode", name, state, makeResourceOptions(options, id));
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
    public static VirtualNode get(String name, Output<String> id, @Nullable VirtualNodeState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VirtualNode(name, id, state, options);
    }
}
