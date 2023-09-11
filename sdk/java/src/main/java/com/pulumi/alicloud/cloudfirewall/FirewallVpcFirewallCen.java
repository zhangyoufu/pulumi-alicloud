// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudfirewall;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallCenArgs;
import com.pulumi.alicloud.cloudfirewall.inputs.FirewallVpcFirewallCenState;
import com.pulumi.alicloud.cloudfirewall.outputs.FirewallVpcFirewallCenLocalVpc;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Firewall Vpc Firewall Cen resource.
 * 
 * For information about Cloud Firewall Vpc Firewall Cen and how to use it, see [What is Vpc Firewall Cen](https://www.alibabacloud.com/help/en/cloud-firewall/latest/createvpcfirewallcenconfigure).
 * 
 * &gt; **NOTE:** Available since v1.194.0.
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
 * import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallCen;
 * import com.pulumi.alicloud.cloudfirewall.FirewallVpcFirewallCenArgs;
 * import com.pulumi.alicloud.cloudfirewall.inputs.FirewallVpcFirewallCenLocalVpcArgs;
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
 *         var default_ = new FirewallVpcFirewallCen(&#34;default&#34;, FirewallVpcFirewallCenArgs.builder()        
 *             .cenId(&#34;cen-cjok7uyb5w2b27573v&#34;)
 *             .localVpc(FirewallVpcFirewallCenLocalVpcArgs.builder()
 *                 .networkInstanceId(&#34;vpc-a2d4wzzfuumzuq6uog5w4&#34;)
 *                 .build())
 *             .memberUid(&#34;1415189284827022&#34;)
 *             .status(&#34;open&#34;)
 *             .vpcFirewallName(&#34;tf-vpc-firewall-name&#34;)
 *             .vpcRegion(&#34;ap-south-1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Cloud Firewall Vpc Firewall Cen can be imported using the id, e.g.
 * 
 * ```sh
 *  $ pulumi import alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen")
public class FirewallVpcFirewallCen extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the CEN instance.
     * 
     */
    @Export(name="cenId", type=String.class, parameters={})
    private Output<String> cenId;

    /**
     * @return The ID of the CEN instance.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }
    /**
     * Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
     * 
     */
    @Export(name="connectType", type=String.class, parameters={})
    private Output<String> connectType;

    /**
     * @return Intercommunication type, value: expressconnect: Express Channel cen: Cloud Enterprise Network
     * 
     */
    public Output<String> connectType() {
        return this.connectType;
    }
    /**
     * The language type of the requested and received messages. Valid values:
     * 
     */
    @Export(name="lang", type=String.class, parameters={})
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language type of the requested and received messages. Valid values:
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The details of the VPC. See `local_vpc` below.
     * 
     */
    @Export(name="localVpc", type=FirewallVpcFirewallCenLocalVpc.class, parameters={})
    private Output<FirewallVpcFirewallCenLocalVpc> localVpc;

    /**
     * @return The details of the VPC. See `local_vpc` below.
     * 
     */
    public Output<FirewallVpcFirewallCenLocalVpc> localVpc() {
        return this.localVpc;
    }
    /**
     * The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
     * 
     */
    @Export(name="memberUid", type=String.class, parameters={})
    private Output</* @Nullable */ String> memberUid;

    /**
     * @return The UID of the member account (other Alibaba Cloud account) of the current Alibaba cloud account.
     * 
     */
    public Output<Optional<String>> memberUid() {
        return Codegen.optional(this.memberUid);
    }
    /**
     * Firewall switch status.
     * 
     */
    @Export(name="status", type=String.class, parameters={})
    private Output<String> status;

    /**
     * @return Firewall switch status.
     * 
     */
    public Output<String> status() {
        return this.status;
    }
    /**
     * VPC firewall ID
     * 
     */
    @Export(name="vpcFirewallId", type=String.class, parameters={})
    private Output<String> vpcFirewallId;

    /**
     * @return VPC firewall ID
     * 
     */
    public Output<String> vpcFirewallId() {
        return this.vpcFirewallId;
    }
    /**
     * The name of the VPC firewall instance.
     * 
     */
    @Export(name="vpcFirewallName", type=String.class, parameters={})
    private Output<String> vpcFirewallName;

    /**
     * @return The name of the VPC firewall instance.
     * 
     */
    public Output<String> vpcFirewallName() {
        return this.vpcFirewallName;
    }
    /**
     * The ID of the region to which the VPC is created.
     * 
     */
    @Export(name="vpcRegion", type=String.class, parameters={})
    private Output<String> vpcRegion;

    /**
     * @return The ID of the region to which the VPC is created.
     * 
     */
    public Output<String> vpcRegion() {
        return this.vpcRegion;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public FirewallVpcFirewallCen(String name) {
        this(name, FirewallVpcFirewallCenArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public FirewallVpcFirewallCen(String name, FirewallVpcFirewallCenArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public FirewallVpcFirewallCen(String name, FirewallVpcFirewallCenArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen", name, args == null ? FirewallVpcFirewallCenArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private FirewallVpcFirewallCen(String name, Output<String> id, @Nullable FirewallVpcFirewallCenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudfirewall/firewallVpcFirewallCen:FirewallVpcFirewallCen", name, state, makeResourceOptions(options, id));
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
    public static FirewallVpcFirewallCen get(String name, Output<String> id, @Nullable FirewallVpcFirewallCenState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new FirewallVpcFirewallCen(name, id, state, options);
    }
}
