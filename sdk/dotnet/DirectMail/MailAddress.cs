// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.DirectMail
{
    /// <summary>
    /// Provides a Direct Mail Mail Address resource.
    /// 
    /// For information about Direct Mail Mail Address and how to use it, see [What is Mail Address](https://www.aliyun.com/product/directmail).
    /// 
    /// &gt; **NOTE:** Available in v1.134.0+.
    /// 
    /// ## Example Usage
    /// 
    /// Basic Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using AliCloud = Pulumi.AliCloud;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new AliCloud.DirectMail.MailAddress("example", new AliCloud.DirectMail.MailAddressArgs
    ///         {
    ///             AccountName = "example_value@email.com",
    ///             Sendtype = "batch",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// &gt; **Note:**
    /// A maximum of 10 mailing addresses can be added.
    /// Individual users: Up to 10 mailing addresses can be deleted within a month.
    /// Enterprise users: Up to 10 mailing addresses can be deleted within a month.
    /// 
    /// ## Import
    /// 
    /// Direct Mail Mail Address can be imported using the id, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import alicloud:directmail/mailAddress:MailAddress example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:directmail/mailAddress:MailAddress")]
    public partial class MailAddress : Pulumi.CustomResource
    {
        /// <summary>
        /// The sender address. The email address must be filled in the format of account@domain, and only lowercase letters or numbers can be used.
        /// </summary>
        [Output("accountName")]
        public Output<string> AccountName { get; private set; } = null!;

        /// <summary>
        /// Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// Return address.
        /// </summary>
        [Output("replyAddress")]
        public Output<string?> ReplyAddress { get; private set; } = null!;

        /// <summary>
        /// Account type. Valid values: `batch`, `trigger`.
        /// </summary>
        [Output("sendtype")]
        public Output<string> Sendtype { get; private set; } = null!;

        /// <summary>
        /// Account Status freeze: 1, normal: 0.
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;


        /// <summary>
        /// Create a MailAddress resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MailAddress(string name, MailAddressArgs args, CustomResourceOptions? options = null)
            : base("alicloud:directmail/mailAddress:MailAddress", name, args ?? new MailAddressArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MailAddress(string name, Input<string> id, MailAddressState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:directmail/mailAddress:MailAddress", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MailAddress resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MailAddress Get(string name, Input<string> id, MailAddressState? state = null, CustomResourceOptions? options = null)
        {
            return new MailAddress(name, id, state, options);
        }
    }

    public sealed class MailAddressArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sender address. The email address must be filled in the format of account@domain, and only lowercase letters or numbers can be used.
        /// </summary>
        [Input("accountName", required: true)]
        public Input<string> AccountName { get; set; } = null!;

        /// <summary>
        /// Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Return address.
        /// </summary>
        [Input("replyAddress")]
        public Input<string>? ReplyAddress { get; set; }

        /// <summary>
        /// Account type. Valid values: `batch`, `trigger`.
        /// </summary>
        [Input("sendtype", required: true)]
        public Input<string> Sendtype { get; set; } = null!;

        public MailAddressArgs()
        {
        }
    }

    public sealed class MailAddressState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The sender address. The email address must be filled in the format of account@domain, and only lowercase letters or numbers can be used.
        /// </summary>
        [Input("accountName")]
        public Input<string>? AccountName { get; set; }

        /// <summary>
        /// Account password. The password must be length 10-20 string, contains numbers, uppercase letters, lowercase letters at the same time.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// Return address.
        /// </summary>
        [Input("replyAddress")]
        public Input<string>? ReplyAddress { get; set; }

        /// <summary>
        /// Account type. Valid values: `batch`, `trigger`.
        /// </summary>
        [Input("sendtype")]
        public Input<string>? Sendtype { get; set; }

        /// <summary>
        /// Account Status freeze: 1, normal: 0.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        public MailAddressState()
        {
        }
    }
}
