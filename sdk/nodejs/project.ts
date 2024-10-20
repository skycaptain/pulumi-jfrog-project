// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides an Artifactory project resource. This can be used to create and manage Artifactory project, maintain users/groups/roles/repos.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as jfrog_project from "@pulumi/jfrog-project";
 *
 * const myproject = new jfrog_project.Project("myproject", {
 *     adminPrivileges: [{
 *         indexResources: true,
 *         manageMembers: true,
 *         manageResources: true,
 *     }],
 *     blockDeploymentsOnLimit: false,
 *     description: "My Project",
 *     displayName: "My Project",
 *     emailNotification: true,
 *     key: "myproj",
 *     maxStorageInGibibytes: 10,
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import jfrog-project:index/project:Project myproject myproj
 * ```
 */
export class Project extends pulumi.CustomResource {
    /**
     * Get an existing Project resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ProjectState, opts?: pulumi.CustomResourceOptions): Project {
        return new Project(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'jfrog-project:index/project:Project';

    /**
     * Returns true if the given object is an instance of Project.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Project {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Project.__pulumiType;
    }

    public readonly adminPrivileges!: pulumi.Output<outputs.ProjectAdminPrivilege[] | undefined>;
    /**
     * Block deployment of artifacts if storage quota is exceeded.
     */
    public readonly blockDeploymentsOnLimit!: pulumi.Output<boolean>;
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Also known as project name on the UI
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Alerts will be sent when reaching 75% and 95% of the storage quota. This serves as a notification only and is not a
     * blocker
     */
    public readonly emailNotification!: pulumi.Output<boolean>;
    /**
     * Project group. Element has one to one mapping with the [JFrog Project Groups
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateGroupinProject)
     *
     * @deprecated Replaced by `jfrog-project.Group` resource. This should not be used in combination with `jfrog-project.Group` resource. Use `useProjectGroupResource` attribute to control which resource manages project roles.
     */
    public readonly groups!: pulumi.Output<outputs.ProjectGroup[] | undefined>;
    /**
     * The Project Key is added as a prefix to resources created within a Project. This field is mandatory and supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter. For example: `us1a-test`.
     */
    public readonly key!: pulumi.Output<string>;
    /**
     * Storage quota in GiB. Must be 1 or larger. Set to -1 for unlimited storage. This is translated to binary bytes for
     * Artifactory API. So for a 1TB quota, this should be set to 1024 (vs 1000) which will translate to 1099511627776 bytes
     * for the API.
     */
    public readonly maxStorageInGibibytes!: pulumi.Output<number>;
    /**
     * Member of the project. Element has one to one mapping with the [JFrog Project Users
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateUserinProject).
     *
     * @deprecated Replaced by `jfrog-project.User` resource. This should not be used in combination with `jfrog-project.User` resource. Use `useProjectUserResource` attribute to control which resource manages project roles.
     */
    public readonly members!: pulumi.Output<outputs.ProjectMember[] | undefined>;
    /**
     * (Optional) List of existing repo keys to be assigned to the project. If you wish to use the alternate method of setting
     * `projectKey` attribute in each `artifactory_*_repository` resource in the `artifactory` provider, you will need to use
     * `lifecycle.ignore_changes` in the `jfrog-project.Project` resource to avoid state drift. ```hcl lifecycle {
     * ignore_changes = [ repos ] } ```
     *
     * @deprecated Replaced by `jfrog-project.Repository` resource. This should not be used in combination with `jfrog-project.Repository` resource. Use `useProjectRepositoryResource` attribute to control which resource manages project repositories.
     */
    public readonly repos!: pulumi.Output<string[] | undefined>;
    /**
     * Project role. Element has one to one mapping with the [JFrog Project Roles
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole)
     *
     * @deprecated Replaced by `jfrog-project.Role` resource. This should not be used in combination with `jfrog-project.Role` resource. Use `useProjectRoleResource` attribute to control which resource manages project roles.
     */
    public readonly roles!: pulumi.Output<outputs.ProjectRole[] | undefined>;
    /**
     * When set to true, this resource will ignore the `group` attributes and allow users to be managed by
     * `jfrog-project.Group` resource instead. Default to `true`.
     */
    public readonly useProjectGroupResource!: pulumi.Output<boolean>;
    /**
     * When set to true, this resource will ignore the `repos` attributes and allow repository to be managed by
     * `jfrog-project.Repository` resource instead. Default to `true`.
     */
    public readonly useProjectRepositoryResource!: pulumi.Output<boolean>;
    /**
     * When set to true, this resource will ignore the `roles` attributes and allow roles to be managed by `jfrog-project.Role`
     * resource instead. Default to `true`.
     */
    public readonly useProjectRoleResource!: pulumi.Output<boolean>;
    /**
     * When set to true, this resource will ignore the `member` attributes and allow users to be managed by
     * `jfrog-project.User` resource instead. Default to `true`.
     */
    public readonly useProjectUserResource!: pulumi.Output<boolean>;

    /**
     * Create a Project resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProjectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ProjectArgs | ProjectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ProjectState | undefined;
            resourceInputs["adminPrivileges"] = state ? state.adminPrivileges : undefined;
            resourceInputs["blockDeploymentsOnLimit"] = state ? state.blockDeploymentsOnLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["emailNotification"] = state ? state.emailNotification : undefined;
            resourceInputs["groups"] = state ? state.groups : undefined;
            resourceInputs["key"] = state ? state.key : undefined;
            resourceInputs["maxStorageInGibibytes"] = state ? state.maxStorageInGibibytes : undefined;
            resourceInputs["members"] = state ? state.members : undefined;
            resourceInputs["repos"] = state ? state.repos : undefined;
            resourceInputs["roles"] = state ? state.roles : undefined;
            resourceInputs["useProjectGroupResource"] = state ? state.useProjectGroupResource : undefined;
            resourceInputs["useProjectRepositoryResource"] = state ? state.useProjectRepositoryResource : undefined;
            resourceInputs["useProjectRoleResource"] = state ? state.useProjectRoleResource : undefined;
            resourceInputs["useProjectUserResource"] = state ? state.useProjectUserResource : undefined;
        } else {
            const args = argsOrState as ProjectArgs | undefined;
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.key === undefined) && !opts.urn) {
                throw new Error("Missing required property 'key'");
            }
            resourceInputs["adminPrivileges"] = args ? args.adminPrivileges : undefined;
            resourceInputs["blockDeploymentsOnLimit"] = args ? args.blockDeploymentsOnLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["emailNotification"] = args ? args.emailNotification : undefined;
            resourceInputs["groups"] = args ? args.groups : undefined;
            resourceInputs["key"] = args ? args.key : undefined;
            resourceInputs["maxStorageInGibibytes"] = args ? args.maxStorageInGibibytes : undefined;
            resourceInputs["members"] = args ? args.members : undefined;
            resourceInputs["repos"] = args ? args.repos : undefined;
            resourceInputs["roles"] = args ? args.roles : undefined;
            resourceInputs["useProjectGroupResource"] = args ? args.useProjectGroupResource : undefined;
            resourceInputs["useProjectRepositoryResource"] = args ? args.useProjectRepositoryResource : undefined;
            resourceInputs["useProjectRoleResource"] = args ? args.useProjectRoleResource : undefined;
            resourceInputs["useProjectUserResource"] = args ? args.useProjectUserResource : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Project.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Project resources.
 */
export interface ProjectState {
    adminPrivileges?: pulumi.Input<pulumi.Input<inputs.ProjectAdminPrivilege>[]>;
    /**
     * Block deployment of artifacts if storage quota is exceeded.
     */
    blockDeploymentsOnLimit?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * Also known as project name on the UI
     */
    displayName?: pulumi.Input<string>;
    /**
     * Alerts will be sent when reaching 75% and 95% of the storage quota. This serves as a notification only and is not a
     * blocker
     */
    emailNotification?: pulumi.Input<boolean>;
    /**
     * Project group. Element has one to one mapping with the [JFrog Project Groups
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateGroupinProject)
     *
     * @deprecated Replaced by `jfrog-project.Group` resource. This should not be used in combination with `jfrog-project.Group` resource. Use `useProjectGroupResource` attribute to control which resource manages project roles.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.ProjectGroup>[]>;
    /**
     * The Project Key is added as a prefix to resources created within a Project. This field is mandatory and supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter. For example: `us1a-test`.
     */
    key?: pulumi.Input<string>;
    /**
     * Storage quota in GiB. Must be 1 or larger. Set to -1 for unlimited storage. This is translated to binary bytes for
     * Artifactory API. So for a 1TB quota, this should be set to 1024 (vs 1000) which will translate to 1099511627776 bytes
     * for the API.
     */
    maxStorageInGibibytes?: pulumi.Input<number>;
    /**
     * Member of the project. Element has one to one mapping with the [JFrog Project Users
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateUserinProject).
     *
     * @deprecated Replaced by `jfrog-project.User` resource. This should not be used in combination with `jfrog-project.User` resource. Use `useProjectUserResource` attribute to control which resource manages project roles.
     */
    members?: pulumi.Input<pulumi.Input<inputs.ProjectMember>[]>;
    /**
     * (Optional) List of existing repo keys to be assigned to the project. If you wish to use the alternate method of setting
     * `projectKey` attribute in each `artifactory_*_repository` resource in the `artifactory` provider, you will need to use
     * `lifecycle.ignore_changes` in the `jfrog-project.Project` resource to avoid state drift. ```hcl lifecycle {
     * ignore_changes = [ repos ] } ```
     *
     * @deprecated Replaced by `jfrog-project.Repository` resource. This should not be used in combination with `jfrog-project.Repository` resource. Use `useProjectRepositoryResource` attribute to control which resource manages project repositories.
     */
    repos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project role. Element has one to one mapping with the [JFrog Project Roles
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole)
     *
     * @deprecated Replaced by `jfrog-project.Role` resource. This should not be used in combination with `jfrog-project.Role` resource. Use `useProjectRoleResource` attribute to control which resource manages project roles.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.ProjectRole>[]>;
    /**
     * When set to true, this resource will ignore the `group` attributes and allow users to be managed by
     * `jfrog-project.Group` resource instead. Default to `true`.
     */
    useProjectGroupResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `repos` attributes and allow repository to be managed by
     * `jfrog-project.Repository` resource instead. Default to `true`.
     */
    useProjectRepositoryResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `roles` attributes and allow roles to be managed by `jfrog-project.Role`
     * resource instead. Default to `true`.
     */
    useProjectRoleResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `member` attributes and allow users to be managed by
     * `jfrog-project.User` resource instead. Default to `true`.
     */
    useProjectUserResource?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a Project resource.
 */
export interface ProjectArgs {
    adminPrivileges?: pulumi.Input<pulumi.Input<inputs.ProjectAdminPrivilege>[]>;
    /**
     * Block deployment of artifacts if storage quota is exceeded.
     */
    blockDeploymentsOnLimit?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    /**
     * Also known as project name on the UI
     */
    displayName: pulumi.Input<string>;
    /**
     * Alerts will be sent when reaching 75% and 95% of the storage quota. This serves as a notification only and is not a
     * blocker
     */
    emailNotification?: pulumi.Input<boolean>;
    /**
     * Project group. Element has one to one mapping with the [JFrog Project Groups
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateGroupinProject)
     *
     * @deprecated Replaced by `jfrog-project.Group` resource. This should not be used in combination with `jfrog-project.Group` resource. Use `useProjectGroupResource` attribute to control which resource manages project roles.
     */
    groups?: pulumi.Input<pulumi.Input<inputs.ProjectGroup>[]>;
    /**
     * The Project Key is added as a prefix to resources created within a Project. This field is mandatory and supports only 2 - 32 lowercase alphanumeric and hyphen characters. Must begin with a letter. For example: `us1a-test`.
     */
    key: pulumi.Input<string>;
    /**
     * Storage quota in GiB. Must be 1 or larger. Set to -1 for unlimited storage. This is translated to binary bytes for
     * Artifactory API. So for a 1TB quota, this should be set to 1024 (vs 1000) which will translate to 1099511627776 bytes
     * for the API.
     */
    maxStorageInGibibytes?: pulumi.Input<number>;
    /**
     * Member of the project. Element has one to one mapping with the [JFrog Project Users
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-UpdateUserinProject).
     *
     * @deprecated Replaced by `jfrog-project.User` resource. This should not be used in combination with `jfrog-project.User` resource. Use `useProjectUserResource` attribute to control which resource manages project roles.
     */
    members?: pulumi.Input<pulumi.Input<inputs.ProjectMember>[]>;
    /**
     * (Optional) List of existing repo keys to be assigned to the project. If you wish to use the alternate method of setting
     * `projectKey` attribute in each `artifactory_*_repository` resource in the `artifactory` provider, you will need to use
     * `lifecycle.ignore_changes` in the `jfrog-project.Project` resource to avoid state drift. ```hcl lifecycle {
     * ignore_changes = [ repos ] } ```
     *
     * @deprecated Replaced by `jfrog-project.Repository` resource. This should not be used in combination with `jfrog-project.Repository` resource. Use `useProjectRepositoryResource` attribute to control which resource manages project repositories.
     */
    repos?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Project role. Element has one to one mapping with the [JFrog Project Roles
     * API](https://www.jfrog.com/confluence/display/JFROG/Artifactory+REST+API#ArtifactoryRESTAPI-AddaNewRole)
     *
     * @deprecated Replaced by `jfrog-project.Role` resource. This should not be used in combination with `jfrog-project.Role` resource. Use `useProjectRoleResource` attribute to control which resource manages project roles.
     */
    roles?: pulumi.Input<pulumi.Input<inputs.ProjectRole>[]>;
    /**
     * When set to true, this resource will ignore the `group` attributes and allow users to be managed by
     * `jfrog-project.Group` resource instead. Default to `true`.
     */
    useProjectGroupResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `repos` attributes and allow repository to be managed by
     * `jfrog-project.Repository` resource instead. Default to `true`.
     */
    useProjectRepositoryResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `roles` attributes and allow roles to be managed by `jfrog-project.Role`
     * resource instead. Default to `true`.
     */
    useProjectRoleResource?: pulumi.Input<boolean>;
    /**
     * When set to true, this resource will ignore the `member` attributes and allow users to be managed by
     * `jfrog-project.User` resource instead. Default to `true`.
     */
    useProjectUserResource?: pulumi.Input<boolean>;
}
