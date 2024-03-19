// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ros
{
    /// <summary>
    /// Provides a ROS Change Set resource.
    /// 
    /// For information about ROS Change Set and how to use it, see [What is Change Set](https://www.alibabacloud.com/help/doc-detail/131051.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.105.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// using Random = Pulumi.Random;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Random.RandomInteger("default", new()
    ///     {
    ///         Max = 99999,
    ///         Min = 10000,
    ///     });
    /// 
    ///     var example = new AliCloud.Ros.ChangeSet("example", new()
    ///     {
    ///         ChangeSetName = "example_value",
    ///         ChangeSetType = "CREATE",
    ///         Description = "Test From Terraform",
    ///         StackName = @default.Result.Apply(result =&gt; $"tf-example-{result}"),
    ///         TemplateBody = "{\"ROSTemplateFormatVersion\":\"2015-09-01\"}",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// ROS Change Set can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:ros/changeSet:ChangeSet example &lt;change_set_id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:ros/changeSet:ChangeSet")]
    public partial class ChangeSet : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        /// </summary>
        [Output("changeSetName")]
        public Output<string> ChangeSetName { get; private set; } = null!;

        /// <summary>
        /// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
        /// </summary>
        [Output("changeSetType")]
        public Output<string?> ChangeSetType { get; private set; } = null!;

        /// <summary>
        /// The description of the change set. The description can be up to 1,024 bytes in length.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Output("disableRollback")]
        public Output<bool?> DisableRollback { get; private set; } = null!;

        /// <summary>
        /// The notification urls.
        /// </summary>
        [Output("notificationUrls")]
        public Output<ImmutableArray<string>> NotificationUrls { get; private set; } = null!;

        /// <summary>
        /// Parameters.
        /// </summary>
        [Output("parameters")]
        public Output<ImmutableArray<Outputs.ChangeSetParameter>> Parameters { get; private set; } = null!;

        /// <summary>
        /// The ram role name.
        /// </summary>
        [Output("ramRoleName")]
        public Output<string?> RamRoleName { get; private set; } = null!;

        /// <summary>
        /// The replacement option.
        /// </summary>
        [Output("replacementOption")]
        public Output<string?> ReplacementOption { get; private set; } = null!;

        /// <summary>
        /// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
        /// </summary>
        [Output("stackId")]
        public Output<string> StackId { get; private set; } = null!;

        /// <summary>
        /// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Output("stackName")]
        public Output<string?> StackName { get; private set; } = null!;

        /// <summary>
        /// The stack policy body.
        /// </summary>
        [Output("stackPolicyBody")]
        public Output<string?> StackPolicyBody { get; private set; } = null!;

        /// <summary>
        /// The stack policy during update body.
        /// </summary>
        [Output("stackPolicyDuringUpdateBody")]
        public Output<string?> StackPolicyDuringUpdateBody { get; private set; } = null!;

        /// <summary>
        /// The stack policy during update url.
        /// </summary>
        [Output("stackPolicyDuringUpdateUrl")]
        public Output<string?> StackPolicyDuringUpdateUrl { get; private set; } = null!;

        /// <summary>
        /// The stack policy url.
        /// </summary>
        [Output("stackPolicyUrl")]
        public Output<string?> StackPolicyUrl { get; private set; } = null!;

        /// <summary>
        /// The status of the change set.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
        /// </summary>
        [Output("templateBody")]
        public Output<string?> TemplateBody { get; private set; } = null!;

        /// <summary>
        /// The template url.
        /// </summary>
        [Output("templateUrl")]
        public Output<string?> TemplateUrl { get; private set; } = null!;

        /// <summary>
        /// Timeout In Minutes.
        /// </summary>
        [Output("timeoutInMinutes")]
        public Output<int> TimeoutInMinutes { get; private set; } = null!;

        /// <summary>
        /// The use previous parameters.
        /// </summary>
        [Output("usePreviousParameters")]
        public Output<bool?> UsePreviousParameters { get; private set; } = null!;


        /// <summary>
        /// Create a ChangeSet resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ChangeSet(string name, ChangeSetArgs args, CustomResourceOptions? options = null)
            : base("alicloud:ros/changeSet:ChangeSet", name, args ?? new ChangeSetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ChangeSet(string name, Input<string> id, ChangeSetState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:ros/changeSet:ChangeSet", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ChangeSet resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ChangeSet Get(string name, Input<string> id, ChangeSetState? state = null, CustomResourceOptions? options = null)
        {
            return new ChangeSet(name, id, state, options);
        }
    }

    public sealed class ChangeSetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        /// </summary>
        [Input("changeSetName", required: true)]
        public Input<string> ChangeSetName { get; set; } = null!;

        /// <summary>
        /// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
        /// </summary>
        [Input("changeSetType")]
        public Input<string>? ChangeSetType { get; set; }

        /// <summary>
        /// The description of the change set. The description can be up to 1,024 bytes in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Input("disableRollback")]
        public Input<bool>? DisableRollback { get; set; }

        [Input("notificationUrls")]
        private InputList<string>? _notificationUrls;

        /// <summary>
        /// The notification urls.
        /// </summary>
        public InputList<string> NotificationUrls
        {
            get => _notificationUrls ?? (_notificationUrls = new InputList<string>());
            set => _notificationUrls = value;
        }

        [Input("parameters")]
        private InputList<Inputs.ChangeSetParameterArgs>? _parameters;

        /// <summary>
        /// Parameters.
        /// </summary>
        public InputList<Inputs.ChangeSetParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ChangeSetParameterArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The ram role name.
        /// </summary>
        [Input("ramRoleName")]
        public Input<string>? RamRoleName { get; set; }

        /// <summary>
        /// The replacement option.
        /// </summary>
        [Input("replacementOption")]
        public Input<string>? ReplacementOption { get; set; }

        /// <summary>
        /// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Input("stackName")]
        public Input<string>? StackName { get; set; }

        /// <summary>
        /// The stack policy body.
        /// </summary>
        [Input("stackPolicyBody")]
        public Input<string>? StackPolicyBody { get; set; }

        /// <summary>
        /// The stack policy during update body.
        /// </summary>
        [Input("stackPolicyDuringUpdateBody")]
        public Input<string>? StackPolicyDuringUpdateBody { get; set; }

        /// <summary>
        /// The stack policy during update url.
        /// </summary>
        [Input("stackPolicyDuringUpdateUrl")]
        public Input<string>? StackPolicyDuringUpdateUrl { get; set; }

        /// <summary>
        /// The stack policy url.
        /// </summary>
        [Input("stackPolicyUrl")]
        public Input<string>? StackPolicyUrl { get; set; }

        /// <summary>
        /// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
        /// </summary>
        [Input("templateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// The template url.
        /// </summary>
        [Input("templateUrl")]
        public Input<string>? TemplateUrl { get; set; }

        /// <summary>
        /// Timeout In Minutes.
        /// </summary>
        [Input("timeoutInMinutes")]
        public Input<int>? TimeoutInMinutes { get; set; }

        /// <summary>
        /// The use previous parameters.
        /// </summary>
        [Input("usePreviousParameters")]
        public Input<bool>? UsePreviousParameters { get; set; }

        public ChangeSetArgs()
        {
        }
        public static new ChangeSetArgs Empty => new ChangeSetArgs();
    }

    public sealed class ChangeSetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.
        /// </summary>
        [Input("changeSetName")]
        public Input<string>? ChangeSetName { get; set; }

        /// <summary>
        /// The type of the change set. Valid values:  CREATE: creates a change set for a new stack. UPDATE: creates a change set for an existing stack. IMPORT: creates a change set for a new stack or an existing stack to import non-ROS-managed resources. If you create a change set for a new stack, ROS creates a stack that has a unique stack ID. The stack is in the REVIEW_IN_PROGRESS state until you execute the change set.  You cannot use the UPDATE type to create a change set for a new stack or the CREATE type to create a change set for an existing stack.
        /// </summary>
        [Input("changeSetType")]
        public Input<string>? ChangeSetType { get; set; }

        /// <summary>
        /// The description of the change set. The description can be up to 1,024 bytes in length.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Specifies whether to disable rollback on stack creation failure. Default value: false.  Valid values:  true: disables rollback on stack creation failure. false: enables rollback on stack creation failure. Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Input("disableRollback")]
        public Input<bool>? DisableRollback { get; set; }

        [Input("notificationUrls")]
        private InputList<string>? _notificationUrls;

        /// <summary>
        /// The notification urls.
        /// </summary>
        public InputList<string> NotificationUrls
        {
            get => _notificationUrls ?? (_notificationUrls = new InputList<string>());
            set => _notificationUrls = value;
        }

        [Input("parameters")]
        private InputList<Inputs.ChangeSetParameterGetArgs>? _parameters;

        /// <summary>
        /// Parameters.
        /// </summary>
        public InputList<Inputs.ChangeSetParameterGetArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.ChangeSetParameterGetArgs>());
            set => _parameters = value;
        }

        /// <summary>
        /// The ram role name.
        /// </summary>
        [Input("ramRoleName")]
        public Input<string>? RamRoleName { get; set; }

        /// <summary>
        /// The replacement option.
        /// </summary>
        [Input("replacementOption")]
        public Input<string>? ReplacementOption { get; set; }

        /// <summary>
        /// The ID of the stack for which you want to create the change set. ROS generates the change set by comparing the stack information with the information that you submit, such as a modified template or different inputs.
        /// </summary>
        [Input("stackId")]
        public Input<string>? StackId { get; set; }

        /// <summary>
        /// The name of the stack for which you want to create the change set.  The name can be up to 255 characters in length and can contain digits, letters, hyphens (-), and underscores (_). It must start with a digit or letter.  Note This parameter takes effect only when ChangeSetType is set to CREATE or IMPORT.
        /// </summary>
        [Input("stackName")]
        public Input<string>? StackName { get; set; }

        /// <summary>
        /// The stack policy body.
        /// </summary>
        [Input("stackPolicyBody")]
        public Input<string>? StackPolicyBody { get; set; }

        /// <summary>
        /// The stack policy during update body.
        /// </summary>
        [Input("stackPolicyDuringUpdateBody")]
        public Input<string>? StackPolicyDuringUpdateBody { get; set; }

        /// <summary>
        /// The stack policy during update url.
        /// </summary>
        [Input("stackPolicyDuringUpdateUrl")]
        public Input<string>? StackPolicyDuringUpdateUrl { get; set; }

        /// <summary>
        /// The stack policy url.
        /// </summary>
        [Input("stackPolicyUrl")]
        public Input<string>? StackPolicyUrl { get; set; }

        /// <summary>
        /// The status of the change set.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The structure that contains the template body. The template body must be 1 to 524,288 bytes in length.  If the length of the template body is longer than required, we recommend that you add parameters to the HTTP POST request body to avoid request failures due to excessive length of URLs.  You can specify one of TemplateBody or TemplateURL parameters, but you cannot specify both of them.
        /// </summary>
        [Input("templateBody")]
        public Input<string>? TemplateBody { get; set; }

        /// <summary>
        /// The template url.
        /// </summary>
        [Input("templateUrl")]
        public Input<string>? TemplateUrl { get; set; }

        /// <summary>
        /// Timeout In Minutes.
        /// </summary>
        [Input("timeoutInMinutes")]
        public Input<int>? TimeoutInMinutes { get; set; }

        /// <summary>
        /// The use previous parameters.
        /// </summary>
        [Input("usePreviousParameters")]
        public Input<bool>? UsePreviousParameters { get; set; }

        public ChangeSetState()
        {
        }
        public static new ChangeSetState Empty => new ChangeSetState();
    }
}
