// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.vpc;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.vpc.PeerConnectionArgs;
import com.pulumi.alicloud.vpc.inputs.PeerConnectionState;
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
 * Provides a VPC Peer Connection resource.
 * 
 * For information about VPC Peer Connection and how to use it, see [What is Peer Connection](https://www.alibabacloud.com/help/en/virtual-private-cloud/latest/createvpcpeer).
 * 
 * &gt; **NOTE:** Available since v1.186.0.
 * 
 * ## Example Usage
 * 
 * Basic Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.vpc.PeerConnection;
 * import com.pulumi.alicloud.vpc.PeerConnectionArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         final var acceptingRegion = config.get("acceptingRegion").orElse("cn-beijing");
 *         var localVpc = new Network("localVpc", NetworkArgs.builder()        
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var acceptingVpc = new Network("acceptingVpc", NetworkArgs.builder()        
 *             .vpcName("terraform-example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var defaultPeerConnection = new PeerConnection("defaultPeerConnection", PeerConnectionArgs.builder()        
 *             .peerConnectionName("terraform-example")
 *             .vpcId(localVpc.id())
 *             .acceptingAliUid(default_.id())
 *             .acceptingRegionId(acceptingRegion)
 *             .acceptingVpcId(acceptingVpc.id())
 *             .description("terraform-example")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * VPC Peer Connection can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:vpc/peerConnection:PeerConnection example &lt;id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:vpc/peerConnection:PeerConnection")
public class PeerConnection extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
     * - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
     * - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
     * &gt; **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
     * 
     */
    @Export(name="acceptingAliUid", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> acceptingAliUid;

    /**
     * @return The ID of the Alibaba Cloud account (primary account) of the receiving end of the VPC peering connection to be created.
     * - Enter the ID of your Alibaba Cloud account to create a peer-to-peer connection to the VPC account.
     * - Enter the ID of another Alibaba Cloud account to create a cross-account VPC peer-to-peer connection.
     * &gt; **NOTE:**  If the recipient account is a RAM user (sub-account), enter the ID of the Alibaba Cloud account corresponding to the RAM user.
     * 
     */
    public Output<Optional<Integer>> acceptingAliUid() {
        return Codegen.optional(this.acceptingAliUid);
    }
    /**
     * The region ID of the recipient of the VPC peering connection to be created.
     * - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
     * - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
     * 
     */
    @Export(name="acceptingRegionId", refs={String.class}, tree="[0]")
    private Output<String> acceptingRegionId;

    /**
     * @return The region ID of the recipient of the VPC peering connection to be created.
     * - When creating a VPC peer-to-peer connection in the same region, enter the same region ID as the region ID of the initiator.
     * - When creating a cross-region VPC peer-to-peer connection, enter a region ID that is different from the region ID of the initiator.
     * 
     */
    public Output<String> acceptingRegionId() {
        return this.acceptingRegionId;
    }
    /**
     * The VPC ID of the receiving end of the VPC peer connection.
     * 
     */
    @Export(name="acceptingVpcId", refs={String.class}, tree="[0]")
    private Output<String> acceptingVpcId;

    /**
     * @return The VPC ID of the receiving end of the VPC peer connection.
     * 
     */
    public Output<String> acceptingVpcId() {
        return this.acceptingVpcId;
    }
    /**
     * The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
     * 
     */
    @Export(name="bandwidth", refs={Integer.class}, tree="[0]")
    private Output<Integer> bandwidth;

    /**
     * @return The bandwidth of the VPC peering connection to be modified. Unit: Mbps. The value range is an integer greater than 0.
     * 
     */
    public Output<Integer> bandwidth() {
        return this.bandwidth;
    }
    /**
     * The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
     * 
     */
    @Export(name="createTime", refs={String.class}, tree="[0]")
    private Output<String> createTime;

    /**
     * @return The creation time of the VPC peer connection. Use UTC time in the format `YYYY-MM-DDThh:mm:ssZ`.
     * 
     */
    public Output<String> createTime() {
        return this.createTime;
    }
    /**
     * The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     * 
     */
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the VPC peer connection to be created.It must be 2 to 256 characters in length and must start with a letter or Chinese, but cannot start with `http://` or `https://`.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * Whether to PreCheck only this request. Default value: `false`. Valid values:
     * 
     */
    @Export(name="dryRun", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> dryRun;

    /**
     * @return Whether to PreCheck only this request. Default value: `false`. Valid values:
     * 
     */
    public Output<Optional<Boolean>> dryRun() {
        return Codegen.optional(this.dryRun);
    }
    /**
     * The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
     * 
     */
    @Export(name="peerConnectionName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> peerConnectionName;

    /**
     * @return The name of the VPC peer connection. The name of the resource. The name must be 2 to 128 characters in length, and must start with a letter. It can contain digits, underscores (_), and hyphens (-).
     * 
     */
    public Output<Optional<String>> peerConnectionName() {
        return Codegen.optional(this.peerConnectionName);
    }
    /**
     * The ID of the resource group.
     * 
     */
    @Export(name="resourceGroupId", refs={String.class}, tree="[0]")
    private Output<String> resourceGroupId;

    /**
     * @return The ID of the resource group.
     * 
     */
    public Output<String> resourceGroupId() {
        return this.resourceGroupId;
    }
    /**
     * The status of the VPC peer connection.
     * 
     */
    @Export(name="status", refs={String.class}, tree="[0]")
    private Output<String> status;

    /**
     * @return The status of the VPC peer connection.
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
     * The ID of the requester VPC.
     * 
     */
    @Export(name="vpcId", refs={String.class}, tree="[0]")
    private Output<String> vpcId;

    /**
     * @return The ID of the requester VPC.
     * 
     */
    public Output<String> vpcId() {
        return this.vpcId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public PeerConnection(String name) {
        this(name, PeerConnectionArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public PeerConnection(String name, PeerConnectionArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public PeerConnection(String name, PeerConnectionArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/peerConnection:PeerConnection", name, args == null ? PeerConnectionArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private PeerConnection(String name, Output<String> id, @Nullable PeerConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:vpc/peerConnection:PeerConnection", name, state, makeResourceOptions(options, id));
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
    public static PeerConnection get(String name, Output<String> id, @Nullable PeerConnectionState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new PeerConnection(name, id, state, options);
    }
}
