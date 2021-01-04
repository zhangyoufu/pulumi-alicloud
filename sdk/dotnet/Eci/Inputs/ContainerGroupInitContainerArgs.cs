// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.AliCloud.Eci.Inputs
{

    public sealed class ContainerGroupInitContainerArgs : Pulumi.ResourceArgs
    {
        [Input("args")]
        private InputList<string>? _args;

        /// <summary>
        /// The arguments passed to the commands.
        /// </summary>
        public InputList<string> Args
        {
            get => _args ?? (_args = new InputList<string>());
            set => _args = value;
        }

        [Input("commands")]
        private InputList<string>? _commands;

        /// <summary>
        /// The commands run by the init container.
        /// </summary>
        public InputList<string> Commands
        {
            get => _commands ?? (_commands = new InputList<string>());
            set => _commands = value;
        }

        /// <summary>
        /// The amount of CPU resources allocated to the container.
        /// </summary>
        [Input("cpu")]
        public Input<double>? Cpu { get; set; }

        [Input("environmentVars")]
        private InputList<Inputs.ContainerGroupInitContainerEnvironmentVarArgs>? _environmentVars;

        /// <summary>
        /// The structure of environmentVars.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerEnvironmentVarArgs> EnvironmentVars
        {
            get => _environmentVars ?? (_environmentVars = new InputList<Inputs.ContainerGroupInitContainerEnvironmentVarArgs>());
            set => _environmentVars = value;
        }

        /// <summary>
        /// The number GPUs.
        /// </summary>
        [Input("gpu")]
        public Input<int>? Gpu { get; set; }

        /// <summary>
        /// The image of the container.
        /// </summary>
        [Input("image")]
        public Input<string>? Image { get; set; }

        /// <summary>
        /// The restart policy of the image.
        /// </summary>
        [Input("imagePullPolicy")]
        public Input<string>? ImagePullPolicy { get; set; }

        /// <summary>
        /// The amount of memory resources allocated to the container.
        /// </summary>
        [Input("memory")]
        public Input<double>? Memory { get; set; }

        /// <summary>
        /// The name of the mounted volume.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ports")]
        private InputList<Inputs.ContainerGroupInitContainerPortArgs>? _ports;

        /// <summary>
        /// The structure of port.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerPortArgs> Ports
        {
            get => _ports ?? (_ports = new InputList<Inputs.ContainerGroupInitContainerPortArgs>());
            set => _ports = value;
        }

        [Input("ready")]
        public Input<bool>? Ready { get; set; }

        [Input("restartCount")]
        public Input<int>? RestartCount { get; set; }

        [Input("volumeMounts")]
        private InputList<Inputs.ContainerGroupInitContainerVolumeMountArgs>? _volumeMounts;

        /// <summary>
        /// The structure of volumeMounts.
        /// </summary>
        public InputList<Inputs.ContainerGroupInitContainerVolumeMountArgs> VolumeMounts
        {
            get => _volumeMounts ?? (_volumeMounts = new InputList<Inputs.ContainerGroupInitContainerVolumeMountArgs>());
            set => _volumeMounts = value;
        }

        /// <summary>
        /// The working directory of the container.
        /// </summary>
        [Input("workingDir")]
        public Input<string>? WorkingDir { get; set; }

        public ContainerGroupInitContainerArgs()
        {
        }
    }
}
