// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface ProjectAdminPrivilege {
    /**
     * Enables a project admin to define the resources to be indexed by Xray
     */
    indexResources: pulumi.Input<boolean>;
    /**
     * Allows the Project Admin to manage Platform users/groups as project members with different roles.
     */
    manageMembers: pulumi.Input<boolean>;
    /**
     * Allows the Project Admin to manage resources - repositories, builds and Pipelines resources on the project level.
     */
    manageResources: pulumi.Input<boolean>;
}

export interface ProjectGroup {
    /**
     * Must be existing Artifactory group
     */
    name: pulumi.Input<string>;
    /**
     * List of pre-defined Project or custom roles
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ProjectMember {
    /**
     * Must be existing Artifactory user
     */
    name: pulumi.Input<string>;
    /**
     * List of pre-defined Project or custom roles
     */
    roles: pulumi.Input<pulumi.Input<string>[]>;
}

export interface ProjectRole {
    /**
     * List of pre-defined actions (READ*REPOSITORY, ANNOTATE*REPOSITORY, DEPLOY*CACHE*REPOSITORY, DELETE*OVERWRITE*REPOSITORY, MANAGE*XRAY*MD*REPOSITORY, READ*RELEASE*BUNDLE, ANNOTATE*RELEASE*BUNDLE, CREATE*RELEASE*BUNDLE, DISTRIBUTE*RELEASE*BUNDLE, DELETE*RELEASE*BUNDLE, MANAGE*XRAY*MD*RELEASE*BUNDLE, READ*BUILD, ANNOTATE*BUILD, DEPLOY*BUILD, DELETE*BUILD, MANAGE*XRAY*MD*BUILD, READ*SOURCES*PIPELINE, TRIGGER*PIPELINE, READ*INTEGRATIONS*PIPELINE, READ*POOLS*PIPELINE, MANAGE*INTEGRATIONS*PIPELINE, MANAGE*SOURCES*PIPELINE, MANAGE*POOLS*PIPELINE, TRIGGER*SECURITY, ISSUES*SECURITY, LICENCES*SECURITY, REPORTS*SECURITY, WATCHES*SECURITY, POLICIES*SECURITY, RULES*SECURITY, MANAGE*MEMBERS, MANAGE*RESOURCES)
     */
    actions: pulumi.Input<pulumi.Input<string>[]>;
    description?: pulumi.Input<string>;
    /**
     * A repository can be available in different environments. Members with roles defined in the set environment will have access to the repository. List of pre-defined environments (DEV, PROD)
     */
    environments: pulumi.Input<pulumi.Input<string>[]>;
    name: pulumi.Input<string>;
    /**
     * Type of role. Only "CUSTOM" is supported
     */
    type?: pulumi.Input<string>;
}