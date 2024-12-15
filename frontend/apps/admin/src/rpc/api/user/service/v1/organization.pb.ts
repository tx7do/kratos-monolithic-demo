// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: user/service/v1/organization.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 部门 */
export interface Organization {
  /** 部门ID */
  id?:
    | number
    | null
    | undefined;
  /** 部门名称 */
  name?:
    | string
    | null
    | undefined;
  /** 排序号 */
  orderNo?:
    | number
    | null
    | undefined;
  /** 状态 */
  status?:
    | string
    | null
    | undefined;
  /** 备注 */
  remark?:
    | string
    | null
    | undefined;
  /** 父节点ID */
  parentId?:
    | number
    | null
    | undefined;
  /** 子节点树 */
  children: Organization[];
  /** 创建时间 */
  createTime?:
    | Timestamp
    | null
    | undefined;
  /** 更新时间 */
  updateTime?:
    | Timestamp
    | null
    | undefined;
  /** 删除时间 */
  deleteTime?: Timestamp | null | undefined;
}

/** 部门列表 - 答复 */
export interface ListOrganizationResponse {
  items: Organization[];
  total: number;
}

/** 部门数据 - 请求 */
export interface GetOrganizationRequest {
  id: number;
}

/** 创建部门 - 请求 */
export interface CreateOrganizationRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  org: Organization | null;
}

/** 更新部门 - 请求 */
export interface UpdateOrganizationRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  org:
    | Organization
    | null;
  /** 要更新的字段列表 */
  updateMask:
    | string[]
    | null;
  /** 如果设置为true的时候，资源不存在则会新增(插入)，并且在这种情况下`updateMask`字段将会被忽略。 */
  allowMissing?: boolean | null | undefined;
}

/** 删除部门 - 请求 */
export interface DeleteOrganizationRequest {
  /** 操作用户ID */
  operatorId?: number | null | undefined;
  id: number;
}

/** 部门管理服务 */
export interface OrganizationService {
  /** 查询部门列表 */
  ListOrganization(request: PagingRequest): Promise<ListOrganizationResponse>;
  /** 查询部门详情 */
  GetOrganization(request: GetOrganizationRequest): Promise<Organization>;
  /** 创建部门 */
  CreateOrganization(request: CreateOrganizationRequest): Promise<Empty>;
  /** 更新部门 */
  UpdateOrganization(request: UpdateOrganizationRequest): Promise<Empty>;
  /** 删除部门 */
  DeleteOrganization(request: DeleteOrganizationRequest): Promise<Empty>;
}
