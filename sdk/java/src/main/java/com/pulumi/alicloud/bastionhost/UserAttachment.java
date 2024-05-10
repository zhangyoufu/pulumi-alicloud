// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.bastionhost;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.bastionhost.UserAttachmentArgs;
import com.pulumi.alicloud.bastionhost.inputs.UserAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a Bastion Host User Attachment resource to add user to one user group.
 * 
 * &gt; **NOTE:** Available since v1.134.0.
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
 * import com.pulumi.alicloud.bastionhost.User;
 * import com.pulumi.alicloud.bastionhost.UserArgs;
 * import com.pulumi.alicloud.bastionhost.UserAttachment;
 * import com.pulumi.alicloud.bastionhost.UserAttachmentArgs;
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
 *         final var name = config.get("name").orElse("tf_example");
 *         final var default = AlicloudFunctions.getZones(GetZonesArgs.builder()
 *             .availableResourceCreation("VSwitch")
 *             .build());
 * 
 *         final var defaultGetNetworks = VpcFunctions.getNetworks(GetNetworksArgs.builder()
 *             .nameRegex("^default-NODELETING$")
 *             .cidrBlock("10.4.0.0/16")
 *             .build());
 * 
 *         final var defaultGetSwitches = VpcFunctions.getSwitches(GetSwitchesArgs.builder()
 *             .cidrBlock("10.4.0.0/24")
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .zoneId(default_.zones()[0].id())
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup("defaultSecurityGroup", SecurityGroupArgs.builder()        
 *             .vpcId(defaultGetNetworks.applyValue(getNetworksResult -> getNetworksResult.ids()[0]))
 *             .build());
 * 
 *         var defaultInstance = new Instance("defaultInstance", InstanceArgs.builder()        
 *             .description(name)
 *             .licenseCode("bhah_ent_50_asset")
 *             .planCode("cloudbastion")
 *             .storage("5")
 *             .bandwidth("5")
 *             .period("1")
 *             .vswitchId(defaultGetSwitches.applyValue(getSwitchesResult -> getSwitchesResult.ids()[0]))
 *             .securityGroupIds(defaultSecurityGroup.id())
 *             .build());
 * 
 *         var defaultUserGroup = new UserGroup("defaultUserGroup", UserGroupArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .userGroupName(name)
 *             .build());
 * 
 *         var localUser = new User("localUser", UserArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .mobileCountryCode("CN")
 *             .mobile("13312345678")
 *             .password("YourPassword-123")
 *             .source("Local")
 *             .userName(String.format("%s_local_user", name))
 *             .build());
 * 
 *         var defaultUserAttachment = new UserAttachment("defaultUserAttachment", UserAttachmentArgs.builder()        
 *             .instanceId(defaultInstance.id())
 *             .userGroupId(defaultUserGroup.userGroupId())
 *             .userId(localUser.userId())
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
 * Bastion Host User Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:bastionhost/userAttachment:UserAttachment example &lt;instance_id&gt;:&lt;user_group_id&gt;:&lt;user_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:bastionhost/userAttachment:UserAttachment")
public class UserAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the user group to add the user&#39;s bastion host ID of.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Specifies the user group to add the user&#39;s bastion host ID of.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Specifies the user group to which you want to add the user ID.
     * 
     */
    @Export(name="userGroupId", refs={String.class}, tree="[0]")
    private Output<String> userGroupId;

    /**
     * @return Specifies the user group to which you want to add the user ID.
     * 
     */
    public Output<String> userGroupId() {
        return this.userGroupId;
    }
    /**
     * Specify that you want to add to the policy attached to the user group ID. This includes response parameters in a Json-formatted string supports up to set up 100 USER ID.
     * 
     */
    @Export(name="userId", refs={String.class}, tree="[0]")
    private Output<String> userId;

    /**
     * @return Specify that you want to add to the policy attached to the user group ID. This includes response parameters in a Json-formatted string supports up to set up 100 USER ID.
     * 
     */
    public Output<String> userId() {
        return this.userId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserAttachment(String name) {
        this(name, UserAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserAttachment(String name, UserAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserAttachment(String name, UserAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/userAttachment:UserAttachment", name, args == null ? UserAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserAttachment(String name, Output<String> id, @Nullable UserAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:bastionhost/userAttachment:UserAttachment", name, state, makeResourceOptions(options, id));
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
    public static UserAttachment get(String name, Output<String> id, @Nullable UserAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserAttachment(name, id, state, options);
    }
}
