// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs
{
    public static class GetCapacityReservations
    {
        /// <summary>
        /// This data source provides Ecs Capacity Reservation available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
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
        ///     var @default = AliCloud.Ecs.GetCapacityReservations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_ecs_capacity_reservation.Default.Id,
        ///         },
        ///         NameRegex = alicloud_ecs_capacity_reservation.Default.Name,
        ///         InstanceType = "ecs.c6.large",
        ///         Platform = "linux",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudEcsCapacityReservationExampleId"] = @default.Apply(@default =&gt; @default.Apply(getCapacityReservationsResult =&gt; getCapacityReservationsResult.Reservations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetCapacityReservationsResult> InvokeAsync(GetCapacityReservationsArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCapacityReservationsResult>("alicloud:ecs/getCapacityReservations:getCapacityReservations", args ?? new GetCapacityReservationsArgs(), options.WithDefaults());

        /// <summary>
        /// This data source provides Ecs Capacity Reservation available to the user.
        /// 
        /// &gt; **NOTE:** Available in 1.195.0+
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
        ///     var @default = AliCloud.Ecs.GetCapacityReservations.Invoke(new()
        ///     {
        ///         Ids = new[]
        ///         {
        ///             alicloud_ecs_capacity_reservation.Default.Id,
        ///         },
        ///         NameRegex = alicloud_ecs_capacity_reservation.Default.Name,
        ///         InstanceType = "ecs.c6.large",
        ///         Platform = "linux",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["alicloudEcsCapacityReservationExampleId"] = @default.Apply(@default =&gt; @default.Apply(getCapacityReservationsResult =&gt; getCapacityReservationsResult.Reservations[0]?.Id)),
        ///     };
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetCapacityReservationsResult> Invoke(GetCapacityReservationsInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCapacityReservationsResult>("alicloud:ecs/getCapacityReservations:getCapacityReservations", args ?? new GetCapacityReservationsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCapacityReservationsArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacityReservationIds")]
        private List<string>? _capacityReservationIds;
        public List<string> CapacityReservationIds
        {
            get => _capacityReservationIds ?? (_capacityReservationIds = new List<string>());
            set => _capacityReservationIds = value;
        }

        [Input("ids")]
        private List<string>? _ids;

        /// <summary>
        /// A list of Capacity Reservation IDs.
        /// </summary>
        public List<string> Ids
        {
            get => _ids ?? (_ids = new List<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        [Input("instanceType")]
        public string? InstanceType { get; set; }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public string? OutputFile { get; set; }

        /// <summary>
        /// The payment type of the resource. value range `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("paymentType")]
        public string? PaymentType { get; set; }

        /// <summary>
        /// platform of the capacity reservation , value range `windows`, `linux`, `all`.
        /// </summary>
        [Input("platform")]
        public string? Platform { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public string? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the capacity reservation. value range `All`, `Pending`, `Preparing`, `Prepared`, `Active`, `Released`.
        /// </summary>
        [Input("status")]
        public string? Status { get; set; }

        [Input("tags")]
        private Dictionary<string, object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public Dictionary<string, object> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, object>());
            set => _tags = value;
        }

        public GetCapacityReservationsArgs()
        {
        }
        public static new GetCapacityReservationsArgs Empty => new GetCapacityReservationsArgs();
    }

    public sealed class GetCapacityReservationsInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("capacityReservationIds")]
        private InputList<string>? _capacityReservationIds;
        public InputList<string> CapacityReservationIds
        {
            get => _capacityReservationIds ?? (_capacityReservationIds = new InputList<string>());
            set => _capacityReservationIds = value;
        }

        [Input("ids")]
        private InputList<string>? _ids;

        /// <summary>
        /// A list of Capacity Reservation IDs.
        /// </summary>
        public InputList<string> Ids
        {
            get => _ids ?? (_ids = new InputList<string>());
            set => _ids = value;
        }

        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        [Input("instanceType")]
        public Input<string>? InstanceType { get; set; }

        /// <summary>
        /// A regex string to filter results by Group Metric Rule name.
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        /// <summary>
        /// File name where to save data source results (after running `pulumi preview`).
        /// </summary>
        [Input("outputFile")]
        public Input<string>? OutputFile { get; set; }

        /// <summary>
        /// The payment type of the resource. value range `PostPaid`, `PrePaid`.
        /// </summary>
        [Input("paymentType")]
        public Input<string>? PaymentType { get; set; }

        /// <summary>
        /// platform of the capacity reservation , value range `windows`, `linux`, `all`.
        /// </summary>
        [Input("platform")]
        public Input<string>? Platform { get; set; }

        /// <summary>
        /// The resource group id.
        /// </summary>
        [Input("resourceGroupId")]
        public Input<string>? ResourceGroupId { get; set; }

        /// <summary>
        /// The status of the capacity reservation. value range `All`, `Pending`, `Preparing`, `Prepared`, `Active`, `Released`.
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputMap<object>? _tags;

        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public InputMap<object> Tags
        {
            get => _tags ?? (_tags = new InputMap<object>());
            set => _tags = value;
        }

        public GetCapacityReservationsInvokeArgs()
        {
        }
        public static new GetCapacityReservationsInvokeArgs Empty => new GetCapacityReservationsInvokeArgs();
    }


    [OutputType]
    public sealed class GetCapacityReservationsResult
    {
        public readonly ImmutableArray<string> CapacityReservationIds;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A list of Capacity Reservation IDs.
        /// </summary>
        public readonly ImmutableArray<string> Ids;
        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        public readonly string? InstanceType;
        public readonly string? NameRegex;
        /// <summary>
        /// A list of name of Capacity Reservations.
        /// </summary>
        public readonly ImmutableArray<string> Names;
        public readonly string? OutputFile;
        /// <summary>
        /// The payment type of the resource
        /// </summary>
        public readonly string? PaymentType;
        /// <summary>
        /// platform of the capacity reservation.
        /// </summary>
        public readonly string? Platform;
        /// <summary>
        /// A list of Capacity Reservation Entries. Each element contains the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCapacityReservationsReservationResult> Reservations;
        /// <summary>
        /// The resource group id
        /// </summary>
        public readonly string? ResourceGroupId;
        /// <summary>
        /// The status of the capacity reservation.
        /// </summary>
        public readonly string? Status;
        /// <summary>
        /// A mapping of tags to assign to the Capacity Reservation.
        /// </summary>
        public readonly ImmutableDictionary<string, object>? Tags;

        [OutputConstructor]
        private GetCapacityReservationsResult(
            ImmutableArray<string> capacityReservationIds,

            string id,

            ImmutableArray<string> ids,

            string? instanceType,

            string? nameRegex,

            ImmutableArray<string> names,

            string? outputFile,

            string? paymentType,

            string? platform,

            ImmutableArray<Outputs.GetCapacityReservationsReservationResult> reservations,

            string? resourceGroupId,

            string? status,

            ImmutableDictionary<string, object>? tags)
        {
            CapacityReservationIds = capacityReservationIds;
            Id = id;
            Ids = ids;
            InstanceType = instanceType;
            NameRegex = nameRegex;
            Names = names;
            OutputFile = outputFile;
            PaymentType = paymentType;
            Platform = platform;
            Reservations = reservations;
            ResourceGroupId = resourceGroupId;
            Status = status;
            Tags = tags;
        }
    }
}
