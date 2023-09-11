// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.CR.Inputs
{

    public sealed class ChainChainConfigNodeNodeConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("denyPolicies")]
        private InputList<Inputs.ChainChainConfigNodeNodeConfigDenyPolicyGetArgs>? _denyPolicies;

        /// <summary>
        /// Blocking rules for scanning nodes in delivery chain nodes. See `deny_policy` below. **Note:** When `node_name` is `VULNERABILITY_SCANNING`, the parameters in `deny_policy` need to be filled in.
        /// </summary>
        public InputList<Inputs.ChainChainConfigNodeNodeConfigDenyPolicyGetArgs> DenyPolicies
        {
            get => _denyPolicies ?? (_denyPolicies = new InputList<Inputs.ChainChainConfigNodeNodeConfigDenyPolicyGetArgs>());
            set => _denyPolicies = value;
        }

        public ChainChainConfigNodeNodeConfigGetArgs()
        {
        }
        public static new ChainChainConfigNodeNodeConfigGetArgs Empty => new ChainChainConfigNodeNodeConfigGetArgs();
    }
}
