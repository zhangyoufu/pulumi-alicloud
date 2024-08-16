// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Ecs.Outputs
{

    [OutputType]
    public sealed class GetCapacityReservationsReservationResult
    {
        /// <summary>
        /// Capacity Reservation id
        /// </summary>
        public readonly string CapacityReservationId;
        /// <summary>
        /// Capacity reservation service name.
        /// </summary>
        public readonly string CapacityReservationName;
        /// <summary>
        /// description of the capacity reservation instance
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// end time of the capacity reservation. the capacity reservation will be  released at the end time automatically if set. otherwise it will last until manually released
        /// </summary>
        public readonly string EndTime;
        /// <summary>
        /// Release mode of capacity reservation service. Value range:Limited: release at specified time. The EndTime parameter must be specified at the same time.Unlimited: manual release. No time limit.
        /// </summary>
        public readonly string EndTimeType;
        /// <summary>
        /// The ID of the Capacity Reservation.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The total number of instances that need to be reserved within the capacity reservation
        /// </summary>
        public readonly string InstanceAmount;
        /// <summary>
        /// Instance type. Currently, you can only set the capacity reservation service for one instance type.
        /// </summary>
        public readonly string InstanceType;
        /// <summary>
        /// The type of private resource pool generated after the capacity reservation service takes effect. Value range:Open: Open mode.Target: dedicated mode.Default value: Open
        /// </summary>
        public readonly string MatchCriteria;
        /// <summary>
        /// The payment type of the resource. value range `PostPaid`, `PrePaid`.
        /// </summary>
        public readonly string PaymentType;
        /// <summary>
        /// platform of the capacity reservation , value range `windows`, `linux`, `all`.
        /// </summary>
        public readonly string Platform;
        /// <summary>
        /// The resource group id.
        /// </summary>
        public readonly string ResourceGroupId;
        /// <summary>
        /// time of the capacity reservation which become active
        /// </summary>
        public readonly string StartTime;
        /// <summary>
        /// The capacity is scheduled to take effect. Possible values:-Now: Effective immediately.-Later: the specified time takes effect.
        /// </summary>
        public readonly string StartTimeType;
        /// <summary>
        /// The status of the capacity reservation. value range `All`, `Pending`, `Preparing`, `Prepared`, `Active`, `Released`.
        /// </summary>
        public readonly string Status;
        /// <summary>
        /// The tag of the resource.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Tags;
        /// <summary>
        /// This parameter is under test and is not yet open for use.
        /// </summary>
        public readonly string TimeSlot;
        /// <summary>
        /// The ID of the zone in the region to which the capacity reservation service belongs. Currently, it is only supported to create a capacity reservation service in one zone.
        /// </summary>
        public readonly ImmutableArray<string> ZoneIds;

        [OutputConstructor]
        private GetCapacityReservationsReservationResult(
            string capacityReservationId,

            string capacityReservationName,

            string description,

            string endTime,

            string endTimeType,

            string id,

            string instanceAmount,

            string instanceType,

            string matchCriteria,

            string paymentType,

            string platform,

            string resourceGroupId,

            string startTime,

            string startTimeType,

            string status,

            ImmutableDictionary<string, string>? tags,

            string timeSlot,

            ImmutableArray<string> zoneIds)
        {
            CapacityReservationId = capacityReservationId;
            CapacityReservationName = capacityReservationName;
            Description = description;
            EndTime = endTime;
            EndTimeType = endTimeType;
            Id = id;
            InstanceAmount = instanceAmount;
            InstanceType = instanceType;
            MatchCriteria = matchCriteria;
            PaymentType = paymentType;
            Platform = platform;
            ResourceGroupId = resourceGroupId;
            StartTime = startTime;
            StartTimeType = startTimeType;
            Status = status;
            Tags = tags;
            TimeSlot = timeSlot;
            ZoneIds = zoneIds;
        }
    }
}
