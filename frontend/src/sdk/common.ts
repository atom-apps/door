import { Loader } from "./axios"

export interface QueryFilter {
    name?: string
    age?: string
}
export interface PageFilterInterface {
    page?: number
    limit?: number
}

export class PageFilter {
    filter: PageFilterInterface

    constructor(filter?: PageFilterInterface) {
        this.filter = filter || {}
    }

    static default(): PageFilter {
        return new PageFilter({
            page: 1,
            limit: 10
        })
    }

    page(page: number): PageFilter {
        this.filter.page = page
        return this
    }

    limit(limit: number): PageFilter {
        this.filter.limit = limit
        return this
    }

    get(): PageFilterInterface {
        return this.filter
    }

}


export interface SortFilterInterface {
    asc?: string
    desc?: string
}

export class SortFilter {
    private ascFields: string[] = []
    private descFields: string[] = []

    static default(): SortFilter {
        return new SortFilter()
    }

    asc(...asc: string[]): SortFilter {
        this.ascFields.push(...asc)
        this.ascFields = Array.from(new Set(this.ascFields));
        return this
    }
    desc(...fields: string[]): SortFilter {
        this.descFields.push(...fields)
        this.descFields = Array.from(new Set(this.descFields));
        return this
    }

    get(): SortFilterInterface {
        const sortFilter: SortFilterInterface = {};

        if (this.ascFields.length > 0) {
            sortFilter.asc = this.ascFields.join(',');
        }

        if (this.descFields.length > 0) {
            sortFilter.desc = this.descFields.join(',');
        }

        return sortFilter;
    }
}

export interface OAuth2Token {
    accessToken: string;
    tokenType?: string;
    refreshToken?: string;
    expiry?: Date;
}


export abstract class BaseRequest {
    public get loading(): boolean {
        return Loader.loading
    }

    // pageFilter
    private _pageFilter?: PageFilterInterface | undefined
    public get pageFilter(): PageFilterInterface | undefined {
        return this._pageFilter
    }
    public set pageFilter(value: PageFilterInterface | undefined) {
        this._pageFilter = value
    }


    // sort filter
    private _sortFilter?: SortFilterInterface | undefined
    public get sortFilter(): SortFilterInterface | undefined {
        return this._sortFilter
    }
    public set sortFilter(value: SortFilterInterface | undefined) {
        this._sortFilter = value
    }
}

export interface GlobalVar {
}

export interface AuthScopeResponse {
    code: string;
    scope: string;
}