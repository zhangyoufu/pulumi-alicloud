// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Arms
{
    public static class GetAlertContacts
    {
        /// <summary>
        /// This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.129.0+.
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
        ///     var ids = AliCloud.Arms.GetAlertContacts.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Arms.GetAlertContacts.Invoke(new()
        ///     {
        ///         NameRegex = "^my-AlertContact",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["armsAlertContactId1"] = ids.Apply(getAlertContactsResult =&gt; getAlertContactsResult.Contacts[0]?.Id),
        ///         ["armsAlertContactId2"] = nameRegex.Apply(getAlertContactsResult =&gt; getAlertContactsResult.Contacts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetAlertContactsResult> InvokeAsync(GetAlertContactsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAlertContactsResult>("alicloud:arms/getAlertContacts:getAlertContacts", args ?? new GetAlertContactsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides the Arms Alert Contacts of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.129.0+.
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
        ///     var ids = AliCloud.Arms.GetAlertContacts.Invoke();
        /// 
        ///     var nameRegex = AliCloud.Arms.GetAlertContacts.Invoke(new()
        ///     {
        ///         NameRegex = "^my-AlertContact",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["armsAlertContactId1"] = ids.Apply(getAlertContactsResult =&gt; getAlertContactsResult.Contacts[0]?.Id),
        ///         ["armsAlertContactId2"] = nameRegex.Apply(getAlertContactsResult =&gt; getAlertContactsResult.Contacts[0]?.Id),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetAlertContactsResult> Invoke(GetAlertContactsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAlertContactsResult>("alicloud:arms/getAlertContacts:getAlertContacts", args ?? new GetAlertContactsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAlertContactsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alert contact.
        /// </summary>
        [Input("alertContactName")]
        public string? AlertContactName { get; set; }

        /// <summary>
        /// The email address of the alert contact.
        /// </summary>
        [Input("email")]
        public string? Email { get; set; }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Alert Contact IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Alert Contact name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The mobile number of the alert contact.
        /// </summary>
        [Input("phoneNum")]
        public string? PhoneNum { get; set; }

        public GetAlertContactsArgs()
        {
        }
        public static new GetAlertContactsArgs Empty => new GetAlertContactsArgs();
    }

    public sealed class GetAlertContactsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the alert contact.
        /// </summary>
        [Input("alertContactName")]
        public Input<string>? AlertContactName { get; set; }

        /// <summary>
        /// The email address of the alert contact.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Alert Contact IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// A regex string to filter results by Alert Contact name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The mobile number of the alert contact.
        /// </summary>
        [Input("phoneNum")]
        public Input<string>? PhoneNum { get; set; }

        public GetAlertContactsInvokeArgs()
        {
        }
        public static new GetAlertContactsInvokeArgs Empty => new GetAlertContactsInvokeArgs();
    }


    [OutputType]
    public sealed class GetAlertContactsResult
    {
        public readonly string? AlertContactName;
        public readonly ImmutableArray<Outputs.GetAlertContactsContactResult> Contacts;
        public readonly string? Email;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> Ids;
        public readonly string? NameRegex;
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        public readonly string? PhoneNum;

        [OutputConstructor]
        private GetAlertContactsResult(
            string? alertContactName,

            ImmutableArray<Outputs.GetAlertContactsContactResult> contacts,

            string? email,

            string id,

            ImmutableArray<string> ids,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? phoneNum)
        {
            AlertContactName = alertContactName;
            Contacts = contacts;
            Email = email;
            Id = id;
            Ids = ids;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PhoneNum = phoneNum;
        }
    }
}
