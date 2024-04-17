// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.alicloud.ens;

import com.pulumi.alicloud.Utilities;
import com.pulumi.alicloud.ens.InstanceSecurityGroupAttachmentArgs;
import com.pulumi.alicloud.ens.inputs.InstanceSecurityGroupAttachmentState;
import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a ENS Instance Security Group Attachment resource. Unbind instance and security group.
 * 
 * For information about ENS Instance Security Group Attachment and how to use it, see [What is Instance Security Group Attachment](https://www.alibabacloud.com/help/en/).
 * 
 * &gt; **NOTE:** Available since v1.216.0.
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
 * import com.pulumi.alicloud.ens.Instance;
 * import com.pulumi.alicloud.ens.InstanceArgs;
 * import com.pulumi.alicloud.ens.inputs.InstanceSystemDiskArgs;
 * import com.pulumi.alicloud.ens.SecurityGroup;
 * import com.pulumi.alicloud.ens.SecurityGroupArgs;
 * import com.pulumi.alicloud.ens.InstanceSecurityGroupAttachment;
 * import com.pulumi.alicloud.ens.InstanceSecurityGroupAttachmentArgs;
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
 *         final var name = config.get(&#34;name&#34;).orElse(&#34;terraform-example&#34;);
 *         var default_ = new Instance(&#34;default&#34;, InstanceArgs.builder()        
 *             .systemDisk(InstanceSystemDiskArgs.builder()
 *                 .size(&#34;20&#34;)
 *                 .build())
 *             .scheduleAreaLevel(&#34;Region&#34;)
 *             .imageId(&#34;centos_6_08_64_20G_alibase_20171208&#34;)
 *             .paymentType(&#34;Subscription&#34;)
 *             .instanceType(&#34;ens.sn1.stiny&#34;)
 *             .password(&#34;12345678ABCabc&#34;)
 *             .amount(&#34;1&#34;)
 *             .period(&#34;1&#34;)
 *             .internetMaxBandwidthOut(&#34;10&#34;)
 *             .publicIpIdentification(true)
 *             .ensRegionId(&#34;cn-chenzhou-telecom_unicom_cmcc&#34;)
 *             .periodUnit(&#34;Month&#34;)
 *             .build());
 * 
 *         var defaultSecurityGroup = new SecurityGroup(&#34;defaultSecurityGroup&#34;, SecurityGroupArgs.builder()        
 *             .description(&#34;InstanceSecurityGroupAttachment_Description&#34;)
 *             .securityGroupName(name)
 *             .build());
 * 
 *         var defaultInstanceSecurityGroupAttachment = new InstanceSecurityGroupAttachment(&#34;defaultInstanceSecurityGroupAttachment&#34;, InstanceSecurityGroupAttachmentArgs.builder()        
 *             .instanceId(default_.id())
 *             .securityGroupId(defaultSecurityGroup.id())
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * ENS Instance Security Group Attachment can be imported using the id, e.g.
 * 
 * ```sh
 * $ pulumi import alicloud:ens/instanceSecurityGroupAttachment:InstanceSecurityGroupAttachment example &lt;instance_id&gt;:&lt;security_group_id&gt;
 * ```
 * 
 */
@ResourceType(type="alicloud:ens/instanceSecurityGroupAttachment:InstanceSecurityGroupAttachment")
public class InstanceSecurityGroupAttachment extends com.pulumi.resources.CustomResource {
    /**
     * Instance ID.
     * 
     */
    @Export(name="instanceId", refs={String.class}, tree="[0]")
    private Output<String> instanceId;

    /**
     * @return Instance ID.
     * 
     */
    public Output<String> instanceId() {
        return this.instanceId;
    }
    /**
     * Security group ID.
     * 
     */
    @Export(name="securityGroupId", refs={String.class}, tree="[0]")
    private Output<String> securityGroupId;

    /**
     * @return Security group ID.
     * 
     */
    public Output<String> securityGroupId() {
        return this.securityGroupId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public InstanceSecurityGroupAttachment(String name) {
        this(name, InstanceSecurityGroupAttachmentArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public InstanceSecurityGroupAttachment(String name, InstanceSecurityGroupAttachmentArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public InstanceSecurityGroupAttachment(String name, InstanceSecurityGroupAttachmentArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/instanceSecurityGroupAttachment:InstanceSecurityGroupAttachment", name, args == null ? InstanceSecurityGroupAttachmentArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private InstanceSecurityGroupAttachment(String name, Output<String> id, @Nullable InstanceSecurityGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("alicloud:ens/instanceSecurityGroupAttachment:InstanceSecurityGroupAttachment", name, state, makeResourceOptions(options, id));
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
    public static InstanceSecurityGroupAttachment get(String name, Output<String> id, @Nullable InstanceSecurityGroupAttachmentState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new InstanceSecurityGroupAttachment(name, id, state, options);
    }
}
