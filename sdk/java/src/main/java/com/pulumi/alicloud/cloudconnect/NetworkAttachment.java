// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.cloudconnect;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.cloudconnect.NetworkAttachmentArgs;
import com.pulumi.alicloud.cloudconnect.inputs.NetworkAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Cloud Connect Network Attachment resource. This topic describes how to associate a Smart Access Gateway (SAG) instance with a network instance. You must associate an SAG instance with a network instance if you want to connect the SAG to Alibaba Cloud. You can connect an SAG to Alibaba Cloud through a leased line, the Internet, or the active and standby links.
 * 
 * For information about Cloud Connect Network Attachment and how to use it, see [What is Cloud Connect Network Attachment](https://www.alibabacloud.com/help/en/smart-access-gateway/latest/bindsmartaccessgateway).
 * 
 * &gt; **NOTE:** Available since v1.64.0.
 * 
 * &gt; **NOTE:** Only the following regions support. [`cn-shanghai`, `cn-shanghai-finance-1`, `cn-hongkong`, `ap-southeast-1`, `ap-southeast-2`, `ap-southeast-3`, `ap-southeast-5`, `ap-northeast-1`, `eu-central-1`]
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
 * import com.pulumi.alicloud.cloudconnect.Network;
 * import com.pulumi.alicloud.cloudconnect.NetworkArgs;
 * import com.pulumi.alicloud.cloudconnect.NetworkAttachment;
 * import com.pulumi.alicloud.cloudconnect.NetworkAttachmentArgs;
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
 *         final var sagId = config.get(&#34;sagId&#34;).orElse(&#34;sag-9bifkf***&#34;);
 *         var defaultNetwork = new Network(&#34;defaultNetwork&#34;, NetworkArgs.builder()        
 *             .description(name)
 *             .cidrBlock(&#34;192.168.0.0/24&#34;)
 *             .isDefault(true)
 *             .build());
 * 
 *         var defaultNetworkAttachment = new NetworkAttachment(&#34;defaultNetworkAttachment&#34;, NetworkAttachmentArgs.builder()        
 *             .ccnId(defaultNetwork.id())
 *             .sagId(sagId)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * The Cloud Connect Network Attachment can be imported using the instance_id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:cloudconnect/networkAttachment:NetworkAttachment example ccn-abc123456:sag-abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:cloudconnect/networkAttachment:NetworkAttachment")
public class NetworkAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The ID of the CCN instance.
     * 
     */
    @Export(name="ccnId", refs={String.class}, tree="[0]")
    private Output<String> ccnId;

    /**
     * @return The ID of the CCN instance.
     * 
     */
    public Output<String> ccnId() {
        return this.ccnId;
    }
    /**
     * The ID of the Smart Access Gateway instance.
     * 
     */
    @Export(name="sagId", refs={String.class}, tree="[0]")
    private Output<String> sagId;

    /**
     * @return The ID of the Smart Access Gateway instance.
     * 
     */
    public Output<String> sagId() {
        return this.sagId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public NetworkAttachment(String name) {
        this(name, NetworkAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public NetworkAttachment(String name, NetworkAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public NetworkAttachment(String name, NetworkAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudconnect/networkAttachment:NetworkAttachment", name, args == null ? NetworkAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private NetworkAttachment(String name, Output<String> id, @Nullable NetworkAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:cloudconnect/networkAttachment:NetworkAttachment", name, state, makeResourceOptions(options, id));
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
    public static NetworkAttachment get(String name, Output<String> id, @Nullable NetworkAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new NetworkAttachment(name, id, state, options);
    }
}
