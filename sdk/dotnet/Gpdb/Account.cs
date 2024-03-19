// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Gpdb
{
    /// <summary>
    /// Provides a GPDB Account resource.
    /// 
    /// For information about GPDB Account and how to use it, see [What is Account](https://www.alibabacloud.com/help/doc-detail/86924.htm).
    /// 
    /// &gt; **NOTE:** Available since v1.142.0.
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
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var config = new Config();
    ///     var name = config.Get("name") ?? "terraform-example";
    ///     var defaultZones = AliCloud.Gpdb.GetZones.Invoke();
    /// 
    ///     var defaultNetworks = AliCloud.Vpc.GetNetworks.Invoke(new()
    ///     {
    ///         NameRegex = "^default-NODELETING$",
    ///     });
    /// 
    ///     var defaultSwitches = AliCloud.Vpc.GetSwitches.Invoke(new()
    ///     {
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///     });
    /// 
    ///     var defaultInstance = new AliCloud.Gpdb.Instance("defaultInstance", new()
    ///     {
    ///         DbInstanceCategory = "HighAvailability",
    ///         DbInstanceClass = "gpdb.group.segsdx1",
    ///         DbInstanceMode = "StorageElastic",
    ///         Description = name,
    ///         Engine = "gpdb",
    ///         EngineVersion = "6.0",
    ///         ZoneId = defaultZones.Apply(getZonesResult =&gt; getZonesResult.Ids[0]),
    ///         InstanceNetworkType = "VPC",
    ///         InstanceSpec = "2C16G",
    ///         MasterNodeNum = 1,
    ///         PaymentType = "PayAsYouGo",
    ///         PrivateIpAddress = "1.1.1.1",
    ///         SegStorageType = "cloud_essd",
    ///         SegNodeNum = 4,
    ///         StorageSize = 50,
    ///         VpcId = defaultNetworks.Apply(getNetworksResult =&gt; getNetworksResult.Ids[0]),
    ///         VswitchId = defaultSwitches.Apply(getSwitchesResult =&gt; getSwitchesResult.Ids[0]),
    ///         IpWhitelists = new[]
    ///         {
    ///             new AliCloud.Gpdb.Inputs.InstanceIpWhitelistArgs
    ///             {
    ///                 SecurityIpList = "127.0.0.1",
    ///             },
    ///         },
    ///     });
    /// 
    ///     var defaultAccount = new AliCloud.Gpdb.Account("defaultAccount", new()
    ///     {
    ///         AccountName = "tf_example",
    ///         DbInstanceId = defaultInstance.Id,
    ///         AccountPassword = "Example1234",
    ///         AccountDescription = "tf_example",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// GPDB Account can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:gpdb/account:Account example &lt;db_instance_id&gt;:&lt;account_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:gpdb/account:Account")]
    public partial class Account : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the account.
        /// * Starts with a letter.
        /// * Does not start with `http://` or `https://`.
        /// * Contains letters, underscores (_), hyphens (-), or digits.
        /// * Be 2 to 256 characters in length.
        /// </summary>
        [Output("accountDescription")]
        public Output<string?> AccountDescription { get; private set; } = null!;

        /// <summary>
        /// The name of the account. The account name must be unique and meet the following requirements:
        /// * Starts with a letter.
        /// * Contains only lowercase letters, digits, or underscores (_).
        /// * Be up to 16 characters in length.
        /// * Contains no reserved keywords.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
        /// </summary>
        [Output("accountPassword")]
        public Output<string> AccountPassword { get; private set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Output("dbInstanceId")]
        public Output<string> DbInstanceId { get; private set; } = null!;

        /// <summary>
        /// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a Account resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Account(string name, AccountArgs args, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/account:Account", name, args ?? new AccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Account(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:gpdb/account:Account", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Account resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Account Get(string name, Input<string> id, AccountState? state = null, CustomResourceOptions? options = null)
        {
            return new Account(name, id, state, options);
        }
    }

    public sealed class AccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// * Starts with a letter.
        /// * Does not start with `http://` or `https://`.
        /// * Contains letters, underscores (_), hyphens (-), or digits.
        /// * Be 2 to 256 characters in length.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account. The account name must be unique and meet the following requirements:
        /// * Starts with a letter.
        /// * Contains only lowercase letters, digits, or underscores (_).
        /// * Be up to 16 characters in length.
        /// * Contains no reserved keywords.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
        /// </summary>
        [Input("accountPassword", required: true)]
        public Input<string> AccountPassword { get; set; } = null!;

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId", required: true)]
        public Input<string> DbInstanceId { get; set; } = null!;

        public AccountArgs()
        {
        }
        public static new AccountArgs Empty => new AccountArgs();
    }

    public sealed class AccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the account.
        /// * Starts with a letter.
        /// * Does not start with `http://` or `https://`.
        /// * Contains letters, underscores (_), hyphens (-), or digits.
        /// * Be 2 to 256 characters in length.
        /// </summary>
        [Input("accountDescription")]
        public Input<string>? AccountDescription { get; set; }

        /// <summary>
        /// The name of the account. The account name must be unique and meet the following requirements:
        /// * Starts with a letter.
        /// * Contains only lowercase letters, digits, or underscores (_).
        /// * Be up to 16 characters in length.
        /// * Contains no reserved keywords.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// The password of the account. The password must be 8 to 32 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include `! @ # $ %!^(MISSING) &amp; * ( ) _ + - =`.
        /// </summary>
        [Input("accountPassword")]
        public Input<string>? AccountPassword { get; set; }

        /// <summary>
        /// The ID of the instance.
        /// </summary>
        [Input("dbInstanceId")]
        public Input<string>? DbInstanceId { get; set; }

        /// <summary>
        /// The status of the account. Valid values: `Active`, `Creating` and `Deleting`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public AccountState()
        {
        }
        public static new AccountState Empty => new AccountState();
    }
}
