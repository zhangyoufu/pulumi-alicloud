// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Oss.Inputs
{

    public sealed class BucketCorsRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedHeaders")]
        private InputList<string>? _allowedHeaders;

        /// <summary>
        /// Specifies which headers are allowed.
        /// </summary>
        public InputList<string> AllowedHeaders
        {
            get => _allowedHeaders ?? (_allowedHeaders = new InputList<string>());
            set => _allowedHeaders = value;
        }

        [Input("allowedMethods", required: true)]
        private InputList<string>? _allowedMethods;

        /// <summary>
        /// Specifies which methods are allowed. Can be GET, PUT, POST, DELETE or HEAD.
        /// </summary>
        public InputList<string> AllowedMethods
        {
            get => _allowedMethods ?? (_allowedMethods = new InputList<string>());
            set => _allowedMethods = value;
        }

        [Input("allowedOrigins", required: true)]
        private InputList<string>? _allowedOrigins;

        /// <summary>
        /// Specifies which origins are allowed.
        /// </summary>
        public InputList<string> AllowedOrigins
        {
            get => _allowedOrigins ?? (_allowedOrigins = new InputList<string>());
            set => _allowedOrigins = value;
        }

        [Input("exposeHeaders")]
        private InputList<string>? _exposeHeaders;

        /// <summary>
        /// Specifies expose header in the response.
        /// </summary>
        public InputList<string> ExposeHeaders
        {
            get => _exposeHeaders ?? (_exposeHeaders = new InputList<string>());
            set => _exposeHeaders = value;
        }

        /// <summary>
        /// Specifies time in seconds that browser can cache the response for a preflight request.
        /// </summary>
        [Input("maxAgeSeconds")]
        public Input<int>? MaxAgeSeconds { get; set; }

        public BucketCorsRuleArgs()
        {
        }
        public static new BucketCorsRuleArgs Empty => new BucketCorsRuleArgs();
    }
}
