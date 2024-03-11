// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.pvtz;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.pvtz.ZoneAttachmentArgs;
import com.pulumi.alicloud.pvtz.inputs.ZoneAttachmentState;
import com.pulumi.alicloud.pvtz.outputs.ZoneAttachmentVpc;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * Using `vpc_ids` to attach being in same region several vpc instances to a private zone
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.pvtz.Zone;
 * import com.pulumi.alicloud.pvtz.ZoneArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.pvtz.ZoneAttachment;
 * import com.pulumi.alicloud.pvtz.ZoneAttachmentArgs;
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
 *         var zone = new Zone(&#34;zone&#34;, ZoneArgs.builder()        
 *             .zoneName(&#34;foo.example.com&#34;)
 *             .build());
 * 
 *         var first = new Network(&#34;first&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-first-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var second = new Network(&#34;second&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-second-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var zone_attachment = new ZoneAttachment(&#34;zone-attachment&#34;, ZoneAttachmentArgs.builder()        
 *             .zoneId(zone.id())
 *             .vpcIds(            
 *                 first.id(),
 *                 second.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Using `vpcs` to attach being in same region several vpc instances to a private zone
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.pvtz.Zone;
 * import com.pulumi.alicloud.pvtz.ZoneArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.pvtz.ZoneAttachment;
 * import com.pulumi.alicloud.pvtz.ZoneAttachmentArgs;
 * import com.pulumi.alicloud.pvtz.inputs.ZoneAttachmentVpcArgs;
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
 *         var zone = new Zone(&#34;zone&#34;, ZoneArgs.builder()        
 *             .zoneName(&#34;foo.example.com&#34;)
 *             .build());
 * 
 *         var first = new Network(&#34;first&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-first-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var second = new Network(&#34;second&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-second-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var zone_attachment = new ZoneAttachment(&#34;zone-attachment&#34;, ZoneAttachmentArgs.builder()        
 *             .zoneId(zone.id())
 *             .vpcs(            
 *                 ZoneAttachmentVpcArgs.builder()
 *                     .vpcId(first.id())
 *                     .build(),
 *                 ZoneAttachmentVpcArgs.builder()
 *                     .vpcId(second.id())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * Using `vpcs` to attach being in different regions several vpc instances to a private zone
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.alicloud.pvtz.Zone;
 * import com.pulumi.alicloud.pvtz.ZoneArgs;
 * import com.pulumi.alicloud.vpc.Network;
 * import com.pulumi.alicloud.vpc.NetworkArgs;
 * import com.pulumi.alicloud.Provider;
 * import com.pulumi.alicloud.ProviderArgs;
 * import com.pulumi.alicloud.pvtz.ZoneAttachment;
 * import com.pulumi.alicloud.pvtz.ZoneAttachmentArgs;
 * import com.pulumi.alicloud.pvtz.inputs.ZoneAttachmentVpcArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var zone = new Zone(&#34;zone&#34;, ZoneArgs.builder()        
 *             .zoneName(&#34;foo.example.com&#34;)
 *             .build());
 * 
 *         var first = new Network(&#34;first&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-first-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/12&#34;)
 *             .build());
 * 
 *         var second = new Network(&#34;second&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-second-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build());
 * 
 *         var eu = new Provider(&#34;eu&#34;, ProviderArgs.builder()        
 *             .region(&#34;eu-central-1&#34;)
 *             .build());
 * 
 *         var third = new Network(&#34;third&#34;, NetworkArgs.builder()        
 *             .vpcName(&#34;the-third-vpc&#34;)
 *             .cidrBlock(&#34;172.16.0.0/16&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .provider(alicloud.eu())
 *                 .build());
 * 
 *         var zone_attachment = new ZoneAttachment(&#34;zone-attachment&#34;, ZoneAttachmentArgs.builder()        
 *             .zoneId(zone.id())
 *             .vpcs(            
 *                 ZoneAttachmentVpcArgs.builder()
 *                     .vpcId(first.id())
 *                     .build(),
 *                 ZoneAttachmentVpcArgs.builder()
 *                     .vpcId(second.id())
 *                     .build(),
 *                 ZoneAttachmentVpcArgs.builder()
 *                     .regionId(&#34;eu-central-1&#34;)
 *                     .vpcId(third.id())
 *                     .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Private Zone attachment can be imported using the id(same with `zone_id`), e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:pvtz/zoneAttachment:ZoneAttachment example abc123456
 * ```
 * 
 */
@ResourceType(type="alicloud:pvtz/zoneAttachment:ZoneAttachment")
public class ZoneAttachment extends com.pulumi.resources.CustomResource {
    /**
     * The language of code.
     * 
     */
    @Export(name="lang", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lang;

    /**
     * @return The language of code.
     * 
     */
    public Output<Optional<String>> lang() {
        return Codegen.optional(this.lang);
    }
    /**
     * The user custom IP address.
     * 
     */
    @Export(name="userClientIp", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userClientIp;

    /**
     * @return The user custom IP address.
     * 
     */
    public Output<Optional<String>> userClientIp() {
        return Codegen.optional(this.userClientIp);
    }
    /**
     * The id List of the VPC with the same region, for example:[&#34;vpc-1&#34;,&#34;vpc-2&#34;].
     * 
     */
    @Export(name="vpcIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> vpcIds;

    /**
     * @return The id List of the VPC with the same region, for example:[&#34;vpc-1&#34;,&#34;vpc-2&#34;].
     * 
     */
    public Output<List<String>> vpcIds() {
        return this.vpcIds;
    }
    /**
     * See `vpcs` below.Recommend to use `vpcs`.
     * 
     */
    @Export(name="vpcs", refs={List.class,ZoneAttachmentVpc.class}, tree="[0,1]")
    private Output<List<ZoneAttachmentVpc>> vpcs;

    /**
     * @return See `vpcs` below.Recommend to use `vpcs`.
     * 
     */
    public Output<List<ZoneAttachmentVpc>> vpcs() {
        return this.vpcs;
    }
    /**
     * The name of the Private Zone Record.
     * 
     */
    @Export(name="zoneId", refs={String.class}, tree="[0]")
    private Output<String> zoneId;

    /**
     * @return The name of the Private Zone Record.
     * 
     */
    public Output<String> zoneId() {
        return this.zoneId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ZoneAttachment(String name) {
        this(name, ZoneAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ZoneAttachment(String name, ZoneAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ZoneAttachment(String name, ZoneAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, args == null ? ZoneAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ZoneAttachment(String name, Output<String> id, @Nullable ZoneAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:pvtz/zoneAttachment:ZoneAttachment", name, state, makeResourceOptions(options, id));
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
    public static ZoneAttachment get(String name, Output<String> id, @Nullable ZoneAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ZoneAttachment(name, id, state, options);
    }
}
