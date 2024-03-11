// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.ActionTrail
{
    public static class GetGlobalEventsStorageRegion
    {
        /// <summary>
        /// This data source provides the Actiontrail Global Events Storage Region of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.201.0+.
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
        ///     var @default = AliCloud.ActionTrail.GetGlobalEventsStorageRegion.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudActiontrailGlobalEventsStorageRegion1"] = @default.Apply(@default =&gt; @default.Apply(getGlobalEventsStorageRegionResult =&gt; getGlobalEventsStorageRegionResult.StorageRegion)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetGlobalEventsStorageRegionResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetGlobalEventsStorageRegionResult>("alicloud:actiontrail/getGlobalEventsStorageRegion:getGlobalEventsStorageRegion", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// This data source provides the Actiontrail Global Events Storage Region of the current Alibaba Cloud user.
        /// 
        /// &gt; **NOTE:** Available in v1.201.0+.
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
        ///     var @default = AliCloud.ActionTrail.GetGlobalEventsStorageRegion.Invoke();
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudActiontrailGlobalEventsStorageRegion1"] = @default.Apply(@default =&gt; @default.Apply(getGlobalEventsStorageRegionResult =&gt; getGlobalEventsStorageRegionResult.StorageRegion)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetGlobalEventsStorageRegionResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetGlobalEventsStorageRegionResult>("alicloud:actiontrail/getGlobalEventsStorageRegion:getGlobalEventsStorageRegion", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetGlobalEventsStorageRegionResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string StorageRegion;

        [OutputConstructor]
        private GetGlobalEventsStorageRegionResult(
            string id,

            string storageRegion)
        {
            Id = id;
            StorageRegion = storageRegion;
        }
    }
}
