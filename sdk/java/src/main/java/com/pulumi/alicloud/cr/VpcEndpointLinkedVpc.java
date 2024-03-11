// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cr;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cr.VpcEndpointLinkedVpcArgs;
import com.pulumi.alicloud.cr.inputs.VpcEndpointLinkedVpcState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a CR Vpc Endpoint Linked Vpc resource.
 * 
 * For information about CR Vpc Endpoint Linked Vpc and how to use it, see [What is Vpc Endpoint Linked Vpc](https://www.alibabacloud.com/help/en/acr/developer-reference/api-cr-2018-12-01-createinstancevpcendpointlinkedvpc).
 * 
 * &gt; **NOTE:** Available since v1.199.0.
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
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.Switch;
 * import com.pulumi.alicloud.vpc.SwitchArgs;
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstance;
 * import com.pulumi.alicloud.cr.RegistryEnterpriseInstanceArgs;
 * import com.pulumi.alicloud.cr.VpcEndpointLinkedVpc;
 * import com.pulumi.alicloud.cr.VpcEndpointLinkedVpcArgs;
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
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
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
 *         var defaultRegistryEnterpriseInstance = new RegistryEnterpriseInstance(&#34;defaultRegistryEnterpriseInstance&#34;, RegistryEnterpriseInstanceArgs.builder()        
 *             .paymentType(&#34;Subscription&#34;)
 *             .period(1)
 *             .renewPeriod(0)
 *             .renewalStatus(&#34;ManualRenewal&#34;)
 *             .instanceType(&#34;Advanced&#34;)
 *             .instanceName(name)
 *             .build());
 * 
 *         var defaultVpcEndpointLinkedVpc = new VpcEndpointLinkedVpc(&#34;defaultVpcEndpointLinkedVpc&#34;, VpcEndpointLinkedVpcArgs.builder()        
 *             .instanceId(defaultRegistryEnterpriseInstance.id())
 *             .vpcId(defaultNetwork.id())
 *             .vswitchId(defaultSwitch.id())
 *             .moduleName(&#34;Registry&#34;)
 *             .enableCreateDnsRecordInPvzt(true)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * CR Vpc Endpoint Linked Vpc can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc example &lt;instance_id&gt;:&lt;vpc_id&gt;:&lt;vswitch_id&gt;:&lt;module_name&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc")
public class VpcEndpointLinkedVpc extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
     * 
     */
    @Export(name="enableCreateDnsRecordInPvzt", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enableCreateDnsRecordInPvzt;

    /**
     * @return Specifies whether to automatically create an Alibaba Cloud DNS PrivateZone record. Valid Values:
     * 
     */
    public Output<Optional<Boolean>> enableCreateDnsRecordInPvzt() {
        return Codegen.optional(this.enableCreateDnsRecordInPvzt);
    }
    /**
     * The ID of the instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The name of the module that you want to access. Valid Values:
     * 
     */
    @Export(name="moduleName", refs={String.class}, tree="[0]")
    private Output<String> moduleName;

    /**
     * @return The name of the module that you want to access. Valid Values:
     * 
     */
    public Output<String> moduleName() {
        return this.moduleName;
    }
    /**
     * The status of the Vpc Endpoint Linked Vpc.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the Vpc Endpoint Linked Vpc.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * The ID of the VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }
    /**
     * The ID of the vSwitch.
     * 
     */
    @Export(name="vswitchId", refs={String.class}, tree="[0]")
    private Output<String> vswitchId;

    /**
     * @return The ID of the vSwitch.
     * 
     */
    public Output<String> vswitchId() {
        return this.vswitchId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public VpcEndpointLinkedVpc(String name) {
        this(name, VpcEndpointLinkedVpcArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public VpcEndpointLinkedVpc(String name, VpcEndpointLinkedVpcArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public VpcEndpointLinkedVpc(String name, VpcEndpointLinkedVpcArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc", name, args == null ? VpcEndpointLinkedVpcArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private VpcEndpointLinkedVpc(String name, Output<String> id, @Nullable VpcEndpointLinkedVpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cr/vpcEndpointLinkedVpc:VpcEndpointLinkedVpc", name, state, makeResourceOptions(options, id));
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
    public static VpcEndpointLinkedVpc get(String name, Output<String> id, @Nullable VpcEndpointLinkedVpcState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new VpcEndpointLinkedVpc(name, id, state, options);
    }
}
