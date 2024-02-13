// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.bastionhost.UserGroupArgs;
import com.pulumi.alicloud.bastionhost.inputs.UserGroupState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a Bastion Host User Group resource.
 * 
 * For information about Bastion Host User Group and how to use it, see [What is User Group](https://www.alibabacloud.com/help/doc-detail/204596.htm).
 * 
 * &gt; **NOTE:** Available since v1.132.0.
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
 * import com.pulumi.alicloud.AlicloudFunctions;
 * import com.pulumi.alicloud.inputs.GetZonesArgs;
 * import com.pulumi.alicloud.vpc.VpcFunctions;
 * import com.pulumi.alicloud.vpc.inputs.GetNetworksArgs;
 * import com.pulumi.alicloud.vpc.inputs.GetSwitchesArgs;
 * import com.pulumi.alicloud.ecs.SecurityGroup;
 * import com.pulumi.alicloud.ecs.SecurityGroupArgs;
 * import com.pulumi.alicloud.bastionhost.Instance;
 * import com.pulumi.alicloud.bastionhost.InstanceArgs;
 * import com.pulumi.alicloud.bastionhost.UserGroup;
 * import com.pulumi.alicloud.bastionhost.UserGroupArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;tf_example&#34;);
 *         final var defaultZones = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation(&#34;VSwitch&#34;)
 *             .build());
 * 
 *         final var defaultNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex(&#34;^default-NODELETING$&#34;)
 *             .cidrBlock(&#34;10.4.0.0/16&#34;)
 *             .build());
 * 
 *         final var defaultSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .cidrBlock(&#34;10.4.0.0/24&#34;)
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .zoneId(defaultZones.applyValue(getZonesResult -&gt; getZonesResult.zones()[0].id()))
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .vpcId(defaultNetworks.applyValue(getNetworksResult -&gt; getNetworksResult.ids()[0]))
 *             .build());
 * 
 *         var defaultInstance = new Instance(&#34;defaultInstance&#34;, InstanceArgs.builder()        
 *             .description(name)
 *             .licenseCode(&#34;bhah_ent_50_asset&#34;)
 *             .planCode(&#34;cloudbastion&#34;)
 *             .storage(&#34;5&#34;)
 *             .bandwidth(&#34;5&#34;)
 *             .period(&#34;1&#34;)
 *             .vswitchId(defaultSwitches.applyValue(getSwitchesResult -&gt; getSwitchesResult.ids()[0]))
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .build());
 * 
 *         var defaultUserGroup = new UserGroup(&#34;defaultUserGroup&#34;, UserGroupArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .userGroupName(name)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bastion Host User Group can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:bastionhost/userGroup:UserGroup example &lt;instance_id&gt;:&lt;user_group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:bastionhost/userGroup:UserGroup")
public class UserGroup extends com.pulumi.resources.CustomResource {
    /**
     * Specify the New Group of Remark Information. Supports up to 500 Characters.
     * 
     */
    @Export(name="comment", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> comment;

    /**
     * @return Specify the New Group of Remark Information. Supports up to 500 Characters.
     * 
     */
    public Output<Optional<String>> comment() {
        return Codegen.optional(this.comment);
    }
    /**
     * Specify the New Group of the Bastion Host of Instance Id.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Specify the New Group of the Bastion Host of Instance Id.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The User Group self ID.
     * 
     */
    @Export(name="userGroupId", refs={String.class}, tree="[0]")
    private Output<String> userGroupId;

    /**
     * @return The User Group self ID.
     * 
     */
    public Output<String> userGroupId() {
        return this.userGroupId;
    }
    /**
     * Specify the New Group Name. Supports up to 128 Characters.
     * 
     */
    @Export(name="userGroupName", refs={String.class}, tree="[0]")
    private Output<String> userGroupName;

    /**
     * @return Specify the New Group Name. Supports up to 128 Characters.
     * 
     */
    public Output<String> userGroupName() {
        return this.userGroupName;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserGroup(String name) {
        this(name, UserGroupArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserGroup(String name, UserGroupArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserGroup(String name, UserGroupArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/userGroup:UserGroup", name, args == null ? UserGroupArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserGroup(String name, Output<String> id, @Nullable UserGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/userGroup:UserGroup", name, state, makeResourceOptions(options, id));
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
    public static UserGroup get(String name, Output<String> id, @Nullable UserGroupState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserGroup(name, id, state, options);
    }
}
