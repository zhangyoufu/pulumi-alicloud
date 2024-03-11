// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.SecurityCenter
{
    /// <summary>
    /// Using this resource can create SecurityCenter service-linked role : `AliyunServiceRolePolicyForSas`.  This Role is a Resource Access Management (RAM) role, which to obtain permissions to access another Alibaba Cloud service.
    /// 
    /// For information about Security Center Service Role and how to use it, see [What is Security Center](https://www.alibabacloud.com/help/en/doc-detail/42302.htm).
    /// 
    /// &gt; **NOTE:** Available in v1.142.0+.
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
    ///     var serviceLinkedRole = new AliCloud.SecurityCenter.ServiceLinkedRole("serviceLinkedRole");
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// SecurityCenter service-linked roles(SLR) can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole example &lt;product_name&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole")]
    public partial class ServiceLinkedRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
        /// </summary>
        [Output("status")]
        public Output<bool> Status { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceLinkedRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceLinkedRole(string name, ServiceLinkedRoleArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole", name, args ?? new ServiceLinkedRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceLinkedRole(string name, Input<string> id, ServiceLinkedRoleState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:securitycenter/serviceLinkedRole:ServiceLinkedRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceLinkedRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceLinkedRole Get(string name, Input<string> id, ServiceLinkedRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceLinkedRole(name, id, state, options);
        }
    }

    public sealed class ServiceLinkedRoleArgs : global::Pulumi.ResourceArgs
    {
        public ServiceLinkedRoleArgs()
        {
        }
        public static new ServiceLinkedRoleArgs Empty => new ServiceLinkedRoleArgs();
    }

    public sealed class ServiceLinkedRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The status of the service Associated role. Valid Values: `true`: Created. `false`: not created.
        /// </summary>
        [Input("status")]
        public Input<bool>? Status { get; set; }

        public ServiceLinkedRoleState()
        {
        }
        public static new ServiceLinkedRoleState Empty => new ServiceLinkedRoleState();
    }
}
