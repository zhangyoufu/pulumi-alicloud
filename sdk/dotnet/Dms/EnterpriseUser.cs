// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Dms
{
    /// <summary>
    /// Provides a DMS Enterprise User resource. For information about Alidms Enterprise User and how to use it, see [What is Resource Alidms Enterprise User](https://www.alibabacloud.com/help/en/dms/developer-reference/api-dms-enterprise-2018-11-01-registeruser).
    /// 
    /// &gt; **NOTE:** Available since v1.90.0.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "tfexamplename";
    ///     var defaultUser = new AliCloud.Ram.User("defaultUser", new()
    ///     {
    ///         DisplayName = name,
    ///         Mobile = "86-18688888888",
    ///         Email = "hello.uuu@aaa.com",
    ///         Comments = "example",
    ///     });
    /// 
    ///     var defaultEnterpriseUser = new AliCloud.Dms.EnterpriseUser("defaultEnterpriseUser", new()
    ///     {
    ///         Uid = defaultUser.Id,
    ///         UserName = name,
    ///         RoleNames = new[]
    ///         {
    ///             "DBA",
    ///         },
    ///         Mobile = "86-18688888888",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// DMS Enterprise User can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:dms/enterpriseUser:EnterpriseUser example 24356xxx
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:dms/enterpriseUser:EnterpriseUser")]
    public partial class EnterpriseUser : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Maximum number of inquiries on the day.
        /// </summary>
        [Output("maxExecuteCount")]
        public Output<int?> MaxExecuteCount { get; private set; } = null!;

        /// <summary>
        /// Query the maximum number of rows on the day.
        /// </summary>
        [Output("maxResultCount")]
        public Output<int?> MaxResultCount { get; private set; } = null!;

        /// <summary>
        /// The DingTalk number or mobile number of the user.
        /// </summary>
        [Output("mobile")]
        public Output<string?> Mobile { get; private set; } = null!;

        /// <summary>
        /// It has been deprecated from 1.100.0 and use `user_name` instead.
        /// </summary>
        [Output("nickName")]
        public Output<string> NickName { get; private set; } = null!;

        /// <summary>
        /// The roles that the user plays.
        /// </summary>
        [Output("roleNames")]
        public Output<ImmutableArray<string>> RoleNames { get; private set; } = null!;

        /// <summary>
        /// The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
        /// </summary>
        [Output("status")]
        public Output<string?> Status { get; private set; } = null!;

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Output("tid")]
        public Output<int?> Tid { get; private set; } = null!;

        /// <summary>
        /// The Alibaba Cloud unique ID (UID) of the user to add.
        /// </summary>
        [Output("uid")]
        public Output<string> Uid { get; private set; } = null!;

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Output("userName")]
        public Output<string> UserName { get; private set; } = null!;


        /// <summary>
        /// Create a EnterpriseUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EnterpriseUser(string name, EnterpriseUserArgs args, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseUser:EnterpriseUser", name, args ?? new EnterpriseUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EnterpriseUser(string name, Input<string> id, EnterpriseUserState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:dms/enterpriseUser:EnterpriseUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EnterpriseUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EnterpriseUser Get(string name, Input<string> id, EnterpriseUserState? state = null, CustomResourceOptions? options = null)
        {
            return new EnterpriseUser(name, id, state, options);
        }
    }

    public sealed class EnterpriseUserArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of inquiries on the day.
        /// </summary>
        [Input("maxExecuteCount")]
        public Input<int>? MaxExecuteCount { get; set; }

        /// <summary>
        /// Query the maximum number of rows on the day.
        /// </summary>
        [Input("maxResultCount")]
        public Input<int>? MaxResultCount { get; set; }

        /// <summary>
        /// The DingTalk number or mobile number of the user.
        /// </summary>
        [Input("mobile")]
        public Input<string>? Mobile { get; set; }

        /// <summary>
        /// It has been deprecated from 1.100.0 and use `user_name` instead.
        /// </summary>
        [Input("nickName")]
        public Input<string>? NickName { get; set; }

        [Input("roleNames")]
        private InputList<string>? _roleNames;

        /// <summary>
        /// The roles that the user plays.
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _roleNames ?? (_roleNames = new InputList<string>());
            set => _roleNames = value;
        }

        /// <summary>
        /// The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Input("tid")]
        public Input<int>? Tid { get; set; }

        /// <summary>
        /// The Alibaba Cloud unique ID (UID) of the user to add.
        /// </summary>
        [Input("uid", required: true)]
        public Input<string> Uid { get; set; } = null!;

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public EnterpriseUserArgs()
        {
        }
        public static new EnterpriseUserArgs Empty => new EnterpriseUserArgs();
    }

    public sealed class EnterpriseUserState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Maximum number of inquiries on the day.
        /// </summary>
        [Input("maxExecuteCount")]
        public Input<int>? MaxExecuteCount { get; set; }

        /// <summary>
        /// Query the maximum number of rows on the day.
        /// </summary>
        [Input("maxResultCount")]
        public Input<int>? MaxResultCount { get; set; }

        /// <summary>
        /// The DingTalk number or mobile number of the user.
        /// </summary>
        [Input("mobile")]
        public Input<string>? Mobile { get; set; }

        /// <summary>
        /// It has been deprecated from 1.100.0 and use `user_name` instead.
        /// </summary>
        [Input("nickName")]
        public Input<string>? NickName { get; set; }

        [Input("roleNames")]
        private InputList<string>? _roleNames;

        /// <summary>
        /// The roles that the user plays.
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _roleNames ?? (_roleNames = new InputList<string>());
            set => _roleNames = value;
        }

        /// <summary>
        /// The state of DMS Enterprise User. Valid values: `NORMAL`, `DISABLE`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// The tenant ID.
        /// </summary>
        [Input("tid")]
        public Input<int>? Tid { get; set; }

        /// <summary>
        /// The Alibaba Cloud unique ID (UID) of the user to add.
        /// </summary>
        [Input("uid")]
        public Input<string>? Uid { get; set; }

        /// <summary>
        /// The nickname of the user.
        /// </summary>
        [Input("userName")]
        public Input<string>? UserName { get; set; }

        public EnterpriseUserState()
        {
        }
        public static new EnterpriseUserState Empty => new EnterpriseUserState();
    }
}
