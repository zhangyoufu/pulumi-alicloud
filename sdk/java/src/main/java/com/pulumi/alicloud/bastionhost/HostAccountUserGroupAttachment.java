// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.bastionhost.HostAccountUserGroupAttachmentArgs;
import com.pulumi.alicloud.bastionhost.inputs.HostAccountUserGroupAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import java.util.List;
import javax.annotation.Nullable;

/**
 * Provides a Bastion Host Host Account Attachment resource to add list host accounts into one user group.
 * 
 * &gt; **NOTE:** Available since v1.135.0.
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
 * import com.pulumi.alicloud.bastionhost.Host;
 * import com.pulumi.alicloud.bastionhost.HostArgs;
 * import com.pulumi.alicloud.bastionhost.HostAccount;
 * import com.pulumi.alicloud.bastionhost.HostAccountArgs;
 * import com.pulumi.alicloud.bastionhost.UserGroup;
 * import com.pulumi.alicloud.bastionhost.UserGroupArgs;
 * import com.pulumi.alicloud.bastionhost.HostAccountUserGroupAttachment;
 * import com.pulumi.alicloud.bastionhost.HostAccountUserGroupAttachmentArgs;
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
 *         var defaultHost = new Host(&#34;defaultHost&#34;, HostArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .hostName(name)
 *             .activeAddressType(&#34;Private&#34;)
 *             .hostPrivateAddress(&#34;172.16.0.10&#34;)
 *             .osType(&#34;Linux&#34;)
 *             .source(&#34;Local&#34;)
 *             .build());
 * 
 *         var defaultHostAccount = new HostAccount(&#34;defaultHostAccount&#34;, HostAccountArgs.builder()        
 *             .hostAccountName(name)
 *             .hostId(defaultHost.hostId())
 *             .instanceId(defaultHost.instanceId())
 *             .protocolName(&#34;SSH&#34;)
 *             .password(&#34;YourPassword12345&#34;)
 *             .build());
 * 
 *         var defaultUserGroup = new UserGroup(&#34;defaultUserGroup&#34;, UserGroupArgs.builder()        
 *             .instanceId(defaultHost.instanceId())
 *             .userGroupName(name)
 *             .build());
 * 
 *         var defaultHostAccountUserGroupAttachment = new HostAccountUserGroupAttachment(&#34;defaultHostAccountUserGroupAttachment&#34;, HostAccountUserGroupAttachmentArgs.builder()        
 *             .instanceId(defaultHost.instanceId())
 *             .userGroupId(defaultUserGroup.userGroupId())
 *             .hostId(defaultHost.hostId())
 *             .hostAccountIds(defaultHostAccount.hostAccountId())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Bastion Host Host Account can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment example &lt;instance_id&gt;:&lt;user_group_id&gt;:&lt;host_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment")
public class HostAccountUserGroupAttachment extends com.pulumi.resources.CustomResource {
    /**
     * A list IDs of the host account.
     * 
     */
    @Export(name="hostAccountIds", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> hostAccountIds;

    /**
     * @return A list IDs of the host account.
     * 
     */
    public Output<List<String>> hostAccountIds() {
        return this.hostAccountIds;
    }
    /**
     * The ID of the host.
     * 
     */
    @Export(name="hostId", refs={String.class}, tree="[0]")
    private Output<String> hostId;

    /**
     * @return The ID of the host.
     * 
     */
    public Output<String> hostId() {
        return this.hostId;
    }
    /**
     * The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return The ID of the Bastionhost instance where you want to authorize the user group to manage the specified hosts and host accounts.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     * 
     */
    @Export(name="userGroupId", refs={String.class}, tree="[0]")
    private Output<String> userGroupId;

    /**
     * @return The ID of the user group that you want to authorize to manage the specified hosts and host accounts.
     * 
     */
    public Output<String> userGroupId() {
        return this.userGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HostAccountUserGroupAttachment(String name) {
        this(name, HostAccountUserGroupAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HostAccountUserGroupAttachment(String name, HostAccountUserGroupAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HostAccountUserGroupAttachment(String name, HostAccountUserGroupAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment", name, args == null ? HostAccountUserGroupAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HostAccountUserGroupAttachment(String name, Output<String> id, @Nullable HostAccountUserGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/hostAccountUserGroupAttachment:HostAccountUserGroupAttachment", name, state, makeResourceOptions(options, id));
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
    public static HostAccountUserGroupAttachment get(String name, Output<String> id, @Nullable HostAccountUserGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HostAccountUserGroupAttachment(name, id, state, options);
    }
}
