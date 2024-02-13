// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CloudSso
{
    /// <summary>
    /// Provides a Cloud SSO Directory resource.
    /// 
    /// For information about Cloud SSO Directory and how to use it, see [What is Directory](https://www.alibabacloud.com/help/en/cloudsso/latest/api-cloudsso-2021-05-15-createdirectory).
    /// 
    /// &gt; **NOTE:** Available since v1.135.0.
    /// 
    /// &gt; **NOTE:** Cloud SSO Only Support `cn-shanghai` And `us-west-1` Region
    /// 
    /// ## Import
    /// 
    /// Cloud SSO Directory can be imported using the id, e.g.
    /// 
    /// ```sh
    /// $ pulumi import alicloud:cloudsso/directory:Directory example &lt;id&gt;
    /// ```
    /// </summary>
    [AliCloudResourceType("alicloud:cloudsso/directory:Directory")]
    public partial class Directory : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        /// </summary>
        [Output("directoryName")]
        public Output<string?> DirectoryName { get; private set; } = null!;

        /// <summary>
        /// The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        /// </summary>
        [Output("mfaAuthenticationStatus")]
        public Output<string> MfaAuthenticationStatus { get; private set; } = null!;

        /// <summary>
        /// The saml identity provider configuration. See `saml_identity_provider_configuration` below.
        /// 
        /// &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        /// </summary>
        [Output("samlIdentityProviderConfiguration")]
        public Output<Outputs.DirectorySamlIdentityProviderConfiguration> SamlIdentityProviderConfiguration { get; private set; } = null!;

        /// <summary>
        /// The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        /// </summary>
        [Output("scimSynchronizationStatus")]
        public Output<string> ScimSynchronizationStatus { get; private set; } = null!;


        /// <summary>
        /// Create a Directory resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Directory(string name, DirectoryArgs? args = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudsso/directory:Directory", name, args ?? new DirectoryArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Directory(string name, Input<string> id, DirectoryState? state = null, CustomResourceOptions? options = null)
            : base("alicloud:cloudsso/directory:Directory", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Directory resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Directory Get(string name, Input<string> id, DirectoryState? state = null, CustomResourceOptions? options = null)
        {
            return new Directory(name, id, state, options);
        }
    }

    public sealed class DirectoryArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        /// </summary>
        [Input("directoryName")]
        public Input<string>? DirectoryName { get; set; }

        /// <summary>
        /// The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        /// </summary>
        [Input("mfaAuthenticationStatus")]
        public Input<string>? MfaAuthenticationStatus { get; set; }

        /// <summary>
        /// The saml identity provider configuration. See `saml_identity_provider_configuration` below.
        /// 
        /// &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        /// </summary>
        [Input("samlIdentityProviderConfiguration")]
        public Input<Inputs.DirectorySamlIdentityProviderConfigurationArgs>? SamlIdentityProviderConfiguration { get; set; }

        /// <summary>
        /// The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        /// </summary>
        [Input("scimSynchronizationStatus")]
        public Input<string>? ScimSynchronizationStatus { get; set; }

        public DirectoryArgs()
        {
        }
        public static new DirectoryArgs Empty => new DirectoryArgs();
    }

    public sealed class DirectoryState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the CloudSSO directory. The length is 2-64 characters, and it can contain lowercase letters, numbers, and dashes (-). It cannot start or end with a dash and cannot have two consecutive dashes. Need to be globally unique, and capitalization is not supported. Cannot start with `d-`.
        /// </summary>
        [Input("directoryName")]
        public Input<string>? DirectoryName { get; set; }

        /// <summary>
        /// The mfa authentication status. Valid values: `Enabled` or `Disabled`. Default to `Enabled`.
        /// </summary>
        [Input("mfaAuthenticationStatus")]
        public Input<string>? MfaAuthenticationStatus { get; set; }

        /// <summary>
        /// The saml identity provider configuration. See `saml_identity_provider_configuration` below.
        /// 
        /// &gt; **NOTE:** The `saml_identity_provider_configuration` will be removed automatically when the resource is deleted, please operate with caution. If there are left more configuration in the directory, please remove them before deleting the directory.
        /// </summary>
        [Input("samlIdentityProviderConfiguration")]
        public Input<Inputs.DirectorySamlIdentityProviderConfigurationGetArgs>? SamlIdentityProviderConfiguration { get; set; }

        /// <summary>
        /// The scim synchronization status. Valid values: `Enabled` or `Disabled`. Default to `Disabled`.
        /// </summary>
        [Input("scimSynchronizationStatus")]
        public Input<string>? ScimSynchronizationStatus { get; set; }

        public DirectoryState()
        {
        }
        public static new DirectoryState Empty => new DirectoryState();
    }
}
