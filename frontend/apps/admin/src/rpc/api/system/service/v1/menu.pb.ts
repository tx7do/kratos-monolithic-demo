// Code generated by protoc-gen-ts_proto. DO NOT EDIT.
// versions:
//   protoc-gen-ts_proto  v2.6.0
//   protoc               unknown
// source: system/service/v1/menu.proto

/* eslint-disable */
import { type Empty } from "../../../google/protobuf/empty.pb";
import { type Timestamp } from "../../../google/protobuf/timestamp.pb";
import { type PagingRequest } from "../../../pagination/v1/pagination.pb";

/** 菜单类型 */
export enum MenuType {
  /** FOLDER - 菜单夹 */
  FOLDER = "FOLDER",
  /** MENU - 菜单项 */
  MENU = "MENU",
  /** BUTTON - 按钮 */
  BUTTON = "BUTTON",
}

/** 菜单 */
export interface Menu {
  id?: number | null | undefined;
  parentId?: number | null | undefined;
  status?: string | null | undefined;
  type?: MenuType | null | undefined;
  children: Menu[];
  path?: string | null | undefined;
  redirect?: string | null | undefined;
  alias?: string | null | undefined;
  name?: string | null | undefined;
  component?: string | null | undefined;
  meta?: RouteMeta | null | undefined;
  createTime?: Timestamp | null | undefined;
  updateTime?: Timestamp | null | undefined;
  deleteTime?: Timestamp | null | undefined;
}

/** 路由元信息 */
export interface RouteMeta {
  activeIcon?: string | null | undefined;
  activePath?: string | null | undefined;
  affixTab?: boolean | null | undefined;
  affixTabOrder?: number | null | undefined;
  authority: string[];
  badge?: string | null | undefined;
  badgeType?: string | null | undefined;
  badgeVariants?: string | null | undefined;
  hideChildrenInMenu?: boolean | null | undefined;
  hideInBreadcrumb?: boolean | null | undefined;
  hideInMenu?: boolean | null | undefined;
  hideInTab?: boolean | null | undefined;
  icon?: string | null | undefined;
  iframeSrc?: string | null | undefined;
  ignoreAccess?: boolean | null | undefined;
  keepAlive?: boolean | null | undefined;
  link?: string | null | undefined;
  loaded?: boolean | null | undefined;
  maxNumOfOpenTab?: number | null | undefined;
  menuVisibleWithForbidden?: boolean | null | undefined;
  openInNewWindow?: boolean | null | undefined;
  order?: number | null | undefined;
  title?: string | null | undefined;
}

/** 查询菜单列表 - 回应 */
export interface ListMenuResponse {
  items: Menu[];
  total: number;
}

/** 查询菜单详情 - 请求 */
export interface GetMenuRequest {
  id: number;
}

/** 创建菜单 - 请求 */
export interface CreateMenuRequest {
  operatorId?: number | null | undefined;
  menu: Menu | null;
}

/** 更新菜单 - 请求 */
export interface UpdateMenuRequest {
  operatorId?: number | null | undefined;
  menu: Menu | null;
  updateMask: string[] | null;
  allowMissing?: boolean | null | undefined;
}

/** 删除菜单 - 请求 */
export interface DeleteMenuRequest {
  operatorId?: number | null | undefined;
  id: number;
}

/** 后台菜单服务 */
export interface MenuService {
  /** 查询菜单列表 */
  ListMenu(request: PagingRequest): Promise<ListMenuResponse>;
  /** 查询菜单详情 */
  GetMenu(request: GetMenuRequest): Promise<Menu>;
  /** 创建菜单 */
  CreateMenu(request: CreateMenuRequest): Promise<Empty>;
  /** 更新菜单 */
  UpdateMenu(request: UpdateMenuRequest): Promise<Empty>;
  /** 删除菜单 */
  DeleteMenu(request: DeleteMenuRequest): Promise<Empty>;
}