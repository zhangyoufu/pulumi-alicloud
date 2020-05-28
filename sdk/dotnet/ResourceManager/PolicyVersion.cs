// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ResourceManager
{
    /// <summary>
    /// Provides a Resource Manager Policy Version resource. 
    /// For information about Resource Manager Policy Version and how to use it, see [What is Resource Manager Policy Version](https://www.alibabacloud.com/help/en/doc-detail/116817.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.84.0+.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var examplePolicy = new AliCloud.ResourceManager.Policy("examplePolicy", new AliCloud.ResourceManager.PolicyArgs
    ///         {
    ///             PolicyName = "tftest",
    ///             PolicyDocument = @"		{
    /// 			""Statement"": [{
    /// 				""Action"": [""oss:*""],
    /// 				""Effect"": ""Allow"",
    /// 				""Resource"": [""acs:oss:*:*:*""]
    /// 			}],
    /// 			""Version"": ""1""
    /// 		}
    /// ",
    ///         });
    ///         var examplePolicyVersion = new AliCloud.ResourceManager.PolicyVersion("examplePolicyVersion", new AliCloud.ResourceManager.PolicyVersionArgs
    ///         {
    ///             PolicyName = examplePolicy.PolicyName,
    ///             PolicyDocument = @"		{
    /// 			""Statement"": [{
    /// 				""Action"": [""oss:*""],
    /// 				""Effect"": ""Allow"",
    /// 				""Resource"": [""acs:oss:*:*:myphotos""]
    /// 			}],
    /// 			""Version"": ""1""
    /// 		}
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class PolicyVersion : Pulumi.CustomResource
    {
        /// <summary>
        /// The time when the policy version was created.
        /// </summary>
        [Output("createDate")]
        public Output<string> CreateDate { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to set the policy version as the default version. Default to `false`. 
        /// </summary>
        [Output("isDefaultVersion")]
        public Output<bool?> IsDefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The content of the policy. The content must be 1 to 2,048 characters in length.
        /// </summary>
        [Output("policyDocument")]
        public Output<string> PolicyDocument { get; private set; } = null!;

        /// <summary>
        /// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
        /// </summary>
        [Output("policyName")]
        public Output<string> PolicyName { get; private set; } = null!;

        /// <summary>
        /// The ID of the policy version.
        /// </summary>
        [Output("versionId")]
        public Output<string> VersionId { get; private set; } = null!;


        /// <summary>
        /// Create a PolicyVersion resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PolicyVersion(string name, PolicyVersionArgs args, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/policyVersion:PolicyVersion", name, args ?? new PolicyVersionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PolicyVersion(string name, Input<string> id, PolicyVersionState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:resourcemanager/policyVersion:PolicyVersion", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PolicyVersion resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PolicyVersion Get(string name, Input<string> id, PolicyVersionState? state = null, CustomResourceOptions? options = null)
        {
            return new PolicyVersion(name, id, state, options);
        }
    }

    public sealed class PolicyVersionArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether to set the policy version as the default version. Default to `false`. 
        /// </summary>
        [Input("isDefaultVersion")]
        public Input<bool>? IsDefaultVersion { get; set; }

        /// <summary>
        /// The content of the policy. The content must be 1 to 2,048 characters in length.
        /// </summary>
        [Input("policyDocument", required: true)]
        public Input<string> PolicyDocument { get; set; } = null!;

        /// <summary>
        /// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
        /// </summary>
        [Input("policyName", required: true)]
        public Input<string> PolicyName { get; set; } = null!;

        public PolicyVersionArgs()
        {
        }
    }

    public sealed class PolicyVersionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time when the policy version was created.
        /// </summary>
        [Input("createDate")]
        public Input<string>? CreateDate { get; set; }

        /// <summary>
        /// Specifies whether to set the policy version as the default version. Default to `false`. 
        /// </summary>
        [Input("isDefaultVersion")]
        public Input<bool>? IsDefaultVersion { get; set; }

        /// <summary>
        /// The content of the policy. The content must be 1 to 2,048 characters in length.
        /// </summary>
        [Input("policyDocument")]
        public Input<string>? PolicyDocument { get; set; }

        /// <summary>
        /// The name of the policy. Name must be 1 to 128 characters in length and can contain letters, digits, and hyphens (-).
        /// </summary>
        [Input("policyName")]
        public Input<string>? PolicyName { get; set; }

        /// <summary>
        /// The ID of the policy version.
        /// </summary>
        [Input("versionId")]
        public Input<string>? VersionId { get; set; }

        public PolicyVersionState()
        {
        }
    }
}
