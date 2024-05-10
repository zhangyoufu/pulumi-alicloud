// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cen;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cen.TransitRouterGrantAttachmentArgs;
import com.pulumi.alicloud.cen.inputs.TransitRouterGrantAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Enterprise Network (CEN) Transit Router Grant Attachment resource.
 * 
 * For information about Cloud Enterprise Network (CEN) Transit Router Grant Attachment and how to use it, see [What is Transit Router Grant Attachment](https://www.alibabacloud.com/help/en/cloud-enterprise-network/latest/grantinstancetotransitrouter).
 * 
 * &gt; **NOTE:** Available since v1.187.0.
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
 * import com.pulumi.alicloud.cen.Instance;
 * import com.pulumi.alicloud.cen.InstanceArgs;
 * import com.pulumi.alicloud.cen.TransitRouterGrantAttachment;
 * import com.pulumi.alicloud.cen.TransitRouterGrantAttachmentArgs;
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
 *         final var default = AlicloudFunctions.getAccount();
 * 
 *         var example = new Network("example", NetworkArgs.builder()        
 *             .vpcName("tf_example")
 *             .cidrBlock("172.17.3.0/24")
 *             .build());
 * 
 *         var exampleInstance = new Instance("exampleInstance", InstanceArgs.builder()        
 *             .cenInstanceName("tf_example")
 *             .description("an example for cen")
 *             .build());
 * 
 *         var exampleTransitRouterGrantAttachment = new TransitRouterGrantAttachment("exampleTransitRouterGrantAttachment", TransitRouterGrantAttachmentArgs.builder()        
 *             .cenId(exampleInstance.id())
 *             .cenOwnerId(default_.id())
 *             .instanceId(example.id())
 *             .instanceType("VPC")
 *             .orderType("PayByCenOwner")
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
 * Cloud Enterprise Network (CEN) Transit Router Grant Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment example &lt;instance_type&gt;:&lt;instance_id&gt;:&lt;cen_owner_id&gt;:&lt;cen_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment")
public class TransitRouterGrantAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
     * 
     */
    @Export(name="cenId", refs={String.class}, tree="[0]")
    private Output<String> cenId;

    /**
     * @return The ID of the Cloud Enterprise Network (CEN) instance to which the transit router belongs.
     * 
     */
    public Output<String> cenId() {
        return this.cenId;
    }
    /**
     * The ID of the Alibaba Cloud account to which the CEN instance belongs.
     * 
     */
    @Export(name="cenOwnerId", refs={String.class}, tree="[0]")
    private Output<String> cenOwnerId;

    /**
     * @return The ID of the Alibaba Cloud account to which the CEN instance belongs.
     * 
     */
    public Output<String> cenOwnerId() {
        return this.cenOwnerId;
    }
    /**
     * The ID of the network instance.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the network instance.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
     * 
     */
    @Export(name="instanceType", refs={String.class}, tree="[0]")
    private Output<String> instanceType;

    /**
     * @return The type of the network instance. Valid values: `VPC`, `ExpressConnect`, `VPN`.
     * 
     */
    public Output<String> instanceType() {
        return this.instanceType;
    }
    /**
     * The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
     * 
     */
    @Export(name="orderType", refs={String.class}, tree="[0]")
    private Output<String> orderType;

    /**
     * @return The entity that pays the fees of the network instance. Valid values: `PayByResourceOwner`, `PayByCenOwner`.
     * 
     */
    public Output<String> orderType() {
        return this.orderType;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public TransitRouterGrantAttachment(String name) {
        this(name, TransitRouterGrantAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public TransitRouterGrantAttachment(String name, TransitRouterGrantAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public TransitRouterGrantAttachment(String name, TransitRouterGrantAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, args == null ? TransitRouterGrantAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private TransitRouterGrantAttachment(String name, Output<String> id, @Nullable TransitRouterGrantAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cen/transitRouterGrantAttachment:TransitRouterGrantAttachment", name, state, makeResourceOptions(options, id));
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
    public static TransitRouterGrantAttachment get(String name, Output<String> id, @Nullable TransitRouterGrantAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new TransitRouterGrantAttachment(name, id, state, options);
    }
}
