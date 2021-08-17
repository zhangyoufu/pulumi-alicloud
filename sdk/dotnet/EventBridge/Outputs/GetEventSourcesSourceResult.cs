// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.EventBridge.Outputs
{

    [OutputType]
    public sealed class GetEventSourcesSourceResult
    {
        /// <summary>
        /// The detail describe of event source.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The code name of event source.
        /// </summary>
        public readonly string EventSourceName;
        /// <summary>
        /// The config of external data source.
        /// </summary>
        public readonly ImmutableDictionary<string, object> ExternalSourceConfig;
        /// <summary>
        /// The type of external data source.
        /// </summary>
        public readonly string ExternalSourceType;
        /// <summary>
        /// The ID of the Event Source.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Whether to connect to an external data source.
        /// </summary>
        public readonly bool LinkedExternalSource;
        public readonly string Type;

        [OutputConstructor]
        private GetEventSourcesSourceResult(
            string description,

            string eventSourceName,

            ImmutableDictionary<string, object> externalSourceConfig,

            string externalSourceType,

            string id,

            bool linkedExternalSource,

            string type)
        {
            Description = description;
            EventSourceName = eventSourceName;
            ExternalSourceConfig = externalSourceConfig;
            ExternalSourceType = externalSourceType;
            Id = id;
            LinkedExternalSource = linkedExternalSource;
            Type = type;
        }
    }
}
